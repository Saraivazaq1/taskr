package models

// Definição de tarefa
type Tarefa struct {
	Id        int    `json:"Id"`
	Descricao string `json:"Descricao"`
	Feita     bool   `json:"Feita"`
}
