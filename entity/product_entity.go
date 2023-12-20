package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `gorm:"primarykey"`
	Title       string    `gorm:"unique"`
	Description string
	Rating      float64
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (m *Product) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return nil
}
