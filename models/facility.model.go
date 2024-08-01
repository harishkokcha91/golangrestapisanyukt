package models

type Facility struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`      // Name of the dharmshala/marriage home/PG/community center
	Caretaker string `json:"caretaker"` // Name of the caretaker
	Contact   string `json:"contact"`   // Contact number
	Address   string `json:"address"`   // Address
}

// func (Facility) TableName() string {
// 	return "facilities" // Ensure this matches your database table name
// }
