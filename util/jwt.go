package util

import (
	"go-risky/types"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userInput types.User) (string, error) {
	var sampleSecretKey = []byte("SecretYouShouldHide")
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["userId"] = userInput.ID
	claims["email"] = userInput.Email

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

// TODO: This needs work on claim validation
func ValidateJWT(tokenString string) (bool, error) {
	var sampleSecretKey = []byte("SecretYouShouldHide")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return sampleSecretKey, nil
	})

	if err != nil {
		return false, err
	}

	if token.Valid {
		return true, nil
	}

	return false, nil
}
