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

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()

	itemRouter := r.Group("/items")
	itemRouter.GET("", itemController.FindAll)
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)

	authRouter := r.Group("/auth")
	authRouter.POST("/sinup", authController.SignUp)
	authRouter.POST("/login", authController.LogIn)

	r.Run("localhost:8080")
}
