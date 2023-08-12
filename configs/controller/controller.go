package controller

import (
	"log"

	"github.com/go-chi/chi/v5"

	"ecommerce/configs/repository"
	"ecommerce/modules/admin"
	"ecommerce/modules/auth"
	"ecommerce/modules/product"
	"ecommerce/modules/user"
)

type Controllers struct {
	productController *product.Controller
	adminController   *admin.Controller
	userController    *user.Controller
	authController    *auth.Controller
}

func InitControllers(repo *repository.Repository, r chi.Router) *Controllers {
	log.Println("Init Controllers Successfully! 🚀")

	return &Controllers{
		productController: product.InitController(r, repo.ProductRepo),
		adminController:   admin.InitController(r, repo.AdminRepo),
		userController:    user.InitController(r, repo.UserRepo),
		authController:    auth.InitController(r, repo.UserRepo),
	}
}
