package auth

import (
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xadhithiyan/videon/config"
)

func CreateJWT(key string, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Env.JWTExpirationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(userID),
		"exp":    time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AuthenticateJwt(tokenStr string) (int, bool) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Env.JWTSecrect), nil
	})

	if err != nil || !token.Valid {
		return 0, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, false
	}

	userIDStr, ok := claims["userID"].(string)
	if !ok {
		log.Print("Error: userID claim is not a string")
		return 0, false
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Print("Error converting userID to integer: ", err)
		return 0, false
	}

	return userID, true
}
