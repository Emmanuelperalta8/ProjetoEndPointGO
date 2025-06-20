package models

type Pedido struct {
	ID       uint `gorm:"primaryKey"`
	CoffeeID uint
	Detalhes string
}
