package models

import (
	"time"
)

type User struct {
	User_id  int64      `json:"user_id" gorm:"primaryKey"`
	Email    string     `json:"email" gorm:"not null"`
	Password string     `json:"password" gorm:"not null"`
	Fullname string     `json:"fullname" gorm:"not null"`
	Nickname string     `json:"nickname" gorm:"not null"`
	Phone    string     `json:"phone" gorm:"not null"`
	Birthday *time.Time `json:"birthday"  gorm:"not null"`
	Role     string     `json:"role" gorm:"not null"`
	Picture  string     `json:"picture" gorm:"not null"`
	Point    string     `json:"point" gorm:"default:0"`
	Money    string     `json:"money" gorm:"default:0"`
}
