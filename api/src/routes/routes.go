package routes

import (
	"ALP/src/controllers"
	"ALP/src/utils"
	"net/http"
)

func SetupRoutes(){
	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		utils.SendResponse(w, http.StatusOK, map[string]any{ "status": "ok" }) 
	})

	http.HandleFunc("GET /users/{id}", controllers.GetUserById)
	http.HandleFunc("GET /users", controllers.GetUsers)
	http.HandleFunc("POST /users", controllers.CreateUser)
	http.HandleFunc("PATCH /users/{id}", controllers.UpdateUser)
	http.HandleFunc("DELETE /users/{id}", controllers.DeleteUser)
}