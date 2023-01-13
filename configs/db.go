package configs

import (
	"fmt"
	"os"
	"time"

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
	//Drop All Rows
	if err := DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Subject{}, &models.UserTeacher{}, &models.Course{}).Error; err != nil {
		panic(err)
	}

	//Subjects
	var subject = &[]models.Subject{
		{
			//SubjectID: 1
			Subject_title:   "คณิตศาสตร์",
			Class_level:     1,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/546/546743.png",
		},
		{
			//SubjectID: 2
			Subject_title:   "คณิตศาสตร์",
			Class_level:     2,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/546/546743.png",
		},
		{
			//SubjectID: 3
			Subject_title:   "คณิตศาสตร์",
			Class_level:     3,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/546/546743.png",
		},
		{
			//SubjectID: 4
			Subject_title:   "คณิตศาสตร์",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3153/3153038.png",
		},
		{
			//SubjectID: 5
			Subject_title:   "คณิตศาสตร์",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3153/3153038.png",
		},
		{
			//SubjectID: 6
			Subject_title:   "คณิตศาสตร์",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3153/3153038.png",
		},
		{
			//SubjectID: 7
			Subject_title:   "แคลคูลัส",
			Class_level:     7,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/2106/2106460.png",
		},
		{
			//SubjectID: 8
			Subject_title:   "วิทยาศาสตร์",
			Class_level:     1,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1048/1048971.png",
		},
		{
			//SubjectID: 9
			Subject_title:   "วิทยาศาสตร์",
			Class_level:     2,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1048/1048971.png",
		},
		{
			//SubjectID: 10
			Subject_title:   "วิทยาศาสตร์",
			Class_level:     3,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1048/1048971.png",
		},
		{
			//SubjectID: 11
			Subject_title:   "ประวัติศาตร์",
			Class_level:     1,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1183/1183801.png",
		},
		{
			//SubjectID: 12
			Subject_title:   "ประวัติศาตร์",
			Class_level:     2,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1183/1183801.png",
		},
		{
			//SubjectID: 13
			Subject_title:   "ประวัติศาตร์",
			Class_level:     3,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1183/1183801.png",
		},
		{
			//SubjectID: 14
			Subject_title:   "ภาษาไทย",
			Class_level:     1,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3054/3054084.png",
		},
		{
			//SubjectID: 15
			Subject_title:   "ภาษาไทย",
			Class_level:     2,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3054/3054084.png",
		},
		{
			//SubjectID: 16
			Subject_title:   "ภาษาไทย",
			Class_level:     3,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3054/3054084.png",
		},
		{
			//SubjectID: 17
			Subject_title:   "ภาษาไทย",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3054/3054084.png",
		},
		{
			//SubjectID: 18
			Subject_title:   "ภาษาไทย",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3054/3054084.png",
		},
		{
			//SubjectID: 19
			Subject_title:   "ภาษาไทย",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3054/3054084.png",
		},
		{
			//SubjectID: 20
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     1,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			//SubjectID: 21
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     2,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			//SubjectID: 22
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     3,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			//SubjectID: 23
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			//SubjectID: 24
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			//SubjectID: 25
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			//SubjectID: 26
			Subject_title:   "ภาษาอังกฤษ",
			Class_level:     7,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/299/299688.png",
		},
		{
			//SubjectID: 27
			Subject_title:   "สังคมศึกษา",
			Class_level:     1,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/921/921490.png",
		},
		{
			//SubjectID: 28
			Subject_title:   "สังคมศึกษา",
			Class_level:     2,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/921/921490.png",
		},
		{
			//SubjectID: 29
			Subject_title:   "สังคมศึกษา",
			Class_level:     3,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/921/921490.png",
		},
		{
			//SubjectID: 30
			Subject_title:   "สังคมศึกษา",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/921/921490.png",
		},
		{
			//SubjectID: 31
			Subject_title:   "สังคมศึกษา",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/921/921490.png",
		},
		{
			//SubjectID: 32
			Subject_title:   "สังคมศึกษา",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/921/921490.png",
		},
		{
			//SubjectID: 33
			Subject_title:   "ฟิสิกส์",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3247/3247965.png",
		},
		{
			//SubjectID: 34
			Subject_title:   "ฟิสิกส์",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3247/3247965.png",
		},
		{
			//SubjectID: 35
			Subject_title:   "ฟิสิกส์",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/3247/3247965.png",
		},
		{
			//SubjectID: 36
			Subject_title:   "ฟิสิกส์",
			Class_level:     7,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/7275/7275469.png",
		},
		{
			//SubjectID: 37
			Subject_title:   "เคมี",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1530/1530141.png",
		},
		{
			//SubjectID: 38
			Subject_title:   "เคมี",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1530/1530141.png",
		},
		{
			//SubjectID: 39
			Subject_title:   "เคมี",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1530/1530141.png",
		},
		{
			//SubjectID: 40
			Subject_title:   "เคมี",
			Class_level:     7,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1055/1055613.png",
		},
		{
			//SubjectID: 41
			Subject_title:   "ชีววิทยา",
			Class_level:     4,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1186/1186505.png",
		},
		{
			//SubjectID: 42
			Subject_title:   "ชีววิทยา",
			Class_level:     5,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1186/1186505.png",
		},
		{
			//SubjectID: 43
			Subject_title:   "ชีววิทยา",
			Class_level:     6,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1186/1186505.png",
		},
		{
			//SubjectID: 44
			Subject_title:   "ชีววิทยา",
			Class_level:     7,
			Subject_picture: "https://cdn-icons-png.flaticon.com/512/1186/1186523.png",
		},
	}
	if err := DB.Model(&models.Subject{}).Create(&subject).Error; err != nil {
		panic(err)
	}

	fmt.Println("Subject created")

	//Teacher

	var teacher = &[]models.User{
		{
			Email:    "teacher1@mail.com", //ครูคณิต
			Password: "1234",
			Fullname: "สุนีย์ คิดดี",
			Nickname: "ครูกุ๊กไก่",
			Birthday: time.Now(),
			Role:     "Teacher",
			Picture:  "https://shorturl.asia/CZMPx",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "https://shorturl.asia/SecRv",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูสอนพิเศษตามบ้าน สอนเด็กเล็ก ชั้นอนุบาล-ประถมต้น",
						Evidence: "https://shorturl.asia/4iuPM",
					},
				},
			},
		},
		{
			Email:    "teacher2@mail.com",
			Password: "1234",
			Fullname: "วรรณษา ทองสุก",
			Nickname: "ครูเดือน",
			Birthday: time.Now(),
			Role:     "Tutor",
			Picture:  "https://shorturl.asia/PoI4T",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "https://shorturl.asia/SecRv",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูสอนพิเศษตามบ้าน สอนเด็กเล็ก ชั้นอนุบาล-ประถมต้น",
						Evidence: "https://shorturl.asia/4iuPM",
					},
				},
			},
		},
		{
			Email:    "teacher3@mail.com",
			Password: "1234",
			Fullname: "สุชาติ ม่วงดี",
			Nickname: "ครูเมฆ",
			Birthday: time.Now(),
			Role:     "Teacher",
			Picture:  "https://shorturl.asia/3SM9q",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "https://shorturl.asia/SecRv",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูสอนพิเศษตามบ้าน สอนเด็กเล็ก ชั้นอนุบาล-ประถมต้น",
						Evidence: "https://shorturl.asia/4iuPM",
					},
				},
			},
		},
		{
			Email:    "teacher4@mail.com", //ครูคณิต
			Password: "1234",
			Fullname: "สิริยากร ผ่องพันธ์",
			Nickname: "ครูดาว",
			Birthday: time.Now(),
			Role:     "Teacher",
			Picture:  "https://shorturl.asia/FRYPW",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "https://shorturl.asia/SecRv",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูจิตอาสา สอนหนังสือให้เด็กเล็ก",
						Evidence: "https://shorturl.asia/V7k58",
					},
				},
			},
		},
		{
			Email:    "teacher5@mail.com", //ครูคณิต
			Password: "1234",
			Fullname: "นนทการ ศรีสุข",
			Nickname: "ครูสิงโต",
			Birthday: time.Now(),
			Role:     "Teacher",
			Picture:  "https://shorturl.asia/X1iyt",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "https://shorturl.asia/SecRv",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูสอนพิเศษตามบ้าน สอนเด็กเล็ก ชั้นอนุบาล-ประถมต้น",
						Evidence: "https://shorturl.asia/4iuPM",
					},
				},
			},
		},
		{
			Email:    "teacher6@mail.com", //ครูอังกฤษ
			Password: "1234",
			Fullname: "นาตาชา เเจ่มใส",
			Nickname: "ครูข้าวต้ม",
			Birthday: time.Now(),
			Role:     "Teacher",
			Picture:  "https://shorturl.asia/rxl37",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "https://shorturl.asia/SecRv",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูจิตอาสา สอนหนังสือให้เด็กเล็ก",
						Evidence: "https://shorturl.asia/V7k58",
					},
				},
			},
		},
		{
			Email:    "teacher7@mail.com",
			Password: "1234",
			Fullname: "สมเกียรติ มั่งมี",
			Nickname: "ครูเฟียส",
			Birthday: time.Now(),
			Role:     "Teacher",
			Picture:  "https://shorturl.asia/iKC41",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "https://shorturl.asia/SecRv",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูสอนพิเศษตามบ้าน สอนเด็กเล็ก ชั้นอนุบาล-ประถมต้น",
						Evidence: "https://shorturl.asia/4iuPM",
					},
				},
			},
		},
		{
			Email:    "teacher8@mail.com",
			Password: "1234",
			Fullname: "เเพรวา เเจ่มเเจ้ง",
			Nickname: "ครูปลา",
			Birthday: time.Now(),
			Role:     "Teacher",
			Picture:  "https://shorturl.asia/ri1Av",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "https://shorturl.asia/SecRv",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูจิตอาสา สอนหนังสือให้เด็กเล็ก",
						Evidence: "https://shorturl.asia/V7k58",
					},
				},
			},
		},
		{
			Email:    "teacher9@mail.com",
			Password: "1234",
			Fullname: "ภูวินทร์ สีม่วง",
			Nickname: "ครูติ๊ก",
			Birthday: time.Now(),
			Role:     "Teacher",
			Picture:  "https://shorturl.asia/Ckv3h",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "https://shorturl.asia/SecRv",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูสอนพิเศษตามบ้าน สอนเด็กเล็ก ชั้นอนุบาล-ประถมต้น",
						Evidence: "https://shorturl.asia/4iuPM",
					},
				},
			},
		},
		{
			Email:    "teacher10@mail.com", //ครูอังกฤษ
			Password: "1234",
			Fullname: "ฟ้าใส ชื่นมื่น",
			Nickname: "ครูฟ้า",
			Birthday: time.Now(),
			Role:     "Teacher",
			Picture:  "https://shorturl.asia/q7tQd",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "https://shorturl.asia/SecRv",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูจิตอาสา สอนหนังสือให้เด็กเล็ก",
						Evidence: "https://shorturl.asia/V7k58",
					},
				},
			},
		},
	}
	if err := DB.Model(&models.UserTeacher{}).Create(&teacher).Error; err != nil {
		panic(err)
	}

	fmt.Println("Teacher created")

	//Course
	var course = &[]models.Course{
		{
			//Course_1
			SubjectID: 1,
			Videos: &[]models.Video{
				{
					Video_name:  "จำนวนเต็มเเละการบวกจำนวนเต็ม",
					Picture:     "https://i.ytimg.com/vi/unXOuA0PkCQ/maxresdefault.jpg",
					Description: "จำนวนเต็ม ม.1 EP.1/3 | ทบทวนพื้นฐานจำนวนเต็ม และ การบวกจำนวนเต็ม เป็นคลิปสอนตั้งแต่พื้นฐานของจำนวนเต็ม ม.1 และการบวกจำนวนเต็มแบบละเอียด ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_1%2FVideo_1%2F%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%95%E0%B9%87%E0%B8%A1%20%E0%B8%A1.1%20EP.1_3%20_%20%E0%B8%97%E0%B8%9A%E0%B8%97%E0%B8%A7%E0%B8%99%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%90%E0%B8%B2%E0%B8%99%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%95%E0%B9%87%E0%B8%A1%20%E0%B9%81%E0%B8%A5%E0%B8%B0%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%9A%E0%B8%A7%E0%B8%81%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%95%E0%B9%87%E0%B8%A1.mp4?alt=media&token=b1e9102e-d19f-4b58-9849-3ec18a809300",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_1%2FVideo_1%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%95%E0%B9%87%E0%B8%A1%20%E0%B8%A1.1%20(course_1-Video_1).pdf?alt=media&token=20bb4339-a353-4c68-918b-c6a39cc7ba27",
					Exercises: &[]models.Exercise{
						{
							Question: "1.จงหาผลบวกของ 15 + 43",
							Choices: &[]models.Choice{
								{
									Title:   "55",
									Correct: false,
								},
								{
									Title:   "56",
									Correct: false,
								},
								{
									Title:   "57",
									Correct: false,
								},
								{
									Title:   "58",
									Correct: true,
								},
							},
						},
						{
							Question: "2.จงหาผลบวกของ (-8) + (-6)",
							Choices: &[]models.Choice{
								{
									Title:   "-13",
									Correct: false,
								},
								{
									Title:   "-14",
									Correct: true,
								},
								{
									Title:   "14",
									Correct: false,
								},
								{
									Title:   "13",
									Correct: false,
								},
							},
						},
						{
							Question: "3.จงหาผลบวกของ 38 + (-49)",
							Choices: &[]models.Choice{
								{
									Title:   "9",
									Correct: false,
								},
								{
									Title:   "-9",
									Correct: false,
								},
								{
									Title:   "11",
									Correct: false,
								},
								{
									Title:   "-11",
									Correct: true,
								},
							},
						},
						{
							Question: "4.จงหาผลบวกของ (-25) + 32",
							Choices: &[]models.Choice{
								{
									Title:   "9",
									Correct: false,
								},
								{
									Title:   "12",
									Correct: false,
								},
								{
									Title:   "7",
									Correct: true,
								},
								{
									Title:   "-7",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "การลบจำนวนเต็ม",
					Picture:     "https://i.ytimg.com/vi/608EppPrPrQ/maxresdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอนตั้งแต่พื้นฐานของจำนวนเต็ม ม.1 และการลบจำนวนเต็มแบบละเอียด  ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_1%2FVideo_2%2F%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%95%E0%B9%87%E0%B8%A1%20%E0%B8%A1.1%20EP.2_3%20_%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A5%E0%B8%9A%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%95%E0%B9%87%E0%B8%A1.mp4?alt=media&token=0318a318-5440-4fbf-ae6b-60d84f6bba54",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_1%2FVideo_1%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%95%E0%B9%87%E0%B8%A1%20%E0%B8%A1.1%20(course_1-Video_1).pdf?alt=media&token=20bb4339-a353-4c68-918b-c6a39cc7ba27",
					Exercises: &[]models.Exercise{
						{
							Question: "1.จงหาผลลบของ 8 - 12",
							Choices: &[]models.Choice{
								{
									Title:   "-7",
									Correct: false,
								},
								{
									Title:   "6",
									Correct: false,
								},
								{
									Title:   "4",
									Correct: false,
								},
								{
									Title:   "-4",
									Correct: true,
								},
							},
						},
						{
							Question: "2.จงหาผลลบของ -9 - 12",
							Choices: &[]models.Choice{
								{
									Title:   "20",
									Correct: false,
								},
								{
									Title:   "-21",
									Correct: true,
								},
								{
									Title:   "21",
									Correct: false,
								},
								{
									Title:   "-20",
									Correct: false,
								},
							},
						},
						{
							Question: "3.จงหาผลบวกของ -8 - (-10)",
							Choices: &[]models.Choice{
								{
									Title:   "2",
									Correct: true,
								},
								{
									Title:   "-2",
									Correct: false,
								},
								{
									Title:   "3",
									Correct: false,
								},
								{
									Title:   "5",
									Correct: false,
								},
							},
						},
						{
							Question: "4.จงหาผลลบของ -4 - (-8)",
							Choices: &[]models.Choice{
								{
									Title:   "-8",
									Correct: false,
								},
								{
									Title:   "8",
									Correct: false,
								},
								{
									Title:   "-4",
									Correct: false,
								},
								{
									Title:   "4",
									Correct: true,
								},
							},
						},
					},
				},
				{
					Video_name:  "การคูณจำนวนเต็ม การหารจำนวนเต็ม สมบัติของจำนวนเต็ม",
					Price:       25,
					Picture:     "https://shorturl.asia/uv4mi",
					Description: "จำนวนเต็ม ม.1 EP.3/3 | ทบทวนพื้นฐานจำนวนเต็ม และ การการคูณจำนวนเต็ม การหารจำนวนเต็ม สมบัติของจำนวนเต็ม ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_1%2FVideo_3%2F%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%95%E0%B9%87%E0%B8%A1%20%E0%B8%A1.1%20EP.3_3%20_%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%84%E0%B8%B9%E0%B8%93%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%95%E0%B9%87%E0%B8%A1%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%AB%E0%B8%B2%E0%B8%A3%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%95%E0%B9%87%E0%B8%A1%20%E0%B8%AA%E0%B8%A1%E0%B8%9A%E0%B8%B1%E0%B8%95%E0%B8%B4%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%95%E0%B9%87%E0%B8%A1.mp4?alt=media&token=69e04ff4-bd48-400b-b34b-219c50d3bff1",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_1%2FVideo_1%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%95%E0%B9%87%E0%B8%A1%20%E0%B8%A1.1%20(course_1-Video_1).pdf?alt=media&token=20bb4339-a353-4c68-918b-c6a39cc7ba27",
					Exercises: &[]models.Exercise{
						{
							Question: "1.หาผลคูณของ 8 x 4",
							Choices: &[]models.Choice{
								{
									Title:   "8",
									Correct: false,
								},
								{
									Title:   "16",
									Correct: false,
								},
								{
									Title:   "24",
									Correct: false,
								},
								{
									Title:   "32",
									Correct: true,
								},
							},
						},
						{
							Question: "2.จงหาผลหารของ 27 ÷ 9",
							Choices: &[]models.Choice{
								{
									Title:   "9",
									Correct: false,
								},
								{
									Title:   "6",
									Correct: false,
								},
								{
									Title:   "3",
									Correct: true,
								},
								{
									Title:   "12",
									Correct: false,
								},
							},
						},
						{
							Question: "3.ข้อใดเป็นสมบัติการสลับที่ของการบวก",
							Choices: &[]models.Choice{
								{
									Title:   "13 + 9 = 9 + 13",
									Correct: true,
								},
								{
									Title:   "(3 + 5) + 4 = 3 + (5 + 4)",
									Correct: false,
								},
								{
									Title:   "9 x (4 + 5) = (9 x 4) + (9 x 5)",
									Correct: false,
								},
								{
									Title:   "30 x 1 = 30",
									Correct: false,
								},
							},
						},
						{
							Question: "4.ข้อใดเป็นสมบัติการเเจกเเจง",
							Choices: &[]models.Choice{
								{
									Title:   "25 + 8 = 8 + 25",
									Correct: false,
								},
								{
									Title:   "(9 + 25) + 2 = 9 + (25 + 2)",
									Correct: false,
								},
								{
									Title:   "(4 x 5) x 7 = 4 x (5 x 7)",
									Correct: false,
								},
								{
									Title:   "[12 + (-4)] x (-3) = [12 x (-3)] + [(-4) x (-3)]",
									Correct: true,
								},
							},
						},
					},
				},
			},
			TeacherID:   1,
			Course_name: "รู้จักกับจำนวนเต็ม",
			Picture:     "https://shorturl.asia/vlAMb",
			Description: "หลักการพื้นฐานสำคัญและเป็นจุดเริ่มต้นของหลักการอื่นๆในวิชาคณิตศาสตร์ โดยในชั้นม.ต้น เนื้อหาจะครอบคลุมตั้งแต่ “จำนวนเต็ม” คืออะไร มีกี่อย่างและมีอะไรบ้าง รวมไปถึงการเปรียบเทียบ และการบวก ลบ คูณ หารจำนวนเต็ม",
		},
		{
			//Course_2
			SubjectID: 2,
			Videos: &[]models.Video{
				{
					Video_name:  "การแยกตัวประกอบของพหุนามดีกรีสอง (พหุนามดีกรีสองตัวแปรเดียว)",
					Picture:     "https://www.tertututor.com/images/cover/factorize-degree-2.png",
					Description: "การเเยกตัวประกอบ ม.2 | ทบทวนการแยกตัวประกอบของพหุนามดีกรีสอง (พหุนามดีกรีสองตัวแปรเดียว) แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนคณิตศาสตร์ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_2%2FVideo_1%2F(1)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%84%E0%B8%93%E0%B8%B4%E0%B8%95%E0%B8%A8%E0%B8%B2%E0%B8%AA%E0%B8%95%E0%B8%A3%E0%B9%8C%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.2%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%81%E0%B8%A2%E0%B8%81%E0%B8%95%E0%B8%B1%E0%B8%A7%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B8%81%E0%B8%AD%E0%B8%9A%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%9E%E0%B8%AB%E0%B8%B8%E0%B8%99%E0%B8%B2%E0%B8%A1%E0%B8%94%E0%B8%B5%E0%B8%81%E0%B8%A3%E0%B8%B5%E0%B8%AA%E0%B8%AD%E0%B8%87%20(%E0%B8%9E%E0%B8%AB%E0%B8%B8%E0%B8%99%E0%B8%B2%E0%B8%A1%E0%B8%94%E0%B8%B5%E0%B8%81%E0%B8%A3%E0%B8%B5%E0%B8%AA%E0%B8%AD%E0%B8%87%E0%B8%95%E0%B8%B1%E0%B8%A7%E0%B9%81%E0%B8%9B%E0%B8%A3%E0%B9%80%E0%B8%94%E0%B8%B5%E0%B8%A2%E0%B8%A7).mp4?alt=media&token=6ba44259-2a95-43c8-b730-c6cf57500472",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_2%2FVideo_1%2F%E0%B8%97%E0%B8%9A%E0%B8%97%E0%B8%A7%E0%B8%99%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%81%E0%B8%A2%E0%B8%81%E0%B8%95%E0%B8%B1%E0%B8%A7%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B8%81%E0%B8%AD%E0%B8%9A%E0%B8%9E%E0%B8%AB%E0%B8%B8%E0%B8%99%E0%B8%B2%E0%B8%A1%E0%B8%94%E0%B8%B5%E0%B8%81%E0%B8%A3%E0%B8%B5%E0%B8%AA%E0%B8%AD%E0%B8%87.pdf?alt=media&token=3e5a0224-158b-4092-8b39-3273c2cc096a",
					Exercises: &[]models.Exercise{
						{
							Question: "1.จงแยกตัวประกอบของ x กำลังสอง + 3x - 4",
							Choices: &[]models.Choice{
								{
									Title:   "(x-3)(x+4)",
									Correct: false,
								},
								{
									Title:   "(x+3)(x+4)",
									Correct: false,
								},
								{
									Title:   "(x-2)(x+4))",
									Correct: false,
								},
								{
									Title:   "(x-1)(x+4)",
									Correct: true,
								},
							},
						},
						{
							Question: "2.จงแยกตัวประกอบของ x กำลังสอง + 10x + 25",
							Choices: &[]models.Choice{
								{
									Title:   "(x-5)(x-5)",
									Correct: false,
								},
								{
									Title:   "(x+5)(x+5)",
									Correct: true,
								},
								{
									Title:   "(x-5)(x+5)",
									Correct: false,
								},
								{
									Title:   "(x+5)(x-5)",
									Correct: false,
								},
							},
						},
						{
							Question: "3.จงแยกตัวประกอบของ 2x กำลังสอง - 7x + 6",
							Choices: &[]models.Choice{
								{
									Title:   "(2x-3)(x-2)",
									Correct: true,
								},
								{
									Title:   "(2x-7)(x-6)",
									Correct: false,
								},
								{
									Title:   "(2x+7)(x-6)",
									Correct: false,
								},
								{
									Title:   "(2x+3)(x+2)",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "การแยกตัวประกอบของพหุนามดีกรีสอง (การแจกแจงและดึงตัวร่วม)",
					Picture:     "https://www.tertututor.com/images/cover/factorize-degree-2.png",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอนการเเยกตัวประกอบของพหุนามดีกรีสอง (การเเจกเเจงเเละดึงตัวร่วม) ม.2 แบบละเอียด เข้าใจง่าย  ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_2%2FVideo_2%2F(2)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%84%E0%B8%93%E0%B8%B4%E0%B8%95%E0%B8%A8%E0%B8%B2%E0%B8%AA%E0%B8%95%E0%B8%A3%E0%B9%8C%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.2%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%81%E0%B8%A2%E0%B8%81%E0%B8%95%E0%B8%B1%E0%B8%A7%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B8%81%E0%B8%AD%E0%B8%9A%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%9E%E0%B8%AB%E0%B8%B8%E0%B8%99%E0%B8%B2%E0%B8%A1%E0%B8%94%E0%B8%B5%E0%B8%81%E0%B8%A3%E0%B8%B5%E0%B8%AA%E0%B8%AD%E0%B8%87%20(%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%81%E0%B8%88%E0%B8%81%E0%B9%81%E0%B8%88%E0%B8%87%E0%B9%81%E0%B8%A5%E0%B8%B0%E0%B8%94%E0%B8%B6%E0%B8%87%E0%B8%95%E0%B8%B1%E0%B8%A7%E0%B8%A3%E0%B9%88%E0%B8%A7%E0%B8%A1).mp4?alt=media&token=d1551c9a-77de-43f6-81d1-9fded5e7ff10",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_2%2FVideo_2%2F%E0%B8%97%E0%B8%9A%E0%B8%97%E0%B8%A7%E0%B8%99%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%81%E0%B8%A2%E0%B8%81%E0%B8%95%E0%B8%B1%E0%B8%A7%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B8%81%E0%B8%AD%E0%B8%9A%E0%B8%9E%E0%B8%AB%E0%B8%B8%E0%B8%99%E0%B8%B2%E0%B8%A1%E0%B8%94%E0%B8%B5%E0%B8%81%E0%B8%A3%E0%B8%B5%E0%B8%AA%E0%B8%AD%E0%B8%87.pdf?alt=media&token=ad4e23d4-05da-4fc4-ba5c-9621413a2d21",
					Exercises: &[]models.Exercise{
						{
							Question: "1.จงเเยกตัวประกอบของ x กำลังสอง + 2x",
							Choices: &[]models.Choice{
								{
									Title:   "x(x+5)",
									Correct: false,
								},
								{
									Title:   "x(x+4)",
									Correct: false,
								},
								{
									Title:   "x(x+3)",
									Correct: false,
								},
								{
									Title:   "x(x+2)",
									Correct: true,
								},
							},
						},
						{
							Question: "2.จงเเยกตัวประกอบของ x กำลังสอง - 25x",
							Choices: &[]models.Choice{
								{
									Title:   "x(x+25)",
									Correct: false,
								},
								{
									Title:   "x(x-25)",
									Correct: true,
								},
								{
									Title:   "x(x - (-25))",
									Correct: false,
								},
								{
									Title:   "x(x + (-25))",
									Correct: false,
								},
							},
						},
						{
							Question: "3.จงเเยกตัวประกอบของ x กำลังสาม - 4x กำลังสอง",
							Choices: &[]models.Choice{
								{
									Title:   "x กำลังสอง (x-4)",
									Correct: true,
								},
								{
									Title:   "x กำลังสอง (x-2)",
									Correct: false,
								},
								{
									Title:   "x กำลังสาม (x-4)",
									Correct: false,
								},
								{
									Title:   "x กำลังสาม (x-2)",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "การแยกตัวประกอบของพหุนามดีกรีสองที่เป็นกำลังสองสมบูรณ์",
					Price:       25,
					Picture:     "https://www.tertututor.com/images/cover/factorize-degree-2.png",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอนการแยกตัวประกอบของพหุนามดีกรีสองที่เป็นกำลังสองสมบูรณ์ ม.2 แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_2%2FVideo_3%2F(3)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%84%E0%B8%93%E0%B8%B4%E0%B8%95%E0%B8%A8%E0%B8%B2%E0%B8%AA%E0%B8%95%E0%B8%A3%E0%B9%8C%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.2%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%81%E0%B8%A2%E0%B8%81%E0%B8%95%E0%B8%B1%E0%B8%A7%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B8%81%E0%B8%AD%E0%B8%9A%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%9E%E0%B8%AB%E0%B8%B8%E0%B8%99%E0%B8%B2%E0%B8%A1%E0%B8%94%E0%B8%B5%E0%B8%81%E0%B8%A3%E0%B8%B5%E0%B8%AA%E0%B8%AD%E0%B8%87%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%80%E0%B8%9B%E0%B9%87%E0%B8%99%E0%B8%81%E0%B8%B3%E0%B8%A5%E0%B8%B1%E0%B8%87%E0%B8%AA%E0%B8%AD%E0%B8%87%E0%B8%AA%E0%B8%A1%E0%B8%9A%E0%B8%B9%E0%B8%A3%E0%B8%93%E0%B9%8C-(480p).mp4?alt=media&token=6174a660-d3a6-431e-acd1-76413cdc8cee",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_2%2FVideo_3%2F%E0%B8%97%E0%B8%9A%E0%B8%97%E0%B8%A7%E0%B8%99%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%81%E0%B8%A2%E0%B8%81%E0%B8%95%E0%B8%B1%E0%B8%A7%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B8%81%E0%B8%AD%E0%B8%9A%E0%B8%9E%E0%B8%AB%E0%B8%B8%E0%B8%99%E0%B8%B2%E0%B8%A1%E0%B8%94%E0%B8%B5%E0%B8%81%E0%B8%A3%E0%B8%B5%E0%B8%AA%E0%B8%AD%E0%B8%87.pdf?alt=media&token=eb006ae1-21ee-45e2-a804-4c449a0caedd",
					Exercises: &[]models.Exercise{
						{
							Question: "1.จงเเยกตัวประกอบของ x กำลังสอง + 12x + 36",
							Choices: &[]models.Choice{
								{
									Title:   "(x+3)(x+3)",
									Correct: false,
								},
								{
									Title:   "(x+3)(x+6)",
									Correct: false,
								},
								{
									Title:   "(x+12)(x+3)",
									Correct: false,
								},
								{
									Title:   "(x+6)(x+6)",
									Correct: true,
								},
							},
						},
						{
							Question: "2.จงเเยกตัวประกอบของ x กำลังสอง - 10x + 25",
							Choices: &[]models.Choice{
								{
									Title:   "(x-5)(x-10)",
									Correct: false,
								},
								{
									Title:   "(x+5)(x-10)",
									Correct: false,
								},
								{
									Title:   "(x-5)(x-5)",
									Correct: true,
								},
								{
									Title:   "(x+5)(x+5)",
									Correct: false,
								},
							},
						},
						{
							Question: "3.จงเเยกตัวประกอบของ 9x กำลังสอง + 30x + 25",
							Choices: &[]models.Choice{
								{
									Title:   "(3x+5)(3x+5)",
									Correct: true,
								},
								{
									Title:   "(3x+10)(3x+5)",
									Correct: false,
								},
								{
									Title:   "(3x+15)(3x-5)",
									Correct: false,
								},
								{
									Title:   "(3x-5)(3x-5)",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   4,
			Course_name: "คณิตศาสตร์ ชั้น ม.2 เรื่อง การแยกตัวประกอบ",
			Picture:     "https://tuenongfree.xyz/wp-content/uploads/2019/12/การแยกตัวประกอบ.png",
			Description: "“การเเยกตัวประกอบ” คืออะไร คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับการเเยกตัวประกอบ ซึ่งเป็นสิ่งสำคัญมากในคณิตศาสตร์ ได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_3
			SubjectID: 3,
			Videos: &[]models.Video{
				{
					Video_name:  "สรุปสูตร 11 สูตรพื้นที่ สี่เหลี่ยม สามเหลี่ยม วงกลม",
					Picture:     "https://i.ytimg.com/vi/8ER5qRRpyZs/maxresdefault.jpg",
					Description: "คณิตศาสตร์ ชั้น ม.3 เรื่อง พื้นที่ และ ปริมาตร | สรุปสูตร 11 สูตรพื้นที่ สี่เหลี่ยม สามเหลี่ยม วงกลม แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนคณิตศาสตร์ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_3%2FVideo_1%2F(1)%20%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%E0%B8%AA%E0%B8%B9%E0%B8%95%E0%B8%A3%2011%20%E0%B8%AA%E0%B8%B9%E0%B8%95%E0%B8%A3%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%97%E0%B8%B5%E0%B9%88%20%E0%B8%AA%E0%B8%B5%E0%B9%88%E0%B9%80%E0%B8%AB%E0%B8%A5%E0%B8%B5%E0%B9%88%E0%B8%A2%E0%B8%A1%20%E0%B8%AA%E0%B8%B2%E0%B8%A1%E0%B9%80%E0%B8%AB%E0%B8%A5%E0%B8%B5%E0%B9%88%E0%B8%A2%E0%B8%A1%20%E0%B8%A7%E0%B8%87%E0%B8%81%E0%B8%A5%E0%B8%A1%20%E0%B9%84%E0%B8%A1%E0%B9%88%E0%B8%87%E0%B8%87%E0%B9%81%E0%B8%99%E0%B9%88%E0%B8%99%E0%B8%AD%E0%B8%99.mp4?alt=media&token=55d0c899-1c0a-470f-9c3b-26b097f94cf9",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_3%2FVideo_1%2F%E0%B8%AA%E0%B8%B9%E0%B8%95%E0%B8%A3%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%AB%E0%B8%B2%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%80%E0%B9%80%E0%B8%A5%E0%B8%B0%E0%B8%9B%E0%B8%A3%E0%B8%B4%E0%B8%A1%E0%B8%B2%E0%B8%95%E0%B8%A3.pdf?alt=media&token=511316ad-0d83-4320-8a48-05abb78b39df",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดเป็นสูตรการหาพื้นที่ของรูปสามเหลี่ยม",
							Choices: &[]models.Choice{
								{
									Title:   "2 x ฐาน x สูง",
									Correct: false,
								},
								{
									Title:   "3 x ฐาน x สูง",
									Correct: false,
								},
								{
									Title:   "1/2 x ฐาน x สูง",
									Correct: true,
								},
								{
									Title:   "1/3 x ฐาน x สูง",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดเป็นสูตรการหาพื้นที่ของรูปสี่เหลี่ยมจัตตุรัส",
							Choices: &[]models.Choice{
								{
									Title:   "กว้าง x ยาว",
									Correct: false,
								},
								{
									Title:   "ด้าน x ด้าน",
									Correct: true,
								},
								{
									Title:   "ฐาน x สูง",
									Correct: false,
								},
								{
									Title:   "1/2 x ผลคูณเส้นทเเยงมุม",
									Correct: false,
								},
							},
						},
						{
							Question: "3.ข้อใดเป็นสูตรการหาพื้นที่ของรูปวงกลม",
							Choices: &[]models.Choice{
								{
									Title:   "πr²",
									Correct: true,
								},
								{
									Title:   "πr3",
									Correct: false,
								},
								{
									Title:   "2πr",
									Correct: false,
								},
								{
									Title:   "3πr",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "สรุป ทุกสูตร ปริมาตรและพื้นที่ผิว 10 สูตร",
					Picture:     "https://i.ytimg.com/vi/mZrplz8NoNY/maxresdefault.jpg",
					Description: "คณิตศาสตร์ ชั้น ม.3 เรื่อง พื้นที่ และ ปริมาตร | สรุป ทุกสูตร ปริมาตรและพื้นที่ผิว 10 สูตร แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนคณิตศาสตร์ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_3%2FVideo_2%2F(2)%20%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%20%E0%B8%97%E0%B8%B8%E0%B8%81%E0%B8%AA%E0%B8%B9%E0%B8%95%E0%B8%A3%20%E0%B8%9B%E0%B8%A3%E0%B8%B4%E0%B8%A1%E0%B8%B2%E0%B8%95%E0%B8%A3%E0%B9%81%E0%B8%A5%E0%B8%B0%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B8%9C%E0%B8%B4%E0%B8%A7%2010%20%E0%B8%AA%E0%B8%B9%E0%B8%95%E0%B8%A3%E0%B9%83%E0%B8%99%207%20%E0%B8%99%E0%B8%B2%E0%B8%97%E0%B8%B5%20%E0%B8%97%E0%B8%B1%E0%B9%89%E0%B8%87%20%E0%B8%A1.2%20%E0%B9%81%E0%B8%A5%E0%B8%B0%20%E0%B8%A1.3.mp4?alt=media&token=d0e2cbe7-eafa-4712-b819-c89f79322cc8",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_3%2FVideo_2%2F%E0%B8%AA%E0%B8%B9%E0%B8%95%E0%B8%A3%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%AB%E0%B8%B2%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%80%E0%B9%80%E0%B8%A5%E0%B8%B0%E0%B8%9B%E0%B8%A3%E0%B8%B4%E0%B8%A1%E0%B8%B2%E0%B8%95%E0%B8%A3.pdf?alt=media&token=0f5a880e-2d45-43e3-a0ab-b3510b046e06",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดคือความหมายของปริมาตร",
							Choices: &[]models.Choice{
								{
									Title:   "ปริมาณของรูปร่างของสิ่งมีชีวิต",
									Correct: false,
								},
								{
									Title:   "ปริมาณของรูปร่างของสิ่งของ",
									Correct: false,
								},
								{
									Title:   "ปริมาณของรูปทรงสองมิติ",
									Correct: false,
								},
								{
									Title:   "ปริมาณของรูปทรงสามมิติ",
									Correct: true,
								},
							},
						},
						{
							Question: "2.ข้อใดคือความหมายของพื้นที่ผิว",
							Choices: &[]models.Choice{
								{
									Title:   "พื้นที่ข้างในทั้งหมดของวัตถุ",
									Correct: false,
								},
								{
									Title:   "พื้นที่รอบนอกทั้งหมดของวัตถุ",
									Correct: true,
								},
								{
									Title:   "พื้นที่ทั้งหมดของวัตถุ",
									Correct: false,
								},
								{
									Title:   "พื้นที่ส่วนที่สัมผัสกับพื้นดินของวัตถุ",
									Correct: false,
								},
							},
						},
						{
							Question: "3.ข้อใดเป็นสูตรการหาพื้นที่ผิวของปริซึม",
							Choices: &[]models.Choice{
								{
									Title:   "พื้นที่ผิวข้าง + พื้นที่ฐานทั้งสอง",
									Correct: true,
								},
								{
									Title:   "พื้นที่ผิวข้าง x พื้นที่ฐานทั้งสอง",
									Correct: false,
								},
								{
									Title:   "พื้นที่ฐาน x สูง",
									Correct: false,
								},
								{
									Title:   "พื้นที่ฐาน + สูง",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "ปริมาตร ปริซึม เเละการหาปริมาตรของปริซึม",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/3GYgq9C2aYQ/maxresdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอนการหาปริมาตรของปริซึม แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_3%2FVideo_3%2F(3)%20%E0%B8%9B%E0%B8%A3%E0%B8%B4%E0%B8%A1%E0%B8%B2%E0%B8%95%E0%B8%A3%20%E0%B8%9B%E0%B8%A3%E0%B8%B4%E0%B8%8B%E0%B8%B6%E0%B8%A1%20%E0%B8%A1.2%20-%203.1%E0%B8%84%20%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%20%E0%B9%82%E0%B8%88%E0%B8%97%E0%B8%A2%E0%B9%8C%20%E0%B9%80%E0%B8%89%E0%B8%A5%E0%B8%A2%20%E0%B8%AB%E0%B8%B2%E0%B8%9B%E0%B8%A3%E0%B8%B4%E0%B8%A1%E0%B8%B2%E0%B8%95%E0%B8%A3%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%9B%E0%B8%A3%E0%B8%B4%E0%B8%8B%E0%B8%B6%E0%B8%A1.mp4?alt=media&token=da25b9db-4297-4faa-83ba-cb65fc448787",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_3%2FVideo_3%2F%E0%B8%AA%E0%B8%B9%E0%B8%95%E0%B8%A3%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%AB%E0%B8%B2%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%80%E0%B9%80%E0%B8%A5%E0%B8%B0%E0%B8%9B%E0%B8%A3%E0%B8%B4%E0%B8%A1%E0%B8%B2%E0%B8%95%E0%B8%A3.pdf?alt=media&token=07d86434-e350-4be6-bf7e-7c7860796623",
					Exercises: &[]models.Exercise{
						{
							Question: "1.จงหาพื้นที่ฐานของรูปสี่เหลี่ยมมุมฉาก โดยมีความกว้าง 3.5 เเละ ความยาว 4",
							Choices: &[]models.Choice{
								{
									Title:   "11",
									Correct: false,
								},
								{
									Title:   "12",
									Correct: false,
								},
								{
									Title:   "13",
									Correct: false,
								},
								{
									Title:   "14",
									Correct: true,
								},
							},
						},
						{
							Question: "2.จงหาปริมาตรของปริซึมสี่เหลี่ยมมุมฉาก โดยมีพื้นฐาน 14 เเละ ความสูงเป็น 6",
							Choices: &[]models.Choice{
								{
									Title:   "82",
									Correct: false,
								},
								{
									Title:   "83",
									Correct: false,
								},
								{
									Title:   "84",
									Correct: true,
								},
								{
									Title:   "85",
									Correct: false,
								},
							},
						},
						{
							Question: "3.จงหาปริมาตรของปริซึมสี่เหลี่ยมคางหมู โดยมีพื้นฐาน 128 เเละ ความสูงเป็น 12",
							Choices: &[]models.Choice{
								{
									Title:   "1536",
									Correct: true,
								},
								{
									Title:   "1230",
									Correct: false,
								},
								{
									Title:   "1150",
									Correct: false,
								},
								{
									Title:   "1520",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   5,
			Course_name: "คณิตศาสตร์ ชั้น ม.3 เรื่อง พื้นที่ และ ปริมาตร",
			Picture:     "https://altv-cms-images.s3.amazonaws.com/content/image_1621574007157380482.jpg",
			Description: "“พื้นที่ เเละ ปริมาตร” คืออะไร คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ พื้นที่ และ ปริมาตร ต่าง ๆ เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_4
			SubjectID: 4,
			Videos: &[]models.Video{
				{
					Video_name:  "เซต คืออะไร แบบแจกแจงสมาชิก ทำอย่างไร",
					Picture:     "https://i.ytimg.com/vi/5ZegR_M8tJs/maxresdefault.jpg",
					Description: "คณิตศาสตร์ ชั้น ม.4 เรื่อง เซต | เซต คืออะไร แบบแจกแจงสมาชิก ทำอย่างไร แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนคณิตศาสตร์ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_4%2FVideo_1%2F(1)%20%E0%B9%80%E0%B8%8B%E0%B8%95%20%E0%B8%A1.4%20-%201.1%E0%B8%81%20%E0%B9%80%E0%B8%8B%E0%B8%95%20%E0%B8%84%E0%B8%B7%E0%B8%AD%E0%B8%AD%E0%B8%B0%E0%B9%84%E0%B8%A3%20%E0%B9%81%E0%B8%9A%E0%B8%9A%E0%B9%81%E0%B8%88%E0%B8%81%E0%B9%81%E0%B8%88%E0%B8%87%E0%B8%AA%E0%B8%A1%E0%B8%B2%E0%B8%8A%E0%B8%B4%E0%B8%81%20%E0%B8%97%E0%B8%B3%E0%B8%AD%E0%B8%A2%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B9%84%E0%B8%A3%20%E0%B9%83%E0%B8%99%207%20%E0%B8%99%E0%B8%B2%E0%B8%97%E0%B8%B5%20(Step1_7).mp4?alt=media&token=1cbb3945-4772-440e-8614-540adfb83fe8",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_4%2FVideo_1%2F%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%20%E0%B9%80%E0%B8%8B%E0%B8%84%20%E0%B8%A1.4.pdf?alt=media&token=51f769af-0ad1-488f-85a0-93d194f7ece4",
					Exercises: &[]models.Exercise{
						{
							Question: "1.เซตต่อไปนี้ข้อใดเป็นเซตของสระในภาษาอังกฤษในรูปเเบบเเจกเเจงสมาชิก",
							Choices: &[]models.Choice{
								{
									Title:   "{1,2,3,4,5}",
									Correct: false,
								},
								{
									Title:   "[1,2,3,4,5]",
									Correct: false,
								},
								{
									Title:   "{a,e,i,o,u}",
									Correct: true,
								},
								{
									Title:   "[a,e,i,o,u]",
									Correct: false,
								},
							},
						},
						{
							Question: "2.จงเขียน {x|x เป็นจำนวนเต็มบวกที่มากกว่า 3 เเละน้อยกว่า 10} ในรูปเเบบการเเจกเเจงสมาชิก",
							Choices: &[]models.Choice{
								{
									Title:   "[4,5,6,7,8,9]",
									Correct: false,
								},
								{
									Title:   "{4,5,6,7,8,9}",
									Correct: true,
								},
								{
									Title:   "{3,4,5,6,7,8,9,10}",
									Correct: false,
								},
								{
									Title:   "[3,4,5,6,7,8,9,10]",
									Correct: false,
								},
							},
						},
						{
							Question: "3.จงเขียน {x|x เป็นจำนวนเต็มที่มากกว่า -3 เเละน้อยกว่า 3} ในรูปเเบบการเเจกเเจงสมาชิก",
							Choices: &[]models.Choice{
								{
									Title:   "{-2,-1,0,1,2}",
									Correct: true,
								},
								{
									Title:   "[-2,-1,0,1,2]",
									Correct: false,
								},
								{
									Title:   "{-3,-2,-1,0,1,2,3}",
									Correct: false,
								},
								{
									Title:   "[-3,-2,-1,0,1,2,3]",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "วิธีเขียน เซต แบบบอกเงื่อนไข ",
					Picture:     "https://i.ytimg.com/vi/5ZegR_M8tJs/maxresdefault.jpg",
					Description: "คณิตศาสตร์ ชั้น ม.4 เรื่อง เซต | วิธีเขียน เซต แบบบอกเงื่อนไข แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนคณิตศาสตร์ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_4%2FVideo_2%2F(2)%20%E0%B9%80%E0%B8%8B%E0%B8%95%20%E0%B8%A1.4%20-%201.1%E0%B8%81%20%E0%B8%A7%E0%B8%B4%E0%B8%98%E0%B8%B5%E0%B9%80%E0%B8%82%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B9%80%E0%B8%8B%E0%B8%95%20%E0%B9%81%E0%B8%9A%E0%B8%9A%E0%B8%9A%E0%B8%AD%E0%B8%81%E0%B9%80%E0%B8%87%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%99%E0%B9%84%E0%B8%82%20%E0%B9%83%E0%B8%99%205%20%E0%B8%99%E0%B8%B2%E0%B8%97%E0%B8%B5%20(Step2_7).mp4?alt=media&token=dca63e24-8f4f-4c5c-9b7d-ac7497504098",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_4%2FVideo_2%2F%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%20%E0%B9%80%E0%B8%8B%E0%B8%84%20%E0%B8%A1.4.pdf?alt=media&token=1122685e-ad29-46d8-a254-84ab67bfc0b7",
					Exercises: &[]models.Exercise{
						{
							Question: "1.จงเขียน {1,3,5,7,9} เเบบบอกเงื่อนไขสมาชิก",
							Choices: &[]models.Choice{
								{
									Title:   "{x|x เป็นจำนวนคี่ที่มากกว่า 1 เเละน้อยกว่า 9}",
									Correct: false,
								},
								{
									Title:   "{x|x เป็นจำนวนคู่ที่น้อยกว่า 10}",
									Correct: false,
								},
								{
									Title:   "{x|x เป็นจำนวนเต็มบวกที่มากกว่า 1 เเละน้อยกว่า 9}",
									Correct: false,
								},
								{
									Title:   "{x|x เป็นจำนวนคี่ที่น้อยกว่า 10}",
									Correct: true,
								},
							},
						},
						{
							Question: "2.จงเขียน {...,-2,-1,0,1,2,...} เเบบบอกเงื่อนไขสมาชิก",
							Choices: &[]models.Choice{
								{
									Title:   "{x|x เป็นจำนวนคี่}",
									Correct: false,
								},
								{
									Title:   "{x|x เป็นจำนวนเต็ม}",
									Correct: true,
								},
								{
									Title:   "{x|x เป็นจำนวนคู่}",
									Correct: false,
								},
								{
									Title:   "{x|x เป็นจำนวนเต็มลบ}",
									Correct: false,
								},
							},
						},
						{
							Question: "3.จงเขียน {1,4,9,16,25,36,...} เเบบบอกเงื่อนไขสมาชิก",
							Choices: &[]models.Choice{
								{
									Title:   "{x|x = n กำลังสอง เมื่อ n เป็นจำนวนนับ}",
									Correct: true,
								},
								{
									Title:   "{x|x เป็นจำนวนคี่}",
									Correct: false,
								},
								{
									Title:   "{x|x เป็นจำนวนคู่}",
									Correct: false,
								},
								{
									Title:   "{x|x เป็นจำนวนเต็มลบ}",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "วิธีหา สับเซต ทั้งหมดของเซตต่อไปนี้",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/5elJcMDdsMw/maxresdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอนวิธีหา สับเซต แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_4%2FVideo_3%2F(3)%20%E0%B9%80%E0%B8%8B%E0%B8%95%20%E0%B8%A1.4%20-%201.1%E0%B8%82%20%E0%B8%A7%E0%B8%B4%E0%B8%98%E0%B8%B5%E0%B8%AB%E0%B8%B2%20%E0%B8%AA%E0%B8%B1%E0%B8%9A%E0%B9%80%E0%B8%8B%E0%B8%95%20%E0%B8%97%E0%B8%B1%E0%B9%89%E0%B8%87%E0%B8%AB%E0%B8%A1%E0%B8%94%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B9%80%E0%B8%8B%E0%B8%95%E0%B8%95%E0%B9%88%E0%B8%AD%E0%B9%84%E0%B8%9B%E0%B8%99%E0%B8%B5%E0%B9%89%20%E0%B9%83%E0%B8%99%206%20%E0%B8%99%E0%B8%B2%E0%B8%97%E0%B8%B5%20(Step3_7).mp4?alt=media&token=c331a48a-2788-47a0-a24f-bd11f241207a",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_4%2FVideo_3%2F%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%20%E0%B9%80%E0%B8%8B%E0%B8%84%20%E0%B8%A1.4.pdf?alt=media&token=7649b4da-01ef-4fc7-afc5-a91544aaedfd",
					Exercises: &[]models.Exercise{
						{
							Question: "1.จงหาสับเซตทั้งหมดของ {1}",
							Choices: &[]models.Choice{
								{
									Title:   "{1}",
									Correct: false,
								},
								{
									Title:   "{}",
									Correct: false,
								},
								{
									Title:   "{Ø}",
									Correct: false,
								},
								{
									Title:   "{1}{Ø}",
									Correct: true,
								},
							},
						},
						{
							Question: "2.จงหาสับเซตทั้งหมดของ {1,2}",
							Choices: &[]models.Choice{
								{
									Title:   "{Ø}",
									Correct: false,
								},
								{
									Title:   "{1}{2}",
									Correct: false,
								},
								{
									Title:   "{1}{2}{1,2}{Ø}",
									Correct: true,
								},
								{
									Title:   "{1}{2}{1,2}",
									Correct: false,
								},
							},
						},
						{
							Question: "3.จงหาสับเซตทั้งหมดของ {a,b,c}",
							Choices: &[]models.Choice{
								{
									Title:   "{a}{b}{c}{a,b}{a,c}{b,c}{a,b,c}{Ø}",
									Correct: true,
								},
								{
									Title:   "{a}{b}{c}",
									Correct: false,
								},
								{
									Title:   "{a}{b}{c}{a,b}{a,c}{b,c}{a,b,c}",
									Correct: false,
								},
								{
									Title:   "{Ø}",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   5,
			Course_name: "คณิตศาสตร์ ชั้น ม.3 เรื่อง เซต",
			Picture:     "https://i.ytimg.com/vi/5ZegR_M8tJs/maxresdefault.jpg",
			Description: "“เซต” คืออะไร คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ เซต เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_5
			SubjectID: 5,
			Videos: &[]models.Video{
				{
					Video_name:  "ลำดับ - สรุป ลำดับเลขคณิต และ ลำดับเรขาคณิต",
					Picture:     "https://i.ytimg.com/vi/InWJsdKh7Fs/maxresdefault.jpg",
					Description: "คณิตศาสตร์ ชั้น ม.5 เรื่อง ลำดับเลขคณิต และ ลำดับเรขาคณิต | ลำดับเลขคณิต และ ลำดับเรขาคณิต แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนคณิตศาสตร์ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_5%2FVideo_1%2F(1)%20%E0%B8%A5%E0%B8%B3%E0%B8%94%E0%B8%B1%E0%B8%9A%20-%20%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%20%E0%B8%A5%E0%B8%B3%E0%B8%94%E0%B8%B1%E0%B8%9A%E0%B9%80%E0%B8%A5%E0%B8%82%E0%B8%84%E0%B8%93%E0%B8%B4%E0%B8%95%20%E0%B9%81%E0%B8%A5%E0%B8%B0%20%E0%B8%A5%E0%B8%B3%E0%B8%94%E0%B8%B1%E0%B8%9A%E0%B9%80%E0%B8%A3%E0%B8%82%E0%B8%B2%E0%B8%84%E0%B8%93%E0%B8%B4%E0%B8%95%20%E0%B8%A1.5.mp4?alt=media&token=7d4240ba-4972-42be-a315-0fb182aabad5",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_5%2FVideo_1%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B8%A5%E0%B8%B3%E0%B8%94%E0%B8%B1%E0%B8%9A%E0%B9%80%E0%B8%A5%E0%B8%82%E0%B8%B2%E0%B8%84%E0%B8%93%E0%B8%B4%E0%B8%95%20%E0%B9%80%E0%B9%80%E0%B8%A5%E0%B8%B0%E0%B8%AD%E0%B8%99%E0%B8%B8%E0%B8%81%E0%B8%A3%E0%B8%A1.pdf?alt=media&token=336588f6-402e-4c80-bff6-d656c87f3a6c",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดคือความหมายของลำดับจำกัด",
							Choices: &[]models.Choice{
								{
									Title:   "ลำดับที่ไม่ทราบพจน์สุดท้าย (ไม่สามารถนับจำนวนพจน์ได้)",
									Correct: false,
								},
								{
									Title:   "ลำดับที่ไม่ทราบพจน์สุดท้าย (สามารถนับจำนวนพจน์ได้)",
									Correct: false,
								},
								{
									Title:   "ลำดับที่ทราบพจน์สุดท้าย (สามารถนับจำนวนพจน์ได้)",
									Correct: true,
								},
								{
									Title:   "ลำดับที่ทราบพจน์สุดท้าย (ไม่สามารถจำนวนพจน์ไม่ได้)",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดคือความหมายของลำดับอนันต์",
							Choices: &[]models.Choice{
								{
									Title:   "ลำดับที่ไม่ทราบพจน์สุดท้าย (สามารถนับจำนวนพจน์ได้)",
									Correct: false,
								},
								{
									Title:   "ลำดับที่ไม่ทราบพจน์สุดท้าย (ไม่สามารถนับจำนวนพจน์ได้)",
									Correct: true,
								},
								{
									Title:   "ลำดับที่ทราบพจน์สุดท้าย (สามารถนับจำนวนพจน์ได้)",
									Correct: false,
								},
								{
									Title:   "ลำดับที่ทราบพจน์สุดท้าย (ไม่สามารถนับจำนวนพจน์ได้)",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "ซิกม่า คืออะไร สูตร ใช้ยังไง ",
					Picture:     "https://i.ytimg.com/vi/5ZegR_M8tJs/maxresdefault.jpg",
					Description: "คณิตศาสตร์ ชั้น ม.5 เรื่อง ลำดับเลขคณิต | ซิกม่า คืออะไร สูตร ใช้ยังไง แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนคณิตศาสตร์ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_5%2FVideo_2%2F(2)%20%E0%B8%A5%E0%B8%B3%E0%B8%94%E0%B8%B1%E0%B8%9A%E0%B9%81%E0%B8%A5%E0%B8%B0%E0%B8%AD%E0%B8%99%E0%B8%B8%E0%B8%81%E0%B8%A3%E0%B8%A1%20-%20%E0%B8%8B%E0%B8%B4%E0%B8%81%E0%B8%A1%E0%B9%88%E0%B8%B2%20%E0%B8%84%E0%B8%B7%E0%B8%AD%E0%B8%AD%E0%B8%B0%E0%B9%84%E0%B8%A3%20%E0%B8%AA%E0%B8%B9%E0%B8%95%E0%B8%A3%20%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%A2%E0%B8%B1%E0%B8%87%E0%B9%84%E0%B8%87%20%E0%B8%AA%E0%B8%B3%E0%B8%AB%E0%B8%A3%E0%B8%B1%E0%B8%9A%E0%B8%99%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B9%86%20%E0%B8%A1.5%20%E0%B9%81%E0%B8%A5%E0%B8%B0%20%E0%B8%A1.6.mp4?alt=media&token=16193183-2067-44b3-bfad-73360afb3a78",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_5%2FVideo_2%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B8%A5%E0%B8%B3%E0%B8%94%E0%B8%B1%E0%B8%9A%E0%B9%80%E0%B8%A5%E0%B8%82%E0%B8%B2%E0%B8%84%E0%B8%93%E0%B8%B4%E0%B8%95%20%E0%B9%80%E0%B9%80%E0%B8%A5%E0%B8%B0%E0%B8%AD%E0%B8%99%E0%B8%B8%E0%B8%81%E0%B8%A3%E0%B8%A1.pdf?alt=media&token=5e66186c-e70b-4dd6-baa2-1313bd8b35c8",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ซัมเมชั่น หรือ ซิกม่า คืออะไร ",
							Choices: &[]models.Choice{
								{
									Title:   "เป็นเครื่องหมายที่ใช้ในการหาผลหารของอนุกรมหรือลำดับที่มีรูปทั่วไป",
									Correct: false,
								},
								{
									Title:   "เป็นเครื่องหมายที่ใช้ในการหาผลคูณของอนุกรมหรือลำดับที่มีรูปทั่วไป",
									Correct: false,
								},
								{
									Title:   "เป็นเครื่องหมายที่ใช้ในการหาผลลบของอนุกรมหรือลำดับที่มีรูปทั่วไป",
									Correct: false,
								},
								{
									Title:   "เป็นเครื่องหมายที่ใช้ในการหาผลรวมของอนุกรมหรือลำดับที่มีรูปทั่วไปเป็นพหุนามกำลังไม่เกินสาม",
									Correct: true,
								},
							},
						},
						{
							Question: "2.ข้อใดไม่ใช่สมบัติของซิกม่า",
							Choices: &[]models.Choice{
								{
									Title:   "เมื่อใดก็ตามที่มีซิกม่าเเล้วมี c เป็นค่าคงที่ มีค่าเท่ากับ c = nc",
									Correct: false,
								},
								{
									Title:   "เมื่อใดก็ตามที่มีซิกม่าเเล้วมี c เป็นค่าคงที่ มีค่าเท่ากับ c = n+c",
									Correct: true,
								},
								{
									Title:   "เมื่อใดก็ตามที่มีซิกม่าเเล้วมี c เป็นค่าคงที่คูณกับ ai สามารถทำการดึง c ออกมาเเละใส่ซิกม่าไว้ภายในได้",
									Correct: false,
								},
								{
									Title:   "เมื่อใดก็ตามที่มีซิกม่าของผลบวกหรือผลลบของค่าใดก็ตามที่ไม่ใช่ค่าคงที่ สามารถเเยกซิกม่าออกได้",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "6 สูตรต้องรู้ ! อนุกรมเลขคณิต อนุกรมเรขาคณิต",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/BO7b9oasVNg/maxresdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอน 6 สูตรต้องรู้ของ อนุกรมเลขคณิตเเละอนุกรมเรขาคณิต แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_5%2FVideo_3%2F(3)%20%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%E0%B8%AA%E0%B8%B9%E0%B8%95%E0%B8%A3%206%20%E0%B8%AA%E0%B8%B9%E0%B8%95%E0%B8%A3%E0%B8%95%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%A3%E0%B8%B9%E0%B9%89%20!%20%E0%B8%AD%E0%B8%99%E0%B8%B8%E0%B8%81%E0%B8%A3%E0%B8%A1%E0%B9%80%E0%B8%A5%E0%B8%82%E0%B8%84%E0%B8%93%E0%B8%B4%E0%B8%95%20%E0%B8%AD%E0%B8%99%E0%B8%B8%E0%B8%81%E0%B8%A3%E0%B8%A1%E0%B9%80%E0%B8%A3%E0%B8%82%E0%B8%B2%E0%B8%84%E0%B8%93%E0%B8%B4%E0%B8%95%20_%20%E0%B8%A1.5.mp4?alt=media&token=92ebe87f-df48-46dc-807c-47a3a10b7a99",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_5%2FVideo_3%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B8%A5%E0%B8%B3%E0%B8%94%E0%B8%B1%E0%B8%9A%E0%B9%80%E0%B8%A5%E0%B8%82%E0%B8%B2%E0%B8%84%E0%B8%93%E0%B8%B4%E0%B8%95%20%E0%B9%80%E0%B9%80%E0%B8%A5%E0%B8%B0%E0%B8%AD%E0%B8%99%E0%B8%B8%E0%B8%81%E0%B8%A3%E0%B8%A1.pdf?alt=media&token=ef6400e4-13d2-486e-9bf9-3a02d0892f72",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดเป็นสูตรอนุกรมเลขคณิต",
							Choices: &[]models.Choice{
								{
									Title:   "sn = n*(a1 - an)",
									Correct: false,
								},
								{
									Title:   "sn = n/2(a1 - an)",
									Correct: false,
								},
								{
									Title:   "sn = n*(2a1 + (n+1)d)",
									Correct: false,
								},
								{
									Title:   "sn = n/2(2a1 + (n-1)d)",
									Correct: true,
								},
							},
						},
						{
							Question: "2.อนุกรมเรขาคณิตเกิดจากอะไร",
							Choices: &[]models.Choice{
								{
									Title:   "เกิดจากการเอาเลขคณิตมาบวกกัน",
									Correct: false,
								},
								{
									Title:   "เกิดจากการเอาเลขคณิตมาลบกัน",
									Correct: false,
								},
								{
									Title:   "เกิดจากการเอาลำดับเลขาคณิตมาบวกกัน",
									Correct: true,
								},
								{
									Title:   "เกิดจากการเอาลำดับเลขาคณิตมาลบกัน",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   5,
			Course_name: "คณิตศาสตร์ ชั้น ม.5 เรื่อง สรุป ลำดับเลขคณิต และ ลำดับเรขาคณิต",
			Picture:     "https://i.ytimg.com/vi/InWJsdKh7Fs/maxresdefault.jpg",
			Description: "“ลำดับเลขคณิต และ ลำดับเรขาคณิต” คืออะไร คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ ลำดับเลขคณิต และ ลำดับเรขาคณิต เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_6
			SubjectID: 6,
			Videos: &[]models.Video{
				{
					Video_name:  "ความหมายของสถิติ",
					Picture:     "https://i.ytimg.com/vi/InWJsdKh7Fs/maxresdefault.jpg",
					Description: "คณิตศาสตร์ ชั้น ม.6 เรื่อง สถิติเบื้องต้น | ความหมายของสถิติ แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนคณิตศาสตร์ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_6%2FVideo_1%2F(1)%20%E0%B8%AA%E0%B8%96%E0%B8%B4%E0%B8%95%E0%B8%B4%E0%B9%80%E0%B8%9A%E0%B8%B7%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%95%E0%B9%89%E0%B8%99%20%E0%B8%A1.6%20EP.1_10%20%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%AB%E0%B8%A1%E0%B8%B2%E0%B8%A2%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%AA%E0%B8%96%E0%B8%B4%E0%B8%95%E0%B8%B4.mp4?alt=media&token=3ed59f97-1477-490e-96d9-38196a85a621",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_6%2FVideo_1%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B8%AA%E0%B8%96%E0%B8%B7%E0%B8%95%E0%B8%B4%E0%B9%80%E0%B8%9A%E0%B8%B7%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%95%E0%B9%89%E0%B8%99.pdf?alt=media&token=e0c2fd0f-3cdf-4b87-873f-8aa8ac15d5d3",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อมูลประชากรคืออะไร",
							Choices: &[]models.Choice{
								{
									Title:   "ข้อมูลส่วนหนึ่งที่สุ่มจากข้อมูลทั้งหมดเพื่อมาเพื่อศึกษา",
									Correct: false,
								},
								{
									Title:   "ข้อมูลทั้งหมดที่เราจะศึกษา",
									Correct: true,
								},
								{
									Title:   "ค่าของข้อมูลที่สรุปรวบยอดจากข้อมูลทั้งหมด",
									Correct: false,
								},
								{
									Title:   "ถูกทุกข้อ",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อมูลตัวอย่างคืออะไร",
							Choices: &[]models.Choice{
								{
									Title:   "ข้อมูลทั้งหมดที่เราจะศึกษา",
									Correct: false,
								},
								{
									Title:   "ข้อมูลส่วนหนึ่งของข้อมูลประชากรที่เราสุ่มมาเพื่อทำการศึกษา",
									Correct: true,
								},
								{
									Title:   "ค่าของข้อมูลที่สรุปรวบยอดจากข้อมูลประชากร",
									Correct: false,
								},
								{
									Title:   "ถูกทุกข้อ",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ค่าสถิติคืออะไร",
							Choices: &[]models.Choice{
								{
									Title:   "ข้อมูลทั้งหมดที่เราจะศึกษา",
									Correct: false,
								},
								{
									Title:   "ข้อมูลส่วนหนึ่งของข้อมูลประชากรที่เราสุ่มมาเพื่อทำการศึกษา",
									Correct: false,
								},
								{
									Title:   "ค่าสถิติของข้อมูลตัวอย่างหรือค่าของข้อมูลที่สรุปรวบยอดมาจากข้อมูลประชากร",
									Correct: true,
								},
								{
									Title:   "ถูกทุกข้อ",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "การแจกแจงความถี่#1",
					Picture:     "https://i.ytimg.com/vi/OUvKWJ1tny4/mqdefault.jpg",
					Description: "คณิตศาสตร์ ชั้น ม.6 เรื่อง สถิติเบื้องต้น | การแจกแจงความถี่ แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนคณิตศาสตร์ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_6%2FVideo_2%2F(2)%20%E0%B8%AA%E0%B8%96%E0%B8%B4%E0%B8%95%E0%B8%B4%E0%B9%80%E0%B8%9A%E0%B8%B7%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%95%E0%B9%89%E0%B8%99%20%E0%B8%A1.6%20EP.2_10%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%81%E0%B8%88%E0%B8%81%E0%B9%81%E0%B8%88%E0%B8%87%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%96%E0%B8%B5%E0%B9%88%231.mp4?alt=media&token=037301b1-65d5-4cb0-8bbe-8d42e66e8efe",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_6%2FVideo_2%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B8%AA%E0%B8%96%E0%B8%B7%E0%B8%95%E0%B8%B4%E0%B9%80%E0%B8%9A%E0%B8%B7%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%95%E0%B9%89%E0%B8%99.pdf?alt=media&token=9bfb3add-5d27-41c8-982b-8d6dede10eb5",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดไม่ใช่ประเภทการเเจกเเจงความถี่",
							Choices: &[]models.Choice{
								{
									Title:   "เเจกเเจงตามค่าข้อมูล",
									Correct: false,
								},
								{
									Title:   "เเจกเเจงเป็นช่วงคะเเนน",
									Correct: false,
								},
								{
									Title:   "เเจกเเจงเป็นช่วงอันตรภาคชั้น",
									Correct: false,
								},
								{
									Title:   "เเจกเเจงจากการคาดเดา",
									Correct: true,
								},
							},
						},
						{
							Question: "2.การเเจกเเจงความถี่ตามค่าข้อมูล เหมาะกับเเบบใด",
							Choices: &[]models.Choice{
								{
									Title:   "การที่มีข้อมูลจำนวนมาก",
									Correct: false,
								},
								{
									Title:   "การที่มีข้อมูลจำนวนไม่มากเเละมีข้อมูลที่ซ้ำกันอยู่",
									Correct: true,
								},
								{
									Title:   "การที่มีข้อมูลจำนวนมากเเละมีข้อมูลซ้ำกันอยู่น้อย",
									Correct: false,
								},
								{
									Title:   "ผิดทุกข้อ",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  " การแจกแจงความถี่#2",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/X_mX9Npqh7I/mqdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอน การเเจกเเจงความถี่ของสถิติ แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_6%2FVideo_3%2F(3)%20%E0%B8%AA%E0%B8%96%E0%B8%B4%E0%B8%95%E0%B8%B4%E0%B9%80%E0%B8%9A%E0%B8%B7%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%95%E0%B9%89%E0%B8%99%20%E0%B8%A1.6%20EP.3_10%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%81%E0%B8%88%E0%B8%81%E0%B9%81%E0%B8%88%E0%B8%87%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%96%E0%B8%B5%E0%B9%88%232.mp4?alt=media&token=f9987a8f-e1aa-4891-b1cf-71b64500c0b8",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_6%2FVideo_3%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B8%AA%E0%B8%96%E0%B8%B7%E0%B8%95%E0%B8%B4%E0%B9%80%E0%B8%9A%E0%B8%B7%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%95%E0%B9%89%E0%B8%99.pdf?alt=media&token=02055cf6-05eb-4605-9ac7-a30e3ed1003e",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดคือความหมายของความถี่สัมพัทธ์",
							Choices: &[]models.Choice{
								{
									Title:   "อัตราส่วนระหว่างความถี่ของค่านั้นหรือของอันตรภาคชั้นนั้นกับผลคูณของ ความถี่ทั้งหมด",
									Correct: false,
								},
								{
									Title:   "อัตราส่วนระหว่างความถี่ของค่านั้นหรือของอันตรภาคชั้นนั้นกับผลหารของ ความถี่ทั้งหมด",
									Correct: false,
								},
								{
									Title:   "อัตราส่วนระหว่างความถี่ของค่านั้นหรือของอันตรภาคชั้นนั้นกับผลลบของ ความถี่ทั้งหมด",
									Correct: false,
								},
								{
									Title:   "อัตราส่วนระหว่างความถี่ของค่านั้นหรือของอันตรภาคชั้นนั้นกับผลรวมของ ความถี่ทั้งหมด",
									Correct: true,
								},
							},
						},
						{
							Question: "2.ข้อใดคือความหมายของความถี่สะสมสัมพัทธ์",
							Choices: &[]models.Choice{
								{
									Title:   "อัตราส่วนระหว่างความถี่สะสมของอันตรภาคชั้นนั้นกับผลคูณของความถี่ทั้งหมด",
									Correct: false,
								},
								{
									Title:   "อัตราส่วนระหว่างความถี่สะสมของอันตรภาคชั้นนั้นกับผลหารของความถี่ทั้งหมด",
									Correct: false,
								},
								{
									Title:   "อัตราส่วนระหว่างความถี่สะสมของอันตรภาคชั้นนั้นกับผลรวมของความถี่ทั้งหมด",
									Correct: true,
								},
								{
									Title:   "อัตราส่วนระหว่างความถี่สะสมของอันตรภาคชั้นนั้นกับผลลบของความถี่ทั้งหมด",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   1,
			Course_name: "คณิตศาสตร์ ชั้น ม.6 เรื่อง สถิติเบื้องต้น",
			Picture:     "https://i.ytimg.com/vi/iR7yM4uYynA/hqdefault.jpg",
			Description: "“สถิติ” คืออะไร คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ สถิติเบื้องต้น เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_7
			SubjectID: 20,
			Videos: &[]models.Video{
				{
					Video_name:  "การใช้รูปประโยคคำสั่ง คำขอร้อง คำแนะนำ+การใช้ Can/ Could/ Should",
					Picture:     "https://i.ytimg.com/vi/-9j-BHYh-Nw/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ ชั้น ม.1 เรื่อง การใช้รูปประโยคคำสั่ง คำขอร้อง คำแนะนำ+การใช้ Can/ Could/ Should | แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาอังกฤษ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_7%2FVideo_1%2F(1)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.1%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%A3%E0%B8%B9%E0%B8%9B%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B9%82%E0%B8%A2%E0%B8%84%E0%B8%84%E0%B8%B3%E0%B8%AA%E0%B8%B1%E0%B9%88%E0%B8%87%20%E0%B8%84%E0%B8%B3%E0%B8%82%E0%B8%AD%E0%B8%A3%E0%B9%89%E0%B8%AD%E0%B8%87%20%E0%B8%84%E0%B8%B3%E0%B9%81%E0%B8%99%E0%B8%B0%E0%B8%99%E0%B8%B3%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20Can_%20Could_%20Should.mp4?alt=media&token=e34da608-95b9-404b-980f-30357e0c96d1",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_7%2FVideo_1%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%A1.1%20%E0%B8%9B%E0%B8%B9%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%90%E0%B8%B2%E0%B8%99%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9.pdf?alt=media&token=eb5a8eb8-c50d-4af0-b243-9cb4ed192d04",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดไม่ใช่ประโยคคำสั่ง (Imperative)",
							Choices: &[]models.Choice{
								{
									Title:   "Sit down.",
									Correct: false,
								},
								{
									Title:   "Be quiet, please.",
									Correct: true,
								},
								{
									Title:   "Run!",
									Correct: false,
								},
								{
									Title:   "Get out",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดไม่ใช่ประโยคขอร้อง (Request)",
							Choices: &[]models.Choice{
								{
									Title:   "Please wait here.",
									Correct: false,
								},
								{
									Title:   "Be quiet!",
									Correct: true,
								},
								{
									Title:   "Can you come to school this weekend?",
									Correct: false,
								},
								{
									Title:   "Could you buy me some chocolate, please?",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดเป็นประโยคเเนะนำ (Suggestion)",
							Choices: &[]models.Choice{
								{
									Title:   "Run!",
									Correct: false,
								},
								{
									Title:   "Please wait here.",
									Correct: false,
								},
								{
									Title:   "You should read some books.",
									Correct: true,
								},
								{
									Title:   "Can you come to school this weekend?",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "การแนะนำตนเอง + การใช้ like/love/enjoy + V. ing",
					Picture:     "https://i.ytimg.com/vi/SsJq-prcUMA/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ ชั้น ม.1 เรื่อง การแนะนำตนเอง + การใช้ like/love/enjoy + V. ing |  แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาอังกฤษ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_7%2FVideo_2%2F(2)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.1%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%81%E0%B8%99%E0%B8%B0%E0%B8%99%E0%B8%B3%E0%B8%95%E0%B8%99%E0%B9%80%E0%B8%AD%E0%B8%87%20%20%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20like_love_enjoy%20%20%20V.%20ing.mp4?alt=media&token=f87535a7-9486-4301-a058-3b38fa0bee86",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_7%2FVideo_2%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%A1.1%20%E0%B8%9B%E0%B8%B9%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%90%E0%B8%B2%E0%B8%99%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9.pdf?alt=media&token=3fde2467-1068-447b-955b-11d3797f4483",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ประโยคใดเป็นการเเนะนำชื่อตัวเองให้กับผู้อื่น",
							Choices: &[]models.Choice{
								{
									Title:   "I love you",
									Correct: false,
								},
								{
									Title:   "I'm 21 year old",
									Correct: false,
								},
								{
									Title:   "I live in Bangkok",
									Correct: false,
								},
								{
									Title:   "My name is Natacha.",
									Correct: true,
								},
							},
						},
						{
							Question: "2.ประโยคใดเป็นการบอกอายุให้กับผู้อื่น",
							Choices: &[]models.Choice{
								{
									Title:   "I love you",
									Correct: false,
								},
								{
									Title:   "I'm 21 year old",
									Correct: true,
								},
								{
									Title:   "I live in Bangkok",
									Correct: false,
								},
								{
									Title:   "My name is Natacha.",
									Correct: false,
								},
							},
						},
						{
							Question: "3.ประโยคข้อใดผิด",
							Choices: &[]models.Choice{
								{
									Title:   "He likes playing tennis.",
									Correct: false,
								},
								{
									Title:   "I don't like make my bed.",
									Correct: true,
								},
								{
									Title:   "Do you love calculating?",
									Correct: false,
								},
								{
									Title:   "They enjoy watching this program.",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "การใช้ Signal words : First, Second, Then, Next etc",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/SgTvu1U3A40/maxresdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอน การใช้ Signal words : First, Second, Then, Next etc แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_7%2FVideo_3%2F(3)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.1%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20Signal%20words%20_%20First%2C%20Second%2C%20Then%2C%20Next%20etc..mp4?alt=media&token=64aacec8-c13b-498b-8876-71efe2aa7856",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_7%2FVideo_3%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%A1.1%20%E0%B8%9B%E0%B8%B9%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%90%E0%B8%B2%E0%B8%99%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9.pdf?alt=media&token=2e97c7e5-7125-41d7-b4e8-781d39ca58a6",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดไม่ใช่การเล่าเรื่องเป็นประโยคถัดไป",
							Choices: &[]models.Choice{
								{
									Title:   "To begin with",
									Correct: false,
								},
								{
									Title:   "then",
									Correct: false,
								},
								{
									Title:   "next",
									Correct: false,
								},
								{
									Title:   "Firstly",
									Correct: true,
								},
							},
						},
						{
							Question: "2.ข้อใดไม่ใช่การเล่าเรื่องเเบบเป็นลำดับชั้น",
							Choices: &[]models.Choice{
								{
									Title:   "First of all",
									Correct: false,
								},
								{
									Title:   "Secondly",
									Correct: false,
								},
								{
									Title:   "To begin with",
									Correct: true,
								},
								{
									Title:   "Finally",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   6,
			Course_name: "ภาษาอังกฤษ ชั้น ม.1 เรื่อง ปูพื้นฐานภาษาอังกฤษ",
			Picture:     "https://nockacademy.com/wp-content/uploads/2021/04/Suggesting-Profile.png",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ ภาษาอังกฤษพื้นฐานในชีวิตประจำวัน เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_8
			SubjectID: 21,
			Videos: &[]models.Video{
				{
					Video_name:  "การใช้ Was/ were ในการสร้างประโยคคำถามในรูปแบบต่าง ๆ",
					Picture:     "https://i.ytimg.com/vi/-9j-BHYh-Nw/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ ชั้น ม.2 เรื่อง การใช้ Was/ were ในการสร้างประโยคคำถามในรูปแบบต่าง ๆ | แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาอังกฤษ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_8%2FVideo_1%2F(1)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.2%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20Was_%20were%20%E0%B9%83%E0%B8%99%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%AA%E0%B8%A3%E0%B9%89%E0%B8%B2%E0%B8%87%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B9%82%E0%B8%A2%E0%B8%84%E0%B8%84%E0%B8%B3%E0%B8%96%E0%B8%B2%E0%B8%A1%E0%B9%83%E0%B8%99%E0%B8%A3%E0%B8%B9%E0%B8%9B%E0%B9%81%E0%B8%9A%E0%B8%9A%E0%B8%95%E0%B9%88%E0%B8%B2%E0%B8%87%20%E0%B9%86.mp4?alt=media&token=9c098abd-f674-4017-a462-bcb13abe867a",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_8%2FVideo_1%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%A1.2%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20was_were.pdf?alt=media&token=af19c6c9-d2f0-4c02-a281-914eaffb950b",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดใช้ was/were ถูกต้อง",
							Choices: &[]models.Choice{
								{
									Title:   "Was you angry at your friend yesterday?",
									Correct: false,
								},
								{
									Title:   "Were you busy last week?",
									Correct: true,
								},
								{
									Title:   "Were he nervous?",
									Correct: false,
								},
								{
									Title:   "ผิดทุกข้อ",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดผิด",
							Choices: &[]models.Choice{
								{
									Title:   "was he at a concert.",
									Correct: false,
								},
								{
									Title:   "was they busy angry at your friend yesterday?",
									Correct: true,
								},
								{
									Title:   "were they in a dorm?",
									Correct: false,
								},
								{
									Title:   "was christ on the phone min ago?",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "การใช้ Tenses : present simple/ present continuous",
					Picture:     "https://i.ytimg.com/vi/YMpclowjr2o/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ ชั้น ม.2 เรื่อง การใช้ Tenses : present simple/ present continuous |  แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาอังกฤษ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_8%2FVideo_2%2F(2)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.2%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20Tenses%20_%20present%20simple_%20present%20continuous.mp4?alt=media&token=1e242d13-c65b-4499-ac1b-8fbe05a39897",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_8%2FVideo_2%2F%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%A1.2%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20tenses%20present%20simple%20present%20continuous.pdf?alt=media&token=7f8bddc8-0064-46ed-8fa3-dae664c5ea26",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดผิด",
							Choices: &[]models.Choice{
								{
									Title:   "I read books every night.",
									Correct: false,
								},
								{
									Title:   "I am reading books now.",
									Correct: false,
								},
								{
									Title:   "I don't read books every night.",
									Correct: false,
								},
								{
									Title:   "I'm not read book now.",
									Correct: true,
								},
							},
						},
						{
							Question: "2.ข้อใดถูกต้อง",
							Choices: &[]models.Choice{
								{
									Title:   "I reading books every night.",
									Correct: false,
								},
								{
									Title:   "I'm not reading book now",
									Correct: true,
								},
								{
									Title:   "I don't reading books every night.",
									Correct: false,
								},
								{
									Title:   "I am read books now.",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "การใช้ Passive Voice ที่ใช้ในโครงสร้างประโยคง่ายๆ",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/OQrYoFu3twk/maxresdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอน การใช้ Passive Voice ที่ใช้ในโครงสร้างประโยคง่ายๆ แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_8%2FVideo_3%2F(3)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.2%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20Passive%20Voice%20%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B9%83%E0%B8%99%E0%B9%82%E0%B8%84%E0%B8%A3%E0%B8%87%E0%B8%AA%E0%B8%A3%E0%B9%89%E0%B8%B2%E0%B8%87%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B9%82%E0%B8%A2%E0%B8%84%E0%B8%87%E0%B9%88%E0%B8%B2%E0%B8%A2%E0%B9%86.mp4?alt=media&token=21da2674-ea16-4954-a4d8-8ec00ee3dfed",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_8%2FVideo_3%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%A1.2%20Passive%20Voice.pdf?alt=media&token=409b434e-8e79-446c-aed6-8e0ae588c132",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดผิดโครงสร้าง Passive Voice",
							Choices: &[]models.Choice{
								{
									Title:   "I write a sentence.",
									Correct: false,
								},
								{
									Title:   "Credits are given instead of cash if you decide to return your order.",
									Correct: false,
								},
								{
									Title:   "Cacaol from Africa is sold at a low price.",
									Correct: false,
								},
								{
									Title:   "Cacao from Africa sold at a low price.",
									Correct: true,
								},
							},
						},
						{
							Question: "2.ข้อใดถูกโครงสร้าง Passive Voice",
							Choices: &[]models.Choice{
								{
									Title:   "A sentence is write",
									Correct: false,
								},
								{
									Title:   "We cleaned a room.",
									Correct: false,
								},
								{
									Title:   "A room is cleaned by us.",
									Correct: true,
								},
								{
									Title:   "She doesn't written a newpaper.",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   6,
			Course_name: "ภาษาอังกฤษ ชั้น ม.2 เรื่อง พื้นฐานภาษาอังกฤษที่จำเป็น",
			Picture:     "https://i.ytimg.com/vi/ofHyaV_vNmg/maxresdefault.jpg",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ ภาษาอังกฤษพื้นฐานในชีวิตประจำวัน เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_9
			SubjectID: 22,
			Videos: &[]models.Video{
				{
					Video_name:  "การใช้ Present continuous",
					Picture:     "https://i.ytimg.com/vi/Pjbe9jnGxKA/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ ชั้น ม.3 เรื่อง การใช้ Present continuous | แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาอังกฤษ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_9%2FVideo_1%2F(1)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.3%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20Present%20continuous.mp4?alt=media&token=d5d112b0-5fde-4228-9e75-5ea111e87aca",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_9%2FVideo_1%2F%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%A1.3%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20Present%20continuous.pdf?alt=media&token=a0db894c-f963-43c1-8b3b-16ac75c978f3",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดผิดหลักโครงสร้าง Present Continuous",
							Choices: &[]models.Choice{
								{
									Title:   "I am teaching at the moment.",
									Correct: false,
								},
								{
									Title:   "They are not play football.",
									Correct: true,
								},
								{
									Title:   "We are studying English with teacher",
									Correct: false,
								},
								{
									Title:   "Is your mum cooking now?",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดถูกต้องหลักโครงสร้าง Present Continuous",
							Choices: &[]models.Choice{
								{
									Title:   "They are not play football.",
									Correct: false,
								},
								{
									Title:   "We are studying english with my teacher.",
									Correct: true,
								},
								{
									Title:   "I am teach at the moment.",
									Correct: false,
								},
								{
									Title:   "Is your mum cook now?",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "การใช้ Present perfect",
					Picture:     "https://i.ytimg.com/vi/rS0Kx-4kSoQ/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ ชั้น ม.3 เรื่อง การใช้ Tenses : Present perfect |  แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาอังกฤษ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_9%2FVideo_2%2F(2)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.3%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20Present%20perfect.mp4?alt=media&token=149534ca-6436-4d5e-b5b8-5390c76cb198",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_9%2FVideo_2%2F%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%A1.3%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20Present%20perfect.pdf?alt=media&token=77a2b3dc-7efa-4468-9c72-a3eefa885571",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดผิดหลักโครงสร้าง Present perfect",
							Choices: &[]models.Choice{
								{
									Title:   "She has lived in indonesia for 3 years.",
									Correct: false,
								},
								{
									Title:   "He has just left the airport.",
									Correct: false,
								},
								{
									Title:   "We have waited for the bus since 9 o'clock.",
									Correct: false,
								},
								{
									Title:   "She has work in the bank for five years.",
									Correct: true,
								},
							},
						},
						{
							Question: "2.ข้อใดถูกหลักโครงสร้าง Present perfect",
							Choices: &[]models.Choice{
								{
									Title:   "It has rain a lot this year.",
									Correct: false,
								},
								{
									Title:   "I have worked hard this week.",
									Correct: true,
								},
								{
									Title:   "We haven't see her today.",
									Correct: false,
								},
								{
									Title:   "She has work in the bank for five years.",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "การใช้ Past simple",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/J-SNIypNDaw/maxresdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอน การใช้ Past simple ที่ใช้ในโครงสร้างประโยคง่ายๆ แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_9%2FVideo_3%2F(3)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.3%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20Past%20simple.mp4?alt=media&token=4a0797ba-b488-4fbd-9b9f-77ebee000f21",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_9%2FVideo_3%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20past%20simple.pdf?alt=media&token=47382447-9a11-4e4c-92bf-b2adfba17012",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดผิดหลักโครงสร้าง Past simple",
							Choices: &[]models.Choice{
								{
									Title:   "My father died last year.",
									Correct: false,
								},
								{
									Title:   "He lived in Fiji in 1976.",
									Correct: false,
								},
								{
									Title:   "We crossed the channel yesterday.",
									Correct: false,
								},
								{
									Title:   "I live in Thailand until last year.",
									Correct: true,
								},
							},
						},
						{
							Question: "2.ข้อใดถูกหลักโครงสร้าง Past simple",
							Choices: &[]models.Choice{
								{
									Title:   "He is a Taiwanese singer. She lives from 1953 to 1995.",
									Correct: false,
								},
								{
									Title:   "The civil was in America start in 1861",
									Correct: false,
								},
								{
									Title:   "I went to school yesterday.",
									Correct: true,
								},
								{
									Title:   "It was warm, so I taking off my coat.",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   6,
			Course_name: "ภาษาอังกฤษ ชั้น ม.3 เรื่อง Tense",
			Picture:     "https://www.learnneo.in.th/wp-content/uploads/2021/03/12tenses.jpg",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ ภาษาอังกฤษพื้นฐานในเรื่อง tense เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_10
			SubjectID: 23,
			Videos: &[]models.Video{
				{
					Video_name:  "การใช้ Simple Present",
					Picture:     "https://i.ytimg.com/vi/7Gbw16OKOpY/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ ชั้น ม.4 เรื่อง การใช้ Simple Present | แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาอังกฤษ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_10%2FVideo_1%2F(1)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.4%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20Simple%20Present.mp4?alt=media&token=9858a02b-dda6-4290-8a73-d0484d020904",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_10%2FVideo_1%2F%E0%B8%96%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%A1.4%20Present-Simple.pdf?alt=media&token=99fc85e9-e04b-405b-93ff-54b4e916f3bf",
					Exercises: &[]models.Exercise{
						{
							Question: "1. Choose the correct answer. I ..... a bike after school every day.",
							Choices: &[]models.Choice{
								{
									Title:   "ride",
									Correct: true,
								},
								{
									Title:   "rides",
									Correct: false,
								},
								{
									Title:   "riding",
									Correct: false,
								},
								{
									Title:   "drive",
									Correct: false,
								},
							},
						},
						{
							Question: "2.Choose the correct answer. We ..... snowman in winter.",
							Choices: &[]models.Choice{
								{
									Title:   "making",
									Correct: false,
								},
								{
									Title:   "makes",
									Correct: false,
								},
								{
									Title:   "do",
									Correct: false,
								},
								{
									Title:   "make",
									Correct: true,
								},
							},
						},
						{
							Question: "3.Choose the correct answer. My sister ..... shopping at the weekends.",
							Choices: &[]models.Choice{
								{
									Title:   "make",
									Correct: false,
								},
								{
									Title:   "do",
									Correct: false,
								},
								{
									Title:   "does",
									Correct: true,
								},
								{
									Title:   "dos",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "การใช้ Present Progressive",
					Picture:     "https://i.ytimg.com/vi/vN3NGH2eV4g/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ ชั้น ม.4 เรื่อง การใช้ Tenses : Present Progressive |  แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาอังกฤษ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_10%2FVideo_2%2F(2)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.4%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20Present%20Progressive.mp4?alt=media&token=c1c88f2a-b3f1-40be-9207-24dfd9c02d9f",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_10%2FVideo_2%2F%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%A1.4%20Present%20Progressive.pdf?alt=media&token=b271a0fb-66be-4644-8961-d07ef4eb4f3b",
					Exercises: &[]models.Exercise{
						{
							Question: "1. Choose the correct answer. They are ..... now.",
							Choices: &[]models.Choice{
								{
									Title:   "drive",
									Correct: false,
								},
								{
									Title:   "drived",
									Correct: false,
								},
								{
									Title:   "drives",
									Correct: false,
								},
								{
									Title:   "driving",
									Correct: true,
								},
							},
						},
						{
							Question: "2.Choose the correct answer. What is she ..... at the moment?",
							Choices: &[]models.Choice{
								{
									Title:   "do",
									Correct: false,
								},
								{
									Title:   "doing",
									Correct: true,
								},
								{
									Title:   "does",
									Correct: false,
								},
								{
									Title:   "done",
									Correct: false,
								},
							},
						},
						{
							Question: "3.Choose the correct answer. She is ..... plants. ",
							Choices: &[]models.Choice{
								{
									Title:   "grow",
									Correct: false,
								},
								{
									Title:   "growing",
									Correct: true,
								},
								{
									Title:   "growed",
									Correct: false,
								},
								{
									Title:   "grows",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Video_name:  "การใช้ Present Perfect",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/DBkX-4lorGQ/maxresdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอน การใช้ Present Perfect ที่ใช้ในโครงสร้างประโยคง่ายๆ แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_10%2FVideo_3%2F(3)%20%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.4%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20Present%20Perfect.mp4?alt=media&token=156b9336-c00c-4b8f-9790-11cc8e550750",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_10%2FVideo_3%2F%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%A1.3%20%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%20Present%20perfect.pdf?alt=media&token=9d984d32-dc60-498c-8e06-6f9592e405ad",
					Exercises: &[]models.Exercise{
						{
							Question: "1.Choose the correct answer. She has ..... Italian since 2019. ",
							Choices: &[]models.Choice{
								{
									Title:   "study",
									Correct: false,
								},
								{
									Title:   "studied",
									Correct: true,
								},
								{
									Title:   "ถูกทั้งข้อ 1 เเละ 2",
									Correct: false,
								},
								{
									Title:   "ผิดทั้งข้อ 1 เเละ 2",
									Correct: false,
								},
							},
						},
						{
							Question: "2.Choose the correct answer. Haven't you ..... anything?",
							Choices: &[]models.Choice{
								{
									Title:   "eat",
									Correct: false,
								},
								{
									Title:   "eated",
									Correct: false,
								},
								{
									Title:   "eaten",
									Correct: true,
								},
								{
									Title:   "ผิดทุกข้อ",
									Correct: false,
								},
							},
						},
						{
							Question: "3.Choose the correct answer. He has ..... me about how to lose his weight.",
							Choices: &[]models.Choice{
								{
									Title:   "tell",
									Correct: false,
								},
								{
									Title:   "talk",
									Correct: false,
								},
								{
									Title:   "told",
									Correct: true,
								},
								{
									Title:   "talked",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   10,
			Course_name: "ภาษาอังกฤษ ชั้น ม.4 เรื่อง Tense พื้นฐาน",
			Picture:     "https://shorturl.asia/ZFCWJ",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ ภาษาอังกฤษพื้นฐานในเรื่อง tense เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
	}
	if err := DB.Model(&models.Course{}).Create(&course).Error; err != nil {
		panic(err)
	}

	fmt.Println("Course created")
}

func WipeData() {
	DB.Migrator().DropTable(&models.User{}, &models.UserTeacher{}, &models.Experience{}, &models.Comment{}, &models.Course{}, &models.History{}, &models.Post{}, &models.Review_Video{}, &models.Review_Tutor{}, &models.Reward_Info{}, &models.Reward_Item{}, &models.Subject{}, &models.Video{}, &models.Exercise{}, &models.Choice{})
}

func MigrateData() {
	DB.AutoMigrate(&models.User{}, &models.UserTeacher{}, &models.Experience{}, &models.Comment{}, &models.Course{}, &models.History{}, &models.Post{}, &models.Review_Video{}, &models.Review_Tutor{}, &models.Reward_Info{}, &models.Reward_Item{}, &models.Subject{}, &models.Video{}, &models.Exercise{}, &models.Choice{})
}

//Psychological_test: "https://testyourself.psychtests.com/bin/report?req=Mnw0MTc2fDEyMDA3NzM1fDF8MQ==",
