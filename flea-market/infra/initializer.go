package infra

import (
	"log"

	"github.com/joho/godotenv"
)

func Initialize() {
	err := godotenv.Load() // .envファイルを読み込み'='の前後をkey,valueとして環境変数にする
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
