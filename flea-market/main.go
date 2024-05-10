package main

import (
	"github.com/gin-gonic/gin"

	"flea-market/controllers"
	"flea-market/infra"
	"flea-market/models"
	"flea-market/repositories"
	"flea-market/services"
)

func main() {
	infra.Initialize()
	items := []models.Item{
		{ID: 1, Name: "item1", Price: 100, Description: "first item", SoldOut: false},
		{ID: 2, Name: "item2", Price: 200, Description: "second item", SoldOut: true},
		{ID: 3, Name: "item3", Price: 300, Description: "third item", SoldOut: false},
	}

	IItemMemoryRepository := repositories.NewItemMemoryRepository(items)
	IItemService := services.NewItemService(IItemMemoryRepository)
	IItemController := controllers.NewItemController(IItemService)

	router := gin.Default()

	router.GET("/items", IItemController.FindAll)
	router.GET("/items/:id", IItemController.FindById)
	router.POST("/items", IItemController.Create)
	router.PUT("/items/:id", IItemController.Update)
	router.DELETE("/items/:id", IItemController.Delete)

	router.Run("localhost:8080")
}
