package models

type Video struct {
	Video_id    int32           `json:"video_id" gorm:"primaryKey;type:int"`                                                        // Video_id is the id of the video
	CourseID    int32           `json:"course_id" gorm:"index;type:int;not null" validate:"required,number"`                        // CourseID is the id of the course
	Course      *Course         `json:"course" gorm:"not null;foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`    // Course is the Couurse of the video
	Class_level int8            `json:"class_level" gorm:"not null;type:int" validate:"required,number"`                            // Class_level is the level of the class
	Video_name  string          `json:"video_name" gorm:"not null;type:varchar(1000)" validate:"required,max=1000"`                 // Video_name is the name of the video
	Price       int32           `json:"price" gorm:"not null;type:int" validate:"number"`                                           // Price is the price of the video
	Picture     string          `json:"picture" gorm:"not null;type:varchar(5000)" validate:"required,max=5000"`                    // Picture is a link to the picture
	Description string          `json:"description" gorm:"not null;type:text" validate:"required"`                                  // Description is the description of the video
	Url         string          `json:"url" gorm:"not null;type:varchar(5000)" validate:"required,max=5000"`                        // link file
	Sheet       string          `json:"sheet" gorm:"not null;type:varchar(5000)" validate:"max=5000"`                               // link file
	Reviews     *[]Review_Video `json:"reviews" gorm:"foreignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`            // Reviews is the reviews of the video
	Exercises   *[]Exercise     `json:"exercises" gorm:"not null;foreignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Exercises is the exercises of the video
}
