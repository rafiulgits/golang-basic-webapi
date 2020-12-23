package main

import (
	"golearn/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	HttpGET    = "GET"
	HttpPOST   = "POST"
	HttpPUT    = "PUT"
	HttpDELETE = "DELETE"
)

func main() {
	//creaing router to handle route paths of http server
	router := chi.NewRouter()

	// map endpoints in router with functional handlers
	router.Get("/", handlers.GetItem)
	router.Post("/", handlers.CreateItem)
	router.Put("/", handlers.UpdateItem)
	router.Delete("/", handlers.DeleteItem)

	//map endpoints in router with sub-router handlers
	userHandler := handlers.NewUserHandler()
	router.Route("/users", userHandler.Handle)

	// :8090 meaning serve all hosts available in local machine
	// or simply 0.0.0.0:8090
	// using router to handle http server requests
	http.ListenAndServe(":8090", router)
}
