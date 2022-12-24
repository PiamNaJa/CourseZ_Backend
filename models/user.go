package models

import (
	"database/sql"
	"time"
)

type User struct {
	User_id   int32         `json:"user_id" gorm:"primaryKey;type:int"`
	Email     string        `json:"email" gorm:"not null;type:varchar(100)"`
	Password  string        `json:"password" gorm:"not null;type:varchar(255)"`
	Fullname  string        `json:"fullname" gorm:"not null;type:varchar(100)"`
	Nickname  string        `json:"nickname" gorm:"not null;type:varchar(100)"`
	Phone     string        `json:"phone" gorm:"not null;type:varchar(10)"`
	Birthday  time.Time     `json:"birthday"  gorm:"not null;type:date"`
	Role      string        `json:"role" gorm:"not null;type:varchar(20)"`
	Picture   string        `json:"picture" gorm:"not null;type:varchar(255)"`
	Point     int32         `json:"point" gorm:"type:money;default:0"`
	Money     int32         `json:"money" gorm:"type:money;default:0"`
	TeacherID sql.NullInt32 `json:"teacher_id" gorm:"index;type:int"`
	Teacher   *UserTeacher  `json:"teacher" gorm:"foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserTeacher struct {
	Teacher_id         int32        `json:"teacher_id" gorm:"primaryKey;type:int"`
	Teacher_license    string       `json:"teacher_license" gorm:"not null;type:varchar(255)"`
	Transcript         string       `json:"transcript" gorm:"not null;type:varchar(255)"`
	Id_card            string       `json:"id_card" gorm:"not null;type:varchar(255)"`
	Psychological_test string       `json:"psychological_test" gorm:"not null;type:varchar(255)"`
	ExperienceID       int32        `json:"experience_id" gorm:"index;type:int"`
	Experience         []Experience `json:"experience" gorm:"foreignKey:Experience_id;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Experience struct {
	Experience_id int32  `json:"experience_id" gorm:"primaryKey;type:int"`
	Title         string `json:"title" gorm:"not null;type:varchar(100)"`
	Evidence      string `json:"evidence" gorm:"not null;type:varchar(255)"`
}
