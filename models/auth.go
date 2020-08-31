package models

type Register struct {
	Name       string
	Email      string
	Password   string
	ProfileImg string
	About      string
}

type Login struct {
	Email    string
	Password string
}
