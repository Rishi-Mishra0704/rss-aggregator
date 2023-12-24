package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	router := chi.NewRouter()
	port := os.Getenv("PORT")

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	v1Router := chi.NewRouter()

	v1Router.HandleFunc("/ready", handlerReady)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("Listening on port %s", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
