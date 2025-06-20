package dto

type CafeResponseDto struct {
	ID        uint     `json:"id"`
	Nome      string   `json:"nome"`
	Tipo      string   `json:"tipo"`
	Preco     float64  `json:"preco"`
	Descricao string   `json:"descricao"`
	Tags      []string `json:"tags"`
}
