package controllers

import (
"net/http"
"encoding/json"
"ALP/src/models"
"ALP/src/services"
)

func GetUser(w http.ResponseWriter, r *http.Request){
	// TODO: implementar
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.CreateUser(user)

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User created successfully",
			"user": user,
		})
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
