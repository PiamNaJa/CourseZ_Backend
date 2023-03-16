package ai

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/PiamNaJa/CourseZ_Backend/models"
	"gorm.io/gorm"
)

func K_cluster(db *gorm.DB, k int) {
	var users []models.User
	db.Preload("CourseHistory").Preload("PaidVideos").Preload("LikeVideos").Preload("LikeCourses").Preload("LikeCourses.Subject").Find(&users)
	var courses = make([][]models.Course, len(users))
	for i := 0; i < len(users); i++ {
		for j := 0; j < len(users[i].LikeCourses); j++ {
			courses[i] = append(courses[i], *users[i].LikeCourses[j])
		}
		for j := 0; j < len(users[i].LikeVideos); j++ {
			var course models.Course
			db.Preload("Subject").First(&course, users[i].LikeVideos[j].CourseID)
			courses[i] = append(courses[i], course)
		}
		for j := 0; j < len(users[i].PaidVideos); j++ {
			var course models.Course
			db.Preload("Subject").First(&course, users[i].PaidVideos[j].CourseID)
			courses[i] = append(courses[i], course)
		}
		for j := 0; j < len(*users[i].CourseHistory); j++ {
			var course models.Course
			db.Preload("Subject").First(&course, (*users[i].CourseHistory)[j].CourseID)
			courses[i] = append(courses[i], course)
		}
	}
	for i := 0; i < len(courses); i++ {
		for j := 0; j < len(courses[i]); j++ {
			fmt.Println(users[i].User_id, courses[i][j].Category)
		}
	}
}

// func distance(a, b []float64) float64 {
// 	var sum float64
// 	for i := 0; i < len(a); i++ {
// 		sum += (a[i] - b[i]) * (a[i] - b[i])
// 	}
// 	return sum
// }

func RandomLikeCourse(db *gorm.DB) {
	var users []models.User
	db.Preload("LikeCourses").Find(&users)
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < len(users); i++ {
		ranInt := make([]int, 11)
		for k := range ranInt {
			ranInt[k] = rand.Intn(20) + 1
		}
		var ranCourse []*models.Course
		db.Find(&ranCourse, ranInt)
		users[i].LikeCourses = append(users[i].LikeCourses, ranCourse...)
	}
	db.Save(&users)
}
