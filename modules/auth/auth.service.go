package auth

import (
	"ecommerce/errors"
	"ecommerce/modules/user"
	"ecommerce/utils"
	t "ecommerce/utils/token"
	"net/http"
	"reflect"
)

type Service struct {
	// repo Repository
	userRepo user.Repository
}

func (s *Service) handleSignIn(body *user.LoginUserReq) (interface{}, error) {
	op := errors.Op("auth.service.handleSignIn")

	var data struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	}

	params := &user.QueryParams{
		TableName: "users",
		Columns:   []string{"id", "password"},
		Where:     "email ILIKE $1",
		Args:      []interface{}{body.Email},
	}

	if err := s.userRepo.DetailByConditions(&data, params); err != nil {
		return nil, err
	}

	if reflect.ValueOf(data).IsZero() {
		return nil, errors.E(op, http.StatusBadRequest, "account not found")
	}

	match := utils.CompareHashPassword(body.Password, data.Password)
	if !match {
		return nil, errors.E(op, http.StatusBadRequest, "wrong password!")
	}

	return map[string]string{"access_token": t.SignToken(data.ID)}, nil
}

func (s *Service) handleSignUp(body *user.SignUpUserReq) (interface{}, error) {
	op := errors.Op("auth.service.handleSignUp")

	var data struct {
		ID string `json:"id"`
	}

	params := &user.QueryParams{
		TableName: "users",
		Columns:   []string{"id"},
		Where:     "email ILIKE $1",
		Args:      []interface{}{body.Email},
	}

	if err := s.userRepo.DetailByConditions(&data, params); err != nil {
		return nil, err
	}

	if !reflect.ValueOf(data).IsZero() {
		return nil, errors.E(op, http.StatusBadRequest, "account have exists")
	}

	body.Password = utils.HashPassword(body.Password)

	id, err := s.userRepo.InsertOne(*body)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id": id,
	}, nil
}
