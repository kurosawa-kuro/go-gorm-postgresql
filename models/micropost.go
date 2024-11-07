package models

import "gorm.io/gorm"

type Micropost struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt を含む
	Title      string `gorm:"not null"`
}
