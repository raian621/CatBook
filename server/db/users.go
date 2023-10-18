package db

import "catbook.com/auth"

func CreateUser(userCreds auth.UserCredentials) (err error) {
	return err
}

func DeleteUser(userCreds auth.UserCredentials) error {
	return nil
}

func UserExists(username string) bool {
	return false
}
