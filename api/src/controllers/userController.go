package controllers

import (
	"ALP/src/models"
	"ALP/src/services"
	"ALP/src/utils"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
)

func GetUserById(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Parâmetro de rota ID necessário", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserById(id)
	if err == nil {
		utils.SendResponse(w, http.StatusOK, user)
	} else if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		utils.SendResponse(w, http.StatusOK, users)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	var data models.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if data.Name == "" {
		http.Error(w, "Campo 'nome' é obrigatório", http.StatusBadRequest)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.CreateUser(data)
	if err == nil {
		utils.SendResponse(w, http.StatusCreated, map[string]any{
			"message": "Usuário criado com sucesso",
			"user": user,
		})
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Parâmetro de rota ID necessário", http.StatusBadRequest)
		return
	}

	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if data.Name == "" {
		http.Error(w, "Campo 'nome' é obrigatório", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.UpdateUser(id, data)
	if err == nil {
		utils.SendResponse(w, http.StatusOK, map[string]any{
			"message": "Usuário atualizado com sucesso",
			"user": user,
		})
	} else if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Parâmetro de rota ID necessário", http.StatusBadRequest)
		return
	}

	user, err := services.DeleteUser(id)
	if err == nil {
		utils.SendResponse(w, http.StatusOK, map[string]any{
			"message": "Usuário removido com sucesso",
			"user": user,
		})
	} else if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
