package models

type Event struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string
	Date     string
	Location string
	Ticket   int
}
