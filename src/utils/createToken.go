package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID int) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	expMinutesStr := os.Getenv("JWT_EXPIRE_MINUTES")

	// default 60 minutes if env not set or invalid
	expMinutes, err := strconv.Atoi(expMinutesStr)
	if err != nil || expMinutes <= 0 {
		expMinutes = 60
	}

	expDuration := time.Duration(expMinutes) * time.Minute

	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprint(userID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expDuration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(jwtSecret))
}
