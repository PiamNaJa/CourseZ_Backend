package models

type Course struct {
	Course_id   int32    `json:"course_id" gorm:"primaryKey;type:int"`
	SubjectID   int      `json:"subject_id" gorm:"index;type:int"`
	Subject     *Subject `json:"subject" gorm:"not null;foreignKey:SubjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Course_name string   `json:"course_name" gorm:"not null;type:varchar(100)"`
	Class_level int8     `json:"class_level" gorm:"not null;type:int"`
	Picture     string   `json:"picture" gorm:"not null;type:varchar(255)"`
	Description string   `json:"description" gorm:"not null;type:text"`
}
