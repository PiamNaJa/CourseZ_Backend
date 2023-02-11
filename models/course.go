package models

import "time"

type Course struct {
	Course_id   int32     `json:"course_id" gorm:"primaryKey;type:int"`                                                      // Course_id is the id of the course
	Course_name string    `json:"course_name" gorm:"not null;type:varchar(1000)" validate:"required,max=1000"`               // Course_name is the name of the course
	SubjectID   int32     `json:"subject_id" gorm:"index;type:int;not null" validate:"required,number"`                      // SubjectID is the id of the subject
	Subject     *Subject  `json:"subject" gorm:"not null;foreignKey:SubjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Subject is the subject of the course
	Videos      *[]Video  `json:"videos" gorm:"not null;foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`  // Videos is the videos of the course
	TeacherID   int32     `json:"teacher_id" gorm:"index;type:int;not null" validate:"required,number"`                      // TeacherID is the id of the teacher
	Picture     string    `json:"picture" gorm:"not null;type:varchar(5000)" validate:"required,max=5000"`                   // Picture is a link to the picture
	Description string    `json:"description" gorm:"not null;type:text" validate:"required"`                                 // Description is the description of the course
	CreatedAt   time.Time `json:"created_at" gorm:"not null;"`
	Like        int32     `json:"like" gorm:"not null;type:int;default:0"`
}
