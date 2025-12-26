package models

type Booking struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	UserID  uint `json:"user_id"`
	EventID uint `json:"event_id"`
	Qty     int  `json:"qty"`
}

