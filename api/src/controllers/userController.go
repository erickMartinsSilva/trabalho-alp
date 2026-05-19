package controllers

import (
	"ALP/src/models"
	"ALP/src/services"
	"encoding/json"
	"net/http"
)

func GetUserById(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "User ID parameter missing", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(user)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(users)
	}
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.UpdateUser(user)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User updated successfully",
			"user": user,
		})
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "User ID parameter missing", http.StatusBadRequest)
		return
	}

	err := services.DeleteUser(id)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User deleted successfully",
			"id": id,
		})
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
