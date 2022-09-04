package models

type Wallet struct {
	ID      int
	UserID  int
	User    User `gorm:"foreignkey:UserID" json:"-"`
	Balance float64
}
