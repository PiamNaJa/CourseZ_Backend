package models

type Post struct {
	Post_id      int32      `json:"post_id" gorm:"primaryKey;type:int"`                                                         // Post_id is the id of the post
	SubjectID    int32      `json:"subject_id" gorm:"index;type:int;not null" validate:"required,number"`                       // SubjectID is the id of the subject
	Subject      *Subject   `json:"subject" gorm:"not null;foreignKey:SubjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Subject is the subject of the post
	UserID       int32      `json:"user_id" gorm:"index;type:int" validate:"required,number"`                                   // UserID is the id of the user
	User         *User      `json:"user" gorm:"not null;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`       // User is the user that made the post
	Class_level  int8       `json:"class_level" gorm:"not null" validate:"required,number"`                                     // Class_level is the level of the class
	Caption      string     `json:"caption" gorm:"not null;type:text" validate:"required"`                                      // Caption is the caption of the post
	Post_picture string     `json:"post_picture" gorm:"not null;type:varchar(5000)" validate:"required,max=5000"`                 // Post_picture is a link to the picture
	Comments     *[]Comment `json:"comments" gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`            // Comments is the comments of the post
}
