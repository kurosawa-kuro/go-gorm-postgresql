package repositories

import (
	"go-gorm-postgresql/models"

	"gorm.io/gorm"
)

type MicropostRepository struct {
	db *gorm.DB
}

func NewMicropostRepository(db *gorm.DB) *MicropostRepository {
	return &MicropostRepository{db: db}
}

func (r *MicropostRepository) Migrate() error {
	return r.db.AutoMigrate(&models.Micropost{})
}

func (r *MicropostRepository) Create(micropost *models.Micropost) error {
	return r.db.Create(micropost).Error
}

func (r *MicropostRepository) FindFirst() (models.Micropost, error) {
	var post models.Micropost
	err := r.db.First(&post).Error
	return post, err
}

func (r *MicropostRepository) Update(post *models.Micropost, title string) error {
	return r.db.Model(post).Update("Title", title).Error
}

func (r *MicropostRepository) FindAll() ([]models.Micropost, error) {
	var posts []models.Micropost
	err := r.db.Find(&posts).Error
	return posts, err
}

func (r *MicropostRepository) Delete(post *models.Micropost) error {
	return r.db.Delete(post).Error
}

func (r *MicropostRepository) DeleteAll() error {
	return r.db.Unscoped().Where("1 = 1").Delete(&models.Micropost{}).Error
}
