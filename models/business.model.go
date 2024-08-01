package models

import (
	"time"

	"github.com/google/uuid"
)

type Business struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key" json:"id,omitempty"`
	Name         string    `gorm:"not null" json:"name,omitempty"`
	Description  string    `gorm:"type:text" json:"description,omitempty"`
	CategoryID   uuid.UUID `gorm:"type:uuid" json:"category_id,omitempty"`
	OwnerName    string    `gorm:"not null" json:"owner_name,omitempty"`
	PhoneNumber  string    `json:"phone_number,omitempty"`
	Email        string    `gorm:"unique" json:"email,omitempty"`
	Website      string    `json:"website,omitempty"`
	AddressLine1 string    `gorm:"not null" json:"address_line_1,omitempty"`
	AddressLine2 string    `json:"address_line_2,omitempty"`
	City         string    `gorm:"not null" json:"city,omitempty"`
	State        string    `gorm:"not null" json:"state,omitempty"`
	PostalCode   string    `gorm:"not null" json:"postal_code,omitempty"`
	Country      string    `gorm:"not null" json:"country,omitempty"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

type Image struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id,omitempty"`
	BusinessID  uuid.UUID `gorm:"type:uuid" json:"business_id,omitempty"`
	URL         string    `gorm:"not null" json:"url,omitempty"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at,omitempty"`
}

type BusinessReview struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key" json:"id,omitempty"`
	BusinessID   uuid.UUID `gorm:"type:uuid" json:"business_id,omitempty"`
	ReviewerName string    `gorm:"not null" json:"reviewer_name,omitempty"`
	Rating       int       `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating,omitempty"`
	Comment      string    `gorm:"type:text" json:"comment,omitempty"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at,omitempty"`
}

type BusinessHour struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key" json:"id,omitempty"`
	BusinessID uuid.UUID `gorm:"type:uuid" json:"business_id,omitempty"`
	DayOfWeek  string    `gorm:"not null" json:"day_of_week,omitempty"`
	OpenTime   time.Time `gorm:"not null" json:"open_time,omitempty"`
	CloseTime  time.Time `gorm:"not null" json:"close_time,omitempty"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type Category struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id,omitempty"`
	Name        string    `gorm:"unique;not null" json:"name,omitempty"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type CreateBusinessRequest struct {
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description,omitempty"`
	CategoryID   uuid.UUID `json:"category_id" binding:"required"`
	OwnerName    string    `json:"owner_name" binding:"required"`
	PhoneNumber  string    `json:"phone_number,omitempty"`
	Email        string    `json:"email,omitempty"`
	Website      string    `json:"website,omitempty"`
	AddressLine1 string    `json:"address_line_1" binding:"required"`
	AddressLine2 string    `json:"address_line_2,omitempty"`
	City         string    `json:"city" binding:"required"`
	State        string    `json:"state" binding:"required"`
	PostalCode   string    `json:"postal_code" binding:"required"`
	Country      string    `json:"country" binding:"required"`
}

type UpdateBusinessRequest struct {
	Name         *string    `json:"name,omitempty"`
	Description  *string    `json:"description,omitempty"`
	CategoryID   *uuid.UUID `json:"category_id,omitempty"`
	OwnerName    *string    `json:"owner_name,omitempty"`
	PhoneNumber  *string    `json:"phone_number,omitempty"`
	Email        *string    `json:"email,omitempty"`
	Website      *string    `json:"website,omitempty"`
	AddressLine1 *string    `json:"address_line_1,omitempty"`
	AddressLine2 *string    `json:"address_line_2,omitempty"`
	City         *string    `json:"city,omitempty"`
	State        *string    `json:"state,omitempty"`
	PostalCode   *string    `json:"postal_code,omitempty"`
	Country      *string    `json:"country,omitempty"`
}
