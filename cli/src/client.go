package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const defaultHost = "http://localhost:5080"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Used by create, update and delete responses: { "message": "...", "user": {...} }
type UserResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

func baseURL() string {
	if h := os.Getenv("LB_HOST"); h != "" {
		return strings.TrimRight(h, "/")
	}
	return defaultHost
}

func doRequest(method, path string, body any) ([]byte, int, error) {
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, 0, fmt.Errorf("falha ao serializar o corpo da requisição: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, baseURL()+path, bodyReader)
	if err != nil {
		return nil, 0, fmt.Errorf("falha ao construir a requisição: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("falha ao executar a requisição: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("falha ao ler a resposta: %w", err)
	}
	return data, resp.StatusCode, nil
}

func printUser(u User) {
	fmt.Printf("  ID:   %d\n", u.ID)
	fmt.Printf("  Nome: %s\n", u.Name)
}

func printError(status int, body []byte) {
	fmt.Fprintf(os.Stderr, "Erro %d: %s\n", status, strings.TrimSpace(string(body)))
}
