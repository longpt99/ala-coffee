// Main function
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"ecommerce/configs"
	"ecommerce/configs/database"
	"ecommerce/configs/repository"
	"ecommerce/configs/router"
	"ecommerce/utils/validate"

	_ "ecommerce/docs"
)

// @title My API
// @version 1.0
// @description This is a sample API for demonstration purposes.
// @BasePath /
func main() {
	// Load env config
	timeWait := 10 * time.Second

	err := configs.InitConfig()
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	err = validate.RegisterValidation()
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeWait)
	defer cancel()

	// Load database
	store := database.InitDatabase(ctx)
	defer store.Close()

	repo := repository.InitRepositories(store)
	p := configs.Env.Port
	r := router.New(repo)
	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", p),
		Handler:           r,
		ReadHeaderTimeout: 1 * time.Second,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	// Create a channel to receive signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Start the server in a separate goroutine
	go func() {
		log.Printf("Server is listening on %s\n", s.Addr)

		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for a signal to shutdown the server
	sig := <-signalCh
	log.Printf("[Server] Received signal: %v\n", sig)

	// Shutdown the server gracefully
	if err := s.Shutdown(ctx); err != nil {
		log.Printf("[Server] Shutdown Failed: %v\n", err)
		return
	}

	log.Println("[Server] Shutdown Gracefully! 🚀")
}

func errorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v\n", r)
				debug.PrintStack()
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResponse{"Internal Server Error"})
			}
		}()

		next.ServeHTTP(w, r)
		// if err != nil {
		// 	log.Printf("Error: %v\n", err)
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	json.NewEncoder(w).Encode(ErrorResponse{"Bad Request"})
		// }
	})
}

// ErrorResponse describe error model
type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}
