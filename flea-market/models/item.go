package models

// import "gorm.io/gorm"

// for ItemMemoryRepository
type Item struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	SoldOut     bool   `json:"soldOut"`
}

// // for ItemRepository
// type Item struct {
// 	gorm.Model
// 	Name        string `gorm:"not null"`
// 	Price       uint   `gorm:"not null"`
// 	Description string
// 	SoldOut     bool `gorm:"not null; default:false"`
// }
