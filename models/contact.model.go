// models/contact.go
package models

type Contact struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	Organization string `json:"organization"`
	Role         string `json:"role"`
	Notes        string `json:"notes"`
	CreatedBy    uint   `json:"created_by"` // ID of the user who created the contact
}
