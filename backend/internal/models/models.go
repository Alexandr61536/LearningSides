package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Login    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null;default:user"`
	Deleted  bool   `gorm:"not null;default:false"`
}
