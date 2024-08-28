package models

type NotasResponse struct {
	Notas    map[int]int `json:"notas,omitempty"`
	Mensagem string      `json:"mensagem,omitempty"`
}
