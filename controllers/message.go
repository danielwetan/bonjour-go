package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielwetan/bonjour-go/models"

	"github.com/danielwetan/bonjour-go/helpers"
)

func Message(w http.ResponseWriter, r *http.Request) {
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
				w.Header().Set("Content-Type", "application/json")
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
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(res)
			}
		default:
			body := "Invalid URL"
			res := helpers.ResponseMsg(false, body)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
		}

	} else if r.Method == "POST" {

	} else {
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}

}
