package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielwetan/bonjour-go/helpers"
	"github.com/danielwetan/bonjour-go/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		type LoginData struct {
			email    string
			password string
		}

		db, err := helpers.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		var result = models.User{}
		email, password := r.FormValue("email"), r.FormValue("password")
		err = db.
			QueryRow("SELECT * FROM users WHERE email = ?", email).
			Scan(&result.ID, &result.Name, &result.Email, &result.Password, &result.ProfileImg, &result.About, &result.CreatedAt, &result.UpdatedAt)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Compare password dengan bcrypt
		match := CheckPasswordHash(password, result.Password)
		// Jika true maka berikan response
		if match {
			type Res struct {
				Status bool        `json:"status"`
				Body   models.User `json:"body"`
			}
			res := &Res{
				Status: true,
				Body: models.User{
					ID:         result.ID,
					Name:       result.Name,
					Email:      result.Email,
					ProfileImg: result.ProfileImg,
					About:      result.About,
					CreatedAt:  result.CreatedAt,
					UpdatedAt:  result.UpdatedAt,
				},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
		} else {
			body := "Username or password is wrong"
			res := helpers.ResponseFailed(body)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
		}
	} else {
		http.Error(w, "Something error", http.StatusBadRequest)
	}
}
