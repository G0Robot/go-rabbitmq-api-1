package routes

import (
	"go-rabbitmq-api-1/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(r *mux.Router) {
	// Route Handlers / Endpoints
	// HandleFunc sets up the routers and which methods will handle them
	r.HandleFunc("/api/messages", controllers.SendMessage).Methods("POST")
}
