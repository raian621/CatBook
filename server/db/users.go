package db

import (
	"database/sql"
	"errors"

	"catbook.com/auth"
)

func CreateUser(userCreds *auth.UserCredentials, email string, database *sql.DB) (err error) {
	if UserOrEmailExists(userCreds.Username, email, database) {
		err = errors.New("User already exists!")
	}
	return err
}

func DeleteUser(userCreds *auth.UserCredentials, database *sql.DB) error {
	return nil
}

func UserOrEmailExists(username, email string, database *sql.DB) bool {
	_, err := database.Query(`
		SELECT username FROM users WHERE username=? OR email=?;
	`, username, email)

	if err == sql.ErrNoRows {
		return false
	}
	return true
}
