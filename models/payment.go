package models

type Payment struct {
	Payment_id int32        `json:"payment_id" gorm:"primaryKey;type:int"`
	UserID     int32        `json:"payee(user_id)" gorm:"index;type:int" validate:"required,number"`
	Payee      *User        `json:"payee" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeacherID  int32        `json:"recipient(teacher_id)" gorm:"index;type:int" validate:"required,number"`
	Recipient  *UserTeacher `json:"recipient" gorm:"not null;foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Money      int32        `json:"money" gorm:"type:int;default:0"`
	Text       string       `json:"text" gorm:"not null;type:varchar(100)" validate:"required,max=100"`
	VideoID    int32        `json:"video_id" gorm:"index;type:int" validate:"required,number"`
	Video      *Video       `json:"video" gorm:"foreignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  int64        `json:"created_at" gorm:"autoCreateTime"`
}
