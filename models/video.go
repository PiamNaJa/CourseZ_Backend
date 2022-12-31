package models

type Video struct {
	Video_id    int32   `json:"video_id" gorm:"primaryKey;type:int"`
	CourseID    int32   `json:"course_id" gorm:"index;type:int"`
	Course      *Course `json:"course" gorm:"not null;foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Video_name  string  `json:"video_name" gorm:"not null;type:varchar(100)"`
	Price       int32   `json:"price" gorm:"not null;type:int"`
	Picture     string  `json:"picture" gorm:"not null;type:varchar(255)"`
	Description string  `json:"description" gorm:"not null;type:text"`
	Url         string  `json:"url" gorm:"not null;type:varchar(255)"`
}
