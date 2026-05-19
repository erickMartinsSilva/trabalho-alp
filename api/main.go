package main

import (
	"ALP/src/routes"
	"net/http"
	"ALP/src/db"
)

func main() {
	db.Init()
	defer db.Instance.Close()
	routes.SetupRoutes()
	http.ListenAndServe(":8080", nil)
}