package models

type Subject struct {
	Subject_id      int32  `json:"subject_id" gorm:"primaryKey;type:int"` // Subject_id is the id of the subject
	Subject_title   string `json:"subject_title" gorm:"not null;type:varchar(255)" validate:"required,max=255"` // Subject_title is the title of the subject
	Class_level     int8   `json:"class_level" gorm:"not null;type:int" validate:"required,number"` // Class_level is the level of the class
	Subject_picture string `json:"subject_picture" gorm:"not null;type:varchar(5000)" validate:"required,max=5000"` // Subject_picture is a link to the picture
}
