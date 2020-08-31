package models

type Contact struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	ProfileImg string `json:"profile_img"`
	About      string `json:"about"`
}
