package model

import (
	"time"

	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/util"
	"gorm.io/gorm"
)

type Cart struct {
	ID          string         `json:"id" gorm:"primaryKey"`
	ProductCode string         `json:"product_code"`
	ProductName string         `json:"product_name"`
	Quantity    int64          `json:"quantity"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (u *Cart) BeforeCreate(tx *gorm.DB) error {
	u.ID = util.GenerateUUID()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}
