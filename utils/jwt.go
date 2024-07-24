package utils

import (
	"davinci-chat/config"
	"davinci-chat/consts"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func MakeJwt(userName string, userEmail string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userName
	claims["email"] = userEmail
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(config.JWTSecretKey))
}

func DecodeJwt(tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.JWTSecretKey), nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", errors.New("invalid token")
	}

	name, okName := claims["name"].(string)
	email, okEmail := claims["email"].(string)
	if !okName || !okEmail {
		return "", "", errors.New("invalid token claims")
	}

	return name, email, nil
}

func MakeJwtCookie(token string) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	cookie.Secure = true
	cookie.SameSite = "Lax"
	if runEnv := config.GetRunEnv(); runEnv == consts.Production {
		cookie.Domain = ".songmingi.com"
	}

	return cookie
}

func GetUserName(c *websocket.Conn) (string, error) {
	tokenString := c.Cookies("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(config.JWTSecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userName, ok := claims["name"].(string)
		if !ok {
			return "", errors.New("no user name in token")
		}
		return userName, nil
	}

	return "", errors.New("can not get user name from token")
}
