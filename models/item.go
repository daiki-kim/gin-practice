package models

type Item struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	SoldOut     bool   `json:"soldOut"`
}
