package models

type Review_Video struct {
	Review_video_id int32   `json:"review_video_id" gorm:"primaryKey;type:int"`
	VideoID         int32   `json:"video_id" gorm:"index;type:int;not null" validate:"required,number"` // VideoID is the id of the video
	Rating          float32 `json:"rating" gorm:"not null;type:float" validate:"required,number"`       // Rating is the rating of the course
	Comment         string  `json:"comment" gorm:"not null;type:text" validate:"required"`
	CreatedAt       int64   `json:"created_at" gorm:"autoCreateTime"`
}

type Review_Tutor struct {
	Review_tutor_id int32   `json:"review_tutor_id" gorm:"primaryKey;type:int"`
	TeacherID       int32   `json:"teacher_id" gorm:"index;type:int;not null" validate:"required,number"` // TeacherID is the id of the teacher
	Rating          float32 `json:"rating" gorm:"not null;type:float" validate:"required,number"`         // Rating is the rating of the course
	Comment         string  `json:"comment" gorm:"not null;type:text" validate:"required"`
	CreatedAt       int64   `json:"created_at" gorm:"autoCreateTime"`
}
