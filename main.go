package main

import (
	"fmt"
	"go-gorm-postgresql/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 環境変数の読み込み
	if err := godotenv.Load(".env.development"); err != nil {
		log.Fatal("Error loading .env.development file")
	}

	// データベース接続
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}

	// マイグレーション
	db.AutoMigrate(&models.Micropost{})

	// 作成
	micropost := models.Micropost{
		Title: "最初の投稿",
	}
	db.Create(&micropost)

	// 読み取り
	var post models.Micropost
	db.First(&post) // ID=1の投稿を取得
	fmt.Printf("取得した投稿: %v\n", post)

	// 更新
	db.Model(&post).Update("Title", "更新された投稿")

	// 複数の投稿を取得
	var posts []models.Micropost
	db.Find(&posts)
	for _, p := range posts {
		fmt.Printf("投稿: %v\n", p)
	}

	// 削除
	db.Delete(&post)
}
