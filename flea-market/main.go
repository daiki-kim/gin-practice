package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"flea-market/infra"

	"flea-market/controllers"
	"flea-market/middlewares"

	// "flea-market/models"
	"flea-market/repositories"
	"flea-market/services"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	// IItemMemoryRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	r.Use(cors.Default())

	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")

	itemRouter.GET("", itemController.FindAll)
	itemRouterWithAuth.GET("/:id", itemController.FindById)
	itemRouterWithAuth.POST("", itemController.Create)
	itemRouterWithAuth.PUT("/:id", itemController.Update)
	itemRouterWithAuth.DELETE("/:id", itemController.Delete)

	authRouter.POST("/sinup", authController.SignUp)
	authRouter.POST("/login", authController.LogIn)

	return r
}

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	// items := []models.Item{
	// 	{ID: 1, Name: "item1", Price: 100, Description: "first item", SoldOut: false},
	// 	{ID: 2, Name: "item2", Price: 200, Description: "second item", SoldOut: true},
	// 	{ID: 3, Name: "item3", Price: 300, Description: "third item", SoldOut: false},
	// }

	r := setupRouter(db)
	r.Run("localhost:8080")
}
