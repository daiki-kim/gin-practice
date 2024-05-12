package main

import (
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"flea-market/infra"
	"flea-market/models"
)

func TestMain(m *testing.M) { // "Test~()"がテスト関数として認識される
	if err := godotenv.Load(".env.test"); err != nil { //  テスト用の`.env.test`を読み込み環境変数として保存
		log.Fatalln("Error loading .env.test file")
	}
	code := m.Run() // ファイル内のテスト関数が全て呼び出される

	os.Exit(code) // テストの終了
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "test item 1", Price: 1000, Description: "", SoldOut: false, UserID: 1},
		{Name: "test item 2", Price: 2000, Description: "test2", SoldOut: true, UserID: 1},
		{Name: "test item 3", Price: 3000, Description: "test3", SoldOut: false, UserID: 2},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
	}

	for _, user := range users {
		db.Create(&user)
	}
	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}
