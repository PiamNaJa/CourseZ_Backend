package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user_id *int32, role *string, teacher_id *int32) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id":    &user_id,
		"role":       &role,
		"teacher_id": &teacher_id,
		"exp":        expireTime.Unix(),
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return token, err
}

func GenerateRefreshToken(user_id *int32, role *string, teacher_id *int32) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(72 * time.Hour)

	claims := jwt.MapClaims{
		"user_id":    &user_id,
		"role":       &role,
		"teacher_id": &teacher_id,
		"exp":        expireTime.Unix(),
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refeshToken, err := tokenClaims.SignedString([]byte(os.Getenv("JWT_REFRESH_Secret")))

	return refeshToken, err
}
func GetClaims(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	return claims, err
}
