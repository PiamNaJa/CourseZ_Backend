package models

type Review_Video struct {
	Review_video_id int32  `json:"review_course_id" gorm:"primaryKey;type:int"`
	VideoID         int32  `json:"course_id" gorm:"index;type:int"`
	Rating          int32  `json:"rating" gorm:"not null;type:int"`
	Comment         string `json:"comment" gorm:"not null;type:text"`
}

type Review_Tutor struct {
	Review_tutor_id int32  `json:"review_tutor_id" gorm:"primaryKey;type:int"`
	TeacherID       int32  `json:"teacher_id" gorm:"index;type:int"`
	Rating          int32  `json:"rating" gorm:"not null;type:int"`
	Comment         string `json:"comment" gorm:"not null;type:text"`
}
