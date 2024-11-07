package seeders

import (
	"go-gorm-postgresql/models"
	"go-gorm-postgresql/repositories"
)

func SeedMicroposts(repo *repositories.MicropostRepository) error {
	microposts := []models.Micropost{
		{Title: "最初の投稿"},
		{Title: "2番目の投稿"},
		{Title: "3番目の投稿"},
	}

	for _, micropost := range microposts {
		if err := repo.Create(&micropost); err != nil {
			return err
		}
	}
	return nil
}
