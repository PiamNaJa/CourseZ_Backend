package repository

import (
	"time"

	"github.com/PiamNaJa/CourseZ_Backend/models"
	"gorm.io/gorm"
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
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return &userRopo{db: db}
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
	err := u.db.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRopo) GetAllUser() (*[]models.User, error) {
	var users []models.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}