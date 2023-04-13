package models

type Exercise struct {
	Exercise_id int32     `json:"exercise_id" gorm:"primaryKey;type:int"`
	VideoID     int32     `json:"video_id" gorm:"index;type:int;not null" validate:"required,number"`       // VideoID is the id of the video
	Question    string    `json:"question" gorm:"not null;type:varchar(1000)" validate:"required,max=1000"` // Question is the question of the exercise
	Image       string    `json:"image" gorm:"type:varchar(5000)" validate:"max=5000"`                      // Image is a link to the image
	Choices     *[]Choice `json:"choices" gorm:"not null;foreignKey:ExerciseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
