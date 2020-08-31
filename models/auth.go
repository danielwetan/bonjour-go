package models

type Register struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ProfileImg string `json:"profile_img"`
	About      string `json:"about"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"-"` // remove password from json response
	ProfileImg string `json:"profile_img"`
	About      string `json:"about"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
