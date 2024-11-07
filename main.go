package main

import (
	"fmt"
	"go-gorm-postgresql/config"
	"go-gorm-postgresql/models"
	"go-gorm-postgresql/repositories"
)

func main() {
	// データベース初期化
	db := config.InitDB()

	// リポジトリの作成
	micropostRepo := repositories.NewMicropostRepository(db)

	// マイグレーション
	if err := micropostRepo.Migrate(); err != nil {
		panic("マイグレーションに失敗しました")
	}

	// すべての投稿を削除
	if err := micropostRepo.DeleteAll(); err != nil {
		panic("投稿の削除に失敗しました")
	}

	// 作成
	micropost := &models.Micropost{
		Title: "最初の投稿",
	}
	if err := micropostRepo.Create(micropost); err != nil {
		panic("投稿の作成に失敗しました")
	}

	// 読み取り
	post, err := micropostRepo.FindFirst()
	if err != nil {
		panic("投稿の取得に失敗しました")
	}
	fmt.Printf("取得した投稿: %v\n", post)

	// 更新
	if err := micropostRepo.Update(&post, "更新された投稿"); err != nil {
		panic("投稿の更新に失敗しました")
	}

	// 複数の投稿を取得
	posts, err := micropostRepo.FindAll()
	if err != nil {
		panic("投稿の一覧取得に失敗しました")
	}
	for _, p := range posts {
		fmt.Printf("投稿: %v\n", p)
	}

	// 削除
	if err := micropostRepo.Delete(&post); err != nil {
		panic("投稿の削除に失敗しました")
	}
}
