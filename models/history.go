package models

type History struct {
	History_id int32  `json:"history_id" gorm:"primaryKey;type:int"`
	UserID     int    `json:"user_id" gorm:"index;type:int"`
	User       *User  `json:"user" gorm:"not null;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	VideoID    int    `json:"video_id" gorm:"index;type:int"`
	Video      *Video `json:"video" gorm:"not null;foreignKey:video_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Duration   int32  `json:"duration" gorm:"not null;type:int"`
}
