package main

import (
	"go-rabbitmq-api-1/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	StartWebService()
}

func StartWebService() {
	// init router
	r := mux.NewRouter()
	// call to set up route handlers
	routes.RegisterRoutes(r)

	// Add a listener on port 8000, which needs to use our router
	log.Printf("Webservice up and running...")
	log.Fatal(http.ListenAndServe(":8000", r))
	log.Printf("--------------------")
}
