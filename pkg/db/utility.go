package db

import (
	"fmt"
	"online-pathsaala/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPasword(password string) (hash string, err error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(byte), err
}

func ConcateArray(arr []string) (arrString string) {
	for i, val := range arr {
		comma := ""
		if i < (len(arr) - 1) {
			comma = ", "
		}
		arrString = fmt.Sprintf("%s%s%s", arrString, val, comma)
	}
	return arrString
}

func GetQueryCondition(conditions []string, startingNum int) (conditionString string) {
	for i, condi := range conditions {
		comma := ""
		if i < (len(conditions) - 1) {
			comma = ", "
		}
		conditionString = fmt.Sprintf("%s%s=$%d%s", conditionString, condi, startingNum+i, comma)
	}
	return conditionString
}

func GetQueryInCondition(paramLength int) (arrString string) {
	for i := 1; i <= paramLength; i++ {
		comma := ""
		if i < paramLength {
			comma = ", "
		}
		arrString = fmt.Sprintf("%s$%d%s", arrString, i, comma)
	}
	return
}

func VerifyPassword(hashPassword, providedPassword string) (bool, string) {

	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(providedPassword))
	check := true
	msg := ""
	if err != nil {
		msg = "Invalid provided credentials."
		check = false
	}
	return check, msg
}

func GenerateTokens(email, id string) (signedToken string, err error) {
	SECRET_KEY := os.Getenv("SECRET_KEY")
	claims := model.SignedDetails{
		Email: email,
		Id:    id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	signedToken, err = token.SignedString([]byte(SECRET_KEY))
	return signedToken, err
}

func GenerateInsertValue(fields []string) (keys, placeholders string) {
	for i, field := range fields {
		comma := ""
		if i < (len(fields) - 1) {
			comma = ", "
		}
		keys = fmt.Sprintf("%s%s%s", keys, field, comma)
		placeholders = fmt.Sprintf("%s$%d%s", placeholders, i+1, comma)
	}
	return
}

func GetLanguageArray(languages ...map[string]string) []string {
	languageArray := make([]string, 0)
	for _, lng := range languages {
		for key := range lng {
			exists := false
			for _, l := range languageArray {
				if l == key {
					exists = true
				}
			}
			if !exists {
				languageArray = append(languageArray, key)
			}
		}
	}
	return languageArray
}
