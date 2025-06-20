package models

type Tag struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Nome string `json:"nome"`
}
