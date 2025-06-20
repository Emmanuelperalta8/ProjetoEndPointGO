package dto

type CreateCafeDto struct {
	Nome      string   `json:"nome" binding:"required"`
	Tipo      string   `json:"tipo" binding:"required"`
	Preco     float64  `json:"preco" binding:"required"`
	Descricao string   `json:"descricao" binding:"required"`
	Tags      []string `json:"tags"`
}
