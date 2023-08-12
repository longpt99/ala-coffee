package user

import (
	"ecommerce/errors"
	"ecommerce/middleware"
	res "ecommerce/utils/response"
	t "ecommerce/utils/token"
	"ecommerce/utils/validate"
	"fmt"

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

	r.Route("/admin/users", func(r chi.Router) {
		r.Get("/", h.handlerGetUsers)
		r.Post("/", h.handlerCreateUser)
		r.Get("/{id}", h.handlerGetUser)
		r.Patch("/{id}", h.handlerUpdateUser)
		r.Delete("/{id}", h.handlerDeleteUser)
	})

	r.Route("/user", func(r chi.Router) {
		r.Get("/profile", middleware.AuthMiddleware(h.handleGetProfile))
		r.Patch("/profile", middleware.AuthMiddleware(h.handleUpdateProfile))
	})

	return h
}

func (h *Controller) handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.HandlerGetUsers()

	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}

func (h *Controller) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result, err := h.service.HandlerGetUser(id)

	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}

func (h *Controller) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.HandlerDeleteUser(chi.URLParam(r, "id"))

	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}

func (h *Controller) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	var body CreateUserReq

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

	result := h.service.HandlerCreateUser(&body)
	res.Write(w, r, result, http.StatusOK)
}

func (h *Controller) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	var body CreateUserReq
	// var id = chi.URLParam(r, "id")

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

	result := h.service.HandlerCreateUser(&body)
	res.Write(w, r, result, http.StatusOK)
}

func (h *Controller) handleGetProfile(w http.ResponseWriter, r *http.Request) {
	op := errors.Op("user.controller.handleGetProfile")

	// Retrieve the claims from the request context
	claims := t.GetPayload(r)
	if claims == nil {
		res.WriteError(w, r, errors.E(op, http.StatusBadRequest, "claims not found"))
		return
	}

	result, err := h.service.HandlerGetProfile(claims.ID)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}

func (h *Controller) handleUpdateProfile(w http.ResponseWriter, r *http.Request) {
	op := errors.Op("user.controller.handleUpdateProfile")

	var body UpdateUserProfileReq

	if err := validate.ReadValid(&body, r); err != nil {
		res.WriteError(w, r, err)
		return
	}

	// Retrieve the claims from the request context
	claims, ok := r.Context().Value("claims").(*t.Claims)
	if !ok {
		res.WriteError(w, r, errors.E(op, http.StatusBadRequest, "claims not found"))
		return
	}

	fmt.Println(body)

	result, err := h.service.HandlerUpdateProfile(claims.ID, body)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}
