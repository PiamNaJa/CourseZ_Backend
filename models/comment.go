package models

type Comment struct {
	Comment_id  int32  `json:"comment_id" gorm:"primaryKey;type:int"`                                               // Comment_id is the id of the comment
	PostID      int32  `json:"post_id" gorm:"index;type:int;not null" validate:"required,number"`                   // PostID is the id of the post
	UserID      int32  `json:"user_id" gorm:"index;type:int;not null" validate:"required,number"`                   // UserID is the id of the user
	User        *User  `json:"user" gorm:"not null;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // User is the user that made the comment
	Description string `json:"description" gorm:"not null;type:text" validate:"required"`                           // Description is the description of the comment
}