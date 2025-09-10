package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Login    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
