package helpers

var Query = map[string]string{
	"login":          "SELECT * FROM users WHERE email = ?",
	"register":       "INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
	"getContact":     "SELECT id, name, email, profile_img, about FROM users WHERE NOT id=? ORDER BY name DESC LIMIT 10",
	"getUser":        "SELECT id, name, email, profile_img, about FROM users WHERE id=?",
	"updateUser":     "UPDATE users SET name=?, email=? WHERE id = ?",
	"latestMessages": "SELECT messages.id, messages.receiver_id, messages.sender_id, users.name as sender_name, users.profile_img, messages.message, messages.created_at FROM messages INNER JOIN users ON users.id=messages.sender_id WHERE messages.id IN (SELECT MAX(id) FROM messages WHERE messages.receiver_id=? GROUP BY messages.sender_id) ORDER BY created_at DESC",
	"conversation":   "SELECT * FROM messages WHERE sender_id=? && receiver_id=? OR sender_id=? && receiver_id=? ORDER BY id",
	"postMessage":    "INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)",
}
