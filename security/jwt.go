package security

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
	"timewise/model"
)

//JwtKey secret key for encode and decode JWT
const JwtKey = "3545kdkdajdajdjid393929090"

//GenToken create new token
func GenToken(user model.User) (string, error) {
	jwtExpConfig := os.Getenv("JwtExpires")
	jwtExpValue, _ := strconv.Atoi(jwtExpConfig)
	jwtExpDuration := time.Hour * time.Duration(jwtExpValue)
	claims := &model.JwtCustomClaims{
		UserId: user.UserID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtExpDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
