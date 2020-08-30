package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielwetan/bonjour-go/helpers"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		r.ParseForm()
		fmt.Println(r.FormValue("name"))

		type Response struct {
			status string
			body   string
		}

		db, err := helpers.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		name, email, password := r.FormValue("name"), r.FormValue("email"), r.FormValue("password")
		hash, _ := HashPassword(password)
		_, err = db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", name, email, hash)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data := []struct {
			Status string
			Body   string
		}{
			{"true", "Register success"},
		}

		jsonInBytes, err := json.Marshal(data[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonInBytes)

		fmt.Println("Register success!")
	} else {
		http.Error(w, "Something error", http.StatusBadRequest)
	}
}
