package handlers

import (
	"encoding/json"
	"myproject/models"
	"myproject/services"
	"net/http"
	"strconv"
)

func NotasHandler(w http.ResponseWriter, r *http.Request) {

	valorStr := r.URL.Query().Get("valor")

	valor, err := strconv.Atoi(valorStr)
	if err != nil || valor <= 0 {
		http.Error(w, "Valor inválido", http.StatusBadRequest)
		return
	}

	resultado, disponivel := services.CalcularNotas(valor)
	response := models.NotasResponse{}

	if disponivel {
		response.Notas = resultado
	} else {
		response.Mensagem = "Valor indisponível para saque."
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
