package repository

import (
	"ecommerce/configs/database"
	"ecommerce/modules/admin"
	"ecommerce/modules/product"
	"ecommerce/modules/user"
	"log"
)

type RepositoryConfig interface {
	InitRepositories(*database.PostgresConfig) *Repository
}

type Repository struct {
	ProductRepo product.Repository
	AdminRepo   admin.Repository
	UserRepo    user.Repository
}

func InitRepositories(store *database.PostgresConfig) *Repository {
	log.Println("Init Repositories Successfully! 🚀")

	return &Repository{
		ProductRepo: product.InitRepository(store),
		AdminRepo:   admin.InitRepository(store),
		UserRepo:    user.InitRepository(store),
	}
}
