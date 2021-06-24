package tools

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

const (
	secret = "3jazzhnGyh+q9d46eMrq9YJ07xvURaqWDqjTWL8/5RYEjU6AetMYsg=="
	emailKey = "email"
	iatKey = "iat"
	expKey = "exp"
)

func CreateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims[emailKey] = email
	claims[iatKey] = time.Now().Unix()
	claims[expKey] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return t, err
	}
	return t, nil
}

func DecodeToken(tokenString string) (string, error) {
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return "", errors.Wrapf(err, "%s is expired", tokenString)
			} else {
				return "", errors.Wrapf(err, "%s is invalid", tokenString)
			}
		} else {
			return "", errors.Wrapf(err, "%s is invalid", tokenString)
		}
	}

	if t == nil {
		return "", errors.Errorf("not found token in %s:", tokenString)
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.Errorf("not found claims in %s", tokenString)
	}

	email, ok := claims[emailKey].(string)
	if !ok {
		return "", errors.Errorf("not found %s in %s", emailKey, tokenString)
	}

	return email, nil
}