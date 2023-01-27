package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterStudent(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := constants.Validate.Struct(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		user.Password = string(encryptPassword)

		if err := db.Model(&models.User{}).Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var no_id int32 = -1
		tokenString, err := constants.GenerateToken(&user.User_id, &user.Role, &no_id)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		refeshTokenString, err := constants.GenerateRefreshToken(&user.User_id, &user.Role, &no_id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		c.Set("authorization", tokenString)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"token":        &tokenString,
			"refreshToken": &refeshTokenString,
		})
	}
}

func RegisterTeacher(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *models.User
		var userTeacher *models.UserTeacher
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := constants.Validate.Struct(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		if err := c.BodyParser(&userTeacher); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := constants.Validate.Struct(userTeacher); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		user.Password = string(encryptPassword)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		tx := db.Begin()
		if err := tx.Model(&models.User{}).Create(&user).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		userTeacher.UserID = user.User_id

		if err := tx.Model(&models.UserTeacher{}).Create(&userTeacher).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		tokenString, err := constants.GenerateToken(&user.User_id, &user.Role, &userTeacher.Teacher_id)
		if err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		refeshTokenString, err := constants.GenerateRefreshToken(&user.User_id, &user.Role, &userTeacher.Teacher_id)
		if err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		tx.Commit()
		c.Set("authorization", tokenString)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"token":        &tokenString,
			"refreshToken": &refeshTokenString,
		})
	}
}

func LoginUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *models.User
		var err error

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		password := user.Password

		if err := db.Model(&models.User{}).Preload("Teacher").Where("email = ?", user.Email).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var tokenString string
		var no_id int32 = 0

		if user.Role == "teacher" || user.Role == "Teacher" || user.Role == "tutor" || user.Role == "Tutor" {
			tokenString, err = constants.GenerateToken(&user.User_id, &user.Role, &user.Teacher.Teacher_id)
		} else {
			tokenString, err = constants.GenerateToken(&user.User_id, &user.Role, &no_id)
		}

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var refeshTokenString string

		if user.Role == "teacher" || user.Role == "Teacher" || user.Role == "tutor" || user.Role == "Tutor" {
			refeshTokenString, err = constants.GenerateRefreshToken(&user.User_id, &user.Role, &user.Teacher.Teacher_id)
		} else {
			refeshTokenString, err = constants.GenerateRefreshToken(&user.User_id, &user.Role, &no_id)
		}

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		c.Set("authorization", tokenString)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token":        &tokenString,
			"refreshToken": &refeshTokenString,
		})
	}
}

func Update(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if user.Password != "" {
			encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
			user.Password = string(encryptPassword)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
		}

		if err := db.Model(&models.User{}).Where("user_id = ?", c.Params("id")).Updates(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Update Success",
		})
	}
}

func GetProfile(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *models.User
		if err := db.Model(&models.User{}).Omit("password").Preload("Teacher").Preload("Teacher.Courses").Preload("History").Preload("History.Video").Where("user_id = ?", c.Params("id")).First(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// var logindata = &models.LoginData{
		// 	Fullname: user.Fullname,
		// 	Nickname: user.Nickname,
		// 	Birthday: user.Birthday,
		// 	Role:     user.Role,
		// 	Picture:  user.Picture,
		// 	Point:    user.Point,
		// 	History:  history, //waiting history api
		// 	Teacher:  user.Teacher,
		// }
		if len(*user.History) == 0 {
			user.History = nil
		}

		return c.Status(fiber.StatusOK).JSON(&user)
	}
}

func GetTeacher(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var result []map[string]interface{}
		db.Raw(`SELECT MIN(users.user_id) as user_id, MIN(user_teachers.teacher_id) as teacher_id, MIN(nickname) as nickname, MIN(fullname) as fullname, MIN(class_level) as class_level, MIN(users.picture) as picture, COALESCE(AVG(rating), 0.0) AS rating
		FROM users JOIN user_teachers on users.user_id = user_teachers.user_id 
		JOIN courses ON user_teachers.teacher_id = courses.teacher_id
		JOIN subjects ON courses.subject_id = subjects.subject_id
		LEFT JOIN review_tutors ON user_teachers.teacher_id = review_tutors.teacher_id
		WHERE role = 'Teacher' OR role = 'Tutor'
		GROUP BY users.user_id
		ORDER BY rating DESC`).Find(&result)
		return c.Status(fiber.StatusOK).JSON(&result)
	}
}
func GetTeacherByClassLevel(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
		var result []map[string]interface{}
        db.Table("users").
        Select("users.user_id, user_teachers.teacher_id, nickname, fullname, subjects.class_level, users.picture, COALESCE(AVG(rating), 0.0) AS rating").
        Joins("JOIN user_teachers on users.user_id = user_teachers.user_id").
        Joins("JOIN courses ON user_teachers.teacher_id = courses.teacher_id").
        Joins("JOIN subjects ON courses.subject_id = subjects.subject_id").
        Joins("LEFT JOIN review_tutors ON user_teachers.teacher_id = review_tutors.teacher_id").
        Where("role = 'Teacher' OR role = 'Tutor'").
        Group("users.user_id, user_teachers.teacher_id, subjects.class_level, nickname, fullname, users.picture").
        Having("subjects.class_level = ?", c.Params("class_level")).
        Order("rating DESC").
        Find(&result)
        return c.Status(fiber.StatusOK).JSON(&result)
    }
}