package models

type Discount struct {
	ID            int
	Code          string `gorm:"unique"`
	QuantityLeft  int    `gorm:"check:quantity_left > 0"`
	TotalQuantity int    `gorm:"check:total_quantity > 0 "`
	Amount        float64
}
