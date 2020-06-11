package store

import (
	"user-management-service/model"
)

func (s *Store) GetUsers() ([]model.User, error) {
	users := make([]model.User, 0)
	rows, err := s.db.GetUsers()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := model.User{}
		err = rows.Scan(
			&user.UserID,
			&user.GroupID,
			&user.Email,
			&user.Password,
			&user.Name,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Store) GetUser(userID int) (*model.User, error) {
	row := s.db.GetUser(userID)
	var user model.User
	err := row.Scan(
		&user.UserID,
		&user.GroupID,
		&user.Email,
		&user.Password,
		&user.Name,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Store) CreateUser(user_group_id int, email, password, name string) (model.User, error) {
	row := s.db.CreateUser(user_group_id, email, password, name)

	user := model.User{
		GroupID:  user_group_id,
		Email:    email,
		Password: password,
		Name:     name,
	}

	err := row.Scan(
		&user.UserID,
	)

	return user, err
}

func (s *Store) ModifyUser(userID int, user model.User) error {
	return s.db.ModifyUser(userID, user)

}

func (s *Store) RemoveUser(userID int) error {
	return s.db.RemoveUser(userID)
}
