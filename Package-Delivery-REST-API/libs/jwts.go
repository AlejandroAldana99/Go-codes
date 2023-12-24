package libs

import (
	"fmt"
	"time"

	"github.com/AlejandroAldana99/mvp_api/constants"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/golang-jwt/jwt/v4"
)

func DecodeJWT(token string) (string, string, error) {
	hmacSecret := []byte(constants.Key)

	tokenD, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if claims, ok := tokenD.Claims.(jwt.MapClaims); ok {
		return claims["role"].(string), claims["userid"].(string), nil
	} else {
		return constants.GenericName, "", err
	}
}

func EncodeJWT(id string, role string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": id,
		"role":   role,
		"nbf":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, tErr := token.SignedString([]byte(constants.Key))
	if tErr != nil {
		logger.Error("services", "Login", tErr.Error())
		return ""
	}

	return tokenString
}
