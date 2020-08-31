package helpers

var Query = map[string]string{
	"login":      "SELECT * FROM users WHERE email = ?",
	"register":   "INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
	"getContact": "SELECT id, name, email, profile_img, about FROM users WHERE NOT id=? ORDER BY name DESC LIMIT 10",
	"getUser":    "SELECT id, name, email, profile_img, about FROM users WHERE id=?",
	"updateUser": "UPDATE users SET name=?, email=? WHERE id = ?",
}
