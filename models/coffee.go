package models

type Coffee struct {
	ID        uint     `gorm:"primaryKey" json:"id"`
	Nome      string   `json:"nome"`
	Tipo      string   `json:"tipo"`
	Preco     float64  `json:"preco"`
	Descricao string   `json:"descricao"`
	Tags      []Tag    `gorm:"many2many:coffee_tags;" json:"tags"`
	Pedidos   []Pedido `json:"pedidos,omitempty"`
}
