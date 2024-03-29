package models

type VideoHistory struct {
	History_id int32  `json:"history_id" gorm:"primaryKey;type:int"`
	UserID     int32  `json:"user_id" gorm:"index;type:int;not null" validate:"required,number"`
	VideoID    int32  `json:"video_id" gorm:"index;type:int;not null" validate:"required,number"`
	Video      *Video `json:"video" gorm:"not null;foreignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Duration   int32  `json:"duration" gorm:"not null;type:int"`
	UpdateAt   int64  `json:"update_at" gorm:"autoUpdateTime"`
}

type CourseHistory struct {
	History_id int32   `json:"history_id" gorm:"primaryKey;type:int"`
	UserID     int32   `json:"user_id" gorm:"index;type:int;not null" validate:"required,number"`
	CourseID   int32   `json:"course_id" gorm:"index;type:int;not null" validate:"required,number"`
	Course     *Course `json:"course" gorm:"not null;foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UpdateAt   int64   `json:"update_at" gorm:"autoUpdateTime"`
	Frequency  int32   `json:"frequency" gorm:"not null;type:int"`
}
