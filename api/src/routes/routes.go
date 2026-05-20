package routes

import (
	"net/http"
	"ALP/src/controllers"
)

func SetupRoutes(){
	http.HandleFunc("GET /users/{id}", controllers.GetUserById)
	http.HandleFunc("GET /users", controllers.GetUsers)
	http.HandleFunc("POST /users", controllers.CreateUser)
	http.HandleFunc("PATCH /users/{id}", controllers.UpdateUser)
	http.HandleFunc("DELETE /users/{id}", controllers.DeleteUser)
}