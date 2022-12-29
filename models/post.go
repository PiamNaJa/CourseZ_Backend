package models

type Post struct {
	Post_id      int32    `json:"post_id" gorm:"primaryKey;type:int"`
	SubjectID    int32    `json:"subject_id" gorm:"index;type:int"`
	Subject      *Subject `json:"subject" gorm:"not null;foreignKey:SubjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID       int32    `json:"user_id" gorm:"index;type:int"`
	User         *User    `json:"user" gorm:"not null;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Class_level  int8     `json:"class_level" gorm:"not null"`
	Caption      string   `json:"caption" gorm:"not null;type:text"`
	Post_picture string   `json:"post_picture" gorm:"not null;type:varchar(255)"`
}
