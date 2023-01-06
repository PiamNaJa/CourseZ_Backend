package models

type Exercise struct {
	Exercise_id int32     `json:"exercise_id" gorm:"primaryKey;type:int"`
	VideoID     int32     `json:"video_id" gorm:"index;type:int"`
	Question    string    `json:"question" gorm:"not null;type:varchar(100)"`
	Image       string    `json:"image" gorm:"type:varchar(255)"`
	Choices     *[]Choice `json:"choices" gorm:"not null;foreignKey:ExerciseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
