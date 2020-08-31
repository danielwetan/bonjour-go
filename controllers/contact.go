package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielwetan/bonjour-go/models"

	"github.com/danielwetan/bonjour-go/helpers"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id, _ := r.URL.Query()["id"]

		db, err := helpers.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		rows, err := db.Query(helpers.Query["getContact"], id[0])
		if err != nil {
			panic(err)
		}
		var result []interface{}
		defer rows.Close()
		for rows.Next() {
			data := models.Contact{}
			err = rows.Scan(&data.ID, &data.Name, &data.Email, &data.ProfileImg, &data.About)
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
	} else {
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
