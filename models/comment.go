package models

type Comment struct {
	Comment_id  int32  `json:"comment_id" gorm:"primaryKey;type:int"`
	PostID      int32  `json:"post_id" gorm:"index;type:int"`
	Post        *Post  `json:"post" gorm:"not null;foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID      int32  `json:"user_id" gorm:"index;type:int"`
	User        *User  `json:"user" gorm:"not null;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Description string `json:"description" gorm:"not null;type:text"`
}
