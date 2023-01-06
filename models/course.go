package models

type Course struct {
	Course_id   int32    `json:"course_id" gorm:"primaryKey;type:int"`
	SubjectID   int32    `json:"subject_id" gorm:"index;type:int"`
	Subject     *Subject `json:"subject" gorm:"not null;foreignKey:SubjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Videos      *[]Video `json:"videos" gorm:"not null;foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeacherID   int32    `json:"teacher_id" gorm:"index;type:int"`
	Course_name string   `json:"course_name" gorm:"not null;type:varchar(100)"`
	Picture     string   `json:"picture" gorm:"not null;type:varchar(255)"`
	Description string   `json:"description" gorm:"not null;type:text"`
}
