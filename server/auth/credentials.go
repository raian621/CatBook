package "auth"

type UserCredentials struct {
	username string `json:username`
	password string `json:password`
}