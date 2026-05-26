package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func prompt(label string) string {
	fmt.Print(label)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func printMenu() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║       ALP CLI  —  Usuários       ║")
	fmt.Println("╠══════════════════════════════════╣")
	fmt.Println("║  1. Listar todos os usuários     ║")
	fmt.Println("║  2. Buscar usuário por ID        ║")
	fmt.Println("║  3. Criar usuário                ║")
	fmt.Println("║  4. Atualizar usuário            ║")
	fmt.Println("║  5. Deletar usuário              ║")
	fmt.Println("║  0. Sair                         ║")
	fmt.Println("╚══════════════════════════════════╝")
}

func runMenu() {
	for {
		printMenu()
		op := prompt("Escolha uma opção: ")

		fmt.Println()

		switch op {
		case "1":
			cmdList()
		case "2":
			cmdGet()
		case "3":
			cmdCreate()
		case "4":
			cmdUpdate()
		case "5":
			cmdDelete()
		case "0":
			fmt.Println("Saindo...")
			return
		default:
			fmt.Fprintln(os.Stderr, "Opção inválida. Tente novamente.")
		}

		fmt.Println()
		prompt("Pressione Enter para continuar...")
	}
}

func cmdList() {
	data, status, err := doRequest("GET", "/users", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao conectar:", err)
		return
	}
	if status != http.StatusOK {
		printError(status, data)
		return
	}
	trimmed := strings.TrimSpace(string(data))
	if trimmed == "null" || trimmed == "[]" {
		fmt.Println("Não há usuários para exibir.")
		return
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao interpretar resposta:", err)
		return
	}
	fmt.Printf("%-5s  %s\n", "ID", "Nome")
	fmt.Println(strings.Repeat("─", 30))
	for _, u := range users {
		fmt.Printf("%-5d  %s\n", u.ID, u.Name)
	}
}

func cmdGet() {
	idStr := prompt("ID do usuário: ")
	if _, err := strconv.Atoi(idStr); err != nil {
		fmt.Fprintln(os.Stderr, "Erro: ID deve ser um número.")
		return
	}

	data, status, err := doRequest("GET", "/users/"+idStr, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao conectar:", err)
		return
	}
	if status == http.StatusNotFound {
		fmt.Fprintln(os.Stderr, "Erro: usuário não encontrado.")
		return
	}
	if status != http.StatusOK {
		printError(status, data)
		return
	}

	var u User
	if err := json.Unmarshal(data, &u); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao interpretar resposta:", err)
		return
	}
	printUser(u)
}

func cmdCreate() {
	name := prompt("Nome do usuário: ")
	if name == "" {
		fmt.Fprintln(os.Stderr, "Erro: o nome não pode estar vazio.")
		return
	}

	data, status, err := doRequest("POST", "/users", map[string]string{"name": name})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao conectar:", err)
		return
	}
	if status != http.StatusCreated {
		printError(status, data)
		return
	}

	var resp UserResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao interpretar resposta:", err)
		return
	}
	fmt.Println(resp.Message)
	printUser(resp.User)
}

func cmdUpdate() {
	idStr := prompt("ID do usuário: ")
	if _, err := strconv.Atoi(idStr); err != nil {
		fmt.Fprintln(os.Stderr, "Erro: ID deve ser um número.")
		return
	}

	name := prompt("Novo nome: ")
	if name == "" {
		fmt.Fprintln(os.Stderr, "Erro: o nome não pode estar vazio.")
		return
	}

	data, status, err := doRequest("PATCH", "/users/"+idStr, map[string]string{"name": name})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao conectar:", err)
		return
	}
	if status == http.StatusNotFound {
		fmt.Fprintln(os.Stderr, "Erro: usuário não encontrado.")
		return
	}
	if status != http.StatusOK {
		printError(status, data)
		return
	}

	var resp UserResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao interpretar resposta:", err)
		return
	}
	fmt.Println(resp.Message)
	printUser(resp.User)
}

func cmdDelete() {
	idStr := prompt("ID do usuário: ")
	if _, err := strconv.Atoi(idStr); err != nil {
		fmt.Fprintln(os.Stderr, "Erro: ID deve ser um número.")
		return
	}

	data, status, err := doRequest("DELETE", "/users/"+idStr, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao conectar:", err)
		return
	}
	if status == http.StatusNotFound {
		fmt.Fprintln(os.Stderr, "Erro: usuário não encontrado.")
		return
	}
	if status != http.StatusOK {
		printError(status, data)
		return
	}

	var resp UserResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao interpretar resposta:", err)
		return
	}
	fmt.Println(resp.Message)
	printUser(resp.User)
}
