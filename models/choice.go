package models

type Choice struct {
	Choice_id  int32  `json:"choice_id" gorm:"primaryKey;type:int"`                                  // Choice_id is the id of the choice
	ExerciseID int32  `json:"exercise_id" gorm:"index;type:int;not null" validate:"required,number"` // ExerciseID is the id of the exercise
	Title      string `json:"title" gorm:"not null;type:varchar(100)" validate:"required,max=100"`   // Title is the title of the choice
	Correct    bool   `json:"correct" gorm:"not null;default:false;type:bool"`                       // Correct is the this choice is correct or not
}
