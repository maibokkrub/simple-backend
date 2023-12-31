package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID           int           `gorm:"primaryKey;autoIncrement" json:"-"`
	DisplayName  string        `gorm:"not null" json:"displayName"`
	Email        string        `json:"email" validate:"required,email"`
	AvatarURL    string        `json:"avatar"`
	Appointments []Appointment `gorm:"foreignKey:CreatedBy" json:"-"`
	Comments     []Comment     `gorm:"foreignKey:UserID" json:"-"`
	// Password     string        `gorm:"not null" json:"-"`
}

func GetAllUsers(db *gorm.DB, page int) (*[]User, error) {
	users := []User{}

	res := db.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return &users, nil
}

//
// DB OPS
//
func (user *User) Create(db *gorm.DB) error {
	if tx := db.Create(user); tx.Error != nil {
		return tx.Error
	}

	return nil
}
