package ai

import (
	//"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

type UsersCoursesData struct {
	UserID    int32
	CourseID  int32
	Frequency int32
	IsLike    int32
	IsBuy     int32
}

func TrainData(db *gorm.DB) {
	users_courses_data := []map[string]interface{}{}
	db.Raw(`SELECT 
    COALESCE(ch.user_id, ul.user_user_id, up.user_id) AS user_id, 
    COALESCE(ch.course_id, ul.course_course_id, up.course_id) AS course_id,
    COALESCE(SUM(ch.frequency), 0) AS frequency,
    MAX(CASE WHEN ul.user_user_id IS NOT NULL THEN 1 ELSE 0 END) AS islike,
    MAX(CASE WHEN up.user_id IS NOT NULL THEN 1 ELSE 0 END) AS isbuy
	FROM  course_histories AS ch 
    FULL OUTER JOIN user_likecourses AS ul 
        ON ul.user_user_id = ch.user_id AND ul.course_course_id = ch.course_id
    FULL OUTER JOIN 
        (SELECT DISTINCT up.user_user_id AS user_id, courses.course_id
         FROM user_paidvideos as up 
         JOIN videos ON up.video_video_id = videos.video_id 
         JOIN courses ON videos.course_id = courses.course_id) AS up
        ON up.course_id = ch.course_id AND up.user_id = ch.user_id 
	GROUP BY COALESCE(ch.user_id, ul.user_user_id, up.user_id), COALESCE(ch.course_id, ul.course_course_id, up.course_id)
	ORDER BY user_id, course_id`).Find(&users_courses_data)
	var usersCoursesData []UsersCoursesData = make([]UsersCoursesData, len(users_courses_data))
	for i := 0; i < len(users_courses_data); i++ {
		usersCoursesData[i].UserID = users_courses_data[i]["user_id"].(int32)
		usersCoursesData[i].CourseID = users_courses_data[i]["course_id"].(int32)
		usersCoursesData[i].Frequency = int32(users_courses_data[i]["frequency"].(int64)) 
		usersCoursesData[i].IsLike = users_courses_data[i]["islike"].(int32)
		usersCoursesData[i].IsLike = users_courses_data[i]["isbuy"].(int32)
	}
	fmt.Println(usersCoursesData)
	// jsonBytes, _ := json.Marshal(usersCoursesData)
	// request, err := http.NewRequest("POST", "http://localhost:8080/train", bytes.NewBuffer([]byte(jsonBytes)))
	// request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	// if err != nil {
	// 	panic(err)
	// }
	// client := &http.Client{}
	// response, err := client.Do(request)
	// if err != nil {
	// 	panic(err)
	// }
	// defer response.Body.Close()

	// fmt.Println("response Status:", response.Status)
	// body, _ := io.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))
}

func Predict() (recommendVideoId []int) {
	request, err := http.NewRequest("GET", "http://localhost:8080/predict/1", nil)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	json.Unmarshal(body, &recommendVideoId)
	return recommendVideoId
}

// func RandomLikeCourse(db *gorm.DB) {
// 	var users []models.User
// 	db.Preload("LikeCourses").Find(&users)
// 	rand.NewSource(time.Now().UnixNano())
// 	for i := 0; i < len(users); i++ {
// 		ranInt := make([]int, 11)
// 		for k := range ranInt {
// 			ranInt[k] = rand.Intn(20) + 1
// 		}
// 		var ranCourse []*models.Course
// 		db.Find(&ranCourse, ranInt)
// 		users[i].LikeCourses = append(users[i].LikeCourses, ranCourse...)
// 	}
// 	db.Save(&users)
// }
