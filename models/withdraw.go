package models

type Withdraw struct {
	Withdraw_id int32  `json:"withdraw_id" gorm:"primaryKey;type:int"`
	TeacherID   int32  `json:"recipient(teacher_id)" gorm:"index;type:int" validate:"required,number"`
	Money       int32  `json:"money" gorm:"type:int;default:0 "`
	BankName    string `json:"bankname" gorm:"not null;type:varchar(12)" validate:"required"`
	BankNumber  string `json:"banknumber" gorm:"not null;type:varchar(20)" validate:"required,max=10,min=10"`
	CreatedAt   int64  `json:"created_at" gorm:"autoCreateTime"`
}
