package main

import (
	"fmt"
	"go-gorm-postgresql/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// データベース接続
	dsn := "postgresql://postgres:postgres@localhost:5432/web_app_db_integration_go?sslmode=disable"
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
