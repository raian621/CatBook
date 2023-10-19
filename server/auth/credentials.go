package auth

import (
	"database/sql"

	"catbook.com/util"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegistrationInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func ValidCredentials(
	creds *UserCredentials,
	database *sql.DB,
	userExistsCB func(string) bool,
) (bool, error) {
	if userExistsCB(creds.Username) {
		return false, nil
	}

	row := database.QueryRow(`
		SELECT passhash FROM users WHERE username=?;
	`, creds.Username)

	var passhash string
	if err := row.Scan(&passhash); err == sql.ErrNoRows {
		return false, nil
	}

	return util.MatchHash(creds.Password, passhash)
}
