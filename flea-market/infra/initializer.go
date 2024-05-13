package infra

import (
	"log"

	"github.com/joho/godotenv"
)

/*
本番環境で使う環境変数の読み込み
*/
func Initialize() {
	err := godotenv.Load() // .envファイルを読み込み'='の前後をkey,valueとして環境変数にする
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
