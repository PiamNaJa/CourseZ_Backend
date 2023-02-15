package models

import (
	"time"
)

type User struct {
	User_id     int32        `json:"user_id" gorm:"primaryKey;type:int"`                                                           // User_id is the id of the user
	Email       string       `json:"email" gorm:"unique;not null;type:varchar(100);" validate:"required,email"`                    // Email is the email of the user
	Password    string       `json:"password" gorm:"not null;type:varchar(255)" validate:"required,min=6,max=255"`                 // Password is the password of the user
	Fullname    string       `json:"fullname" gorm:"not null;type:varchar(100)" validate:"required,max=100"`                       // Fullname is the fullname of the user
	Nickname    string       `json:"nickname" gorm:"not null;type:varchar(100)" validate:"required,max=100"`                       // Nickname is the nickname of the user
	Birthday    time.Time    `json:"birthday"  gorm:"not null;type:date"`                                                          // Birthday is the birthday of the user
	Role        string       `json:"role" gorm:"not null;type:varchar(20);default:Student"`                                        // Role is the role of the user
	Picture     string       `json:"picture" gorm:"type:varchar(5000)" validate:"max=5000"`                                        // Picture is a link to the picture
	Point       int32        `json:"point" gorm:"type:int;default:0"`                                                              // Point is the point of the user
	History     *[]History   `json:"history" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`               // History is the history of the user
	Teacher     *UserTeacher `json:"teacher" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`               // Teacher is the teacher of the user
	PaidVideos  []*Video     `json:"paid_videos" gorm:"many2many:user_paidvideos;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`   // PaidVideo is the paid video of the user
	LikeVideos  []*Video     `json:"like_videos" gorm:"many2many:user_likevideos;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`   // LikeVideo is the video user buyed
	LikeCourses []*Course    `json:"like_courses" gorm:"many2many:user_likecourses;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // LikeCourse is the course user likeed
  Tracsaction *[]Payment   `json:"payment" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`             // Tracsaction is the tracsaction of the user
}

type UserTeacher struct {
	Teacher_id         int32           `json:"teacher_id" gorm:"primaryKey;type:int"`                                                // Teacher_id is the id of the teacher
	UserID             int32           `json:"user_id" gorm:"index;type:int"`                                                        // UserID is the id of the user
	Teacher_license    string          `json:"teacher_license" gorm:"not null;type:varchar(5000)" validate:"max=5000"`               // link file
	Transcript         string          `json:"transcript" gorm:"not null;type:varchar(5000)" validate:"max=5000"`                    // link file
	Id_card            string          `json:"id_card" gorm:"not null;type:varchar(5000)" validate:"required,max=5000"`              // link file
	Psychological_test string          `json:"psychological_test" gorm:"not null;type:varchar(255)" validate:"max=255"`              // link file
	Money              int32           `json:"money" gorm:"type:int;default:0"`                                                      // Money is the money of the teacher
	Reviews            *[]Review_Tutor `json:"reviews" gorm:"foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`    // Reviews is the reviews of the teacher
	Experience         *[]Experience   `json:"experience" gorm:"foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Experience is the experience of the teacher
	Courses            *[]Course       `json:"courses" gorm:"foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`    // Courses is the courses of the teacher
	Tracsaction        *[]Withdraw     `json:"withdraw" gorm:"foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`   // Tracsaction is the tracsaction of the teacher
}

type Experience struct {
	Experience_id int32  `json:"experience_id" gorm:"primaryKey;type:int"`                                 // Experience_id is the id of the experience
	TeacherID     int32  `json:"teacher_id" gorm:"index;type:int"`                                         // TeacherID is the id of the teacher
	Title         string `json:"title" gorm:"not null;type:varchar(100)" validate:"required,max=100"`      // Title is the title of the experience
	Evidence      string `json:"evidence" gorm:"not null;type:varchar(5000)" validate:"required,max=5000"` // link file
}

// type LoginData struct {
// 	Fullname string       `json:"fullname"` // Fullname is the fullname of the user
// 	Nickname string       `json:"nickname"` // Nickname is the nickname of the user
// 	Birthday time.Time    `json:"birthday"` // Birthday is the birthday of the user
// 	Role     string       `json:"role"`     // Role is the role of the user
// 	Picture  string       `json:"picture"`  // Picture is a link to the picture
// 	Point    int32        `json:"point"`    // Point is the point of the user
// 	History  *[]History   `json:"history"`  // History is the history of the user
// 	Teacher  *UserTeacher `json:"teacher"`  // Teacher is the teacher of the user
// }
