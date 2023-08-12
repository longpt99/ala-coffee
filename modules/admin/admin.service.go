package admin

import (
	"ecommerce/errors"
	"ecommerce/utils"
	t "ecommerce/utils/token"
	"net/http"
	"reflect"
)

type Service struct {
	repo Repository
}

func (s *Service) HandlerGetAdmins() (interface{}, error) {
	data, err := s.repo.List()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Service) HandlerGetAdmin(id string) (*Admin, error) {
	data, err := s.repo.DetailByID(id)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.E("Admin Not Found", http.StatusBadRequest)
	}

	return data, nil
}

func (s *Service) HandlerDeleteAdmin(id string) (interface{}, error) {
	err := s.repo.Delete(id)

	if err != nil {
		return nil, err
	}

	return map[string]bool{"is_succeed": true}, nil
}

func (s *Service) HandlerCreateAdmin(body *CreateAdminReq) (interface{}, error) {
	params := CreateAdminParams{
		CreateAdminReq: *body,
		Password:       utils.HashPassword("123456"),
	}

	id, err := s.repo.InsertOne(params)

	if err != nil {
		return nil, err
	}

	return map[string]*string{"id": id}, nil
}

func (s *Service) HandlerLoginAdmin(body *LoginAdminReq) (interface{}, error) {
	op := errors.Op("HandlerLoginAdmin")

	var data struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	}

	params := &QueryParams{
		TableName: "admins",
		Columns:   []string{"id", "password"},
		Where:     "email = $1",
		Args:      []interface{}{body.Email},
	}

	if err := s.repo.DetailByConditions(params, &data); err != nil {
		return nil, err
	}

	if reflect.ValueOf(data).IsZero() {
		return nil, errors.E(op, http.StatusBadRequest, "account not found")
	}

	if match := utils.CompareHashPassword(body.Password, data.Password); !match {
		return nil, errors.E(op, http.StatusBadRequest, "wrong password")
	}

	return map[string]string{
		"access_token": t.SignToken(data.ID),
	}, nil
}

func (s *Service) HandlerUpdateAdmin(body *CreateAdminReq) {
	// id := s.repo.InsertOne(body.Name, body.Description)

	// return map[string]string{
	// 	"id": id,
	// }
}
