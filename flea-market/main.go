package main

import (
	"github.com/gin-gonic/gin"

	"flea-market/controllers"
	"flea-market/infra"

	// "flea-market/models"
	"flea-market/repositories"
	"flea-market/services"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	// items := []models.Item{
	// 	{ID: 1, Name: "item1", Price: 100, Description: "first item", SoldOut: false},
	// 	{ID: 2, Name: "item2", Price: 200, Description: "second item", SoldOut: true},
	// 	{ID: 3, Name: "item3", Price: 300, Description: "third item", SoldOut: false},
	// }

	// IItemMemoryRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	router := r.Group("/items")
	router.GET("", itemController.FindAll)
	router.GET("/:id", itemController.FindById)
	router.POST("", itemController.Create)
	router.PUT("/:id", itemController.Update)
	router.DELETE("/:id", itemController.Delete)

	r.Run("localhost:8080")
}