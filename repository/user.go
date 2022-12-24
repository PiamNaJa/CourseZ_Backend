package repository

import (
	"time"

	"github.com/PiamNaJa/CourseZ_Backend/configs"
	"github.com/PiamNaJa/CourseZ_Backend/models"
)

type userRepository interface {
	CreateUser(email *string, password *string, fullname *string, nickname *string, phone *string, birthday *time.Time, Role *string, picture *string, point *int32, money *int32, teacher *models.UserTeacher) (*models.User, error)
	GetAllUser() (*[]models.User, error)
	GetByIDUser(user_id *int32) (*models.User, error)
	DeleteUser(user_id *int32) (error)
	UpdateUser(user_id *int32, email *string, password *string, fullname *string, nickname *string, phone *string, birthday *time.Time, Role *string, picture *string, point *int32, money *int32, teacher *models.UserTeacher) (*models.User, error)
}

type userRopo struct {
	userRepository
}

func NewUserRepository() userRepository {
	return &userRopo{}
}

func (u *userRopo) CreateUser(email *string, password *string, fullname *string, nickname *string, phone *string, birthday *time.Time, Role *string, picture *string, point *int32, money *int32, teacher *models.UserTeacher) (*models.User, error) {
	user := models.User{
		Email:    *email,
		Password: *password,
		Fullname: *fullname,
		Nickname: *nickname,
		Phone:    *phone,
		Birthday: *birthday,
		Role:     *Role,
		Picture:  *picture,
		Point:    *point,
		Money:    *money,
		Teacher:  teacher,
	}
	err := configs.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRopo) GetAllUser() (*[]models.User, error) {
	var users []models.User
	err := configs.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}