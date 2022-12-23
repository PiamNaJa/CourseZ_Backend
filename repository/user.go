package repository

import (
	"time"

	"github.com/PiamNaJa/CourseZ_Backend/configs"
	"github.com/PiamNaJa/CourseZ_Backend/models"
)

type userRepository interface {
	CreateUser(user_id *int64, email *string, password *string, fullname *string, nickname *string, phone *string, birthday *time.Time, Role *string, picture *string, point *string, money *string) (*models.User, error)
	GetAllUser() (*[]models.User, error)
	GetByIDUser(user_id *int64) (*models.User, error)
	DeleteUser(user_id *int64) (error)
	UpdateUser(user_id *int64, email *string, password *string, fullname *string, nickname *string, phone *string, birthday *time.Time, Role *string, picture *string, point *string, money *string) (*models.User, error)
}

type userRopo struct {
	userRepository
}

func NewUserRepository() userRepository {
	return &userRopo{}
}

func (u *userRopo) CreateUser(user_id *int64, email *string, password *string, fullname *string, nickname *string, phone *string, birthday *time.Time, Role *string, picture *string, point *string, money *string) (*models.User, error) {
	user := models.User{
		User_id:  *user_id,
		Email:    *email,
		Password: *password,
		Fullname: *fullname,
		Nickname: *nickname,
		Phone:    *phone,
		Birthday: birthday,
		Role:     *Role,
		Picture:  *picture,
		Point:    *point,
		Money:    *money,
	}
	err := configs.DB.Create(&user).Error
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