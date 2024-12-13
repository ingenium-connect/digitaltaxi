package utils

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func FormatPhoneNumber(msisdn string) string {
	switch {
	case len(msisdn) == 0:
		return msisdn
	case msisdn[0] == '0' && len(msisdn) > 3:
		return "254" + msisdn[1:]
	case msisdn[0] == '+' && len(msisdn) > 3:
		return msisdn[1:]
	case strings.HasPrefix(msisdn, "254") && len(msisdn) > 3:
		return msisdn
	case !IsDigit(msisdn[0]):
		return msisdn
	default:
		return "254" + msisdn
	}
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func RemoveDoubleWhitespace(input string) string {
	// Use a regular expression to replace all whitespace characters with an empty string
	input = strings.ReplaceAll(input, "  ", " ")
	input = strings.ReplaceAll(input, "  ", " ")
	re := regexp.MustCompile(`\s`)

	return re.ReplaceAllString(input, "")
}

func ValidateToken(token string, publicKey string) (interface{}, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)

	if err != nil {
		return "", fmt.Errorf("validate: parse key: %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}

		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	return claims["sub"], nil
}
