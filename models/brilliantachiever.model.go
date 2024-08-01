package models

import (
	"time"
)

type BrilliantAchiever struct {
	ID            uint      `gorm:"primaryKey"`
	Name          string    `gorm:"type:varchar(255)" json:"name"`
	Location      string    `gorm:"type:varchar(255)" json:"location"`
	Achievement   string    `gorm:"type:text" json:"achievement"`
	Image         string    `gorm:"type:varchar(255)" json:"image"`
	CreatedBy     uint      `gorm:"type:integer" json:"created_by"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	Date          time.Time `json:"date"`
	CreatedByUser User      `gorm:"foreignKey:CreatedBy"`
}
