package models

type Review_Course struct {
	Review_course_id      int32   `json:"review_course_id" gorm:"primaryKey;type:int"`
	CourseID              int32   `json:"course_id" gorm:"index;type:int"`
	Course                *Course `json:"course" gorm:"not null;foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rating                int32   `json:"rating" gorm:"not null;type:int"`
	Review_course_comment string  `json:"review_course_comment" gorm:"not null;type:text"`
}

type Review_Tutor struct {
	Review_tutor_id      int32  `json:"review_tutor_id" gorm:"primaryKey;type:int"`
	UserID               int    `json:"user_id" gorm:"index;type:int"`
	User                 *User  `json:"user" gorm:"not null;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rating               int32  `json:"rating" gorm:"not null;type:int"`
	Review_tutor_comment string `json:"review_tutor_comment" gorm:"not null;type:text"`
}
