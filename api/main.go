package main

import (
	"ALP/src/routes"
	"net/http"
)

func main() {
	routes.SetupRoutes()
	http.ListenAndServe(":8080", nil)
}