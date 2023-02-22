package models

type Withdraw struct {
	Withdraw_id int32        `json:"withdraw_id" gorm:"primaryKey;type:int"`
	TeacherID   int32        `json:"recipient(teacher_id)" gorm:"index;type:int" validate:"required,number"`
	Recipient   *UserTeacher `json:"recipient" gorm:"not null;foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Money       int32        `json:"money" gorm:"type:int;default:0 "`
	Text        string       `json:"text" gorm:"not null;type:varchar(100)" validate:"required,max=100"`
	CreatedAt   int64        `json:"created_at" gorm:"autoCreateTime"`
}
