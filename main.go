package main

import (
	"fmt"
	"myproject/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/notas", handlers.NotasHandler)

	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}
