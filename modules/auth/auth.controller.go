package auth

import (
	"ecommerce/errors"
	"ecommerce/modules/user"
	res "ecommerce/utils/response"
	"ecommerce/utils/validate"

	"net/http"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	service *Service
}

func InitController(r chi.Router, userRepo user.Repository) *Controller {
	h := &Controller{
		&Service{
			userRepo,
		},
	}

	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-in", h.handleSignIn)
		r.Post("/sign-up", h.handleSignUp)
	})

	return h
}

func (h *Controller) handleSignIn(w http.ResponseWriter, r *http.Request) {
	op := errors.Op("auth.controller.handleSignIn")

	var body user.LoginUserReq

	if err := validate.ReadValid(&body, r); err != nil {
		res.WriteError(w, r, errors.E(op, err))
		return
	}

	result, err := h.service.handleSignIn(&body)

	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}

func (h *Controller) handleSignUp(w http.ResponseWriter, r *http.Request) {
	op := errors.Op("auth.controller.handleSignUp")

	var body user.SignUpUserReq

	if err := validate.ReadValid(&body, r); err != nil {
		res.WriteError(w, r, errors.E(op, err))
		return
	}

	result, err := h.service.handleSignUp(&body)

	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Write(w, r, result, http.StatusOK)
}

// func (h *Controller) handlerGetUser(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	result, err := h.service.HandlerGetUser(id)

// 	if err != nil {
// 		res.WriteError(w, r, err)
// 		return
// 	}

// 	res.Write(w, r, result, http.StatusOK)
// }

// func (h *Controller) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
// 	result, err := h.service.HandlerDeleteUser(chi.URLParam(r, "id"))

// 	if err != nil {
// 		res.WriteError(w, r, err)
// 		return
// 	}

// 	res.Write(w, r, result, http.StatusOK)
// }

// func (h *Controller) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
// 	var body CreateUserReq

// 	err := json.NewDecoder(r.Body).Decode(&body)
// 	if err != nil {
// 		res.WriteError(w, r, err)
// 		return
// 	}

// 	validate := validator.New()

// 	err = validate.Struct(body)
// 	if err != nil {
// 		res.WriteError(w, r, err)
// 		return
// 	}

// 	result := h.service.HandlerCreateUser(&body)
// 	res.Write(w, r, result, http.StatusOK)
// }

// func (h *Controller) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
// 	var body CreateUserReq
// 	// var id = chi.URLParam(r, "id")

// 	err := json.NewDecoder(r.Body).Decode(&body)
// 	if err != nil {
// 		res.WriteError(w, r, err)
// 		return
// 	}

// 	validate := validator.New()

// 	err = validate.Struct(body)
// 	if err != nil {
// 		// Handle validation errors
// 		res.WriteError(w, r, err)
// 		return
// 	}

// 	result := h.service.HandlerCreateUser(&body)
// 	res.Write(w, r, result, http.StatusOK)
// }
