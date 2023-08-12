package router

import (
	"ecommerce/configs/controller"
	"ecommerce/configs/repository"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "ecommerce/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func New(repo *repository.Repository) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/documentation/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/documentation/doc.json"), // The Swagger JSON endpoint URL
	))

	r.Route("/api/v1", func(r chi.Router) {
		controller.InitControllers(repo, r)
	})

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("[%s]: '%s' has been registered!\n", method, route)
		return nil
	})

	return r
}
