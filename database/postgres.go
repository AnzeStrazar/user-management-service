package database

import (
	"database/sql"
	"fmt"
	"log"
	"user-management-service/model"

	_ "github.com/lib/pq"
)

const (
	createUserTable = `
		CREATE TABLE IF NOT EXISTS "app_user" (
			"user_id" serial PRIMARY KEY,
			"user_group_id" INT REFERENCES "app_group"("group_id"),
			"email" VARCHAR (255)  NOT NULL,
			"password" VARCHAR (255) NOT NULL,
			"name" VARCHAR (255) NOT NULL
		 );`
	createGroupTable = `
		 CREATE TABLE IF NOT EXISTS "app_group" (
			"group_id" serial PRIMARY KEY,
			"name" VARCHAR (255)  NOT NULL
		  );`
	dropUserTable  = `DROP TABLE "app_user";`
	dropGroupTable = `DROP TABLE "app_group";`
	getUsers       = `
		SELECT
			"user_id",
			"user_group_id",
			"email",
			"password",
			"name"
		FROM "app_user";`
	getUser = `
		SELECT
			"user_id",
			"user_group_id",
			"email",
			"password",
			"name"
		FROM "app_user"
		WHERE "user_id" = $1;`
	userInsertion = `
		INSERT INTO "app_user"
			("user_group_id", "email", "password", "name")
		VALUES
			($1, $2, $3, $4)
		RETURNING "user_id";`
	modifyUser = `
		UPDATE "app_user"
		SET	"email" = $2,
			"password" = $3,
			"name" = $4
		WHERE "user_id" = $1;`
	removeUser = `
		DELETE FROM "app_user"
		WHERE "user_id" = $1;`
	getGroups = `
		SELECT
			"group_id",
			"name"
		FROM "app_group";`
	getGroup = `
		SELECT
			"group_id",
			"name"
		FROM "app_group"
		WHERE "group_id" = $1;`
	groupInsertion = `
		INSERT INTO "app_group"
			("name")
		VALUES
			($1)
		RETURNING "group_id";`
	modifyGroup = `
		UPDATE "app_group"
		SET	"name" = $2
		WHERE "group_id" = $1;`
	removeGroup = `
		DELETE FROM "app_group"
		WHERE "group_id" = $1;`
	getGroupWithUsers = `
		SELECT
			"user_id",
			"user_group_id",
			"email",
			"password",
			"name"
		FROM "app_user"
		WHERE "user_group_id" = $1;`
)

type Postgres struct {
	Db *sql.DB
}

func NewPostgres(dbhost, dbport, dbuser, dbpass, dbname string) Postgres {
	postgres := Postgres{}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbhost, dbport,
		dbuser, dbpass, dbname)
	log.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Database connection established!")

	postgres.Db = db

	return postgres
}

func (db *Postgres) CreateUserTable() {
	stmt, err := db.Db.Prepare(createUserTable)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	log.Println("User table created.")
}

func (db *Postgres) CreateGroupTable() {
	stmt, err := db.Db.Prepare(createGroupTable)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	log.Println("Group table created.")
}

func (db *Postgres) DropUserTable() {
	stmt, err := db.Db.Prepare(dropUserTable)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	log.Println("User table dropped.")
}

func (db *Postgres) DropGroupTable() {
	stmt, err := db.Db.Prepare(dropGroupTable)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	log.Println("Group table dropped.")
}

func (db *Postgres) GetUsers() (*sql.Rows, error) {
	return db.Db.Query(getUsers)
}

func (db *Postgres) GetUser(userID int) *sql.Row {
	return db.Db.QueryRow(getUser, userID)
}

func (db *Postgres) CreateUser(user_group_id int, email, password, name string) *sql.Row {
	return db.Db.QueryRow(userInsertion, user_group_id, email, password, name)

}

func (db *Postgres) ModifyUser(userID int, user model.User) error {
	result, err := db.Db.Exec(modifyUser, userID, user.Email, user.Password, user.Name)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return err
	}

	return nil
}

func (db *Postgres) RemoveUser(userID int) error {
	result, err := db.Db.Exec(removeUser, userID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return err
	}

	return nil
}

func (db *Postgres) GetGroups() (*sql.Rows, error) {
	rows, err := db.Db.Query(getGroups)

	return rows, err
}

func (db *Postgres) GetGroup(groupID int) *sql.Row {
	row := db.Db.QueryRow(getGroup, groupID)

	return row
}

func (db *Postgres) CreateGroup(name string) *sql.Row {
	return db.Db.QueryRow(groupInsertion, name)
}

func (db *Postgres) ModifyGroup(groupID int, group model.Group) error {
	result, err := db.Db.Exec(modifyGroup, groupID, group.Name)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return err
	}

	return nil
}

func (db *Postgres) RemoveGroup(groupID int) error {
	result, err := db.Db.Exec(removeGroup, groupID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return err
	}

	return nil
}

func (db *Postgres) GetGroupWithUsers(groupID int) (*sql.Rows, error) {
	return db.Db.Query(getGroupWithUsers, groupID)
}
