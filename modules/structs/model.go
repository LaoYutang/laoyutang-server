package structs

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id        int            `json:"id" gorm:"type:bigint"`
	CreatedAt time.Time      `json:"createdAt" gorm:"type:timestamp"`
	CreatedBy string         `json:"createdBy" gorm:"type:varchar(100)"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:timestamp"`
	UpdatedBy string         `json:"updatedBy" gorm:"type:varchar(100)"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"type:timestamp"`
	DeletedBy string         `json:"deletedBy" gorm:"type:varchar(100)"`
}
