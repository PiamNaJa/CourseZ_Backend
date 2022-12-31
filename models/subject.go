package models

type Subject struct {
	Subject_id      int32  `json:"subject_id" gorm:"primaryKey;type:int"`
	Subject_title   string `json:"subject_title" gorm:"not null;type:varchar(255)"`
	Class_level     int8   `json:"class_level" gorm:"not null;type:int"`
	Subject_picture string `json:"subject_picture" gorm:"not null;type:varchar(255)"`
}
