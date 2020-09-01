package models

type Conversation struct {
	ID         int    `json:"id"`
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Message    string `json:"message"`
	CreatedAt  string `json:"created_at"`
}

type LatestMessages struct {
	ID         int    `json:"id"`
	ReceiverID int    `json:"receiver_id"`
	SenderiID  int    `json:"sender_id"`
	SenderName string `json:"sender_name"`
	ProfileImg string `json:"profile_img"`
	Message    string `json:"message"`
	CreatedAt  string `json:"created_at"`
}
