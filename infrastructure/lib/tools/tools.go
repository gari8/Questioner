package tools

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"

func CreateId() (*string, error) {
	u, err := uuid.NewRandom()
	if err != nil { return nil, err }
	id := u.String()
	return &id, nil
}

func MuskWord(word string) *string {
	musk := strings.Repeat("*", len(word))
	return &musk
}

func DigestWord(word string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(word),12)
	return string(hash), err
}

func CompareHashAndPlane(hash string, plane string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plane))
	if err != nil {
		return false
	}
	return true
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func ConvertStringToTime(timeString string) (*time.Time, error) {
	t, err := time.Parse(TimeLayout, timeString)
	if err != nil { return nil, err }
	return &t, nil
}