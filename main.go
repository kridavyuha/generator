package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

type App struct {
}

func main() {

	app := &App{}

	r := chi.NewRouter()

	// CORS middleware configuration
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"}, // Your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}).Handler)

	// Define routes
	r.Get("/fixtures", app.GetFixtures)
	r.Get("/scores", app.GetScores)
	r.Get("/squad", app.GetSquads)

	if err := http.ListenAndServe(":8081", r); err != nil {
		panic(err)
	}

}
