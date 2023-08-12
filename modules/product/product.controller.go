package product

import (
	res "ecommerce/utils/response"

	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	service *Service
}

func InitController(r chi.Router, repo Repository) *Controller {
	h := &Controller{
		&Service{
			repo,
		},
	}

	r.Route("/products", func(r chi.Router) {
		r.Get("/", h.handlerGetProducts)
		r.Post("/", h.handlerCreateProduct)
		r.Get("/{id}", h.handlerGetProduct)
		r.Patch("/{id}", h.handlerUpdateProduct)
		r.Delete("/{id}", h.handlerDeleteProduct)
	})

	return h
}

// @Router /products [GET]
// @Tags Products
// @Summary Get a list of products
// @Description Get all products
// @Produce json
func (h *Controller) handlerGetProducts(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.HandlerGetProducts()

	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}

// @Router /products/:id [GET]
// @Tags Products
// @Summary Get a product detail
// @Produce json
// @Param	id	path	string  true  "ID"
func (h *Controller) handlerGetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result, err := h.service.HandlerGetProduct(id)

	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}

// @Router /products/:id [DELETE]
// @Tags Products
// @Summary Delete product by ID
// @Produce json
// @Param	id	path	string  true  "ID"
func (h *Controller) handlerDeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	result, err := h.service.HandlerDeleteProduct(id)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}

func (h *Controller) handlerCreateProduct(w http.ResponseWriter, r *http.Request) {
	var body CreateProductReq

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	validate := validator.New()

	err = validate.Struct(body)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	result, err := h.service.HandlerCreateProduct(&body)
	if err != nil {
		// Handle validation errors
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}

func (h *Controller) handlerUpdateProduct(w http.ResponseWriter, r *http.Request) {
	var body CreateProductReq

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	validate := validator.New()

	err = validate.Struct(body)
	if err != nil {
		// Handle validation errors
		res.WriteError(w, r, err)
		return
	}

	result, err := h.service.HandlerCreateProduct(&body)
	if err != nil {
		// Handle validation errors
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}
