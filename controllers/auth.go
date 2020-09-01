package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielwetan/bonjour-go/helpers"
	"github.com/danielwetan/bonjour-go/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	helpers.Headers(&w)

	if r.Method == "POST" {
		r.ParseForm()

		db, err := helpers.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		name, email, password := r.FormValue("name"), r.FormValue("email"), r.FormValue("password")
		hash, _ := helpers.HashPassword(password)
		_, err = db.Exec(helpers.Query["register"], name, email, hash)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		body := "Register success"
		res := helpers.ResponseMsg(true, body)
		json.NewEncoder(w).Encode(res)
	} else {
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	helpers.Headers(&w)

	if r.Method == "POST" {
		r.ParseForm()

		db, err := helpers.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		var result = models.User{}
		email, password := r.FormValue("email"), r.FormValue("password")
		err = db.
			QueryRow(helpers.Query["login"], email).
			Scan(&result.ID, &result.Name, &result.Email, &result.Password, &result.ProfileImg, &result.About, &result.CreatedAt, &result.UpdatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		match := helpers.CheckPasswordHash(password, result.Password)
		if match {
			res := helpers.ResponseMsg(true, result)
			json.NewEncoder(w).Encode(res)
		} else {
			body := "Username or password is wrong"
			res := helpers.ResponseMsg(false, body)
			json.NewEncoder(w).Encode(res)
		}
	} else {
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
	}
}
