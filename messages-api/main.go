package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct usada para enviar/receber mensagens em JSON
type Mensagem struct {
	Texto string `json:"mensagem"`
}

// Handler para GET /mensagem
func getMensagemHandler(w http.ResponseWriter, r *http.Request) {
	msg := Mensagem{Texto: "Olá, Ana! Agora sim, sua API JSON está funcionando 😄"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

// Handler para POST /mensagem
func postMensagemHandler(w http.ResponseWriter, r *http.Request) {
	var msg Mensagem

	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	fmt.Println("Mensagem recebida:", msg.Texto)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":   "sucesso",
		"mensagem": msg.Texto,
	})
}

// Função principal: define a rota e inicia o servidor
func main() {
	http.HandleFunc("/mensagem", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getMensagemHandler(w, r)
		} else if r.Method == http.MethodPost {
			postMensagemHandler(w, r)
		} else {
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("🚀 Servidor rodando em http://localhost:8080/mensagem")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
