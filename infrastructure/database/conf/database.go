package conf

import (
	"faves4/ent"

	_ "github.com/lib/pq"

	"os"
)

var (
	source = os.Getenv("DATABASE_URL")
	driver = os.Getenv("DRIVER")
)

func NewDatabaseConnection() (*ent.Client, error) {
	client, err := ent.Open(driver, source)
	if err != nil {
		return nil, err
	}
	return client, nil
}
