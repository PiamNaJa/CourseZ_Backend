package configs

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/PiamNaJa/CourseZ_Backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func TestChat() {
	var user = &models.User{
		User_id:  11,
		Email:    "p@mail.com",
		Fullname: "PP",
		Nickname: "PP",
		Role:     "Student",
		Password: "$2a$10$cm2Krm4dyA2Xu.qY705/EO96ZudVKv.YLpZDQMBAW2wWNEWOWQ6Ou",
		Picture:  "https://paradepets.com/.image/t_share/MTkxMzY1Nzg4NjczMzIwNTQ2/cutest-dog-breeds-jpg.jpg",
		Point:    999,
	}
	var user2 = &models.User{
		User_id:  12,
		Email:    "aa@mail.com",
		Fullname: "AAA",
		Nickname: "AAA",
		Role:     "Student",
		Password: "$2a$10$cm2Krm4dyA2Xu.qY705/EO96ZudVKv.YLpZDQMBAW2wWNEWOWQ6Ou",
		Picture:  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQj8Sn1pPGFGCvIWyxRnWdpfHCRbg-VL7Lomg&usqp=CAU",
	}
	var inbox = &models.Inbox{
		User1ID:           11,
		User1:             user,
		User2ID:           12,
		User2:             user2,
		LastMessageUserID: 11,
		Last_message:      "สวัสดีค่ะ",
	}
	var chat = &models.ChatRoom{
		Inbox_id: 1,
		Conversations: []*models.Conversation{
			{
				ChatRoom_id: 1,
				Sender_id:   11,
				Message:     "สวัสดีค่ะ",
				CreatedAt:   time.Now().Unix(),
			}},
	}
	DB.Create(user)
	DB.Create(user2)
	DB.Create(inbox)
	DB.Create(chat)
}

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	fmt.Println("Connected to DB")
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
	if err := DB.Create(&subject).Error; err != nil {
		panic(err)
	}

	fmt.Println("Subject created")

	//Teacher

	var teacher = &[]models.User{
		{
			Email:    "teacher1@mail.com", //ครูคณิต
			Password: "$2a$10$cm2Krm4dyA2Xu.qY705/EO96ZudVKv.YLpZDQMBAW2wWNEWOWQ6Ou",
			Fullname: "สุนีย์ คิดดี",
			Nickname: "ครูกุ๊กไก่",
			Role:     "Teacher",
			Picture:  "https://blog.startdee.com/hs-fs/hubfs/1%20Startdee/Blog%20Startdee/%E0%B8%84%E0%B8%B8%E0%B8%93%E0%B8%84%E0%B8%A3%E0%B8%B9%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B9%80%E0%B8%A3%E0%B8%B2/01%20Math/05%20Nook.jpg?width=210&name=05%20Nook.jpg",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "http://edunetwork-kukps.kustaff.ku.ac.th/wp-content/uploads/2020/04/15-Transcript-%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2-%E0%B8%99%E0%B8%B4%E0%B8%AA%E0%B8%B4%E0%B8%95%E0%B8%AA%E0%B8%B3%E0%B9%80%E0%B8%A3%E0%B9%87%E0%B8%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2_Page_1-scaled.jpg",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูสอนพิเศษตามบ้าน สอนเด็กเล็ก ชั้นอนุบาล-ประถมต้น",
						Evidence: "https://img-prod.api-onscene.com/cdn-cgi/image/format=auto,width=1600,height=900/https://sls-prod.api-onscene.com/partner_files/trueidintrend/124879/179988_0.jpg",
					},
				},
			},
		},
		{
			Email:    "teacher2@mail.com", //ครูสังคม
			Password: "$2a$10$cm2Krm4dyA2Xu.qY705/EO96ZudVKv.YLpZDQMBAW2wWNEWOWQ6Ou",
			Fullname: "วรรณษา ทองสุก",
			Nickname: "ครูเดือน",
			Role:     "Tutor",
			Picture:  "https://blog.startdee.com/hs-fs/hubfs/1%20Startdee/Blog%20Startdee/%E0%B8%84%E0%B8%B8%E0%B8%93%E0%B8%84%E0%B8%A3%E0%B8%B9%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B9%80%E0%B8%A3%E0%B8%B2/01%20Math/01%20Nhoi.jpg?width=210&name=01%20Nhoi.jpg",
			Teacher: &models.UserTeacher{
				Psychological_test: "https://testyourself.psychtests.com/testid/4176",
				Id_card:            "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูสอนพิเศษตามบ้าน สอนเด็กเล็ก ชั้นอนุบาล-ประถมต้น",
						Evidence: "https://img-prod.api-onscene.com/cdn-cgi/image/format=auto,width=1600,height=900/https://sls-prod.api-onscene.com/partner_files/trueidintrend/124879/179988_0.jpg",
					},
				},
			},
		},
		{
			Email:    "teacher3@mail.com", //ครูฟิสิก
			Password: "$2a$10$cm2Krm4dyA2Xu.qY705/EO96ZudVKv.YLpZDQMBAW2wWNEWOWQ6Ou",
			Fullname: "สุชาติ ม่วงดี",
			Nickname: "ครูเมฆ",
			Role:     "Teacher",
			Picture:  "https://blog.startdee.com/hs-fs/hubfs/1%20Startdee/Blog%20Startdee/%E0%B8%84%E0%B8%B8%E0%B8%93%E0%B8%84%E0%B8%A3%E0%B8%B9%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B9%80%E0%B8%A3%E0%B8%B2/01%20Math/06%20Ice.jpg?width=210&name=06%20Ice.jpg",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "http://edunetwork-kukps.kustaff.ku.ac.th/wp-content/uploads/2020/04/15-Transcript-%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2-%E0%B8%99%E0%B8%B4%E0%B8%AA%E0%B8%B4%E0%B8%95%E0%B8%AA%E0%B8%B3%E0%B9%80%E0%B8%A3%E0%B9%87%E0%B8%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2_Page_1-scaled.jpg",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูสอนพิเศษตามบ้าน สอนเด็กเล็ก ชั้นอนุบาล-ประถมต้น",
						Evidence: "https://img-prod.api-onscene.com/cdn-cgi/image/format=auto,width=1600,height=900/https://sls-prod.api-onscene.com/partner_files/trueidintrend/124879/179988_0.jpg",
					},
				},
			},
		},
		{
			Email:    "teacher4@mail.com", //ครูคณิต
			Password: "$2a$10$cm2Krm4dyA2Xu.qY705/EO96ZudVKv.YLpZDQMBAW2wWNEWOWQ6Ou",
			Fullname: "สิริยากร ผ่องพันธ์",
			Nickname: "ครูดาว",
			Role:     "Teacher",
			Picture:  "https://blog.startdee.com/hs-fs/hubfs/1%20Startdee/Blog%20Startdee/%E0%B8%84%E0%B8%B8%E0%B8%93%E0%B8%84%E0%B8%A3%E0%B8%B9%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B9%80%E0%B8%A3%E0%B8%B2/03%20Thai/02.png?width=210&name=02.png",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "http://edunetwork-kukps.kustaff.ku.ac.th/wp-content/uploads/2020/04/15-Transcript-%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2-%E0%B8%99%E0%B8%B4%E0%B8%AA%E0%B8%B4%E0%B8%95%E0%B8%AA%E0%B8%B3%E0%B9%80%E0%B8%A3%E0%B9%87%E0%B8%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2_Page_1-scaled.jpg",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูจิตอาสา สอนหนังสือให้เด็กเล็ก",
						Evidence: "https://www.jitarsabank.com/upload/images/20190210/AqRCaPsgcAO0hXXUg6yEA5apjQRfjohC4zdEwDEq.jpeg",
					},
				},
			},
		},
		{
			Email:    "teacher5@mail.com", //ครูคณิต
			Password: "$2a$10$cm2Krm4dyA2Xu.qY705/EO96ZudVKv.YLpZDQMBAW2wWNEWOWQ6Ou",
			Fullname: "นนทการ ศรีสุข",
			Nickname: "ครูสิงโต",
			Role:     "Teacher",
			Picture:  "https://blog.startdee.com/hs-fs/hubfs/1%20Startdee/Blog%20Startdee/%E0%B8%84%E0%B8%B8%E0%B8%93%E0%B8%84%E0%B8%A3%E0%B8%B9%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B9%80%E0%B8%A3%E0%B8%B2/07%20Bio/01.png?width=210&name=01.png",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "http://edunetwork-kukps.kustaff.ku.ac.th/wp-content/uploads/2020/04/15-Transcript-%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2-%E0%B8%99%E0%B8%B4%E0%B8%AA%E0%B8%B4%E0%B8%95%E0%B8%AA%E0%B8%B3%E0%B9%80%E0%B8%A3%E0%B9%87%E0%B8%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2_Page_1-scaled.jpg",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูสอนพิเศษตามบ้าน สอนเด็กเล็ก ชั้นอนุบาล-ประถมต้น",
						Evidence: "https://img-prod.api-onscene.com/cdn-cgi/image/format=auto,width=1600,height=900/https://sls-prod.api-onscene.com/partner_files/trueidintrend/124879/179988_0.jpg",
					},
				},
			},
		},
		{
			Email:    "teacher6@mail.com", //ครูอังกฤษ
			Password: "$2a$10$cm2Krm4dyA2Xu.qY705/EO96ZudVKv.YLpZDQMBAW2wWNEWOWQ6Ou",
			Fullname: "นาตาชา เเจ่มใส",
			Nickname: "ครูข้าวต้ม",
			Role:     "Teacher",
			Picture:  "https://blog.startdee.com/hs-fs/hubfs/1%20Startdee/Blog%20Startdee/%E0%B8%84%E0%B8%B8%E0%B8%93%E0%B8%84%E0%B8%A3%E0%B8%B9%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B9%80%E0%B8%A3%E0%B8%B2/05%20Eng/01.png?width=210&name=01.png",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "http://edunetwork-kukps.kustaff.ku.ac.th/wp-content/uploads/2020/04/15-Transcript-%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2-%E0%B8%99%E0%B8%B4%E0%B8%AA%E0%B8%B4%E0%B8%95%E0%B8%AA%E0%B8%B3%E0%B9%80%E0%B8%A3%E0%B9%87%E0%B8%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2_Page_1-scaled.jpg",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูจิตอาสา สอนหนังสือให้เด็กเล็ก",
						Evidence: "https://www.jitarsabank.com/upload/images/20190210/AqRCaPsgcAO0hXXUg6yEA5apjQRfjohC4zdEwDEq.jpeg",
					},
				},
			},
		},
		{
			Email:    "teacher7@mail.com", //ครูภาษาไทย
			Password: "$2a$10$cm2Krm4dyA2Xu.qY705/EO96ZudVKv.YLpZDQMBAW2wWNEWOWQ6Ou",
			Fullname: "สมเกียรติ มั่งมี",
			Nickname: "ครูเฟียส",
			Role:     "Teacher",
			Picture:  "https://blog.startdee.com/hs-fs/hubfs/1%20Startdee/Blog%20Startdee/%E0%B8%84%E0%B8%B8%E0%B8%93%E0%B8%84%E0%B8%A3%E0%B8%B9%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B9%80%E0%B8%A3%E0%B8%B2/01%20Math/07%20Leng.jpg?width=210&name=07%20Leng.jpg",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "http://edunetwork-kukps.kustaff.ku.ac.th/wp-content/uploads/2020/04/15-Transcript-%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2-%E0%B8%99%E0%B8%B4%E0%B8%AA%E0%B8%B4%E0%B8%95%E0%B8%AA%E0%B8%B3%E0%B9%80%E0%B8%A3%E0%B9%87%E0%B8%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2_Page_1-scaled.jpg",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูสอนพิเศษตามบ้าน สอนเด็กเล็ก ชั้นอนุบาล-ประถมต้น",
						Evidence: "https://img-prod.api-onscene.com/cdn-cgi/image/format=auto,width=1600,height=900/https://sls-prod.api-onscene.com/partner_files/trueidintrend/124879/179988_0.jpg",
					},
				},
			},
		},
		{
			Email:    "teacher8@mail.com", //ครูเคมี
			Password: "$2a$10$cm2Krm4dyA2Xu.qY705/EO96ZudVKv.YLpZDQMBAW2wWNEWOWQ6Ou",
			Fullname: "เเพรวา เเจ่มเเจ้ง",
			Nickname: "ครูปลา",
			Role:     "Teacher",
			Picture:  "https://blog.startdee.com/hs-fs/hubfs/1%20Startdee/Blog%20Startdee/%E0%B8%84%E0%B8%B8%E0%B8%93%E0%B8%84%E0%B8%A3%E0%B8%B9%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B9%80%E0%B8%A3%E0%B8%B2/06%20Sci/01.jpg?width=210&name=01.jpg",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "http://edunetwork-kukps.kustaff.ku.ac.th/wp-content/uploads/2020/04/15-Transcript-%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2-%E0%B8%99%E0%B8%B4%E0%B8%AA%E0%B8%B4%E0%B8%95%E0%B8%AA%E0%B8%B3%E0%B9%80%E0%B8%A3%E0%B9%87%E0%B8%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2_Page_1-scaled.jpg",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูจิตอาสา สอนหนังสือให้เด็กเล็ก",
						Evidence: "https://www.jitarsabank.com/upload/images/20190210/AqRCaPsgcAO0hXXUg6yEA5apjQRfjohC4zdEwDEq.jpeg",
					},
				},
			},
		},
		{
			Email:    "teacher9@mail.com",
			Password: "$2a$10$cm2Krm4dyA2Xu.qY705/EO96ZudVKv.YLpZDQMBAW2wWNEWOWQ6Ou",
			Fullname: "ภูวินทร์ สีม่วง",
			Nickname: "ครูติ๊ก",
			Role:     "Teacher",
			Picture:  "https://blog.startdee.com/hs-fs/hubfs/1%20Startdee/Blog%20Startdee/%E0%B8%84%E0%B8%B8%E0%B8%93%E0%B8%84%E0%B8%A3%E0%B8%B9%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B9%80%E0%B8%A3%E0%B8%B2/02%20Soc/04%20Rabhi.jpg?width=210&name=04%20Rabhi.jpg",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "http://edunetwork-kukps.kustaff.ku.ac.th/wp-content/uploads/2020/04/15-Transcript-%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2-%E0%B8%99%E0%B8%B4%E0%B8%AA%E0%B8%B4%E0%B8%95%E0%B8%AA%E0%B8%B3%E0%B9%80%E0%B8%A3%E0%B9%87%E0%B8%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2_Page_1-scaled.jpg",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูสอนพิเศษตามบ้าน สอนเด็กเล็ก ชั้นอนุบาล-ประถมต้น",
						Evidence: "https://img-prod.api-onscene.com/cdn-cgi/image/format=auto,width=1600,height=900/https://sls-prod.api-onscene.com/partner_files/trueidintrend/124879/179988_0.jpg",
					},
				},
			},
		},
		{
			Email:    "teacher10@mail.com", //ครูอังกฤษ
			Password: "$2a$10$cm2Krm4dyA2Xu.qY705/EO96ZudVKv.YLpZDQMBAW2wWNEWOWQ6Ou",
			Fullname: "ฟ้าใส ชื่นมื่น",
			Nickname: "ครูฟ้า",
			Role:     "Teacher",
			Picture:  "https://blog.startdee.com/hs-fs/hubfs/1%20Startdee/Blog%20Startdee/%E0%B8%84%E0%B8%B8%E0%B8%93%E0%B8%84%E0%B8%A3%E0%B8%B9%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B9%80%E0%B8%A3%E0%B8%B2/06%20Sci/03.png?width=210&name=03.png",
			Teacher: &models.UserTeacher{
				Teacher_license: "https://www.kruachieve.com/wp-content/uploads/2019/12/562000004488703.jpg",
				Transcript:      "http://edunetwork-kukps.kustaff.ku.ac.th/wp-content/uploads/2020/04/15-Transcript-%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2-%E0%B8%99%E0%B8%B4%E0%B8%AA%E0%B8%B4%E0%B8%95%E0%B8%AA%E0%B8%B3%E0%B9%80%E0%B8%A3%E0%B9%87%E0%B8%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2_Page_1-scaled.jpg",
				Id_card:         "http://www.bikes4deal.com/picture_upload/market/full/2021-09/B2ceH51632344286.jpeg",
				Experience: &[]models.Experience{
					{
						Title:    "เป็นครูจิตอาสา สอนหนังสือให้เด็กเล็ก",
						Evidence: "https://www.jitarsabank.com/upload/images/20190210/AqRCaPsgcAO0hXXUg6yEA5apjQRfjohC4zdEwDEq.jpeg",
					},
				},
			},
		},
	}
	if err := DB.Create(&teacher).Error; err != nil {
		panic(err)
	}

	fmt.Println("Teacher created")

	//rewardItem
	var rewardItem = &[]models.Reward_Item{
		{
			Item_name:    "หนังสือเรียนภาษาอังกฤษ Easy Street English",
			Item_title:   "Easy Street English พูดอังกฤษง่ายๆ ใช้งานได้จริง อัดแน่นไปด้วยบทสนทนาภาษาอังกฤษที่ใช้บ่อยในชีวิตประจำวัน การนัดหมาย การเดินทาง การช้อปปิ้ง มื้ออาหาร จองที่พัก เข้าร้านเสริมสวย และอื่นๆ เหมาะสำหรับผู้ที่ต้องการพัฒนาภาษาอังกฤษ ราคา 147 บาท",
			Item_cost:    147,
			Item_picture: "https://promotions.co.th/wp-content/uploads/2021/07/Easy-Street-English.png",
		},
		{
			Item_name:    "หนังสือเรียนภาษาอังกฤษ English From Zero",
			Item_title:   "English From Zero หนังสือ เรียนพูด อ่าน เขียน ภาษาอังกฤษ เริ่มต้นจากศูนย์ พร้อมรวบรวม 22 บทเรียน นับตั้งแต่เรียนตัวอักษร การประสมคำ คำศัพท์ ไวยากรณ์ พูด อ่าน เขียน แปล และอีกมากมาย เรียนรู้ได้ง่าย เข้าใจได้อย่างรวดเร็ว ราคา 201 บาท",
			Item_cost:    201,
			Item_picture: "https://promotions.co.th/wp-content/uploads/2021/07/English-From-Zero.png",
		},
		{
			Item_name:    "หนังสือเรียน 9 สามัญ",
			Item_title:   "หนังสือ เตรียมสอบเข้ม 9 วิชาสามัญ รวมทุกวิชา มั่นใจเต็ม 100",
			Item_cost:    355,
			Item_picture: "https://serazu.com/library/products/2315/XXXXXL9786164871359.jpg",
		},
		{
			Item_name:    "กระเป๋าดินสอ Master Art Funny Deer",
			Item_title:   "กระเป๋าดินสอ Master Art Funny Deer รุ่น MF-C01 (คละลาย)",
			Item_cost:    45,
			Item_picture: "https://www.stationerymine.com/media/catalog/product/cache/cf3f2243ef4940fd5c66f2ff035145ac/2/7/276232_msta_pencil_case_funny_deer_mf-c01_blue_001.jpg",
		},
		{
			Item_name:    "ชุดดินสอสีไม้",
			Item_title:   "ชุดดินสอสีไม้ 36 สี TOMBOW CB-NQ36C",
			Item_cost:    250,
			Item_picture: "https://tombowth.com/wp-content/uploads/2021/05/front-2.jpg",
		},
		{
			Item_name:    "ปฏิทินตั้งโต๊ะ",
			Item_title:   "ปฏิทินตั้งโต๊ะเก๋ๆ น่ารักๆ 2560 ลายการ์ตูน แพนด้า",
			Item_cost:    99,
			Item_picture: "https://www.ubkkikujung.com/shop/ubkkikujung/images/4yx3e5xyc1gifyb4xfd541220162259000260.jpg",
		},
		{
			Item_name:    "เคสไอเเพด Uniq",
			Item_title:   " เคสบาง เบา ด้านหลังเคสเป็นแบบใสโชว์ตัวเครื่อง ตัวเคสแนวพับรูป Y ที่รองรับการตั้งได้ 3 รูปแบบ แนวตั้ง แนวนอนสำหรับการดูหนัง แนวนอนสำหรับการพิมพ์",
			Item_cost:    450,
			Item_picture: "https://media.studio7thailand.com/65430/Uniq-Casing-for-iPad-Air-5-Air-4-10.9-2022-Camden-Northern-Blue-1-square_medium.jpg",
		},
		{
			Item_name:    "หนังสือชีวะเต่าทอง (Biology)",
			Item_title:   " หนังสือสรุปชีววิทยายอด HIT ที่หลายคนบอกต่อว่าเก็งข้อสอบตรงจุดมากๆ มีเนื้อหาอ่านง่ายรวม 23 บท มีแบบฝึกหัดที่ใช้ได้จริง 1,000 ข้อ",
			Item_cost:    750,
			Item_picture: "https://inwfile.com/s-c/rc270s.jpg",
		},
		{
			Item_name:    "ฟิสิกส์ ขนมหวาน เล่มที่ 1",
			Item_title:   "เนื้อหาเรียบเรียงให้เข้าใจง่าย แบบฝึกหัดเรียงจากพื้นฐานไปสู่ขั้นสูง พร้อมเฉลยวิธีคิดแบบละเอียดท้ายเล่ม",
			Item_cost:    255,
			Item_picture: "https://images-se-ed.com/ws/Storage/Originals/978616/321/9786163211514L.jpg?h=fe473fe85ae302d506b6c77853efee11",
		},
		{
			Item_name:    "ฟิสิกส์ ขนมหวาน เล่มที่ 2",
			Item_title:   "เนื้อหาเรียบเรียงให้เข้าใจง่าย แบบฝึกหัดเรียงจากพื้นฐานไปสู่ขั้นสูง พร้อมเฉลยวิธีคิดแบบละเอียดท้ายเล่ม",
			Item_cost:    250,
			Item_picture: "https://images-se-ed.com/ws/Storage/Originals/978999/007/9789990073546L.jpg?h=1cd09dd595cda68636d881367ee37073",
		},
	}
	if err := DB.Create(&rewardItem).Error; err != nil {
		panic(err)
	}

	//Course
	var course = &[]models.Course{
		{
			//Course_1
			SubjectID: 1,
			Videos: &[]models.Video{
				{
					Class_level: 1,
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
					Class_level: 1,
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
					Class_level: 1,
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
					Class_level: 2,
					Price:       29,
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
					Class_level: 2,
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
					Class_level: 2,
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
					Class_level: 3,
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
					Class_level: 3,
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
					Class_level: 3,
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
					Class_level: 4,
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
					Class_level: 4,
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
					Class_level: 4,
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
			Course_name: "คณิตศาสตร์ ชั้น ม.4 เรื่อง เซต",
			Picture:     "https://i.ytimg.com/vi/5ZegR_M8tJs/maxresdefault.jpg",
			Description: "“เซต” คืออะไร คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ เซต เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_5
			SubjectID: 5,
			Videos: &[]models.Video{
				{
					Class_level: 5,
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
					Class_level: 5,
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
					Class_level: 5,
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
					Class_level: 6,
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
					Class_level: 6,
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
					Class_level: 6,
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
					Class_level: 1,
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
					Class_level: 1,
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
					Class_level: 1,
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
					Class_level: 2,
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
					Class_level: 2,
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
					Class_level: 2,
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
					Class_level: 3,
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
					Class_level: 3,
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
					Class_level: 3,
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
					Class_level: 4,
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
					Class_level: 4,
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
					Class_level: 4,
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
		{
			//Course_11
			SubjectID: 24,
			Videos: &[]models.Video{
				{
					Class_level: 5,
					Video_name:  "การใช้ Present perfect",
					Picture:     "https://i.ytimg.com/vi/5a8m3FLHoDY/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ ชั้น ม.5 เรื่อง การใช้ Present perfect | แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาอังกฤษ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_11%2FVideo_1%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.5%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20Present%20perfect.mp4?alt=media&token=e80cbb5f-51eb-4f5a-91e3-1290892f186a",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_11%2FVideo_1%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20present%20perfect.pdf?alt=media&token=1a828bbe-9ba3-4cfd-a6c1-6e1d1eaaf929",
					Exercises: &[]models.Exercise{
						{
							Question: "1.Choose the correct answer. She has ..... in london since 2018.",
							Choices: &[]models.Choice{
								{
									Title:   "lived",
									Correct: true,
								},
								{
									Title:   "live",
									Correct: false,
								},
								{
									Title:   "lives",
									Correct: false,
								},
								{
									Title:   "ผิดทุกข้อ",
									Correct: false,
								},
							},
						},
						{
							Question: "2.Choose the correct answer. I haven't ....... George in ages.",
							Choices: &[]models.Choice{
								{
									Title:   "look",
									Correct: false,
								},
								{
									Title:   "seed",
									Correct: false,
								},
								{
									Title:   "see",
									Correct: false,
								},
								{
									Title:   "seen",
									Correct: true,
								},
							},
						},
						{
							Question: "3.Choose the correct answer. He's already ..... (drink) coffee.",
							Choices: &[]models.Choice{
								{
									Title:   "drink",
									Correct: false,
								},
								{
									Title:   "drank",
									Correct: false,
								},
								{
									Title:   "drunk",
									Correct: true,
								},
								{
									Title:   "drinks",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 5,
					Video_name:  "การใช้ Present Perfect Progressive",
					Picture:     "https://i.ytimg.com/vi/XDNt3cMc8Ls/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ ชั้น ม.5 เรื่อง การใช้ Tenses : Present Perfect Progressive |  แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาอังกฤษ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_11%2FVideo_2%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.5%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20Present%20Perfect%20Progressive.mp4?alt=media&token=9e592919-eed2-43ff-be2f-b35ea861735a",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_11%2FVideo_2%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20Present%20perfect%20progressive..pdf?alt=media&token=e6b48d77-7db8-46cb-af3c-d1cf90aa9cbc",
					Exercises: &[]models.Exercise{
						{
							Question: "1. Choose the correct answer. I have ..... ..... games since last night.",
							Choices: &[]models.Choice{
								{
									Title:   "be play",
									Correct: false,
								},
								{
									Title:   "be playing",
									Correct: false,
								},
								{
									Title:   "been play",
									Correct: false,
								},
								{
									Title:   "been playing",
									Correct: true,
								},
							},
						},
						{
							Question: "2.Choose the correct answer. How long has she ..... ..... ?",
							Choices: &[]models.Choice{
								{
									Title:   "been wait",
									Correct: false,
								},
								{
									Title:   "been waiting",
									Correct: true,
								},
								{
									Title:   "be waiting",
									Correct: false,
								},
								{
									Title:   "be wait",
									Correct: false,
								},
							},
						},
						{
							Question: "3.Choose the correct answer. It's ..... ..... cats and dogs all night. ",
							Choices: &[]models.Choice{
								{
									Title:   "been rain",
									Correct: false,
								},
								{
									Title:   "been raining",
									Correct: true,
								},
								{
									Title:   "be raining",
									Correct: false,
								},
								{
									Title:   "be rain",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 5,
					Video_name:  "การใช้ Passive voice",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/g3o6IY6Yrsw/mqdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอน การใช้ Passive voice แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_11%2FVideo_3%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.5%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20Passive%20voice.mp4?alt=media&token=b47c3897-ae4b-4ddc-b357-92bdeb4751d4",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_11%2FVideo_3%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20Passive%20voice.pdf?alt=media&token=6b244ceb-e242-4779-a75b-c43fd17b5355",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดไม่ใช่ Passive voice ",
							Choices: &[]models.Choice{
								{
									Title:   "Ice-cream is eaten everyday by me.",
									Correct: false,
								},
								{
									Title:   "I eat ice-cream everyday.",
									Correct: true,
								},
								{
									Title:   "Dinner was cooked by Jane yesterday",
									Correct: false,
								},
								{
									Title:   "That car was bought by my mother.",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดเป็น Passive voice",
							Choices: &[]models.Choice{
								{
									Title:   "I eat ice-cream everyday.",
									Correct: false,
								},
								{
									Title:   "Jane feeds the dog",
									Correct: false,
								},
								{
									Title:   "Ice-cream is eaten everyday by me.",
									Correct: true,
								},
								{
									Title:   "Joy will have been feeding the cat",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   10,
			Course_name: "ภาษาอังกฤษ ชั้น ม.5 เรื่อง Tense ",
			Picture:     "http://www.learnersplanet.com/sites/default/files/images/Tenses%20chart-1.png",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ ภาษาอังกฤษพื้นฐานในเรื่อง tense เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_12
			SubjectID: 25,
			Videos: &[]models.Video{
				{
					Class_level: 6,
					Video_name:  "การใช้ Present perfect",
					Picture:     "https://i.ytimg.com/vi/Yi5Z19yolPs/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ ชั้น ม.6 เรื่อง การใช้ Present perfect | แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาอังกฤษ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_12%2FVideo_1%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.6%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20Present%20perfect.mp4?alt=media&token=a87fb784-a035-483c-be44-0ec29a02cf82",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_12%2FVideo_1%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20present%20perfect.pdf?alt=media&token=f09e01cf-687a-4838-afa9-16781186e978",
					Exercises: &[]models.Exercise{
						{
							Question: "1.Conjugate the Verbs. Sheila _____ (not live) in this country for a very long time. ",
							Choices: &[]models.Choice{
								{
									Title:   "hasn't lived",
									Correct: true,
								},
								{
									Title:   "has lived",
									Correct: false,
								},
								{
									Title:   "hasn't live",
									Correct: false,
								},
								{
									Title:   "has live",
									Correct: false,
								},
							},
						},
						{
							Question: "2.Conjugate the Verbs. My father _____ (wear) that jacket for more than twenty years.",
							Choices: &[]models.Choice{
								{
									Title:   "worn",
									Correct: false,
								},
								{
									Title:   "wear",
									Correct: false,
								},
								{
									Title:   "hasn't worn",
									Correct: false,
								},
								{
									Title:   "has worn",
									Correct: true,
								},
							},
						},
					},
				},
				{
					Class_level: 6,
					Video_name:  "การใช้ Present perfect progressive",
					Picture:     "https://i.ytimg.com/vi/ZK8tWKHpj1k/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ ชั้น ม.5 เรื่อง การใช้ Tenses : Present perfect progressive |  แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาอังกฤษ ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_12%2FVideo_2%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.6%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20Present%20perfect%20progressive.mp4?alt=media&token=5f7fa8fb-ff98-45f4-b282-616d37fc5fae",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_12%2FVideo_2%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.6%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20Present%20perfect%20progressive..pdf?alt=media&token=ec19c2ea-3f05-428b-aa66-94e007407ac2",
					Exercises: &[]models.Exercise{
						{
							Question: "1. Choose the correct answer. Teddy  ____ ____ ____ this song for months.",
							Choices: &[]models.Choice{
								{
									Title:   "has be write",
									Correct: false,
								},
								{
									Title:   "has be writing",
									Correct: false,
								},
								{
									Title:   "has been write",
									Correct: false,
								},
								{
									Title:   "has been writing",
									Correct: true,
								},
							},
						},
						{
							Question: "2.Choose the correct answer. Who has she _____ _____ with these days?",
							Choices: &[]models.Choice{
								{
									Title:   "been travel",
									Correct: false,
								},
								{
									Title:   "been travelling",
									Correct: true,
								},
								{
									Title:   "be travelling",
									Correct: false,
								},
								{
									Title:   "be travel",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 6,
					Video_name:  "การใช้ Wish + Past perfect",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/O2B_3GzucHI/maxresdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอน การใช้ Wish + Past perfect แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_12%2FVideo_3%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.6%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20Wish%20%20%20Past%20perfect.mp4?alt=media&token=76219cef-5f0d-450a-ae35-3cfd373ba652",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_12%2FVideo_3%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20Wish%20%20%20Past%20perfect.pdf?alt=media&token=edcb18c8-3cb0-4385-963a-27066944db14",
					Exercises: &[]models.Exercise{
						{
							Question: "1.Choose the correct answer. I wish they _____ the dinner.",
							Choices: &[]models.Choice{
								{
									Title:   "had seen",
									Correct: false,
								},
								{
									Title:   "had made",
									Correct: true,
								},
								{
									Title:   "had had",
									Correct: false,
								},
								{
									Title:   "hadn't broke.",
									Correct: false,
								},
							},
						},
						{
							Question: "2.Choose the correct answer. I wish he _____ there for you.",
							Choices: &[]models.Choice{
								{
									Title:   "had cleaned",
									Correct: false,
								},
								{
									Title:   "hadn't broke",
									Correct: false,
								},
								{
									Title:   "had been",
									Correct: true,
								},
								{
									Title:   "hadn't said",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   10,
			Course_name: "ภาษาอังกฤษ ชั้น ม.6 เรื่อง Tense ",
			Picture:     "https://d1bnvx5vhcnf8w.cloudfront.net/Business%20English%20-%20Coaches/tense.png",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ ภาษาอังกฤษพื้นฐานในเรื่อง tense เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_13
			SubjectID: 14,
			Videos: &[]models.Video{
				{
					Class_level: 1,
					Video_name:  "เสียงพยัญชนะในภาษาไทย",
					Picture:     "https://i.ytimg.com/vi/yXhtEPiJAkY/maxresdefault.jpg",
					Description: "ภาษาไทย ชั้น ม.1 เรื่อง การใช้ เสียงพยัญชนะในภาษาไทย| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาไทย	 ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_13%2FVideo_1%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.1%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B9%80%E0%B8%AA%E0%B8%B5%E0%B8%A2%E0%B8%87%E0%B8%9E%E0%B8%A2%E0%B8%B1%E0%B8%8D%E0%B8%8A%E0%B8%99%E0%B8%B0.mp4?alt=media&token=18488714-5a59-4879-8b65-c5d799528e95",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_13%2FVideo_1%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B9%80%E0%B8%AA%E0%B8%B5%E0%B8%A2%E0%B8%87%E0%B8%9E%E0%B8%A2%E0%B8%B1%E0%B8%8D%E0%B8%8A%E0%B8%99%E0%B8%B0.pdf?alt=media&token=2f1b204c-258f-41b7-8aaf-91d276b9eaae",
					Exercises: &[]models.Exercise{
						{
							Question: "1.พยัญชนะที่มีเสียงเดียวกันท้้งหมดคือข้อใด ",
							Choices: &[]models.Choice{
								{
									Title:   "ฑ ฒ ฐ ธ",
									Correct: true,
								},
								{
									Title:   "ช ซ ฌ ฉ",
									Correct: false,
								},
								{
									Title:   "ก ข ฃ ค",
									Correct: false,
								},
								{
									Title:   "ฉ ส ษ ส",
									Correct: false,
								},
							},
						},
						{
							Question: "2.เสียงพยัญชนะต้นในคำใด เเตกต่างจากคำอื่น",
							Choices: &[]models.Choice{
								{
									Title:   "หู",
									Correct: false,
								},
								{
									Title:   "เฮ",
									Correct: false,
								},
								{
									Title:   "ฮา",
									Correct: false,
								},
								{
									Title:   "หลาย",
									Correct: true,
								},
							},
						},
						{
							Question: "3.เสียงพยัญชนะต้นในคำใดเเตกต่างจากคำอื่น",
							Choices: &[]models.Choice{
								{
									Title:   "อยู่",
									Correct: false,
								},
								{
									Title:   "หญิง",
									Correct: false,
								},
								{
									Title:   "ญาณ",
									Correct: false,
								},
								{
									Title:   "ฌาน",
									Correct: true,
								},
							},
						},
					},
				},
				{
					Class_level: 1,
					Video_name:  "เสียงสระในภาษาไทย",
					Picture:     "https://i.ytimg.com/vi/JxUVunUVZYI/maxresdefault.jpg",
					Description: "ภาษาไทย ชั้น ม.1 เรื่อง เสียงสระในภาษาไทย | แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาไทย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_13%2FVideo_2%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.1%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B9%80%E0%B8%AA%E0%B8%B5%E0%B8%A2%E0%B8%87%E0%B8%AA%E0%B8%A3%E0%B8%B0.mp4?alt=media&token=a46757d8-18c4-4da0-91f2-693daf028dbc",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_13%2FVideo_2%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B9%80%E0%B8%AA%E0%B8%B5%E0%B8%A2%E0%B8%87%E0%B8%AA%E0%B8%A3%E0%B8%B0.pdf?alt=media&token=383e7c02-28c0-41a8-9a19-92a36b1af962",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดเป็นสระเสียงยาวทั้งหมด ",
							Choices: &[]models.Choice{
								{
									Title:   "จุ๊บ จั๊บ จ๋า จ๊ะ",
									Correct: false,
								},
								{
									Title:   "ครู ขา หนู ร้อน",
									Correct: true,
								},
								{
									Title:   "คิด หนัก จัก เศร้า",
									Correct: false,
								},
								{
									Title:   "พาก เพียร เรียน เถอะ",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดต่อไปนี้ประกอบด้วยสระเสียงสั้นทั้งหมด",
							Choices: &[]models.Choice{
								{
									Title:   "เเพทย์ ครู",
									Correct: false,
								},
								{
									Title:   "ถัง กระติก",
									Correct: true,
								},
								{
									Title:   "สมุด ดินสอ",
									Correct: false,
								},
								{
									Title:   "ปากกา ยางลบ",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 1,
					Video_name:  "เสียงวรรณยุกต์ และทบทวนเรื่องการผันวรรณยุกต์",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/4qoDxnZAru8/maxresdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอน เสียงวรรณยุกต์ และทบทวนเรื่องการผันวรรณยุกต์ แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_13%2FVideo_3%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.1%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B9%80%E0%B8%AA%E0%B8%B5%E0%B8%A2%E0%B8%87%E0%B8%A7%E0%B8%A3%E0%B8%A3%E0%B8%93%E0%B8%A2%E0%B8%B8%E0%B8%81%E0%B8%95%E0%B9%8C%20%E0%B9%81%E0%B8%A5%E0%B8%B0%E0%B8%97%E0%B8%9A%E0%B8%97%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%9C%E0%B8%B1%E0%B8%99%E0%B8%A7%E0%B8%A3%E0%B8%A3%E0%B8%93%E0%B8%A2%E0%B8%B8%E0%B8%81%E0%B8%95%E0%B9%8C.mp4?alt=media&token=6e3044be-3a21-4a1e-b8f3-f9b8ae136082",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_13%2FVideo_3%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B9%80%E0%B8%AA%E0%B8%B5%E0%B8%A2%E0%B8%87%E0%B8%A7%E0%B8%A3%E0%B8%A3%E0%B8%93%E0%B8%A2%E0%B8%B8%E0%B8%81%E0%B8%95%E0%B9%8C%20%E0%B9%81%E0%B8%A5%E0%B8%B0%E0%B8%97%E0%B8%9A%E0%B8%97%E0%B8%A7%E0%B8%99%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%9C%E0%B8%B1%E0%B8%99%E0%B8%A7%E0%B8%A3%E0%B8%A3%E0%B8%93%E0%B8%A2%E0%B8%B8%E0%B8%81%E0%B8%95%E0%B9%8C.pdf?alt=media&token=a1cfa19f-a64d-407b-af5e-4ab2cb36a1f1",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดมีเสียงวรรณยุกต์สามัญทั้งหมด",
							Choices: &[]models.Choice{
								{
									Title:   "ใครไปลำปาง",
									Correct: true,
								},
								{
									Title:   "นกเกาะปลายไม้",
									Correct: false,
								},
								{
									Title:   "ยายกิ่งไม่ร้องหรือ",
									Correct: false,
								},
								{
									Title:   "เสือเดินมาโน่นเเล้ว",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดมีเสียงวรรณยุกต์ครบห้าเสียง",
							Choices: &[]models.Choice{
								{
									Title:   "ไปเที่ยวกันนะจ๊ะ",
									Correct: false,
								},
								{
									Title:   "คุณเเม่อยู่ไหมคะ",
									Correct: true,
								},
								{
									Title:   "ดูหนังฟังเพลงกัน",
									Correct: false,
								},
								{
									Title:   "ปรานีชวนพรไปนา",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   7,
			Course_name: "ภาษาไทยง๊ายง่าย ",
			Picture:     "https://sites.google.com/a/samakkhi.ac.th/keiyw-kab-phasa-thiy/_/rsrc/1488422650332/home/20161121160657.png",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ ภาษาไทยพื้นฐานในเรื่องต่าง ๆ เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_14
			SubjectID: 15,
			Videos: &[]models.Video{
				{
					Class_level: 2,
					Video_name:  "คำราชาศัพท์",
					Picture:     "https://i.ytimg.com/vi/oBDM3hrasv0/maxresdefault.jpg",
					Description: "ภาษาไทย ชั้น ม.2 เรื่อง การใช้คำราชาศัพท์| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาไทย	 ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_14%2FVideo_1%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.2%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%84%E0%B8%B3%E0%B8%A3%E0%B8%B2%E0%B8%8A%E0%B8%B2%E0%B8%A8%E0%B8%B1%E0%B8%9E%E0%B8%97%E0%B9%8C.mp4?alt=media&token=7c4db40f-d100-47ab-9a8c-e1a34f80e35e",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_14%2FVideo_1%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%84%E0%B8%B3%E0%B8%A3%E0%B8%B2%E0%B8%8A%E0%B8%B2%E0%B8%A8%E0%B8%B1%E0%B8%9E%E0%B8%97%E0%B9%8C.pdf?alt=media&token=e48d4df6-5904-4878-b289-703379e6ed18",
					Exercises: &[]models.Exercise{
						{
							Question: "1.คำใดไม่มีคำราชาศัพท์",
							Choices: &[]models.Choice{
								{
									Title:   "ทรงประทับ",
									Correct: true,
								},
								{
									Title:   "ทรงพระสรวล",
									Correct: false,
								},
								{
									Title:   "ทรงบำเพ็ญพระราชกุศล",
									Correct: false,
								},
								{
									Title:   "ทรงพระประชวร",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดผิด",
							Choices: &[]models.Choice{
								{
									Title:   "ลิ้น - พระชิวหา",
									Correct: false,
								},
								{
									Title:   "ฟัน - พระทนต์",
									Correct: false,
								},
								{
									Title:   "จมูก - พระชานุ",
									Correct: true,
								},
								{
									Title:   "คอ - พระศอ",
									Correct: false,
								},
							},
						},
						{
							Question: "3.ข้อใดใช้คำราชาศัพท์ถูกต้อง",
							Choices: &[]models.Choice{
								{
									Title:   "ดู - ทรงพระเนตร",
									Correct: false,
								},
								{
									Title:   "ชอบ - ทรงพระโปรด",
									Correct: false,
								},
								{
									Title:   "เขียน - ทรงพระอักษร",
									Correct: true,
								},
								{
									Title:   "พูด - ทรงตรัส",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 2,
					Video_name:  "คำสมาสแบบสมาส และสมาสแบบสนธิ",
					Picture:     "https://i.ytimg.com/vi/iZ6mAvHImNA/maxresdefault.jpg",
					Description: "ภาษาไทย ชั้น ม.2 เรื่อง คำสมาสแบบสมาส และสมาสแบบสนธิ | แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาไทย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_14%2FVideo_2%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.2%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%84%E0%B8%B3%E0%B8%AA%E0%B8%A1%E0%B8%B2%E0%B8%AA%E0%B9%81%E0%B8%9A%E0%B8%9A%E0%B8%AA%E0%B8%A1%E0%B8%B2%E0%B8%AA%20%E0%B9%81%E0%B8%A5%E0%B8%B0%E0%B8%AA%E0%B8%A1%E0%B8%B2%E0%B8%AA%E0%B9%81%E0%B8%9A%E0%B8%9A%E0%B8%AA%E0%B8%99%E0%B8%98%E0%B8%B4.mp4?alt=media&token=d0b3a811-dc5b-41ee-b548-77fc14007fa7",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_14%2FVideo_2%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%84%E0%B8%B3%E0%B8%AA%E0%B8%A1%E0%B8%B2%E0%B8%AA%E0%B9%81%E0%B8%9A%E0%B8%9A%E0%B8%AA%E0%B8%A1%E0%B8%B2%E0%B8%AA%20%E0%B9%81%E0%B8%A5%E0%B8%B0%E0%B8%AA%E0%B8%A1%E0%B8%B2%E0%B8%AA%E0%B9%81%E0%B8%9A%E0%B8%9A%E0%B8%AA%E0%B8%99%E0%B8%98%E0%B8%B4.pdf?alt=media&token=334633df-cd78-452b-b153-c2d5d748dd7f",
					Exercises: &[]models.Exercise{
						{
							Question: "1.คำต่อไปนี้คำใดเป็นคำสมาส ",
							Choices: &[]models.Choice{
								{
									Title:   "สงสาร",
									Correct: false,
								},
								{
									Title:   "จันทร์เพ็ญ",
									Correct: false,
								},
								{
									Title:   "ถิ่นฐาน",
									Correct: false,
								},
								{
									Title:   "วิทยฐานะ",
									Correct: true,
								},
							},
						},
						{
							Question: "2.ข้อใดเป็นคำสมาสทั้งหมด",
							Choices: &[]models.Choice{
								{
									Title:   "อธิการ อัศวมุข มหาบุรุษ",
									Correct: true,
								},
								{
									Title:   "ธนบัตร โภชนาหาร ภัตตาคาร",
									Correct: false,
								},
								{
									Title:   "พฤษภาคม อิสรภาพ สารคดี",
									Correct: false,
								},
								{
									Title:   "จินตกวี กิตติศัพท์ จักขวาพาธ",
									Correct: false,
								},
							},
						},
						{
							Question: "3.คำว่า 'อนามัย' แยกสนธิได้อย่างไร",
							Choices: &[]models.Choice{
								{
									Title:   "อ+นามัย",
									Correct: true,
								},
								{
									Title:   "อนา+มัย",
									Correct: false,
								},
								{
									Title:   "อนา+อมัย",
									Correct: false,
								},
								{
									Title:   "อน+อามัย",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 2,
					Video_name:  "คำภาษาต่างๆที่เข้ามาในภาษาไทย ",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/e_AWrgJY4xU/maxresdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอน คำภาษาต่างๆที่เข้ามาในภาษาไทย แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_14%2FVideo_3%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.2%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%84%E0%B8%B3%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%95%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B9%86%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%80%E0%B8%82%E0%B9%89%E0%B8%B2%E0%B8%A1%E0%B8%B2%E0%B9%83%E0%B8%99%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2.mp4?alt=media&token=0457825f-d02d-43e2-aaa5-ec5537c3792d",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_14%2FVideo_3%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%84%E0%B8%B3%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%95%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B9%86%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%80%E0%B8%82%E0%B9%89%E0%B8%B2%E0%B8%A1%E0%B8%B2%E0%B9%83%E0%B8%99%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2.PDF?alt=media&token=6f673d55-26a2-46a9-9a76-fdd3e532c120",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดไม่ใช่คำภาษาต่างประเทศที่มาจากภาษาญี่ปุ่น",
							Choices: &[]models.Choice{
								{
									Title:   "บะหมี่",
									Correct: true,
								},
								{
									Title:   "ซูชิ",
									Correct: false,
								},
								{
									Title:   "วาซาบิ",
									Correct: false,
								},
								{
									Title:   "ราเม็ง",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดไม่ใช่คำภาษาต่างประเทศที่มาจากภาษาจีน",
							Choices: &[]models.Choice{
								{
									Title:   "เฉาก๋วย",
									Correct: false,
								},
								{
									Title:   "ซูชิ",
									Correct: true,
								},
								{
									Title:   "เต้าหู้",
									Correct: false,
								},
								{
									Title:   "บ๊ะจ่าง",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   7,
			Course_name: "ภาษาไทยน่ารู้ ",
			Picture:     "http://www.kabinburischool.ac.th/wp-content/uploads/2020/05/ปกภาษาไทย-1.jpg",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ ภาษาไทยพื้นฐานในเรื่องต่าง ๆ เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_15
			SubjectID: 16,
			Videos: &[]models.Video{
				{
					Class_level: 3,
					Video_name:  "ประโยคความเดียวแบบซับซ้อน",
					Picture:     "https://i.ytimg.com/vi/V_VExzR4bdg/maxresdefault.jpg",
					Description: "ภาษาไทย ชั้น ม.3 เรื่อง ประโยคความเดียวแบบซับซ้อน| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาไทย	 ",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_15%2FVideo_1%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.3%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B9%82%E0%B8%A2%E0%B8%84%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B9%80%E0%B8%94%E0%B8%B5%E0%B8%A2%E0%B8%A7%E0%B9%81%E0%B8%9A%E0%B8%9A%E0%B8%8B%E0%B8%B1%E0%B8%9A%E0%B8%8B%E0%B9%89%E0%B8%AD%E0%B8%99.mp4?alt=media&token=1109f420-d596-43eb-b8c5-b46d792101cc",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_15%2FVideo_1%2F%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2%20%E0%B8%A1.3%20%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B9%82%E0%B8%A2%E0%B8%84%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B9%80%E0%B8%94%E0%B8%B5%E0%B8%A2%E0%B8%A7%20%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%A3%E0%B8%A7%E0%B8%A1%20%E0%B9%80%E0%B9%80%E0%B8%A5%E0%B8%B0%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%8B%E0%B9%89%E0%B8%AD%E0%B8%99.pdf?alt=media&token=ba6ca43f-1bf8-4d0f-8f21-5ef8e537752c",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ประโยคความเดียวหมายถึงอะไร",
							Choices: &[]models.Choice{
								{
									Title:   "ประโยคที่มีใจความสำคัญเพียงเรื่องเดียว มีภาคประธานภาคเดียว ภาคแสดงเดียว มีกริยาสำคัญเพียงตัวเดียว",
									Correct: true,
								},
								{
									Title:   "ประโยคความเดียวตั้งแต่ ๒ ประโยคขึ้นไปมารวมกัน โดยมีเนื้อความขัดแย้งกัน กริยาในแต่ละประโยคตรงกันข้ามกัน ส่วนใหญ่จะมีสันธาน",
									Correct: false,
								},
								{
									Title:   "ประโยคที่ประกอบด้วยประโยคความเดียว ๒ ประโยคขึ้นไปซ้อนกันอยู่ โดยมีประโยคหนึ่งเป็นประโยคหลัก",
									Correct: false,
								},
								{
									Title:   "ผิดทุกข้อ",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดไม่ใช่ประโยคความเดียว",
							Choices: &[]models.Choice{
								{
									Title:   "นกสวยบินเร็ว",
									Correct: false,
								},
								{
									Title:   "น้องอ่านหนังสือ",
									Correct: false,
								},
								{
									Title:   "แดงซื้อขนมและดำซื้อขนม",
									Correct: true,
								},
								{
									Title:   "วิชัยเป็นนายอำเภอ",
									Correct: false,
								},
							},
						},
						{
							Question: "3.ข้อใดเป็นประโยคความเดียว",
							Choices: &[]models.Choice{
								{
									Title:   "ฝนตกน้ำจึงท่วม",
									Correct: false,
								},
								{
									Title:   "พ่อชอบดูบอลเเละเเม่ชอบดูละคร",
									Correct: false,
								},
								{
									Title:   "ฉันกินมะม่วง",
									Correct: true,
								},
								{
									Title:   "แดงซื้อขนมและดำซื้อขนม",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 3,
					Video_name:  "ประโยคความรวมแบบซับซ้อน",
					Picture:     "https://i.ytimg.com/vi/6wEluUfEqHk/maxresdefault.jpg",
					Description: "ภาษาไทย ชั้น ม.3 เรื่อง ประโยคความรวมแบบซับซ้อน | แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียนภาษาไทย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_15%2FVideo_2%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.3%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B9%82%E0%B8%A2%E0%B8%84%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%A3%E0%B8%A7%E0%B8%A1%E0%B9%81%E0%B8%9A%E0%B8%9A%E0%B8%8B%E0%B8%B1%E0%B8%9A%E0%B8%8B%E0%B9%89%E0%B8%AD%E0%B8%99.mp4?alt=media&token=0685de0b-0cf9-40c1-94e4-8d83a4355775",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_15%2FVideo_2%2F%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2%20%E0%B8%A1.3%20%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B9%82%E0%B8%A2%E0%B8%84%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B9%80%E0%B8%94%E0%B8%B5%E0%B8%A2%E0%B8%A7%20%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%A3%E0%B8%A7%E0%B8%A1%20%E0%B9%80%E0%B9%80%E0%B8%A5%E0%B8%B0%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%8B%E0%B9%89%E0%B8%AD%E0%B8%99.pdf?alt=media&token=4b9de4c1-ee6d-4638-884b-8bd5573d56ef",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ประโยคความรวมหมายถึงอะไร ",
							Choices: &[]models.Choice{
								{
									Title:   "ประโยคที่บอกเล่าถึงเหตุการณ์ของคนสองคน ณ ช่วงเวลาหนึ่ง",
									Correct: false,
								},
								{
									Title:   "ประโยคที่ประกอบด้วยประโยคความเดียว ๒ ประโยคขึ้นไปซ้อนกันอยู่ โดยมีประโยคหนึ่งเป็นประโยคหลัก",
									Correct: false,
								},
								{
									Title:   "ประโยคที่มีใจความสำคัญเพียงเรื่องเดียว มีภาคประธานภาคเดียว ภาคแสดงเดียว มีกริยาสำคัญเพียงตัวเดียว",
									Correct: false,
								},
								{
									Title:   "ประโยคความเดียวตั้งแต่ ๒ ประโยคขึ้นไปมารวมกัน โดยมีเนื้อความขัดแย้งกัน กริยาในแต่ละประโยคตรงกันข้ามกัน ส่วนใหญ่จะมีสันธาน",
									Correct: true,
								},
							},
						},
						{
							Question: "2.ข้อใดไม่ใช่ประโยคความรวม",
							Choices: &[]models.Choice{
								{
									Title:   "ฉันกินมะม่วง",
									Correct: true,
								},
								{
									Title:   "พงศธรจะเป็นวิศวกรหรือสถาปนิก",
									Correct: false,
								},
								{
									Title:   "เธอต้องเขียนเรียงความนะมิฉะนั้นจะไม่ได้คะแนนเพิ่ม",
									Correct: false,
								},
								{
									Title:   "เขาอ่านหนังสือทุกวันดังนั้นเขาจึงประสบผลสำเร็จ",
									Correct: false,
								},
							},
						},
						{
							Question: "3.ข้อใดเป็นประโยคความรวม",
							Choices: &[]models.Choice{
								{
									Title:   "แดงซื้อขนมและดำซื้อขนม",
									Correct: true,
								},
								{
									Title:   "นกสวยบินเร็ว",
									Correct: false,
								},
								{
									Title:   "น้องอ่านหนังสือ",
									Correct: false,
								},
								{
									Title:   "ฉันกินมะม่วง",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 3,
					Video_name:  "ประโยคความซ้อนแบบซับซ้อน",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/u1UgOYCDOsM/maxresdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอน ประโยคความซ้อนแบบซับซ้อน แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_15%2FVideo_3%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2%20%E0%B8%8A%E0%B8%B1%E0%B9%89%E0%B8%99%20%E0%B8%A1.3%20%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87%20%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B9%82%E0%B8%A2%E0%B8%84%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%8B%E0%B9%89%E0%B8%AD%E0%B8%99%E0%B9%81%E0%B8%9A%E0%B8%9A%E0%B8%8B%E0%B8%B1%E0%B8%9A%E0%B8%8B%E0%B9%89%E0%B8%AD%E0%B8%99.mp4?alt=media&token=93dea44d-08a7-430b-b8a9-96ad493bfc1a",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_15%2FVideo_3%2F%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B9%84%E0%B8%97%E0%B8%A2%20%E0%B8%A1.3%20%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B9%82%E0%B8%A2%E0%B8%84%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B9%80%E0%B8%94%E0%B8%B5%E0%B8%A2%E0%B8%A7%20%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%A3%E0%B8%A7%E0%B8%A1%20%E0%B9%80%E0%B9%80%E0%B8%A5%E0%B8%B0%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%8B%E0%B9%89%E0%B8%AD%E0%B8%99.pdf?alt=media&token=84e0aaf5-91c8-4dd8-baa9-01cbb4068bcc",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ประโยคความซ้อนหมายถึงอะไร",
							Choices: &[]models.Choice{
								{
									Title:   "ประโยคที่ประกอบด้วยประโยคความเดียว ๒ ประโยคขึ้นไปซ้อนกันอยู่ โดยมีประโยคหนึ่งเป็นประโยคหลัก",
									Correct: true,
								},
								{
									Title:   "ประโยคที่มีใจความสำคัญเพียงเรื่องเดียว มีภาคประธานภาคเดียว ภาคแสดงเดียว มีกริยาสำคัญเพียงตัวเดียว",
									Correct: false,
								},
								{
									Title:   "ประโยคความเดียวตั้งแต่ ๒ ประโยคขึ้นไปมารวมกัน โดยมีเนื้อความขัดแย้งกัน กริยาในแต่ละประโยคตรงกันข้ามกัน ส่วนใหญ่จะมีสันธาน",
									Correct: false,
								},
								{
									Title:   "ผิดทุกข้อ",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดเป็นประโยคความซ้อน",
							Choices: &[]models.Choice{
								{
									Title:   "น้องอ่านหนังสือ",
									Correct: false,
								},
								{
									Title:   "ฝนตกจนน้ำท่วม ",
									Correct: true,
								},
								{
									Title:   "แดงซื้อขนมและดำซื้อขนม",
									Correct: false,
								},
								{
									Title:   "พ่อชอบดูบอลเเละเเม่ชอบดูละคร",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   7,
			Course_name: "หรรษาภาษาไทย",
			Picture:     "http://www.kabinburischool.ac.th/wp-content/uploads/2020/05/หรรษาภาษาไทย.jpg",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ ภาษาไทยพื้นฐานในเรื่องต่าง ๆ เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_16
			SubjectID: 36,
			Videos: &[]models.Video{
				{
					Class_level: 7,
					Video_name:  "ลิมิตของฟังก์ชัน#1",
					Picture:     "https://i.ytimg.com/vi/YZYk98uKXCA/maxresdefault.jpg",
					Description: "แคลคูลัสเบื้องต้น เรื่อง ลิมิตของฟังก์ชัน| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_16%2FVideo_1%2F%E0%B9%81%E0%B8%84%E0%B8%A5%E0%B8%84%E0%B8%B9%E0%B8%A5%E0%B8%B1%E0%B8%AA%20EP.1_16%20%E0%B8%A5%E0%B8%B4%E0%B8%A1%E0%B8%B4%E0%B8%95%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%9F%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B9%8C%E0%B8%8A%E0%B8%B1%E0%B8%99.mp4?alt=media&token=8fefcca9-86e8-4403-9f03-b9159243e2a3",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_16%2FVideo_1%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B8%A5%E0%B8%B4%E0%B8%A1%E0%B8%B4%E0%B8%95%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%9F%E0%B8%B1%E0%B8%87%E0%B8%8A%E0%B8%B1%E0%B8%99.pdf?alt=media&token=c552ba91-8c27-4ebe-a3b2-7e91c31a0ebb",
					Exercises: &[]models.Exercise{
						{
							Question: "จงหาค่าลิมิตต่อไปนี้",
							Image:    "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_16%2FVideo_1%2FExercise%2F%E0%B9%80%E0%B9%80%E0%B8%9A%E0%B8%9A%E0%B8%9D%E0%B8%B6%E0%B8%81%E0%B8%AB%E0%B8%B1%E0%B8%94%E0%B8%A5%E0%B8%B4%E0%B8%A1%E0%B8%B4%E0%B8%95(1).png?alt=media&token=b52d3cc1-ca7c-46b0-be66-adaff4a66519",
							Choices: &[]models.Choice{
								{
									Title:   "1/6",
									Correct: true,
								},
								{
									Title:   "1/2",
									Correct: false,
								},
								{
									Title:   "1/8",
									Correct: false,
								},
								{
									Title:   "0",
									Correct: false,
								},
							},
						},
						{
							Question: "จงหาค่าลิมิตต่อไปนี้",
							Image:    "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_16%2FVideo_1%2FExercise%2F%E0%B9%80%E0%B9%80%E0%B8%9A%E0%B8%9A%E0%B8%9D%E0%B8%B6%E0%B8%81%E0%B8%AB%E0%B8%B1%E0%B8%94%E0%B8%A5%E0%B8%B4%E0%B8%A1%E0%B8%B4%E0%B8%95(2).png?alt=media&token=3c779dae-cedd-4498-bc90-e9e430f202a2",
							Choices: &[]models.Choice{
								{
									Title:   "-2",
									Correct: false,
								},
								{
									Title:   "-1",
									Correct: false,
								},
								{
									Title:   "0",
									Correct: true,
								},
								{
									Title:   "1",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 7,
					Video_name:  "ลิมิตของฟังก์ชัน#2",
					Picture:     "https://i.ytimg.com/vi/YsJC7nNngXo/mqdefault.jpg",
					Description: "แคลคูลัสเบื้องต้น เรื่อง ลิมิตของฟังก์ชัน| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_16%2FVideo_2%2F%E0%B9%81%E0%B8%84%E0%B8%A5%E0%B8%84%E0%B8%B9%E0%B8%A5%E0%B8%B1%E0%B8%AA%20EP.2_16%20%E0%B8%A5%E0%B8%B4%E0%B8%A1%E0%B8%B4%E0%B8%95%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%9F%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B9%8C%E0%B8%8A%E0%B8%B1%E0%B8%99.mp4?alt=media&token=57ce8e68-2e00-4c4a-bdbc-d0299fb4d1fb",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_16%2FVideo_2%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B8%A5%E0%B8%B4%E0%B8%A1%E0%B8%B4%E0%B8%95%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%9F%E0%B8%B1%E0%B8%87%E0%B8%8A%E0%B8%B1%E0%B8%99.pdf?alt=media&token=f8434e94-3378-41ae-8b8b-754b6602d8ae",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดเป็นสูตรกำลังสองสมบูรณ์",
							Choices: &[]models.Choice{
								{
									Title:   "(หน้า+หลัง)กำลังสอง = หน้ากำลังสอง + 2หน้าหลัง + หลังกำลังสอง ",
									Correct: true,
								},
								{
									Title:   "หน้ากำลังสอง + หลังกำลังสอง = ( หน้า + หลัง )( หน้า - หลัง )",
									Correct: false,
								},
								{
									Title:   "(หน้า + หลัง)กำลังสาม = หน้ากำลังสาม + 3หน้ากำลังสองหลัง + 3หน้ากำลังสองหลัง + หลังกำลังสาม",
									Correct: false,
								},
								{
									Title:   "ผิดทุกข้อ",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดเป็นสูตรผลต่างกำลังสอง",
							Choices: &[]models.Choice{
								{
									Title:   "หน้ากำลังสอง + หลังกำลังสอง = ( หน้า + หลัง )( หน้า - หลัง )",
									Correct: true,
								},
								{
									Title:   "(หน้า+หลัง)กำลังสอง = หน้ากำลังสอง + 2หน้าหลัง + หลังกำลังสอง",
									Correct: false,
								},
								{
									Title:   "(หน้า + หลัง)กำลังสาม = หน้ากำลังสาม + 3หน้ากำลังสองหลัง + 3หน้ากำลังสองหลัง + หลังกำลังสาม",
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
					Class_level: 7,
					Video_name:  "ลิมิตอนันต์",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/P55D83g6zcM/maxresdefault.jpg",
					Description: "สำหรับคลิปนี้ เป็นคลิปสอน ลิมิตอนันต์ แบบละเอียด เข้าใจง่าย",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_16%2FVideo_3%2F%E0%B9%81%E0%B8%84%E0%B8%A5%E0%B8%84%E0%B8%B9%E0%B8%A5%E0%B8%B1%E0%B8%AA%20EP.3_16%20%E0%B8%A5%E0%B8%B4%E0%B8%A1%E0%B8%B4%E0%B8%95%E0%B8%AD%E0%B8%99%E0%B8%B1%E0%B8%99%E0%B8%95%E0%B9%8C.mp4?alt=media&token=b552787a-4299-42fd-8aac-bf23f65ef0c8",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_16%2FVideo_3%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20%E0%B8%A5%E0%B8%B4%E0%B8%A1%E0%B8%B4%E0%B8%95%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%9F%E0%B8%B1%E0%B8%87%E0%B8%8A%E0%B8%B1%E0%B8%99.pdf?alt=media&token=21ff2d94-6aff-41b8-9725-478d20f57d48",
					Exercises: &[]models.Exercise{
						{
							Question: "จงหาค่าลิมิตต่อไปนี้",
							Image:    "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_16%2FVideo_3%2FExercise%2F%E0%B9%80%E0%B9%80%E0%B8%9A%E0%B8%9A%E0%B8%9D%E0%B8%B6%E0%B8%81%E0%B8%AB%E0%B8%B1%E0%B8%94%E0%B8%A5%E0%B8%B4%E0%B8%A1%E0%B8%B4%E0%B8%95(3).png?alt=media&token=8b15f9ff-7b6b-4df2-b384-372cb52e7c8f",
							Choices: &[]models.Choice{
								{
									Title:   "0",
									Correct: true,
								},
								{
									Title:   "1",
									Correct: false,
								},
								{
									Title:   "2",
									Correct: false,
								},
								{
									Title:   "3",
									Correct: false,
								},
							},
						},
						{
							Question: "จงหาค่าลิมิตต่อไปนี้",
							Image:    "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_16%2FVideo_3%2FExercise%2F%E0%B9%80%E0%B9%80%E0%B8%9A%E0%B8%9A%E0%B8%9D%E0%B8%B6%E0%B8%81%E0%B8%AB%E0%B8%B1%E0%B8%94%E0%B8%A5%E0%B8%B4%E0%B8%A1%E0%B8%B4%E0%B8%95(4).png?alt=media&token=d5a0aba9-d35e-4abe-a7ba-4838c1ad7aa0",
							Choices: &[]models.Choice{
								{
									Title:   "1/6",
									Correct: false,
								},
								{
									Title:   "1/8",
									Correct: true,
								},
								{
									Title:   "0",
									Correct: false,
								},
								{
									Title:   "1/2",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   4,
			Course_name: "แคลคูลัสเบื้องต้น",
			Picture:     "https://www.glurgeek.com/wp-content/uploads/2016/10/stock-vector-calculus-law-theory-and-mathematical-formula-equation-doodle-handwriting-icon-in-blackboard-396938626-e1475310187891.jpg",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ แคลคูลัสเบื้องต้นในเรื่องต่าง ๆ เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_17
			SubjectID: 26,
			Videos: &[]models.Video{
				{
					Class_level: 7,
					Video_name:  "เรียนมหาวิทยาลัย ไม่รู้ศัพท์พวกนี้ไม่ได้",
					Picture:     "https://i.ytimg.com/vi/RtUlmytHcEA/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ เรื่อง คำศัพท์ที่ควรรู้| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_17%2FVideo_1%2F%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%E0%B8%A1%E0%B8%AB%E0%B8%B2%E0%B8%A7%E0%B8%B4%E0%B8%97%E0%B8%A2%E0%B8%B2%E0%B8%A5%E0%B8%B1%E0%B8%A2%20%E0%B9%84%E0%B8%A1%E0%B9%88%E0%B8%A3%E0%B8%B9%E0%B9%89%E0%B8%A8%E0%B8%B1%E0%B8%9E%E0%B8%97%E0%B9%8C%E0%B8%9E%E0%B8%A7%E0%B8%81%E0%B8%99%E0%B8%B5%E0%B9%89%E0%B9%84%E0%B8%A1%E0%B9%88%E0%B9%84%E0%B8%94%E0%B9%89.mp4?alt=media&token=815a53c3-1049-4dcb-b243-010c1d3dabeb",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_17%2FVideo_1%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B8%84%E0%B8%B3%E0%B8%A8%E0%B8%B1%E0%B8%9E%E0%B8%97%E0%B9%8C%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%AD%E0%B8%B1%E0%B8%87%E0%B8%81%E0%B8%A4%E0%B8%A9.pdf?alt=media&token=becf7d82-1770-46ae-b52c-8323c73c4ccd",
					Exercises: &[]models.Exercise{
						{
							Question: "ข้อใดคือความหมายของคำว่า Section",
							Image:    "",
							Choices: &[]models.Choice{
								{
									Title:   "กลุ่ม",
									Correct: true,
								},
								{
									Title:   "ถ่ายรูป",
									Correct: false,
								},
								{
									Title:   "ห้องพัก",
									Correct: false,
								},
								{
									Title:   "การเเสดง",
									Correct: false,
								},
							},
						},
						{
							Question: "ข้อใดคือความหมายของคำว่า Tuition",
							Choices: &[]models.Choice{
								{
									Title:   "การนอน",
									Correct: false,
								},
								{
									Title:   "การเรียน",
									Correct: false,
								},
								{
									Title:   "การสอน การอบรม",
									Correct: true,
								},
								{
									Title:   "การงาน",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 7,
					Video_name:  "Past simple VS Past continuous ใช้กับอดีตทั้งคู่ แต่ใช้ต่างกันยังไง",
					Picture:     "https://i.ytimg.com/vi/RHMOjM_IMnQ/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ เรื่อง Past simple VS Past continuous ใช้กับอดีตทั้งคู่ แต่ใช้ต่างกันยังไง| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_17%2FVideo_2%2FPast%20simple%20VS%20Past%20continuous%20%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%81%E0%B8%B1%E0%B8%9A%E0%B8%AD%E0%B8%94%E0%B8%B5%E0%B8%95%E0%B8%97%E0%B8%B1%E0%B9%89%E0%B8%87%E0%B8%84%E0%B8%B9%E0%B9%88%20%E0%B9%81%E0%B8%95%E0%B9%88%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%95%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B8%A2%E0%B8%B1%E0%B8%87%E0%B9%84%E0%B8%87.mp4?alt=media&token=68f354db-f7f2-4653-b8be-9f5262a6b4b3",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_17%2FVideo_2%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20Past%20simple%20VS%20Past%20continuous%20%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%81%E0%B8%B1%E0%B8%9A%E0%B8%AD%E0%B8%94%E0%B8%B5%E0%B8%95%E0%B8%97%E0%B8%B1%E0%B9%89%E0%B8%87%E0%B8%84%E0%B8%B9%E0%B9%88%20%E0%B9%81%E0%B8%95%E0%B9%88%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%95%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B8%A2%E0%B8%B1%E0%B8%87%E0%B9%84%E0%B8%87.pdf?alt=media&token=7bf511a5-1553-470e-8325-6572dcb475b8",
					Exercises: &[]models.Exercise{
						{
							Question: "1. Choose the correct answer. I ______ for exams. (Past continuous tense)",
							Choices: &[]models.Choice{
								{
									Title:   "was studying",
									Correct: true,
								},
								{
									Title:   "studied",
									Correct: false,
								},
								{
									Title:   "was study",
									Correct: false,
								},
								{
									Title:   "study",
									Correct: false,
								},
							},
						},
						{
							Question: "2. Choose the correct answer. She _____ her room. (Past simple tense)",
							Choices: &[]models.Choice{
								{
									Title:   "cleaned",
									Correct: true,
								},
								{
									Title:   "clean",
									Correct: false,
								},
								{
									Title:   "was cleaning",
									Correct: false,
								},
								{
									Title:   "was clean",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 7,
					Video_name:  "Present simple / Present continuous ใช้ต่างกันยังไง",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/wQ0bpzrBz3k/maxresdefault.jpg",
					Description: "ภาษาอังกฤษ เรื่อง Present simple / Present continuous ใช้ต่างกันยังไง | แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_17%2FVideo_3%2FPresent%20simple%20_%20Present%20continuous%20%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%95%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B8%A2%E0%B8%B1%E0%B8%87%E0%B9%84%E0%B8%87.mp4?alt=media&token=df2b476b-68c2-4bbd-84ba-930c6c9c8c0c",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_17%2FVideo_3%2F%E0%B8%8A%E0%B8%B5%E0%B8%97%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%20Present%20simple%20%E0%B9%80%E0%B9%80%E0%B8%A5%E0%B8%B0%20Present%20continuous%20%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%95%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B8%A2%E0%B8%B1%E0%B8%87%E0%B9%84%E0%B8%87.pdf?alt=media&token=9bf21cf1-97df-447e-929a-dd7d9a1b57a7",
					Exercises: &[]models.Exercise{
						{
							Question: "1.Choose the correct answer. I _____ the house now. I'll be there in 20 minutes.",
							Choices: &[]models.Choice{
								{
									Title:   "am leaving",
									Correct: true,
								},
								{
									Title:   "leave",
									Correct: false,
								},
								{
									Title:   "leaving",
									Correct: false,
								},
								{
									Title:   "ผิดทุกข้อ",
									Correct: false,
								},
							},
						},
						{
							Question: "2.Choose the correct answer. We never _____ while on vacation.",
							Choices: &[]models.Choice{
								{
									Title:   "are working",
									Correct: false,
								},
								{
									Title:   "work",
									Correct: true,
								},
								{
									Title:   "are work",
									Correct: false,
								},
								{
									Title:   "working",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   10,
			Course_name: "ภาษาอังกฤษสนุ๊กสนุก",
			Picture:     "https://www.youngciety.com/cms-2016-storage/article/header/hero_English-is-Fun-mb.jpg",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ ภาษาอังกฤษในเรื่องต่าง ๆ เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_18
			SubjectID: 36,
			Videos: &[]models.Video{
				{
					Class_level: 7,
					Video_name:  "ปรับพื้นฐานสำหรับน้อง ป.ตรี | ฟิสิกส์ - แนวราบ 1/2",
					Picture:     "https://i.ytimg.com/vi/s_AA8oYRfk0/maxresdefault.jpg",
					Description: "ฟิสิกส์ เรื่อง ฟิสิกส์ - แนวราบ| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_18%2FVideo_1%2F%E0%B8%9B%E0%B8%A3%E0%B8%B1%E0%B8%9A%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%90%E0%B8%B2%E0%B8%99%E0%B8%AA%E0%B8%B3%E0%B8%AB%E0%B8%A3%E0%B8%B1%E0%B8%9A%E0%B8%99%E0%B9%89%E0%B8%AD%E0%B8%87%20%E0%B8%9B.%E0%B8%95%E0%B8%A3%E0%B8%B5%20_%20%E0%B8%9F%E0%B8%B4%E0%B8%AA%E0%B8%B4%E0%B8%81%E0%B8%AA%E0%B9%8C%20-%20%E0%B9%81%E0%B8%99%E0%B8%A7%E0%B8%A3%E0%B8%B2%E0%B8%9A%201_2.mp4?alt=media&token=58643536-64b3-450e-a0c1-54ecfa966fa1",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_18%2FVideo_1%2F%E0%B8%9B%E0%B8%A3%E0%B8%B1%E0%B8%9A%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%20physic%20by%20SigmaKing.pdf?alt=media&token=e12a2e66-a8a9-40af-ac54-936bd3384e7f",
					Exercises: &[]models.Exercise{
						{
							Question: "ถนนเส้นหนึ่งมีความยาวของเส้นทาง(ระยะทาง)ที่วัตถุเคลื่อนที่จาก A ไป B เป็น 120 เมตร ใช้เวลา 25 วินาที จงหาอัตราเร็วการเคลื่อนที่ของวัตถุจากระยะ A ไป B เป็นเท่าใด",
							Choices: &[]models.Choice{
								{
									Title:   "4.8 m/s",
									Correct: true,
								},
								{
									Title:   "3.8 m/s",
									Correct: false,
								},
								{
									Title:   "2.8 m/s",
									Correct: false,
								},
								{
									Title:   "1.8 m/s",
									Correct: false,
								},
							},
						},
						{
							Question: "จากข้อที่เเล้ว หากวัตถุเคลื่อนที่จาก B ไป C เป็น 80 เมตร ใช้เวลา 35 วินาที จงหาอัตราเร็วการเคลื่อนที่ของวัตถุจากระยะ B ไป C เป็นเท่าใด",
							Choices: &[]models.Choice{
								{
									Title:   "2.1 m/s",
									Correct: false,
								},
								{
									Title:   "2.2 m/s",
									Correct: false,
								},
								{
									Title:   "2.3 m/s",
									Correct: true,
								},
								{
									Title:   "2.4 m/s",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 7,
					Video_name:  "ปรับพื้นฐานสำหรับน้อง ป.ตรี | ฟิสิกส์ - แนวราบ 2/2",
					Picture:     "https://i.ytimg.com/vi/CEeybQsNk4g/maxresdefault.jpg",
					Description: "ฟิสิกส์ เรื่อง ฟิสิกส์ - แนวราบ| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_18%2FVideo_2%2F%E0%B8%9B%E0%B8%A3%E0%B8%B1%E0%B8%9A%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%90%E0%B8%B2%E0%B8%99%E0%B8%AA%E0%B8%B3%E0%B8%AB%E0%B8%A3%E0%B8%B1%E0%B8%9A%E0%B8%99%E0%B9%89%E0%B8%AD%E0%B8%87%20%E0%B8%9B.%E0%B8%95%E0%B8%A3%E0%B8%B5%20_%20%E0%B8%9F%E0%B8%B4%E0%B8%AA%E0%B8%B4%E0%B8%81%E0%B8%AA%E0%B9%8C%20-%20%E0%B9%81%E0%B8%99%E0%B8%A7%E0%B8%A3%E0%B8%B2%E0%B8%9A%202_2.mp4?alt=media&token=af689be7-3cad-4dec-b76f-d26c6666b4af",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_18%2FVideo_2%2F%E0%B8%9B%E0%B8%A3%E0%B8%B1%E0%B8%9A%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%20physic%20by%20SigmaKing.pdf?alt=media&token=354858df-ce77-4ab3-bf8e-9ac413ff3498",
					Exercises: &[]models.Exercise{
						{
							Question: "1.รถยนต์คันหนึ่งขณะเริ่มสังเกตการเคลื่อนที่มีความเร็ว 30 เมตรต่อวินาที เมื่อเวลาผ่านไป 20 วินาที มีความเร็วเป็น 40 เมตรต่อวินาที หลังจากนั้นอีก 15 วินาที รถยนต์คันนั้นจะหยุดการเคลื่อนที่พอดี จงหาความเร่งในช่วง 20 วินาทีแรก",
							Choices: &[]models.Choice{
								{
									Title:   "0.5 m/s กำลังสอง",
									Correct: true,
								},
								{
									Title:   "0.6 m/s กำลังสอง",
									Correct: false,
								},
								{
									Title:   "0.7 m/s กำลังสอง",
									Correct: false,
								},
								{
									Title:   "0.8 m/s กำลังสอง",
									Correct: false,
								},
							},
						},
						{
							Question: "2.รถยนต์คันหนึ่งขณะเริ่มสังเกตการเคลื่อนที่มีความเร็ว 30 เมตรต่อวินาที เมื่อเวลาผ่านไป 20 วินาที มีความเร็วเป็น 40 เมตรต่อวินาที หลังจากนั้นอีก 15 วินาที รถยนต์คันนั้นจะหยุดการเคลื่อนที่พอดี จงหาความเร่งในช่วง 15 วินาทีหลัง",
							Choices: &[]models.Choice{
								{
									Title:   "-2.67 m/s กำลังสอง",
									Correct: true,
								},
								{
									Title:   "-3.67 m/s กำลังสอง",
									Correct: false,
								},
								{
									Title:   "-1.67 m/s กำลังสอง",
									Correct: false,
								},
								{
									Title:   "-4.67 m/s กำลังสอง",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 7,
					Video_name:  "ปรับพื้นฐานสำหรับน้อง ป.ตรี | ฟิสิกส์ - แนวดิ่ง",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/yXSM5ShPshE/maxresdefault.jpg",
					Description: "ฟิสิกส์ เรื่อง ฟิสิกส์ - แนวดิ่ง| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_18%2FVideo_3%2F%E0%B8%9B%E0%B8%A3%E0%B8%B1%E0%B8%9A%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%90%E0%B8%B2%E0%B8%99%E0%B8%AA%E0%B8%B3%E0%B8%AB%E0%B8%A3%E0%B8%B1%E0%B8%9A%E0%B8%99%E0%B9%89%E0%B8%AD%E0%B8%87%20%E0%B8%9B.%E0%B8%95%E0%B8%A3%E0%B8%B5%20_%20%E0%B8%9F%E0%B8%B4%E0%B8%AA%E0%B8%B4%E0%B8%81%E0%B8%AA%E0%B9%8C%20-%20%E0%B9%81%E0%B8%99%E0%B8%A7%E0%B8%94%E0%B8%B4%E0%B9%88%E0%B8%87.mp4?alt=media&token=befed1c9-352b-4db8-80bf-de378fbfd38e",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_18%2FVideo_3%2F%E0%B8%9B%E0%B8%A3%E0%B8%B1%E0%B8%9A%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%20physic%20by%20SigmaKing.pdf?alt=media&token=9a01a32e-96b3-4356-9e82-06620ba530ec",
					Exercises: &[]models.Exercise{
						{
							Question: "โยนลูกบอลขึ้นไปในเเนวดิ่งด้วยความเร็วต้น 4.9 เมตรต่อวินาที นานเท่าใดลูกบอลจึงจะเคลื่อนที่ไปถึงจุดสูงสุด",
							Choices: &[]models.Choice{
								{
									Title:   "0.5 s",
									Correct: true,
								},
								{
									Title:   "0.6 s",
									Correct: false,
								},
								{
									Title:   "0.7 s",
									Correct: false,
								},
								{
									Title:   "0.8 s",
									Correct: false,
								},
							},
						},
						{
							Question: "ยิงวัตถุขึ้นในเเนวดิ่ง จากพื้นด้วยความเร็ว 60 เมตร/วินาที นานเท่าใดวัตถุจึงอยู่สูงจากพื้น 100 เมตร ( g=10 เมตร/วินาทีกำลังสอง)",
							Choices: &[]models.Choice{
								{
									Title:   "2, 17",
									Correct: false,
								},
								{
									Title:   "2, 10",
									Correct: true,
								},
								{
									Title:   "20, 14",
									Correct: false,
								},
								{
									Title:   "2, 15",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   3,
			Course_name: "ปรับพื้นฐานวิชาฟิสิกส์",
			Picture:     "https://panyasociety.com/pages/wp-content/uploads/2020/10/PhysicCover-01_New.png",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับ ฟิสิกส์ในเรื่องต่าง ๆ เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_19
			SubjectID: 30,
			Videos: &[]models.Video{
				{
					Class_level: 4,
					Video_name:  "ความรู้พื้นฐานเกี่ยวกับกฎหมาย",
					Picture:     "https://www.guamattorneygeneral.com/wp-content/uploads/2021/07/14.jpg",
					Description: "สังคมศึกษา เรื่อง ความรู้พื้นฐานเกี่ยวกับกฎหมาย| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_19%2FVideo_1%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%AA%E0%B8%B1%E0%B8%87%E0%B8%84%E0%B8%A1%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2%20_%20%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%A3%E0%B8%B9%E0%B9%89%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%90%E0%B8%B2%E0%B8%99%E0%B9%80%E0%B8%81%E0%B8%B5%E0%B9%88%E0%B8%A2%E0%B8%A7%E0%B8%81%E0%B8%B1%E0%B8%9A%E0%B8%81%E0%B8%8E%E0%B8%AB%E0%B8%A1%E0%B8%B2%E0%B8%A2.mp4?alt=media&token=88e29ed4-2ad7-4a0d-a232-7697aac2efa9",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_19%2FVideo_1%2F%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%A3%E0%B8%B9%E0%B9%89%E0%B9%80%E0%B8%9A%E0%B8%B7%E0%B9%89%E0%B8%AD%E0%B8%87%E0%B8%95%E0%B9%89%E0%B8%99%E0%B9%80%E0%B8%81%E0%B8%B5%E0%B9%88%E0%B8%A2%E0%B8%A7%E0%B8%81%E0%B8%B1%E0%B8%9A%E0%B8%81%E0%B8%8E%E0%B8%AB%E0%B8%A1%E0%B8%B2%E0%B8%A2.pdf?alt=media&token=109ffe57-da89-4cf5-9462-61675ff7f5a8",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดกล่าวถึงความหมายของ กฎหมาย ได้ถูกต้องที่สุด",
							Choices: &[]models.Choice{
								{
									Title:   "เป็นข้อบังคับของรัฐที่ใช้ควบคุมพลเมือง",
									Correct: true,
								},
								{
									Title:   "เป็นกฎเกณฑ์ที่รัฐใช้ควบคุมพลเมือง",
									Correct: false,
								},
								{
									Title:   "เป็นคำสั่งของผู้มีอำนาจสูงสุดในรัฐ",
									Correct: false,
								},
								{
									Title:   "เป็นกฎเกณฑ์ข้อบังคับในสังคม",
									Correct: false,
								},
							},
						},
						{
							Question: "2.กฎหมายใดอยู่ในกฎหมายเอกชน",
							Choices: &[]models.Choice{
								{
									Title:   "พระธรรมศาลยุติธรรม",
									Correct: false,
								},
								{
									Title:   "ประมวลกฎหมายอาญา",
									Correct: false,
								},
								{
									Title:   "ประมวลกฎหมายเเพ่งเเละพาณิชย์",
									Correct: true,
								},
								{
									Title:   "ประมวลกฎหมายวิธีพิจารณาความอาญา",
									Correct: false,
								},
							},
						},
						{
							Question: "3.กฎหมายใดที่เกี่ยวข้องกับชีวิตประจำวันมากที่สุด",
							Choices: &[]models.Choice{
								{
									Title:   "กฎหมายรัฐธรรมนูญ",
									Correct: false,
								},
								{
									Title:   "ประมวลกฎหมายอาญา",
									Correct: false,
								},
								{
									Title:   "กฎหมายเเพ่งเเละกฎหมายอาญา",
									Correct: true,
								},
								{
									Title:   "กฎหมายเเพ่ง",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 4,
					Video_name:  "ระบอบการปกครองและรูปแบบของรัฐ",
					Picture:     "https://thaipublica.org/wp-content/uploads/2020/12/thaipublica-%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%9B%E0%B8%A3%E0%B8%B1%E0%B8%9A%E0%B8%95%E0%B8%B1%E0%B8%A7%E0%B8%97%E0%B8%A8%E0%B8%A7%E0%B8%A3%E0%B8%A3%E0%B8%A921-2-scaled.jpeg",
					Description: "สังคมศึกษา เรื่อง ระบอบการปกครองและรูปแบบของรัฐ| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_19%2FVideo_2%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%AA%E0%B8%B1%E0%B8%87%E0%B8%84%E0%B8%A1%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2%20_%20%E0%B8%A3%E0%B8%B0%E0%B8%9A%E0%B8%AD%E0%B8%9A%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%9B%E0%B8%81%E0%B8%84%E0%B8%A3%E0%B8%AD%E0%B8%87%E0%B9%81%E0%B8%A5%E0%B8%B0%E0%B8%A3%E0%B8%B9%E0%B8%9B%E0%B9%81%E0%B8%9A%E0%B8%9A%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B8%A3%E0%B8%B1%E0%B8%90.mp4?alt=media&token=ae9c9e5b-b465-445b-88dc-ca57bedd9ba0",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_19%2FVideo_2%2F%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%80%E0%B8%A1%E0%B8%B7%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%9B%E0%B8%81%E0%B8%84%E0%B8%A3%E0%B8%AD%E0%B8%87%E0%B9%84%E0%B8%97%E0%B8%A2.pdf?alt=media&token=a3828d12-a1d3-4fa5-809d-a45f7a81d3d8",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดจับคู่ระบบการเมืองการปกครองกับข้อด้อยของระบบไม่ถูกต้อง",
							Choices: &[]models.Choice{
								{
									Title:   "เสรีนิยม กับ ปัญหาขาดเสถียรภาพทางการเมือง",
									Correct: false,
								},
								{
									Title:   "ประชาธิปไตย กับ ปัญหาเสียงข้างมากไม่รับฟังเสียงข้างน้อย",
									Correct: false,
								},
								{
									Title:   "อำนาจนิยม กับ ปัญหาการขาดระเบียบและกติกาในสังคม",
									Correct: true,
								},
								{
									Title:   "สังคมนิยม กับ ปัญหาการลิดรอนเสรีภาพทางเศรษฐกิจ",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ระบอบเสรีประชาธิปไตยมีลักษณะสำคัญดังต่อไปนี้ ยกเว้นข้อใด",
							Choices: &[]models.Choice{
								{
									Title:   "กองทัพมีบทบาทสำคัญในการประกันเสถียรภาพของรัฐบาลที่มาจากการเลือกตั้ง",
									Correct: true,
								},
								{
									Title:   "ลดอำนาจรัฐ เพิ่มอำนาจและการมีส่วนร่วมของประชาชน",
									Correct: false,
								},
								{
									Title:   "ให้เสรีภาพพลเมืองและสิทธิพลเมืองในระดับสูง",
									Correct: false,
								},
								{
									Title:   "ใช้กฎหมายเป็นเครื่องมือสำคัญในการปกครอง",
									Correct: false,
								},
							},
						},
						{
							Question: "3.ข้อใดสอดคล้องกับลักษณะการปกครองระบอบเผด็จการมากที่สุด",
							Choices: &[]models.Choice{
								{
									Title:   "รัฐบาลต้องมีอำนาจมั่นคง",
									Correct: true,
								},
								{
									Title:   "พลเมืองต้องมีหน้าที่ต่อรัฐ",
									Correct: false,
								},
								{
									Title:   "ประชาชนต้องเท่าเทียม",
									Correct: false,
								},
								{
									Title:   "บ้านเมืองต้องสงบเรียบร้อย",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 4,
					Video_name:  "สิทธิมนุษยชนคืออะไร",
					Price:       25,
					Picture:     "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT9SbSLJxDWHuA1ZSb5K2YYRC_VATS22AQZhg&usqp=CAU",
					Description: "สังคมศึกษา เรื่อง สิทธิมนุษยชน| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_19%2FVideo_3%2F%E0%B8%A7%E0%B8%B4%E0%B8%8A%E0%B8%B2%E0%B8%AA%E0%B8%B1%E0%B8%87%E0%B8%84%E0%B8%A1%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2%20_%20%E0%B8%AA%E0%B8%B4%E0%B8%97%E0%B8%98%E0%B8%B4%E0%B8%A1%E0%B8%99%E0%B8%B8%E0%B8%A9%E0%B8%A2%E0%B8%8A%E0%B8%99%E0%B8%84%E0%B8%B7%E0%B8%AD%E0%B8%AD%E0%B8%B0%E0%B9%84%E0%B8%A3.mp4?alt=media&token=cd103a89-7f60-4fc7-8e9a-485fd69124d8",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_19%2FVideo_3%2F%E0%B8%AA%E0%B8%B4%E0%B8%97%E0%B8%98%E0%B8%B4%E0%B8%A1%E0%B8%99%E0%B8%B8%E0%B8%A9%E0%B8%A2%E0%B8%8A%E0%B8%99.pdf?alt=media&token=59808539-f4e8-451b-93e1-efb2167555f9",
					Exercises: &[]models.Exercise{
						{
							Question: "1.สิทธิมนุษยชน มีความหมายว่าอย่างไร",
							Choices: &[]models.Choice{
								{
									Title:   "การละเมิดสิทธิมนุษยชน",
									Correct: false,
								},
								{
									Title:   "ศักดิ์ศรีความเป็นมนุษย์ สิทธิเสรีภาพ และความเสมอภาค",
									Correct: true,
								},
								{
									Title:   "สิทธิในการเรียกร้องผลประโยชน์",
									Correct: false,
								},
								{
									Title:   "การให้ความเคารพบุคคลที่มีอำนาจ",
									Correct: false,
								},
							},
						},
						{
							Question: "2.สิทธิที่ทุกคนต้องมี เป็นรากฐานแนวคิดสิทธิมนุษยชนแบบใด",
							Choices: &[]models.Choice{
								{
									Title:   "สิทธิธรรมชาติ",
									Correct: true,
								},
								{
									Title:   "สิทธิทางกฎหมาย",
									Correct: false,
								},
								{
									Title:   "กฎหมายธรรมชาติ",
									Correct: false,
								},
								{
									Title:   "เสรีภาพในการดำรงชีวิต",
									Correct: false,
								},
							},
						},
						{
							Question: "3.องค์กรใดที่ประกาศใช้ปฎิญญาสากลว่าด้วยสิทธิมนุษยชน",
							Choices: &[]models.Choice{
								{
									Title:   "องค์กรยูเนสโก",
									Correct: false,
								},
								{
									Title:   "องค์กรสหประชาชาติ",
									Correct: true,
								},
								{
									Title:   "องค์กรแอมเนสตี",
									Correct: false,
								},
								{
									Title:   "องค์กร Unicef",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   2,
			Course_name: "สังคมเเสนสนุก",
			Picture:     "https://panyasociety.com/pages/wp-content/uploads/2020/10/Social_Cover-01_New.png",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับเนื้อหาของวิชาสังคมในเรื่องต่าง ๆ เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_20
			SubjectID: 38,
			Videos: &[]models.Video{
				{
					Class_level: 5,
					Video_name:  "แก๊สและสมบัติของแก๊ส ",
					Picture:     "https://i.ytimg.com/vi/-gsi2UhMxpE/maxresdefault.jpg",
					Description: "สังคมศึกษา เรื่อง แก๊สและสมบัติของแก๊ส | แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_20%2FVideo_1%2F%E0%B9%81%E0%B8%81%E0%B9%8A%E0%B8%AA%E0%B9%81%E0%B8%A5%E0%B8%B0%E0%B8%AA%E0%B8%A1%E0%B8%9A%E0%B8%B1%E0%B8%95%E0%B8%B4%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B9%81%E0%B8%81%E0%B9%8A%E0%B8%AA%20%E0%B9%80%E0%B8%84%E0%B8%A1%E0%B8%B5%20%E0%B8%A1.5%20_%E0%B8%89%E0%B8%9A%E0%B8%B1%E0%B8%9A%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%E0%B9%80%E0%B8%95%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%A1%E0%B8%AA%E0%B8%AD%E0%B8%9A%20%E0%B9%80%E0%B8%84%E0%B8%A1%E0%B8%B5%20%E0%B8%A1.5%20%E0%B9%80%E0%B8%97%E0%B8%AD%E0%B8%A1%201.mp4?alt=media&token=559d6491-1b39-47e2-98cd-b4d58c52ec6c",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_20%2FVideo_1%2F%E0%B9%81%E0%B8%81%E0%B9%8A%E0%B8%AA%E0%B9%81%E0%B8%A5%E0%B8%B0%E0%B8%AA%E0%B8%A1%E0%B8%9A%E0%B8%B1%E0%B8%95%E0%B8%B4%E0%B8%82%E0%B8%AD%E0%B8%87%E0%B9%81%E0%B8%81%E0%B9%8A%E0%B8%AA.pdf?alt=media&token=fc23e546-f3f8-4623-9df2-8e497a419bbb",
					Exercises: &[]models.Exercise{
						{
							Question: "1.เมื่ออุณหภูมิและจำนวนโมลของแก๊สคงที่ ปริมาตรของแก๊สจะแปรผกผันกับความดัน เป็นไปตามกฎของใคร",
							Choices: &[]models.Choice{
								{
									Title:   "กฎของบอยล์",
									Correct: true,
								},
								{
									Title:   "กฎของเกย์-ลูสแซก",
									Correct: false,
								},
								{
									Title:   "กฎของชาร์ล",
									Correct: false,
								},
								{
									Title:   "กฎของอาโวกาโดร",
									Correct: false,
								},
							},
						},
						{
							Question: "2.เมื่อใดที่แก๊สจริงจะประพฤติตัวเหมือนแก๊สในอุดมคติ",
							Choices: &[]models.Choice{
								{
									Title:   "อุณหภูมิต่ำ ความดันต่ำ",
									Correct: false,
								},
								{
									Title:   "อุณหภูมิสูง ความดันต่ำ",
									Correct: true,
								},
								{
									Title:   "อุณหภูมิต่ำ ความดันสูง",
									Correct: false,
								},
								{
									Title:   "อุณหภูมิสูง ความดันสูง",
									Correct: false,
								},
							},
						},
						{
							Question: "3.ข้อสรุปใดเกี่ยวกับแก๊สถูกต้อง",
							Choices: &[]models.Choice{
								{
									Title:   "เมื่อ T และ n คงที่ เมื่อ V สูงขึ้น P จะสูงขึ้น",
									Correct: false,
								},
								{
									Title:   "เมื่อ V และ n คงที่ เมื่อ T สูงขึ้น P จะลดลง",
									Correct: false,
								},
								{
									Title:   "เมื่อ T และ n คงที่ เมื่อ V ต่ำลง P จะต่ำลง",
									Correct: false,
								},
								{
									Title:   "เมื่อ V และ n คงที่ เมื่อ T สูงขึ้น P จะสูงขึ้น",
									Correct: true,
								},
							},
						},
					},
				},
				{
					Class_level: 5,
					Video_name:  "สารละลายกรดและเบส ทฤษฎีกรด-เบส",
					Picture:     "https://i.ytimg.com/vi/GMemgCvBQMw/maxresdefault.jpg",
					Description: "เคมี เรื่อง สารละลายกรดและเบส| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_20%2FVideo_2%2F%E0%B8%81%E0%B8%A3%E0%B8%94%20%E0%B9%80%E0%B8%9A%E0%B8%AA%20%E0%B9%80%E0%B8%84%E0%B8%A1%E0%B8%B5%20%E0%B8%A1.5%20_%20%E0%B8%97%E0%B8%A4%E0%B8%A9%E0%B8%8E%E0%B8%B5%E0%B8%81%E0%B8%A3%E0%B8%94-%E0%B9%80%E0%B8%9A%E0%B8%AA.mp4?alt=media&token=8546f85a-b467-4728-9ca6-d564f35d4470",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_20%2FVideo_2%2F%E0%B8%81%E0%B8%A3%E0%B8%94%E0%B9%80%E0%B8%9A%E0%B8%AA%20%E0%B9%80%E0%B8%84%E0%B8%A1%E0%B8%B5%E0%B8%A1.5.pdf?alt=media&token=0679815f-3406-4abc-a890-85b396a282b8",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ข้อใดไม่ใช่ สมบัติของสารละลายที่เป็นกรด",
							Choices: &[]models.Choice{
								{
									Title:   "เป็นสารละลายกรดทำปฏิกิริยากับโลหะ",
									Correct: false,
								},
								{
									Title:   "เป็นสารละลายกรดทำปฏิกิริยากับหินปูน",
									Correct: false,
								},
								{
									Title:   "เปลี่ยนสีกระดาษลิตมัสจากสีแดงเป็นสีน้ำเงิน",
									Correct: true,
								},
								{
									Title:   "เป็นสารละลายกรดมีรสเปรี้ยว",
									Correct: false,
								},
							},
						},
						{
							Question: "2.สารละลายชนิดหนึ่งสามารถเปลี่ยนสีของกระดาษลิตมัสจากสีแดงเป็นสีน้ำเงิน แสดงว่าสารละลายนั้นมีสมบัติตามข้อใด",
							Choices: &[]models.Choice{
								{
									Title:   "เป็นกรด",
									Correct: false,
								},
								{
									Title:   "เป็นกลาง",
									Correct: false,
								},
								{
									Title:   "เป็นเบส",
									Correct: true,
								},
								{
									Title:   "ยังสรุปไม่ได้",
									Correct: false,
								},
							},
						},
						{
							Question: "3.ข้อใดต่อไปนี้มีสมบัติต่างจากพวก",
							Choices: &[]models.Choice{
								{
									Title:   "น้ำยาล้างห้องน้ำ",
									Correct: false,
								},
								{
									Title:   "น้ำมะขาม",
									Correct: false,
								},
								{
									Title:   "น้ำส้มสายชู",
									Correct: false,
								},
								{
									Title:   "น้ำสบู่",
									Correct: true,
								},
							},
						},
					},
				},
				{
					Class_level: 5,
					Video_name:  "อัตราการเกิดปฏิกิริยาเคมี",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/YlYxyymXhlU/mqdefault.jpg",
					Description: "เคมี เรื่อง อัตราการเกิดปฏิกิริยาเคมี| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_20%2FVideo_3%2F%E0%B8%AD%E0%B8%B1%E0%B8%95%E0%B8%A3%E0%B8%B2%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%80%E0%B8%81%E0%B8%B4%E0%B8%94%E0%B8%9B%E0%B8%8F%E0%B8%B4%E0%B8%81%E0%B8%B4%E0%B8%A3%E0%B8%B4%E0%B8%A2%E0%B8%B2%E0%B9%80%E0%B8%84%E0%B8%A1%E0%B8%B5%20%E0%B9%80%E0%B8%84%E0%B8%A1%E0%B8%B5%20%E0%B8%A1.5%20ep.1.mp4?alt=media&token=4a4960f8-b2ac-4ccc-bba3-0e233564947d",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_20%2FVideo_3%2F%E0%B8%AD%E0%B8%B1%E0%B8%95%E0%B8%A3%E0%B8%B2%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%80%E0%B8%81%E0%B8%B4%E0%B8%94%E0%B8%9B%E0%B8%8F%E0%B8%B4%E0%B8%81%E0%B8%B4%E0%B8%A3%E0%B8%B4%E0%B8%A2%E0%B8%B2%E0%B9%80%E0%B8%84%E0%B8%A1%E0%B8%B5.pdf?alt=media&token=7e9a95d0-2b47-475a-ac63-a2ea1378e5ce",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ทฤษฎีใดที่สามารถใช้อธิบายความแตกต่างของอัตราการเกิดปฏิกิริยาเคมี",
							Choices: &[]models.Choice{
								{
									Title:   "ทฤษฎีจลน์ของแก๊ส (kinetic theory of gases)",
									Correct: false,
								},
								{
									Title:   "ทฤษฎีสนามผลึก (Crystal Field Theory)",
									Correct: false,
								},
								{
									Title:   "ทฤษฎีออร์บิทัลเชิงโมเลกุล (molecular orbital)",
									Correct: false,
								},
								{
									Title:   "ทฤษฎีการชน (Collision Theory)",
									Correct: true,
								},
							},
						},
						{
							Question: "2.อัตราการเกิดปฏิกิริยาเคมีเป็นสัดส่วนโดยตรงกับอะไร",
							Choices: &[]models.Choice{
								{
									Title:   "เปอร์เซ็นต์ในการชนกันของอนุภาคที่เป็นผลสำเร็จ",
									Correct: true,
								},
								{
									Title:   "ความถี่ในการชนกันของอนุภาค",
									Correct: true,
								},
								{
									Title:   "อุณหภูมิ",
									Correct: false,
								},
								{
									Title:   "พลังงานศักย์",
									Correct: false,
								},
							},
						},
						{
							Question: "3.พลังงานก่อกัมมันต์ (Activation Energy) คืออะไร",
							Choices: &[]models.Choice{
								{
									Title:   "พลังงานที่น้อยที่สุดที่จะทำให้อิเล็กตรอนหลุดออก",
									Correct: false,
								},
								{
									Title:   "พลังงานที่น้อยที่สุดที่จะทำให้การชนกันแล้วไม่เกิดปฏิกิริยาเคมี",
									Correct: false,
								},
								{
									Title:   "พลังงานที่น้อยที่สุดที่จะทำให้อะตอมหลุดออก",
									Correct: false,
								},
								{
									Title:   "พลังงานจำนวนน้อยที่สุดที่เกิดจากการชนของอนุภาคของสารตั้งต้นแล้วทำปฏิกิริยาเคมี",
									Correct: true,
								},
							},
						},
					},
				},
			},
			TeacherID:   8,
			Course_name: "เคมีหรรษา",
			Picture:     "https://scit.surat.psu.ac.th/chem/media/k2/items/cache/7293a47c0f4cdddd46ff10bcf3d23287_XL.jpg",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับเนื้อหาของวิชาเคมีในเรื่องต่าง ๆ เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
		{
			//Course_21
			SubjectID: 43,
			Videos: &[]models.Video{
				{
					Class_level: 6,
					Video_name:  "ความหลากหลายทางชีวภาพ (อาณาจักรสัตว์) ",
					Picture:     "https://i.ytimg.com/vi/mnN7S8VRwwA/maxresdefault.jpg",
					Description: "ชีวะวิทยา เรื่อง ความหลากหลายทางชีวภาพ | แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_21%2FVideo_1%2F%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%AB%E0%B8%A5%E0%B8%B2%E0%B8%81%E0%B8%AB%E0%B8%A5%E0%B8%B2%E0%B8%A2%E0%B8%97%E0%B8%B2%E0%B8%87%E0%B8%8A%E0%B8%B5%E0%B8%A7%E0%B8%A0%E0%B8%B2%E0%B8%9E.pdf?alt=media&token=df47f647-04b5-4057-b5c4-bac731855be4",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_21%2FVideo_1%2F%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%E0%B8%8A%E0%B8%B5%E0%B8%A7%E0%B8%B0%20%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%AB%E0%B8%A5%E0%B8%B2%E0%B8%81%E0%B8%AB%E0%B8%A5%E0%B8%B2%E0%B8%A2%E0%B8%97%E0%B8%B2%E0%B8%87%E0%B8%8A%E0%B8%B5%E0%B8%A7%E0%B8%A0%E0%B8%B2%E0%B8%9E%20(%E0%B8%AD%E0%B8%B2%E0%B8%93%E0%B8%B2%E0%B8%88%E0%B8%B1%E0%B8%81%E0%B8%A3%E0%B8%AA%E0%B8%B1%E0%B8%95%E0%B8%A7%E0%B9%8C).mp4?alt=media&token=49b4df96-fbde-4ab8-8922-a99b58a16d6e",
					Exercises: &[]models.Exercise{
						{
							Question: "1.การจัดหมวดหมู่ของสิ่งมีชีวิตระดับใดอยู่สูงที่สุด",
							Choices: &[]models.Choice{
								{
									Title:   "super order",
									Correct: false,
								},
								{
									Title:   "order",
									Correct: false,
								},
								{
									Title:   "sub class",
									Correct: true,
								},
								{
									Title:   "super family",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ข้อใดเป็นวิธีการระบุชนิดของสิ่งมีชีวิตที่น่าเชื่อถือน้อยที่สุด",
							Choices: &[]models.Choice{
								{
									Title:   "เปรียบเทียบจากภาพถ่ายในอินเทอเน็ต",
									Correct: true,
								},
								{
									Title:   "สอบถามผู้เชี่ยวชาญเฉพาะ",
									Correct: false,
								},
								{
									Title:   "เปรียบเทียบจากตัวอย่างในพิพิธภัณฑ์",
									Correct: false,
								},
								{
									Title:   "ใช้ไดโคโตมัสคีย์",
									Correct: false,
								},
							},
						},
						{
							Question: "3.ข้อใดไม่จัดเป็นกระบวนการของ Taxonomy",
							Choices: &[]models.Choice{
								{
									Title:   "natural selection",
									Correct: true,
								},
								{
									Title:   "classification",
									Correct: false,
								},
								{
									Title:   "nomenclature",
									Correct: false,
								},
								{
									Title:   "identification",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 6,
					Video_name:  "ระบบขับถ่าย",
					Picture:     "https://i.ytimg.com/vi/jSOeIrhJjQo/maxresdefault.jpg",
					Description: "ชีวะวิทยา เรื่อง ระบบขับถ่าย| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_21%2FVideo_2%2F%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%E0%B8%8A%E0%B8%B5%E0%B8%A7%E0%B8%B0%20%E0%B8%A3%E0%B8%B0%E0%B8%9A%E0%B8%9A%E0%B8%82%E0%B8%B1%E0%B8%9A%E0%B8%96%E0%B9%88%E0%B8%B2%E0%B8%A2.mp4?alt=media&token=a31cea78-90ff-4d03-a80d-c4132c50e705",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_21%2FVideo_2%2F%E0%B8%A3%E0%B8%B0%E0%B8%9A%E0%B8%9A%E0%B8%82%E0%B8%B1%E0%B8%9A%E0%B8%96%E0%B9%88%E0%B8%B2%E0%B8%A2.pdf?alt=media&token=d35d7725-316c-46a3-b7a4-981ee5204056",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ส่วนใดทำหน้าที่เป็นตัวกรองของเสียจำพวกปัสสาวะ",
							Choices: &[]models.Choice{
								{
									Title:   "หน่วยไต",
									Correct: true,
								},
								{
									Title:   "ท่อไต",
									Correct: false,
								},
								{
									Title:   "หลอดไต",
									Correct: false,
								},
								{
									Title:   "ถูกทุกข้อ",
									Correct: false,
								},
							},
						},
						{
							Question: "2.สารใดที่มีการดูดกลับที่บริเวณท่อของหน่วยไต",
							Choices: &[]models.Choice{
								{
									Title:   "เม็ดเลือดแดง",
									Correct: false,
								},
								{
									Title:   "กลูโคส",
									Correct: true,
								},
								{
									Title:   "ยูเรีย",
									Correct: false,
								},
								{
									Title:   "ถูกทั้ง3ข้อ",
									Correct: false,
								},
							},
						},
						{
							Question: "3.การทำงานของระบบขับถ่ายปัสสาวะในข้อใดถูกต้อง",
							Choices: &[]models.Choice{
								{
									Title:   "เลือด --> หน่วยไต --> ท่อปัสสาวะ --> กระเพาะปัสสาวะ",
									Correct: false,
								},
								{
									Title:   "เลือด --> หน่วยไต --> ท่อไต --> กระเพาะปัสสาวะ --> ท่อปัสสาวะ",
									Correct: true,
								},
								{
									Title:   "เลือด --> ท่อไต --> หน่วยไต --> กระเพาะปัสสาวะ --> ท่อปัสสาวะ",
									Correct: false,
								},
								{
									Title:   "เลือด --> กรวยไต --> หน่วยไต --> กระเพาะปัสสาวะ --> ท่อปัสสาวะ",
									Correct: false,
								},
							},
						},
					},
				},
				{
					Class_level: 6,
					Video_name:  "ระบบหมุนเวียนเลือด",
					Price:       25,
					Picture:     "https://i.ytimg.com/vi/v9cqFR7fWBY/maxresdefault.jpg",
					Description: "ชีวะวิทยา เรื่อง ระบบหมุนเวียนเลือด| แบบละเอียดเเบบเข้าใจง่ายเเละสนุกกับการเรียน",
					Url:         "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_21%2FVideo_3%2F%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%E0%B8%8A%E0%B8%B5%E0%B8%A7%E0%B8%B0%20%E0%B8%A3%E0%B8%B0%E0%B8%9A%E0%B8%9A%E0%B8%AB%E0%B8%A1%E0%B8%B8%E0%B8%99%E0%B9%80%E0%B8%A7%E0%B8%B5%E0%B8%A2%E0%B8%99%E0%B9%80%E0%B8%A5%E0%B8%B7%E0%B8%AD%E0%B8%94.mp4?alt=media&token=022fa670-187b-45bd-85ec-0f7f43be3c6b",
					Sheet:       "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Course%2FCourse_21%2FVideo_3%2F%E0%B8%A3%E0%B8%B0%E0%B8%9A%E0%B8%9A%E0%B8%AB%E0%B8%A1%E0%B8%B8%E0%B8%99%E0%B9%80%E0%B8%A7%E0%B8%B5%E0%B8%A2%E0%B8%99%E0%B9%80%E0%B8%A5%E0%B8%B7%E0%B8%AD%E0%B8%94.pdf?alt=media&token=c8dba11d-2cb2-41ce-8b00-27afec16be42",
					Exercises: &[]models.Exercise{
						{
							Question: "1.ระบบหมุนเวียนเลือดสามารถแบ่งได้กี่ระบบ",
							Choices: &[]models.Choice{
								{
									Title:   "5 ระบบ",
									Correct: false,
								},
								{
									Title:   "2 ระบบ",
									Correct: true,
								},
								{
									Title:   "4 ระบบ",
									Correct: false,
								},
								{
									Title:   "10 ระบบ",
									Correct: false,
								},
							},
						},
						{
							Question: "2.ระบบไหลหมุนเวียนเลือดประกอบด้วยอะไรบ้าง",
							Choices: &[]models.Choice{
								{
									Title:   "ระบบหลอดเลือด และระบบหัวใจ",
									Correct: false,
								},
								{
									Title:   "ระบบ เลือดเสียและเลือดดี ระบบน้ำเหลือง",
									Correct: false,
								},
								{
									Title:   "ระบบหัวใจหลอดเลือด",
									Correct: false,
								},
								{
									Title:   "ระบบหัวใจหลอดเลือด และระบบน้ำเหลือง",
									Correct: true,
								},
							},
						},
						{
							Question: "3.หัวใจห้องใดที่ทำหน้าที่สูบฉีดเลือดไปยังส่วนต่างๆ ของร่างกาย",
							Choices: &[]models.Choice{
								{
									Title:   "ห้องบนซ้าย",
									Correct: false,
								},
								{
									Title:   "ห้องล่างซ้าย",
									Correct: true,
								},
								{
									Title:   "ห้องบนขวา",
									Correct: false,
								},
								{
									Title:   "ห้องล่างขวา",
									Correct: false,
								},
							},
						},
					},
				},
			},
			TeacherID:   9,
			Course_name: "สนุกกับชีวะพื้นฐาน",
			Picture:     "https://scontent.fbkk13-3.fna.fbcdn.net/v/t39.30808-6/327189279_511767807512268_1499913387549108745_n.jpg?_nc_cat=110&ccb=1-7&_nc_sid=09cbfe&_nc_eui2=AeHndIyIZ4MSAvGUuXhy-qNDII9m9ojS-gYgj2b2iNL6BuSxegnyw3AJ2T1b-e-05gUPvew0fiGzTa4mTMNofBpI&_nc_ohc=InC-6R7WaEYAX-XyCEV&_nc_ht=scontent.fbkk13-3.fna&oh=00_AfCAbdYmfpBweLozCdw6mulLZHtG1ZSzEs3L31Ot76cZ4g&oe=6428D64C",
			Description: "คอร์สนี้จะพาน้อง ๆ ได้รู้จักกับเนื้อหาของวิชาชีวะในเรื่องต่าง ๆ เเละได้สนุกกับการเรียนรู้เเละเเบบฝึกหัดในคอร์ส มาเริ่มกันเลย!!",
		},
	}
	if err := DB.Create(&course).Error; err != nil {
		panic(err)
	}
	fmt.Println("Course created")

	var review_video = &[]models.Review_Video{
		{
			VideoID: 1,
			Rating:  4,
			Comment: "ขอบคุณมากๆนะคะ หนูตามดูคลิปที่สอนมากๆเลยตัดสินใจไปลงคอร์สเพื่อสนับสนุนและเพื่อการเรียนที่ดีขึ้นของหนู สอนดีและเข้าใจมากๆเลยค่ะ สอนตั้งแต่พื้นฐานเลย ไม่ยากอย่างที่คิด และไม่เครียดด้วย",
		},
		{
			VideoID: 1,
			Rating:  5,
			Comment: "ครูสอนเข้าใจมากค่ะ ขอบคุณมากๆนะคะที่ทำคลิปแบบนี้ เข้าใจขึ้นมากเลย",
		},
		{
			VideoID: 1,
			Rating:  4,
			Comment: "ครูสอนไม่เข้าใจเลยมาดูวิดิโอนี้ เเปปเดียวทำได้เลยครับ ขอบคุณที่ทำวิดิโอนี้มานะครับ จะสอบก็เลยเข้ามาดูขอบคุณครับ",
		},
		{
			VideoID: 2,
			Rating:  3,
			Comment: "คลิปนี้เป็นประโยชน์มากๆๆเลยครับผม สุดยอดครับ เป็นกำลังใจให้ครับผม",
		},
		{
			VideoID: 2,
			Rating:  4,
			Comment: "ดูแล้วเข้าใจขึ้นมากเลย สอนแล้วเข้าใจง่ายกว่าอาจารย์สอนที่โรงเรียนอีก ชอบมากๆค้าบ",
		},
		{
			VideoID: 2,
			Rating:  4,
			Comment: "ขอบคุณมากๆนะคะ ดูแล้วเข้าใจขึ้นเยอะเลยอธิบายดีมากๆ ติดตามต่อไปนะคะ เป็นกำลังใจให้ค่ะ",
		},
		{
			VideoID: 3,
			Rating:  5,
			Comment: "สอนเข้าใจง่ายเดี๋ยวพรุ่งนี้สอบแล้วน่าจะเป็นประโยชน์เยอะมากกกเลยครับ",
		},
		{
			VideoID: 3,
			Rating:  4,
			Comment: "สอนดีเเละเข้าใจมากๆๆๆ ทำให้หนูเข้าใจบทนี้มากขึ้นเนื้อหาครอบคลุม ขอบคุณมากๆเลยนะคะ",
		},
		{
			VideoID: 3,
			Rating:  4,
			Comment: "สอนเข้าใจมากค่ะ ชอบมากค่ะ คือหนูดูมาหลายๆคลิปละไม่เข้าใจ พอมาดูของพี่คือแบบเข้าใจแบบมากๆเลยค่ะ  พี่สอนละเอียด รู้ที่มาของแต่ละตัวมากค่ะ",
		},
		{
			VideoID: 4,
			Rating:  4,
			Comment: "สอนสนุกมากครับ ไม่เครียดเลย แล้วยังเข้าใจง่ายด้วย เป็นกำลังใจให้มาทำคลิปติวแบบนี้เรื่อยๆนะครับ",
		},
		{
			VideoID: 4,
			Rating:  5,
			Comment: "สอนเก่งมากเลยค่ะ ไม่น่าเบื่อเข้าใจมากๆ เรียนแล้วไม่เครียดด้วยสอนตั้งแต่เนื้อหาแรกๆอีก ขอบคุณมากๆเลยนะคะ",
		},
		{
			VideoID: 4,
			Rating:  5,
			Comment: "ขอบคุณพี่มาก ๆ เลยที่ทำให้หนูเข้าใจบทนี้สักที หนูท้อใจมากที่ไม่เข้าใจบทนี้สักที แต่พอมาดูคลิปนี้หนูเข้าใจบทนี้เลยค่ะ ขอบคุณมากๆเลยค่าา",
		},
		{
			VideoID: 5,
			Rating:  4,
			Comment: "ขอบคุณมากๆเลยค้าบบบบบ จากคนที่ไม่มีพื้นฐานคณิตเท่าไหร่เลย แต่พอได้ดูคลิปนี้เเละทำความเข้าใจแล้วรู้เรื่องแบบมากๆ",
		},
		{
			VideoID: 5,
			Rating:  4,
			Comment: "ปีนี้กำลังจะเตรียมสอบแล้วกลัวมากๆแต่พอมาเจอคลิปนี้แล้วใจชื้นเลยค่ะ ขอบคุณที่มาสอนเนื้อหาดีๆให้นะคะ สอนดีมาก เข้าใจขึ้นมากๆเลยค่ะ",
		},
		{
			VideoID: 5,
			Rating:  4,
			Comment: "ขอบคุณมากๆค่ะ เป็นคนที่ตอนเรียนก็ไม่ค่อยเข้าใจ เเต่คลิปนี้สอนดีมาก ตอนนี้เข้าใจมากขึ้นเยอะเลยค่ะ",
		},
		{
			VideoID: 6,
			Rating:  4,
			Comment: "ขอบคุณมากๆนะคะ สอนดีและเข้าใจมากๆเลยค่ะ หนูหาคลิปสอนเพิ่มเติมบทนี้จนมาเจอคลิปนี้ สุดยอดมากๆเลยค่ะ เข้าใจง่ายมากๆเลยขอบคุณนะคะ",
		},
		{
			VideoID: 6,
			Rating:  5,
			Comment: "ขอบคุณมากๆนะคะ ที่โรงเรียนตอนเรียนเรื่องนี้แล้วหนูงงมากๆ เเต่พอเจอคลิปพี่ทีนี่กระจ่างเลย",
		},
		{
			VideoID: 6,
			Rating:  4,
			Comment: "ชอบมากเลยค่ะ พอมาดูคลิปนี้ช่วยเสริมความเข้าใจให้หนูได้มากเลย สอนสนุกมากๆ ไม่น่าเบื่อด้วยค่ะ",
		},
		{
			VideoID: 7,
			Rating:  3,
			Comment: "คุณครูสอนดีมากๆเลยค่ะ แง ขอบคุณนะคะที่ตั้งใจสอน หนูจะตั้งใจเรียนให้สมกับที่ครูตั้งใจสอนเลยค่ะ ชอบเรียนคลิปครูมากๆเลย",
		},
		{
			VideoID: 7,
			Rating:  4,
			Comment: "ชอบมากครับ ถึงจะยังไม่ค่อยเข้าใจเเต่ก็รู้สึกเข้าใจมากขึ้น ทำคลิปสอนดีๆเเบบนี้ไปเรื่อยๆนะครับ",
		},
		{
			VideoID: 7,
			Rating:  5,
			Comment: "วันนี้สอบปลายภาคแล้วตอนเรียนไม่เข้าใจเลย มาดูคลิปนี้ก่อนเข้าสอบเข้าใจเยอะขึ้นมากกกกกก ขอบคุณมากๆครับ",
		},
		{
			VideoID: 8,
			Rating:  5,
			Comment: "ขอบคุณพี่มาก ๆ เลยที่ทำให้หนูเข้าใจบทนี้สักที หนูท้อใจมากที่ไม่เข้าใจบทนี้สักที แต่พอมาดูคลิปนี้หนูเข้าใจบทนี้เลยค่ะ ขอบคุณมากๆเลยค่าา",
		},
		{
			VideoID: 8,
			Rating:  4,
			Comment: "ชอบมากเลยค่ะ พอมาดูคลิปนี้ช่วยเสริมความเข้าใจให้หนูได้มากเลย สอนสนุกมากๆ ไม่น่าเบื่อด้วยค่ะ",
		},
		{
			VideoID: 8,
			Rating:  4,
			Comment: "สอนเก่งมากเลยค่ะ ไม่น่าเบื่อเข้าใจมากๆ เรียนแล้วไม่เครียดด้วยสอนตั้งแต่เนื้อหาแรกๆอีก ขอบคุณมากๆเลยนะคะ",
		},
		{
			VideoID: 9,
			Rating:  4,
			Comment: "สอนสนุกมากครับ ไม่เครียดเลย แล้วยังเข้าใจง่ายด้วย เป็นกำลังใจให้มาทำคลิปติวแบบนี้เรื่อยๆนะครับ",
		},
		{
			VideoID: 9,
			Rating:  4,
			Comment: "ครูสอนไม่เข้าใจเลยมาดูวิดิโอนี้ เเปปเดียวทำได้เลยครับ ขอบคุณที่ทำวิดิโอนี้มานะครับ จะสอบก็เลยเข้ามาดูขอบคุณครับ",
		},
		{
			VideoID: 9,
			Rating:  5,
			Comment: "สอนดีเเละเข้าใจมากๆๆๆ ทำให้หนูเข้าใจบทนี้มากขึ้นเนื้อหาครอบคลุม ขอบคุณมากๆเลยนะคะ",
		},
		{
			VideoID: 10,
			Rating:  3,
			Comment: "ขอบคุณมากๆนะคะ สอนดีและเข้าใจมากๆเลยค่ะ หนูหาคลิปสอนเพิ่มเติมบทนี้จนมาเจอคลิปนี้ สุดยอดมากๆเลยค่ะ เข้าใจง่ายมากๆเลยขอบคุณนะคะ",
		},
		{
			VideoID: 10,
			Rating:  5,
			Comment: "ขอบคุณมากๆนะคะ หนูตามดูคลิปที่สอนมากๆเลยตัดสินใจไปลงคอร์สเพื่อสนับสนุนและเพื่อการเรียนที่ดีขึ้นของหนู สอนดีและเข้าใจมากๆเลยค่ะ",
		},
		{
			VideoID: 10,
			Rating:  4,
			Comment: "สอนเข้าใจมากค่ะ ชอบมากค่ะ คือหนูดูมาหลายๆคลิปละไม่เข้าใจ พอมาดูของพี่คือแบบเข้าใจแบบมากๆเลยค่ะ  พี่สอนละเอียด รู้ที่มาของแต่ละตัวมากค่ะ",
		},
	}
	if err := DB.Create(&review_video).Error; err != nil {
		panic(err)
	}

	fmt.Println("Reviews Video created")

	var review_tutor = &[]models.Review_Tutor{
		{
			TeacherID: 1,
			Rating:    4,
			Comment:   "ขอบคุณครูกุ๊กไก่มากๆเลยนะครับ รู้สึกได้ทบทวนความรู้เดิม จากที่สงสัยวันนี้พอลองมาเรียนกับครูกุ๊กไก่คือหายสงสัยเเล้วครับ ชอบสไตล์การสอนของครูมากครับ เรียนด้วยเเล้วไม่เครียดบวกกับเข้าใจในเนื้อหามากๆครับ",
		},
		{
			TeacherID: 1,
			Rating:    5,
			Comment:   "ชอบครูกุ๊กไก่มากๆ ค่ะ เป็นติวเตอร์ที่สอนเข้าใจ สอนสนุกไม่ง่วงด้วย กำลังสนใจจะซื้อคอร์สอยู่เลยค่ะ เกรดเทอมนี้ต้องดีขึ้นเเน่ๆ มีคอร์สดีๆเเบบนี้ไปเรื่อยๆนะคะ",
		},
		{
			TeacherID: 1,
			Rating:    5,
			Comment:   "ครูกุ๊กไก่เป็นครูที่สอนดีมากกกกกกกก ก.ไก่ล้านตัว เป็นติวเตอร์ที่มากกว่าติวเตอร์ ว่างๆไม่มีอะไรทำดูก็เพลินมากไม่เบื่อเลย สอนสนุก+ได้ความรู้เพิ่มด้วยอีก",
		},
		{
			TeacherID: 1,
			Rating:    4,
			Comment:   "ชอบที่สุดเลยค่ะ สอนสนุกกระชับเข้าใจดีมากๆเลยค่ะ เหมาะกับคนที่ไม่มีเวลาไปเรียนพิเศษข้างนอกบ้าน เรียนที่ไหนก็ได้เป็นกำลังใจให้นะคะทำต่อไปเยอะๆเลยนะคะ",
		},
		{
			TeacherID: 2,
			Rating:    3,
			Comment:   "ดูคลิปที่คุณครูสอนเเล้ว ก็มีงง ๆ บ้าง เเต่ก็เข้าใจมากขึ้นค่ะ อยากให้สอนให้สนุกกว่านี้อีกหน่อยจะทำให้การเรียนสนุกเเละน่าสนใจขึ้นค่ะ เป็นกำลังใจให้นะคะ",
		},
		{
			TeacherID: 2,
			Rating:    3,
			Comment:   "เสียงคุณครูเบาไปหน่อยครับ การพูดก็เนือยๆ ไม่น่าฟัง ปรับวิธีการพูดหน่อยจะดีขึ้นครับ เเต่เนื้อหาที่นำมาสอนดีครับ สู้ๆครับครู",
		},
		{
			TeacherID: 2,
			Rating:    2,
			Comment:   "ตอนดูคลิปที่คุณครูสอน เเรกๆก็ตั้งใจฟังอยู่ครับ เเต่หลังจากนั้นเกือบหลับ เสียงเบามากครับ",
		},
		{
			TeacherID: 3,
			Rating:    5,
			Comment:   "ใครมีคุณครูที่สอนดีแบบนี้ โชคดีมากๆเลย ขอบคุณมากๆนะคะคุณครู เป็นกำลังในการทำคลิปต่อๆไปนะคะ",
		},
		{
			TeacherID: 3,
			Rating:    5,
			Comment:   "สอนเข้าใจมากค่ะ ชอบมากค่ะ คือหนูดูมาหลายๆคลิปละไม่เข้าใจ  พอมาดูของครูเเล้วเข้าใจแบบมากๆเลยค่าาา  พี่สอนละเอียด รู้ที่มาของแต่ละตัวมากค่ะ",
		},
		{
			TeacherID: 3,
			Rating:    4,
			Comment:   "คลิปขอบคุณครูเป็นประโยชน์มากๆๆเลยครับผม สุดยอดครับ เป็นกำลังใจให้กันครับผม",
		},
		{
			TeacherID: 4,
			Rating:    5,
			Comment:   "ชอบครูดาวมากๆ ค่ะ เป็นติวเตอร์ที่สอนเข้าใจ สอนสนุกไม่ง่วงด้วย กำลังสนใจจะซื้อคอร์สอยู่เลยค่ะ เกรดเทอมนี้ต้องดีขึ้นเเน่ๆ มีคอร์สดีๆเเบบนี้ไปเรื่อยๆนะคะ",
		},
		{
			TeacherID: 4,
			Rating:    5,
			Comment:   "หนูดูคลิปของคุณครูแล้วเข้าใจขึ้นมากเลยค่ะ สอนแล้วเข้าใจง่ายกว่าอาจารย์สอนที่โรงเรียนขึ้นอีก ขอบคุณนะคะ ชอบมากๆค่า",
		},
		{
			TeacherID: 4,
			Rating:    5,
			Comment:   "เข้าใจขึ้นเยอะเลยครับครู ตอนเด็กชอบดื้อเรียนเลข ม.ปลายหนีไปศิลป​ภาษา แต่ต้องใช้เลขในการสอบเลื่อนชั้น ได้ครูช่วยได้เยอะเลย",
		},
		{
			TeacherID: 5,
			Rating:    4,
			Comment:   "ขอบคุณมากเลยครับ เมื่อก่อนไม่เอาคณิตเลย พอมาดูคลิปของคุณครู ทำให้ผมเปิดใจกับวิชาคณิตทันทีเลย",
		},
		{
			TeacherID: 5,
			Rating:    4,
			Comment:   "ชอบครูมากๆเลยค่ะ สอนสนุก แล้วก็สอนตั้งแต่พื้นฐานจริงๆ ทำให้คนพื้นไม่แน่นไปต่อในคณิตได้ง่ายขึ้นมากเลยค่ะ ขอบคุณมากๆนะคะ",
		},
		{
			TeacherID: 5,
			Rating:    4,
			Comment:   "สอนเข้าใจง่ายครับ ชัดเจน​ ให้สังเกตุอะไร​ มีลำดับการคิดเป็นอย่างไร​ เยี่ยมเลยครับ​ คุณครูครับ",
		},
		{
			TeacherID: 6,
			Rating:    5,
			Comment:   "คุณครูสอนดีมากเลยค่ะ หนูไม่หลับเลยยย>< เข้าใจมากกว่าครูที่โรงเรียนสอนอีก ขอบคุณมากนะคะ",
		},
		{
			TeacherID: 6,
			Rating:    4,
			Comment:   "สอนเข้าใจมากเลยค่ะ ไม่น่าเบื่อด้วย ขอบคุณมากๆนะคะ",
		},
		{
			TeacherID: 6,
			Rating:    3,
			Comment:   "สอนเข้าใจง่ายครับ ชัดเจน​ ให้สังเกตุอะไร​ มีลำดับการคิดเป็นอย่างไร​ เยี่ยมเลยครับ​ คุณครูครับ",
		},
		{
			TeacherID: 7,
			Rating:    4,
			Comment:   "หนูฟังครูที่โรงเรียนไม่เข้าใจ หนูจะเข้าดูคลิปคุณครูตลอดเลยค่ะ เข้าใจมากๆพูดเข้าใจง่ายชอบมากเลยคะขอบคุณสำหรับความรู้นะคะพี่",
		},
		{
			TeacherID: 7,
			Rating:    4,
			Comment:   "สอนเข้าใจมากเลยค่ะ ไม่น่าเบื่อด้วย ขอบคุณมากๆนะคะ",
		},
		{
			TeacherID: 7,
			Rating:    4,
			Comment:   "คลิปขอบคุณครูเป็นประโยชน์มากๆๆเลยครับผม สุดยอดครับ เป็นกำลังใจให้กันครับผม",
		},
		{
			TeacherID: 8,
			Rating:    2,
			Comment:   "ดูคลิปของคุณครูเเล้วไม่ค่อยเข้าใจเลยค่ะ งงมากกว่าเดิม เเต่เป็นกำลังใจให้นะคะ",
		},
		{
			TeacherID: 8,
			Rating:    3,
			Comment:   "จากที่งงอยู่เเล้ว พอดูคลิปของครูเเล้วงงมากกว่าเดิมอีกครับ เเต่ก็เป็นกำลังใจให้นะครับ",
		},
		{
			TeacherID: 8,
			Rating:    3,
			Comment:   "อาจจะเพิ่งหัดสอน เเต่ยังไงก็สู้ๆนะครับ ขอให้คุณครูสำเร็จในการสอนครับ",
		},
		{
			TeacherID: 9,
			Rating:    4,
			Comment:   "ดูคลิปของคุณครูเเล้วไม่ค่อยเข้าใจเท่าไหร่ เเต่รู้สึกว่าครูตั้งใจสอนดีครับ สู้ๆครับ",
		},
		{
			TeacherID: 9,
			Rating:    3,
			Comment:   "สอนเข้าใจบ้าง ไม่เข้าใจบ้าง เเต่ก็เป็นกำลังใจให้นะครับ",
		},
		{
			TeacherID: 9,
			Rating:    2,
			Comment:   "หนูดูคลิปของคุณครูเเล้วไม่เข้าใจเลยค่ะ ขอโทษนะคะ เเต่ก็เป็นกำลังใจให้นะคะ",
		},
		{
			TeacherID: 10,
			Rating:    4,
			Comment:   "ชอบเทคนิคการติวของครูมากๆเลยค่ะ มันทำให้อยากเรียนเรื่อยๆ ไม่เบื่อเลย",
		},
		{
			TeacherID: 10,
			Rating:    3,
			Comment:   "ขอบคุณมากค่ะกำลังจะสอบพอให้เพื่อนที่ไม่เข้าใจดู เค้าเข้าใจเลยค่ะ ทำผลงานดีๆอย่างนี้ต่อไปนะคะ",
		},
		{
			TeacherID: 10,
			Rating:    4,
			Comment:   "หนูเรียนกับครูที่โรงเรียนไม่รู้เรื่องเลย แถมสอบพรุ่งนี้ด้วย แต่มาเจอคลิปของพี่ หนูเข้าใจเกือบหมดทุกอย่างเลยค่ะ ขอบคุณพี่มาก ๆ นะคะ",
		},
	}
	if err := DB.Create(&review_tutor).Error; err != nil { // 2 8 9
		panic(err)
	}

	fmt.Println("Reviews Tutor created")

	var inbox = &[]models.Inbox{
		{
			User1ID:           1,
			User2ID:           11,
			Last_message:      "สวัสดีค่ะ สงสัยตรงไหน สามารถสอบถามครูได้เลยนะคะ",
			LastMessageUserID: 1,
		},
		{
			User1ID:           2,
			User2ID:           12,
			Last_message:      "สวัสดีค่ะ หากสงสัยเนื้อหาที่คุณครูสอน ถามได้เลยค่ะ",
			LastMessageUserID: 2,
		},
		{
			User1ID:           11,
			User2ID:           5,
			Last_message:      "สวัสดรค่ะคุณครู หนูมีเรื่องอยากสอบถามค่ะ",
			LastMessageUserID: 11,
		},
		{
			User1ID:           12,
			User2ID:           10,
			Last_message:      "สวัสดรค่ะคุณครู เนื้อหาที่คุณครูสอน มีบางช่วงที่หนูยังไม่ค่อยเข้าใจ อยากจะถามเพิ่มเติมค่ะ",
			LastMessageUserID: 12,
		},
	}
	if err := DB.Create(&inbox).Error; err != nil {
		panic(err)
	}

	var chatroom = &[]models.ChatRoom{
		{
			Inbox_id: 2,
			Conversations: []*models.Conversation{
				{
					Sender_id: 1,
					Message:   "สวัสดีค่ะ สงสัยตรงไหน สามารถสอบถามครูได้เลยนะคะ",
					CreatedAt: time.Now().Unix(),
				}},
		},
		{
			Inbox_id: 3,
			Conversations: []*models.Conversation{
				{
					Sender_id: 2,
					Message:   "สวัสดีค่ะ หากสงสัยเนื้อหาที่คุณครูสอน ถามได้เลยค่ะ",
					CreatedAt: time.Now().Unix(),
				}},
		},
		{
			Inbox_id: 4,
			Conversations: []*models.Conversation{
				{
					Sender_id: 11,
					Message:   "สวัสดรค่ะคุณครู หนูมีเรื่องอยากสอบถามค่ะ",
					CreatedAt: time.Now().Unix(),
				}},
		},
		{
			Inbox_id: 5,
			Conversations: []*models.Conversation{
				{
					Sender_id: 12,
					Message:   "สวัสดรค่ะคุณครู เนื้อหาที่คุณครูสอน มีบางช่วงที่หนูยังไม่ค่อยเข้าใจ อยากจะถามเพิ่มเติมค่ะ",
					CreatedAt: time.Now().Unix(),
				}},
		},
	}
	if err := DB.Create(&chatroom).Error; err != nil {
		panic(err)
	}

	var post = &[]models.Post{
		{
			SubjectID:    1,
			UserID:       11,
			CreatedAt:    time.Now().Unix(),
			Caption:      "ข้อไหนง่ายสุดหรอครับ อาจารย์ให้เลือก1ข้อ",
			Post_picture: "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Images%2Fpost_image_2.jpg?alt=media&token=a126d2a4-4702-4851-976f-42b5dfe5ccbd",
			Comments: &[]models.Comment{
				{
					PostID:      1,
					UserID:      1,
					Description: "1-10ง่ายนะ แนะนำว่าทำ1-17 แล้วจับสลากเอาว่าจะได้ข้อไหนไปส่งครูเลยจ้าา",
					CreatedAt:   time.Now().Unix(),
				},
				{
					PostID:      1,
					UserID:      2,
					Description: "ข้อง่ายเยอะนะ เลือกเอาได้เลย ได้ฝึกตัวเองด้วยค่ะ",
					CreatedAt:   time.Now().Unix(),
				},
				{
					PostID:      1,
					UserID:      3,
					Description: "ข้อ 1-10 ข้อไหนก็ได้ คิดว่าง่ายกว่าข้อหลังๆนะ ",
					CreatedAt:   time.Now().Unix(),
				},
			},
		},
		{
			SubjectID:    1,
			UserID:       11,
			CreatedAt:    time.Now().Unix(),
			Caption:      "ข้อนี้มีวิธีคิดยังไงเหรอครับ",
			Post_picture: "https://firebasestorage.googleapis.com/v0/b/coursez-50fb3.appspot.com/o/Images%2Fpost_image_1.jpg?alt=media&token=d8ea6945-e695-4a66-8a88-04a0d7009c26",
			Comments: &[]models.Comment{
				{
					PostID:      2,
					UserID:      4,
					Description: "ข้อนี้ N_2 เป็นจำนวนเต็มบวก แปลว่า 11 ต้องหาร 2N_1-1 ลงตัวค่ะ",
					CreatedAt:   time.Now().Unix(),
				},
				{
					PostID:      2,
					UserID:      5,
					Description: "เอาตัวเลือกหาร 11 แล้วได้เศษ 9 ค่ะ",
					CreatedAt:   time.Now().Unix(),
				},
				{
					PostID:      2,
					UserID:      6,
					Description: "ตามที่ทดมาเลยค่ะ จริง ๆ ข้อนี้มี choice เดียวที่หารด้วย 3 ลงตัว ไม่ต้องใช้เงื่อนไขอื่นคิดเลย",
					CreatedAt:   time.Now().Unix(),
				},
			},
		},
		{
			SubjectID: 8,
			UserID:    12,
			CreatedAt: time.Now().Unix(),
			Caption:   "เอลนีโญ เเละ ลานีญา คืออะไรหรอคะ",
			Comments: &[]models.Comment{
				{
					PostID:      3,
					UserID:      7,
					Description: "ปรากฏการณ์ลานีญาและเอลนีโญ เป็นปรากฏการณ์ธรรมชาติที่เกิดขึ้นจากความผกผันของสภาวะอากาศบริเวณแถบเส้นศูนย์สูตร เหนือมหาสมุทรแปซิฟิก ทำให้การไหลเวียนของน้ำและกระแสลมเกิดความแปรปรวน ซึ่งจะส่งผลกระทบต่อการเปลี่ยนแปลงสภาพอากาศของโลกอย่างรุนแรง สำหรับประเทศไทย ผลกระทบจากปรากฏการณ์ลานีญาอาจทำให้เกิดการสูญเสียในด้านต่างๆ",
					CreatedAt:   time.Now().Unix(),
				},
				{
					PostID:      3,
					UserID:      8,
					Description: "ปรากฏการณ์ เอลนีโญ เกิดจากกระแสลมมีกำลังอ่อนและเปลี่ยนทิศทางพัดจากด้านตะวันออกของมหาสมุทรแปฟิซิกไปด้านตะวันตกของมหาสมุทรแปฟิซิก ทำให้กระแสน้ำอุ่นไหล ไปยังทวีปอเมริกาใต้แทน ด้วยเหตุนี้ภูมิภาคเอเชียตะวันออกเฉียงใต้และออสเตรเลียขาดฝนและเกิดความแห้งแล้ง แต่ชายฝั่งของทวีปอเมริกาใต้กลับมีฝนตกเพิ่มมากขึ้น",
					CreatedAt:   time.Now().Unix(),
				},
				{
					PostID:      3,
					UserID:      8,
					Description: "ปรากฏการณ์ ลานีญา เกิดจากกระแสลมพัดจากด้านตะวันออกของมหาสมุทรแปซิฟิกมายังด้านตะวันตกของมหาสมุทรแปซิฟิกตามเดิม แต่กระแสลมมีความรุนแรงมากกว่าปกติ ทำให้กระแสน้ำอุ่นไหลมายังภูมิภาคเอเชียตะวันออกเฉียงใต้มากขึ้น ส่งผลให้ภูมิภาคเอเชียตะวันออกเฉียงใต้และออสเตรเลียมีระดับน้ำทะเลสูงขึ้นและฝนตกหนักมากกว่าปกติ ในทางตรงข้ามก็เกิดภาวะความแห้งแล้งตามแนวชายฝั่งทวีปอเมริกาใต้",
					CreatedAt:   time.Now().Unix(),
				},
			},
		},
		{
			SubjectID: 8,
			UserID:    12,
			CreatedAt: time.Now().Unix(),
			Caption:   "โลกหมุนรอบตัวเองเป็นเวลาเท่าไหร่หรอคะ และ โลกหมุนรอบดวงอาทิตย์เป็นเวลาเท่าไหร่คะ",
			Comments: &[]models.Comment{
				{
					PostID:      4,
					UserID:      8,
					Description: "โลกหมุนรอบตัวเอง ใช้เวลาประมาณ 24 ชั่วโมงค่ะ เเละ โลกหมุนรอบดวงอาทิตย์ ใช้เวลาประมาณ 365 วันค่ะ",
					CreatedAt:   time.Now().Unix(),
				},
				{
					PostID:      4,
					UserID:      9,
					Description: "โลกหมุนรอบตัวเองใช้เวลาประมาณ 23 ชั่วโมง 56 นาที 4 วินาที",
					CreatedAt:   time.Now().Unix(),
				},
				{
					PostID:      4,
					UserID:      9,
					Description: "โลกหมุนโคจรรอบดวงอาทิตย์ใช้เวลา 365 วัน 5 ชั่วโมง 49 นาที 12 วินาที หรือ 365.2425 วัน",
					CreatedAt:   time.Now().Unix(),
				},
			},
		},
		{
			SubjectID: 14,
			UserID:    11,
			CreatedAt: time.Now().Unix(),
			Caption:   "คำขยาย คืออะไรหรอคะ",
			Comments: &[]models.Comment{
				{
					PostID:      5,
					UserID:      5,
					Description: "คำขยาย คือ คำที่ทำหน้าที่ขยายนาม เป็นคำชนิดต่าง ๆ เช่น คำนาม คำสรรพนาม คำลักษณนาม คำบอกจำนวน เป็นต้น และเมื่อขยายแล้วจะเกิดเป็นกลุ่มคำนามหรือนามวลี",
					CreatedAt:   time.Now().Unix(),
				},
				{
					PostID:      5,
					UserID:      5,
					Description: "คำขยาย คือ คำที่ทำหน้าที่ขยายนาม เช่น นักเรียนที่เรียนเก่งได้รับรางวัล",
					CreatedAt:   time.Now().Unix(),
				},
				{
					PostID:      5,
					UserID:      5,
					Description: "ยกตัวอย่างคำขยายเพิ่มให้นะคะ เช่น ครูชมเชยนักเรียนที่ตั้งใจเรียน เขาเรียนโรงเรียนที่มีชื่อเสียง",
					CreatedAt:   time.Now().Unix(),
				},
			},
		},
		{
			SubjectID: 14,
			UserID:    11,
			CreatedAt: time.Now().Unix(),
			Caption:   "1.สระใดที่มีเสียงเหมือนกับสระอี , 2.กลุ่มคำใดต่อไปนี้เป็นคำขยาย? , คำว่า ขวาง เป็นคำในกลุ่มคำใด? ช่วยหน่อยนะคะ ",
			Comments: &[]models.Comment{
				{
					PostID:      6,
					UserID:      5,
					Description: "1. สระวี , 2.กำลัง, จัก, เคย , 3.คำนาม",
					CreatedAt:   time.Now().Unix(),
				},
				{
					PostID:      6,
					UserID:      5,
					Description: "ข้อ 1 ตอบว่า สระวี ค่ะ ข้อ 2 ตอบว่า กำลัง จัก เคย ค่ะ ข้อ 3 ตอบว่า คำนาม ค่ะ",
					CreatedAt:   time.Now().Unix(),
				},
				{
					PostID:      6,
					UserID:      5,
					Description: "คำตอบตามความคิดเห็นด้านบนเลยค่ะ",
					CreatedAt:   time.Now().Unix(),
				},
			},
		},
	}
	if err := DB.Create(&post).Error; err != nil {
		panic(err)
	}
	rand.NewSource(time.Now().UnixNano())
	var coursehistory []models.CourseHistory
	for i := 0; i < 30; i++ {
		coursehistory = append(coursehistory, models.CourseHistory{
			UserID:    int32(rand.Intn(10) + 1), // 1 - 10
			CourseID:  int32(rand.Intn(21) + 1), // 1 - 10
			Frequency: int32(rand.Intn(10) + 1), // 1 - 10
		})
	}
	if err := DB.Create(&coursehistory).Error; err != nil {
		panic(err)
	}
}

func RandomData() {
	var users []models.User
	DB.Preload("LikeCourses", "PaidVideos").Find(&users)
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < len(users); i++ {
		ranInt := make([]int, 4)
		for k := range ranInt {
			ranInt[k] = rand.Intn(21) + 1
		}
		var ranCourse []*models.Course
		DB.Select("DISTINCT course_id").Find(&ranCourse, ranInt)
		users[i].LikeCourses = append(users[i].LikeCourses, ranCourse...)

		ranInt = make([]int, 10)
		for k := range ranInt {
			ranInt[k] = rand.Intn(63) + 1
		}
		var ranVideos []*models.Video
		DB.Select("DISTINCT video_id, course_id, price").Find(&ranVideos, ranInt)
		for j := 0; j < len(ranVideos); j++ {
			if ranVideos[j].Price >= 0 {
				users[i].PaidVideos = append(users[i].PaidVideos, ranVideos[j])
			} else {
				var userVideosHistory models.VideoHistory
				rowAffect := DB.FirstOrCreate(&userVideosHistory, models.VideoHistory{UserID: users[i].User_id, VideoID: ranVideos[j].Video_id}).RowsAffected
				if rowAffect == 1 {
					userVideosHistory.Duration = int32(rand.Intn(500) + 1)
				}
				var course models.Course
				DB.First(&course, ranVideos[j].CourseID)
				var userCourseHistory models.CourseHistory
				DB.FirstOrCreate(&userCourseHistory, models.CourseHistory{UserID: users[i].User_id, CourseID: course.Course_id})
				userCourseHistory.Frequency++
				DB.Save(&userCourseHistory)
			}

		}

	}
	DB.Save(&users)
}

func WipeData() {
	dbMigrator := DB.Migrator()
	dbMigrator.DropTable("user_paidvideos")
	dbMigrator.DropTable("user_likevideos")
	dbMigrator.DropTable("user_likecourses")
	dbMigrator.DropTable("user_doneexercise")
	dbMigrator.DropTable(&models.User{}, &models.UserTeacher{}, &models.Experience{}, &models.Comment{}, &models.Course{}, &models.VideoHistory{}, &models.CourseHistory{}, &models.Post{}, &models.Review_Video{}, &models.Review_Tutor{}, &models.Reward_Info{}, &models.Reward_Item{}, &models.Subject{}, &models.Video{}, &models.Exercise{}, &models.Choice{}, &models.Payment{}, &models.Withdraw{}, &models.Inbox{}, &models.ChatRoom{}, &models.Conversation{}, &models.Address{})
}

func MigrateData() {
	DB.AutoMigrate(&models.User{}, &models.UserTeacher{}, &models.Experience{}, &models.Comment{}, &models.Course{}, &models.VideoHistory{}, &models.CourseHistory{}, &models.Post{}, &models.Review_Video{}, &models.Review_Tutor{}, &models.Reward_Info{}, &models.Reward_Item{}, &models.Subject{}, &models.Video{}, &models.Exercise{}, &models.Choice{}, &models.Payment{}, &models.Withdraw{}, &models.Inbox{}, &models.ChatRoom{}, &models.Conversation{}, &models.Address{})
}
