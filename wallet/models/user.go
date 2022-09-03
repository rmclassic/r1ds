package models

type User struct {
	ID          int
	PhoneNumber string `gorm:"unique"`
}
