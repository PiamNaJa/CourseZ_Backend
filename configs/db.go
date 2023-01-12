package configs

import (
	"fmt"
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
	DB = db
}

func SeedDB() {
	if err := DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Subject{}).Error; err != nil {
		panic(err)
	}
	var subject = &[]models.Subject{
		{
			Subject_title:   "คณิตศาสตร์",
			Class_level:     1,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/546/546743.png",
		},
		{
			Subject_title:   "คณิตศาสตร์",
			Class_level:     2,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/546/546743.png",
		},
		{
			Subject_title:   "คณิตศาสตร์",
			Class_level:     3,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/546/546743.png",
		},
		{
			Subject_title:   "คณิตศาสตร์",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3153/3153038.png",
		},
		{
			Subject_title:   "คณิตศาสตร์",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3153/3153038.png",
		},
		{
			Subject_title:   "คณิตศาสตร์",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3153/3153038.png",
		},
		{
			Subject_title:   "แคลคูลัส",
			Class_level:     7,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/2106/2106460.png",
		},
		{
			Subject_title:   "วิทยาศาสตร์",
			Class_level:     1,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1048/1048971.png",
		},
		{
			Subject_title:   "วิทยาศาสตร์",
			Class_level:     2,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1048/1048971.png",
		},
		{
			Subject_title:   "วิทยาศาสตร์",
			Class_level:     3,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1048/1048971.png",
		},
		{
			Subject_title:   "ประวัติศาตร์",
			Class_level:     1,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1183/1183801.png",
		},
		{
			Subject_title:   "ประวัติศาตร์",
			Class_level:     2,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1183/1183801.png",
		},
		{
			Subject_title:   "ประวัติศาตร์",
			Class_level:     3,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1183/1183801.png",
		},
		{
			Subject_title:   "ภาษาไทย",
			Class_level:     1,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3054/3054084.png",
		},
		{
			Subject_title:   "ภาษาไทย",
			Class_level:     2,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3054/3054084.png",
		},
		{
			Subject_title:   "ภาษาไทย",
			Class_level:     3,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3054/3054084.png",
		},
		{
			Subject_title:   "ภาษาไทย",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3054/3054084.png",
		},
		{
			Subject_title:   "ภาษาไทย",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3054/3054084.png",
		},
		{
			Subject_title:   "ภาษาไทย",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3054/3054084.png",
		},
		{
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     1,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     2,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     3,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     7,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			Subject_title:   "สังคมศึกษา",
			Class_level:     1,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/921/921490.png",
		},
		{
			Subject_title:   "สังคมศึกษา",
			Class_level:     2,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/921/921490.png",
		},
		{
			Subject_title:   "สังคมศึกษา",
			Class_level:     3,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/921/921490.png",
		},
		{
			Subject_title:   "สังคมศึกษา",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/921/921490.png",
		},
		{
			Subject_title:   "สังคมศึกษา",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/921/921490.png",
		},
		{
			Subject_title:   "สังคมศึกษา",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/921/921490.png",
		},
		{
			Subject_title:   "ฟิสิกส์",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3247/3247965.png",
		},
		{
			Subject_title:   "ฟิสิกส์",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3247/3247965.png",
		},
		{
			Subject_title:   "ฟิสิกส์",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3247/3247965.png",
		},
		{
			Subject_title:   "ฟิสิกส์",
			Class_level:     7,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/7275/7275469.png",
		},
		{
			Subject_title:   "เคมี",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1530/1530141.png",
		},
		{
			Subject_title:   "เคมี",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1530/1530141.png",
		},
		{
			Subject_title:   "เคมี",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1530/1530141.png",
		},
		{
			Subject_title:   "เคมี",
			Class_level:     7,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1055/1055613.png",
		},
		{
			Subject_title:   "ชีววิทยา",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1186/1186505.png",
		},
		{
			Subject_title:   "ชีววิทยา",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1186/1186505.png",
		},
		{
			Subject_title:   "ชีววิทยา",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1186/1186505.png",
		},
		{
			Subject_title:   "ชีววิทยา",
			Class_level:     7,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1186/1186523.png",
		},
	}
	if err := DB.Model(&models.Subject{}).Create(&subject).Error; err != nil {
		panic(err)
	}
	fmt.Println("Subject created")
}

func WipeData() {
	DB.Migrator().DropTable(&models.User{}, &models.UserTeacher{}, &models.Experience{}, &models.Comment{}, &models.Course{}, &models.History{}, &models.Post{}, &models.Review_Video{}, &models.Review_Tutor{}, &models.Reward_Info{}, &models.Reward_Item{}, &models.Subject{}, &models.Video{}, &models.Exercise{}, &models.Choice{})
}

func MigrateData() {
	DB.AutoMigrate(&models.User{}, &models.UserTeacher{}, &models.Experience{}, &models.Comment{}, &models.Course{}, &models.History{}, &models.Post{}, &models.Review_Video{}, &models.Review_Tutor{}, &models.Reward_Info{}, &models.Reward_Item{}, &models.Subject{}, &models.Video{}, &models.Exercise{}, &models.Choice{})
}
