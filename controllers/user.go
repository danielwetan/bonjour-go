package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielwetan/bonjour-go/models"

	"github.com/danielwetan/bonjour-go/helpers"
)

func User(w http.ResponseWriter, r *http.Request) {
	helpers.Headers(&w)
	id := r.URL.Query()["id"]

	db, err := helpers.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	if r.Method == "GET" {
		data := models.Contact{}
		err = db.
			QueryRow(helpers.Query["getUser"], id[0]).
			Scan(&data.ID, &data.Name, &data.Email, &data.ProfileImg, &data.About)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		res := helpers.ResponseMsg(true, data)
		json.NewEncoder(w).Encode(res)

	} else if r.Method == "PATCH" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		_, err = db.Exec(helpers.Query["updateUser"], name, email, id[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		body := "Update success"
		res := helpers.ResponseMsg(true, body)
		json.NewEncoder(w).Encode(res)

	} else {
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
	}
}
