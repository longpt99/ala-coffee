package user

import (
	"errors"
)

type Service struct {
	repo Repository
}

func (s *Service) HandlerGetUsers() (interface{}, error) {
	var data, err = s.repo.List()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Service) HandlerGetUser(id string) (*User, error) {
	var data, err = s.repo.DetailByID(id)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("User Not Found")
	}

	return data, nil
}

func (s *Service) HandlerDeleteUser(id string) (interface{}, error) {
	err := s.repo.Delete(id)

	if err != nil {
		return nil, err
	}

	return map[string]bool{
		"is_succeed": true,
	}, nil
}

func (s *Service) HandlerCreateUser(body *CreateUserReq) interface{} {
	// result := s.repo.InsertOne(body.Name, body.Description)

	return map[string]string{
		"id": "!",
	}
}

func (s *Service) HandlerUpdateUser(body *CreateUserReq) interface{} {
	// result := s.repo.InsertOne(body.Name, body.Description)

	return map[string]string{
		"id": "!",
	}
}

func (s *Service) HandlerGetProfile(id string) (interface{}, error) {
	var data struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	params := &QueryParams{
		TableName: "users",
		Columns:   []string{"id", "name", "email"},
		Where:     "id = $1",
		Args:      []interface{}{id},
	}

	if err := s.repo.DetailByConditions(&data, params); err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *Service) HandlerUpdateProfile(id string, body interface{}) (interface{}, error) {
	var data struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	params := &QueryParams{
		TableName: "users",
		Columns:   []string{"id", "name", "email"},
		Where:     "id = $1",
		Args:      []interface{}{id},
	}

	if err := s.repo.DetailByConditions(&data, params); err != nil {
		return nil, err
	}

	return &data, nil
}
