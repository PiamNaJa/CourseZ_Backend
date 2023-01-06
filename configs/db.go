package configs

import (
	"os"

	"github.com/PiamNaJa/CourseZ_Backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.UserTeacher{}, &models.Experience{}, &models.Comment{}, &models.Course{}, &models.History{}, &models.Post{}, &models.Review_Video{}, &models.Review_Tutor{}, &models.Reward_Info{}, &models.Reward_Item{}, &models.Subject{}, &models.Video{}, &models.Exercise{}, &models.Choice{})
	DB = db
}
