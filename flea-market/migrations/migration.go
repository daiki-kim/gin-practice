package main

import (
	"flea-market/infra"
	"flea-market/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("failed to migrate database")
	}
}
