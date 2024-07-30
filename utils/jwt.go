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

func MakeJwt(userName string, userEmail string, isGuest bool) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userName
	claims["email"] = userEmail

	if isGuest {
		claims["isGuest"] = "true"
	} else {
		claims["isGuest"] = "false"
	}

	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(config.JWTSecretKey))
}

func DecodeJwt(tokenString string) (string, string, bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.JWTSecretKey), nil
	})

	if err != nil {
		return "", "", false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", false, errors.New("invalid token")
	}

	name, okName := claims["name"].(string)
	email, okEmail := claims["email"].(string)
	isGuest, okIsGuest := claims["isGuest"].(string)
	if !okName || !okEmail || !okIsGuest {
		return "", "", false, errors.New("invalid token claims")
	}

	return name, email, isGuest == "true", nil
}

func SetJwtCookie(token string) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	cookie.Secure = true

	runEnv := config.GetRunEnv()
	switch runEnv {
	case consts.Production:
		cookie.SameSite = "None"
		cookie.Domain = "songmingi.com"
	default:
		cookie.SameSite = "Lax"
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

func GetUserEmail(c *websocket.Conn) (string, error) {
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
		userEmail, ok := claims["email"].(string)
		if !ok {
			return "", errors.New("no user email in token")
		}
		return userEmail, nil
	}

	return "", errors.New("can not get user email from token")
}

func GetIsGuest(c *websocket.Conn) (bool, error) {
	tokenString := c.Cookies("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(config.JWTSecretKey), nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		isGuest, ok := claims["isGuest"].(string)
		if !ok {
			return false, errors.New("no user email in token")
		}

		if isGuest == "true" {
			return true, nil
		} else {
			return false, nil
		}
	}

	return false, errors.New("can not get user email from token")
}