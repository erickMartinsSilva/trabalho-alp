package routes

import (
	"net/http"
	"ALP/src/controllers"
)

func SetupRoutes(){
	http.HandleFunc("/users", controllers.CreateUser)
}