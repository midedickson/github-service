package routes

import (
	"github.com/gorilla/mux"
	"github.com/midedickson/github-service/controllers"
)

func ConnectRoutes(r *mux.Router, controller *controllers.Controller) {
	r.HandleFunc("/register", controller.CreateUser).Methods("POST")
	r.HandleFunc("/{owner}/repo/{repo}", controller.GetRepositoryInfo).Methods("GET")
	r.HandleFunc("/{owner}/repo/{repo}/commits", controller.GetRepositoryCommits).Methods("GET")
}
