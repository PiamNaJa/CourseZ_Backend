package ai

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

type UserRating struct {
	UserID  int32
	VideoID int32
	Rating  float64
}

func TrainData(db *gorm.DB) {
	review_videos := []map[string]interface{}{}
	db.Raw("SELECT * FROM video_histories").Find(&review_videos)
	var userRating []UserRating = make([]UserRating, len(review_videos))
	for i := 0; i < len(review_videos); i++ {
		userRating[i].UserID = review_videos[i]["user_id"].(int32)
		userRating[i].VideoID = review_videos[i]["video_id"].(int32)
		userRating[i].Rating = review_videos[i]["rating"].(float64)
	}

	jsonBytes, _ := json.Marshal(userRating)
	request, err := http.NewRequest("POST", "http://localhost:8080/train", bytes.NewBuffer([]byte(jsonBytes)))
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
