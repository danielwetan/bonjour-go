package helpers

var Query = map[string]string{
	"login":    "SELECT * FROM users WHERE email = ?",
	"register": "INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
}
