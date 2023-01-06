package models

type Choice struct {
	Choice_id  int32  `json:"choice_id" gorm:"primaryKey;type:int"`
	ExerciseID int32  `json:"exercise_id" gorm:"index;type:int"`
	Title      string `json:"title" gorm:"not null;type:varchar(100)"`
	Correct    bool   `json:"correct" gorm:"not null;default:false;type:bool"`
}
