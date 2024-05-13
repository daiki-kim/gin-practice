package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"flea-market/dto"
	"flea-market/infra"
	"flea-market/models"
	"flea-market/services"
)

func TestMain(m *testing.M) { // "Test~()"がテスト関数として認識される
	if err := godotenv.Load(".env.test"); err != nil { //  テスト用の`.env.test`を読み込み`ENV`を環境変数として保存
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

/*
test without authorization
*/
func TestFindAll(t *testing.T) {
	router := setup()

	// serve http request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)
	router.ServeHTTP(w, req)

	// get http response
	var res map[string][]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
}

/*
tests with authorization
*/
func TestFindById(t *testing.T) {
	router := setup()

	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items/1", nil)
	req.Header.Set("Authorization", "Bearer "+*token)
	router.ServeHTTP(w, req)

	var res map[string]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "test item 1", res["data"].Name)
}

func TestCreate(t *testing.T) {
	router := setup()

	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	itemInput := dto.CreateItemInput{
		Name:        "test item 4",
		Price:       4000,
		Description: "testing Create()",
	}
	reqestBody, _ := json.Marshal(itemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqestBody))
	req.Header.Set("Authorization", "Bearer "+*token)
	router.ServeHTTP(w, req)

	var res map[string]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, uint(4), res["data"].ID)
}

func TestUpdate(t *testing.T) {
	router := setup()

	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	updateName := "update item 1"

	updateItemInput := dto.UpdateItemInput{
		Name: &updateName,
	}
	requestBody, _ := json.Marshal(updateItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+*token)
	router.ServeHTTP(w, req)

	var res map[string]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, updateName, res["data"].Name)
}

func TestDelete(t *testing.T) {
	router := setup()

	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/items/1", nil)
	req.Header.Set("Authorization", "Bearer "+*token)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSignUp(t *testing.T) {
	router := setup()

	newTestUserInput := dto.SignUpUserInput{
		Email:    "test3@example.com",
		Password: "test3password",
	}

	requestBody, _ := json.Marshal(newTestUserInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

/*
error tests
*/
func TestCreateUnauthorized(t *testing.T) {
	router := setup()

	itemInput := dto.CreateItemInput{
		Name:        "test item 4",
		Price:       4000,
		Description: "testing Create()",
	}
	reqestBody, _ := json.Marshal(itemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqestBody))

	router.ServeHTTP(w, req)

	var res map[string]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
