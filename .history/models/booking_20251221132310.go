package models

type Booking struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	EventID uint
	Qty    int
}
