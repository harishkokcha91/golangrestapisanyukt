package models

import (
	"time"
)

type Profile struct {
	ID          int    `gorm:"primary_key;AUTO_INCREMENT;not null" json:"id,omitempty"`
	UserId      string `gorm:"not null" json:"user_id,omitempty"`
	ProfileFor  string `gorm:"not null" json:"profile_for,omitempty"`
	Gender      string `grom:"not null" json:"gender,omitempty"`
	UserName    string `grom:"not null" json:"user_name,omitempty"`
	DateOfBirth string `json:"date_of_birth,omitempty"`

	//personal profile
	Gotra         string `json:"gotra,omitempty"`
	MotherGotra   string `json:"mother_gotra,omitempty"`
	MotherTongue  string `json:"mother_tongue,omitempty"`
	NativePlace   string `json:"native_place,omitempty"`
	Height        string `json:"height,omitempty"`
	Physique      string `json:"physique,omitempty"`   //{Thin,Slim,Smart,healthy,fatty,strong}
	Complexion    string `json:"complexion,omitempty"` //{fair,white,dark}
	Weight        string `json:"weight,omitempty"`
	BloodGroup    string `json:"blood_group,omitempty"`
	Hobbies       string `json:"hobbies,omitempty"`
	FamilyStatus  string `json:"family_status,omitempty"` //{Genral,middle standard,higher standard}
	PlaceOfBirth  string `json:"place_of_birth,omitempty"`
	TimeOfBirth   string `json:"time_of_birth,omitempty"`
	HaveHoroscope string `json:"have_horoscope,omitempty"` //{yes,no,unknown}
	Manglik       string `json:"manglik,omitempty"`        //{yes,no,unknown}

	//sibling details
	UnmarriedBrothers string `json:"unmarried_brothers,omitempty"`
	MarriedBrothers   string `json:"married_brothers,omitempty"`
	MarriedSisters    string `json:"married_sisters,omitempty"`
	UnmarriedSisters  string `json:"unmarried_sisters,omitempty"`

	//Marital Status / Children

	MaritalStatus string `json:"marital_status,omitempty"`
	NoOfChildren  string `json:"no_of_children,omitempty"`

	//Educational Details
	Academic string `json:"academic,omitempty"`

	//Professional Details
	Profession        string `json:"profession,omitempty"`
	Salary            string `json:"salary,omitempty"`
	Company           string `json:"company,omitempty"`
	OccupationAddress string `json:"occupation_address,omitempty"`

	//Personal Habits
	Smoke           string `json:"smoke,omitempty"`
	Drink           string `json:"drink,omitempty"`
	Diet            string `json:"diet,omitempty"`
	OtherAmusements string `json:"other_amusements,omitempty"`

	//Contact Information
	Address        string `json:"Address,omitempty"`
	City           string `json:"city,omitempty"`
	State          string `json:"state,omitempty"`
	Country        string `json:"country,omitempty"`
	PinCode        string `json:"pin_code,omitempty"`
	EmailAddress   string `json:"email_address,omitempty"`
	ContactNumeber string `json:"contact_number,omitempty"`
	FatherName     string `json:"father_name,omitempty"`
	//Other Details

	AboutSelf         string `json:"about_self,omitempty"`
	AboutFamily       string `json:"about_family,omitempty"`
	AboutPreference   string `json:"about_preference,omitempty"` //life partner
	Photos            string `json:"photos,omitempty"`
	ShowPicturePublic string `json:"show_picture_publiclly,omitempty"` //{yes,no}

	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}
