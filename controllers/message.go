package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielwetan/bonjour-go/models"

	"github.com/danielwetan/bonjour-go/helpers"
)

func Message(w http.ResponseWriter, r *http.Request) {
	helpers.Headers(&w)

	// id, senderId, receiverId := r.URL.Query()["id", "sender_id", "receiver_id"]
	id := r.URL.Query()["id"]
	sender := r.URL.Query()["sender"]
	receiver := r.URL.Query()["receiver"]

	db, err := helpers.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	if r.Method == "GET" {
		// Scan(&data.ID, &data.ReceiverID, &data.SenderiID, &data.SenderName, &data.ProfileImg, &data.Message, &data.CreatedAt)

		// Latest message
		switch {
		case id != nil:
			{
				// Latest Messages
				rows, err := db.Query(helpers.Query["latestMessages"], id[0])
				if err != nil {
					panic(err)
				}
				var result []interface{}
				defer rows.Close()
				for rows.Next() {
					data := models.LatestMessages{}
					err = rows.Scan(&data.ID, &data.ReceiverID, &data.SenderiID, &data.SenderName, &data.ProfileImg, &data.Message, &data.CreatedAt)
					result = append(result, data)
					if err != nil {
						panic(err)
					}
				}
				err = rows.Err()
				if err != nil {
					panic(err)
				}

				res := helpers.ResponseMsg(true, result)
				json.NewEncoder(w).Encode(res)
			}
		case sender != nil && receiver != nil:
			{
				// Conversation
				rows, err := db.Query(helpers.Query["conversation"], sender[0], receiver[0], sender[0], receiver[0])
				if err != nil {
					panic(err)
				}
				var result []interface{}
				defer rows.Close()
				for rows.Next() {
					data := models.Conversation{}
					err = rows.Scan(&data.ID, &data.SenderID, &data.ReceiverID, &data.Message, &data.CreatedAt)
					result = append(result, data)
					if err != nil {
						panic(err)
					}
				}
				err = rows.Err()
				if err != nil {
					panic(err)
				}

				res := helpers.ResponseMsg(true, result)
				json.NewEncoder(w).Encode(res)
			}
		default:
			// Error handling when query params is not exist
			body := "Invalid URL"
			res := helpers.ResponseMsg(false, body)
			json.NewEncoder(w).Encode(res)
		}

	} else if r.Method == "POST" {
		senderID, receiverID, message := r.FormValue("sender_id"), r.FormValue("receiver_id"), r.FormValue("message")
		_, err = db.Exec(helpers.Query["postMessage"], senderID, receiverID, message)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		body := "Send message success"
		res := helpers.ResponseMsg(true, body)
		json.NewEncoder(w).Encode(res)
	} else {
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
	}

}
