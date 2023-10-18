package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseParams struct {
	Provider string
	Hostname string
	Username string
	Password string
	Database string
	Port     int
	SSLMode  string
}

var database *sql.DB

func databaseURL(dbParams *DatabaseParams) (url string) {
	url = fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/sslmode=%s",
		dbParams.Username,
		dbParams.Password,
		dbParams.Hostname,
		dbParams.Port,
		dbParams.Database,
	)

	return url
}

func ConnectToDB(dbParams *DatabaseParams) (err error) {
	url := databaseURL(dbParams)
	database, err = sql.Open(dbParams.Provider, url)

	return err
}

func GetDatabase() *sql.DB {
	return database
}
