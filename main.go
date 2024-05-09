package main

import (
	"github.com/gin-gonic/gin"

	"gin-practice-api/controllers"
	"gin-practice-api/models"
	"gin-practice-api/repositories"
	"gin-practice-api/services"
)

func main() {
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
