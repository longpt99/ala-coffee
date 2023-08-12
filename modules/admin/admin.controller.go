package admin

import (
	"ecommerce/errors"
	"ecommerce/middleware"
	res "ecommerce/utils/response"
	"ecommerce/utils/validate"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	service *Service
}

func InitController(r chi.Router, repo Repository) *Controller {
	handler := &Controller{
		&Service{
			repo,
		}}

	r.Route("/admins", func(r chi.Router) {
		r.Post("/", handler.handlerCreateAdmin)
		r.Get("/{id}", handler.handlerGetAdmin)

		r.Post("/login", handler.handlerLoginAdmin)
		r.Get("/profile", middleware.AuthMiddleware(handler.handlerGetAdmins))
	})

	return handler
}

func (h *Controller) handlerGetAdmins(w http.ResponseWriter, r *http.Request) {
}

func (h *Controller) handlerGetAdmin(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	result, err := h.service.HandlerGetAdmin(id)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}

func (h *Controller) handlerDeleteAdmin(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	result, err := h.service.HandlerDeleteAdmin(id)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusNoContent)
}

func (h *Controller) handlerLoginAdmin(w http.ResponseWriter, r *http.Request) {
	op := errors.Op("admin.controller.handlerLoginAdmin")

	var body LoginAdminReq

	if err := validate.ReadValid(&body, r); err != nil {
		res.WriteError(w, r, errors.E(op, err))
		return
	}

	result, err := h.service.HandlerLoginAdmin(&body)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}

func (h *Controller) handlerUpdateAdmin(w http.ResponseWriter, r *http.Request) {
	var body CreateAdminReq
	// var id = chi.URLParam(r, "id")

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

	result, err := h.service.HandlerCreateAdmin(&body)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusNoContent)
}

func (h *Controller) handlerCreateAdmin(w http.ResponseWriter, r *http.Request) {
	var body CreateAdminReq

	validate := validator.New()

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	err = validate.Struct(body)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	result, err := h.service.HandlerCreateAdmin(&body)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}
