package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/cors"
)

type App struct {
	ch *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

	app := &App{}

	r := chi.NewRouter()

	// initialize a ampq connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	app.ch = ch

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
