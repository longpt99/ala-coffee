package product

import (
	"errors"
)

type Service struct {
	repo Repository
}

func (s *Service) HandlerGetProducts() (interface{}, error) {
	var data, err = s.repo.List()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Service) HandlerGetProduct(id string) (*Product, error) {
	var data, err = s.repo.DetailByID(id)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("Product Not Found")
	}

	return data, nil
}

func (s *Service) HandlerDeleteProduct(id string) (interface{}, error) {
	err := s.repo.Delete(id)

	if err != nil {
		return nil, err
	}

	return map[string]bool{
		"is_succeed": true,
	}, nil
}

func (s *Service) HandlerCreateProduct(body *CreateProductReq) (interface{}, error) {
	id, err := s.repo.InsertOne(body)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id": id,
	}, nil
}

func (s *Service) HandlerUpdateProduct(body *CreateProductReq) interface{} {
	// result := s.repo.InsertOne(body.Name, body.Description)

	return map[string]string{
		"id": "1",
	}
}
