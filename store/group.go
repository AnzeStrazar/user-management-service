package store

import (
	"user-management-service/model"
)

func (s *Store) GetGroups() ([]model.Group, error) {
	groups := make([]model.Group, 0)
	rows, err := s.db.GetGroups()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		group := model.Group{}
		err = rows.Scan(
			&group.GroupID,
			&group.Name,
		)
		if err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (s *Store) GetGroup(groupID int) (*model.Group, error) {
	row := s.db.GetGroup(groupID)
	var group model.Group
	err := row.Scan(
		&group.GroupID,
		&group.Name,
	)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (s *Store) GetGroupWithUsers(groupID int) ([]model.User, error) {
	users := make([]model.User, 0)
	rows, err := s.db.GetGroupWithUsers(groupID)
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

func (s *Store) CreateGroup(name string) (model.Group, error) {
	row := s.db.CreateGroup(name)

	group := model.Group{
		Name: name,
	}

	err := row.Scan(
		&group.GroupID,
	)

	return group, err
}

func (s *Store) ModifyGroup(groupID int, group model.Group) error {
	return s.db.ModifyGroup(groupID, group)
}

func (s *Store) RemoveGroup(groupID int) error {
	return s.db.RemoveGroup(groupID)
}
