package handlers

import (
	"os"
	"time"

	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func RegisterStudent(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := utils.Validate.Struct(user); err != nil {
			return utils.BadRequest(err.Error())
		}

		encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return utils.Unexpected(err.Error())
		}

		user.Password = string(encryptPassword)
		if err := db.Exec("SELECT setval(pg_get_serial_sequence('users', 'user_id'), coalesce(max(user_id)+1, 1), false) FROM users;").Error; err != nil {
			return utils.Unexpected(err.Error())
		}
		if err := db.Create(&user).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		var no_id int32 = -1
		tokenString, err := utils.GenerateToken(&user.User_id, &user.Role, &no_id)

		if err != nil {
			return utils.BadRequest(err.Error())
		}

		refeshTokenString, err := utils.GenerateRefreshToken(&user.User_id, &user.Role, &no_id)
		if err != nil {
			return utils.BadRequest(err.Error())
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
		var user models.User
		var userTeacher models.UserTeacher
		if err := c.BodyParser(&user); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := utils.Validate.Struct(user); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := c.BodyParser(&userTeacher); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := utils.Validate.Struct(userTeacher); err != nil {
			return utils.BadRequest(err.Error())
		}

		encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		user.Password = string(encryptPassword)
		if err != nil {
			return utils.Unexpected(err.Error())
		}

		tx := db.Begin()
		if err := tx.Exec("SELECT setval(pg_get_serial_sequence('users', 'user_id'), coalesce(max(user_id)+1, 1), false) FROM users;").Error; err != nil {
			return utils.Unexpected(err.Error())
		}
		if err := tx.Create(&user).Error; err != nil {
			tx.Rollback()
			return utils.Unexpected(err.Error())
		}

		userTeacher.UserID = user.User_id

		if err := tx.Create(&userTeacher).Error; err != nil {
			tx.Rollback()
			return utils.Unexpected(err.Error())
		}

		tokenString, err := utils.GenerateToken(&user.User_id, &user.Role, &userTeacher.Teacher_id)
		if err != nil {
			tx.Rollback()
			return utils.Unexpected(err.Error())
		}

		refeshTokenString, err := utils.GenerateRefreshToken(&user.User_id, &user.Role, &userTeacher.Teacher_id)
		if err != nil {
			tx.Rollback()
			return utils.Unexpected(err.Error())
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
		var user models.User
		var err error

		if err := c.BodyParser(&user); err != nil {
			return utils.BadRequest(err.Error())
		}

		password := user.Password

		if err := db.Preload("Teacher").Where("email = ?", &user.Email).First(&user).Error; err != nil {
			return utils.HandleFindError(err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return utils.Unexpected(err.Error())
		}

		var tokenString string
		var no_id int32 = 0

		if user.Role == "teacher" || user.Role == "Teacher" || user.Role == "tutor" || user.Role == "Tutor" {
			tokenString, err = utils.GenerateToken(&user.User_id, &user.Role, &user.Teacher.Teacher_id)
		} else {
			tokenString, err = utils.GenerateToken(&user.User_id, &user.Role, &no_id)
		}

		if err != nil {
			return utils.Unexpected(err.Error())
		}

		var refeshTokenString string

		if user.Role == "teacher" || user.Role == "Teacher" || user.Role == "tutor" || user.Role == "Tutor" {
			refeshTokenString, err = utils.GenerateRefreshToken(&user.User_id, &user.Role, &user.Teacher.Teacher_id)
		} else {
			refeshTokenString, err = utils.GenerateRefreshToken(&user.User_id, &user.Role, &no_id)
		}

		if err != nil {
			return utils.Unexpected(err.Error())
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
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return utils.BadRequest(err.Error())
		}

		if user.Password != "" {
			encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
			user.Password = string(encryptPassword)
			return utils.Unexpected(err.Error())
		}

		if err := db.Where("user_id = ?", c.Params("user_id")).Updates(&user).Error; err != nil {
			return utils.HandleFindError(err)
		}

		return c.Status(fiber.StatusOK).JSON("Updated")
	}
}

func GetProfile(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		if err := db.Omit("password").Preload(clause.Associations).Preload("Teacher.Courses").Preload("Teacher.Tracsaction").Preload("CourseHistory.Course").Preload("VideoHistory.Video").Where("user_id = ?", c.Params("user_id")).First(&user).Error; err != nil {
			return utils.HandleFindError(err)
		}
		for i := 0; i < len(*user.CourseHistory); i++ {
			for j := i; j < len(*user.CourseHistory); j++ {
				temp := (*user.CourseHistory)[i]
				if int32((*user.CourseHistory)[i].UpdateAt) > int32((*user.CourseHistory)[j].UpdateAt) {
					(*user.CourseHistory)[i] = (*user.CourseHistory)[j]
					(*user.CourseHistory)[j] = temp
				}
			}
		}
		for i := 0; i < len(*user.VideoHistory); i++ {
			for j := i; j < len(*user.VideoHistory); j++ {
				temp := (*user.VideoHistory)[i]
				if int32((*user.VideoHistory)[i].UpdateAt) > int32((*user.VideoHistory)[j].UpdateAt) {
					(*user.VideoHistory)[i] = (*user.VideoHistory)[j]
					(*user.VideoHistory)[j] = temp
				}
			}
		}
		user.Address = nil
		return c.Status(fiber.StatusOK).JSON(&user)
	}
}

func GetTeacher(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var result []map[string]interface{}
		db.Raw(`SELECT MIN(users.user_id) as user_id, MIN(user_teachers.teacher_id) as teacher_id, MIN(nickname) as nickname, MIN(fullname) as fullname, MIN(class_level) as class_level, MIN(users.picture) as picture, COALESCE(AVG(review_tutors.rating), 0.0) AS rating
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
		if err := db.Table("users").
			Select("users.user_id, user_teachers.teacher_id, nickname, fullname, subjects.class_level, users.picture, COALESCE(AVG(review_tutors.rating), 0.0) AS rating").
			Joins("JOIN user_teachers on users.user_id = user_teachers.user_id").
			Joins("JOIN courses ON user_teachers.teacher_id = courses.teacher_id").
			Joins("JOIN subjects ON courses.subject_id = subjects.subject_id").
			Joins("LEFT JOIN review_tutors ON user_teachers.teacher_id = review_tutors.teacher_id").
			Where("role = 'Teacher' OR role = 'Tutor'").
			Group("users.user_id, user_teachers.teacher_id, subjects.class_level, nickname, fullname, users.picture").
			Having("subjects.class_level = ?", c.Params("class_level")).
			Order("rating DESC").
			Find(&result).Error; err != nil {
			return utils.HandleFindError(err)
		}
		if result == nil {
			result = []map[string]interface{}{}
		}
		return c.Status(fiber.StatusOK).JSON(&result)
	}
}

func GetNewToken(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rT map[string]string
		if err := c.BodyParser(&rT); err != nil {
			return utils.BadRequest(err.Error())
		}
		refreshToken := rT["token"]
		claims := jwt.MapClaims{}
		tkn, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_REFRESH_Secret")), nil
		},
		)
		if err != nil {
			return utils.Unauthorized(err.Error())
		}
		if !tkn.Valid {
			return utils.Unauthorized("Token is not valid")
		}

		// Check if the token has expired
		exp := claims["exp"].(float64)
		if int64(exp) < time.Now().Unix() {
			return utils.Unauthorized("Token is expired")
		}

		user_id := int32(claims["user_id"].(float64))
		role := claims["role"].(string)
		teacher_id := int32(claims["teacher_id"].(float64))
		var user models.User
		if err := db.Where("user_id = ?", &user_id).First(&user).Error; err != nil {
			return utils.HandleFindError(err)
		}
		tokenString, err := utils.GenerateToken(&user_id, &role, &teacher_id)
		if err != nil {
			return utils.Unexpected(err.Error())
		}
		refeshTokenString, err := utils.GenerateRefreshToken(&user_id, &role, &teacher_id)
		if err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token":        &tokenString,
			"refreshToken": &refeshTokenString,
		})
	}
}

func GetTeacherById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		if err := db.Omit("password").Preload("Teacher").Preload("Teacher.Reviews").Preload("Teacher.Courses").Joins("join user_teachers on users.user_id = user_teachers.user_id").Where("teacher_id = ?", c.Params("teacher_id")).First(&user).Error; err != nil {
			return utils.HandleFindError(err)
		}
		return c.Status(fiber.StatusOK).JSON(&user)
	}
}

func CreateOrUpdateUserAddress(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var address models.Address

		user_id, err := c.ParamsInt("user_id")
		if err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.FirstOrCreate(&address, &models.Address{UserID: int32(user_id)}).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		if err := c.BodyParser(&address); err != nil {
			return utils.BadRequest(err.Error())
		}

		address.UserID = int32(user_id)

		if err := utils.Validate.Struct(address); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.Save(&address).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(&address)
	}
}

func GetUserAddress(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var address models.Address
		if err := db.Where("user_id = ?", c.Params("user_id")).First(&address).Error; err != nil {
			return utils.HandleFindError(err)
		}
		return c.Status(fiber.StatusOK).JSON(&address)
	}
}
