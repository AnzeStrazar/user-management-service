package store

import "user-management-service/database"

type Store struct {
	db database.Postgres
}

func NewStore(db database.Postgres) *Store {
	store := &Store{
		db: db,
	}

	return store
}
