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

		// var data = LoginData{}

		// email := r.FormValue("email")
		// fmt.Println(email)
		// stmt, err := db.Prepare("SELECT * FROM users WHERE email=?")
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return
		// }

		type User struct {
			id         int
			name       string
			email      string
			password   string
			profileImg string
			about      string
			createdAt  string
			updatedAt  string
		}

		var result = User{}
		var mail = r.FormValue("email")
		err = db.
			QueryRow("SELECT * FROM users WHERE email = ?", mail).
			Scan(&result.id, &result.name, &result.email, &result.password, &result.profileImg, &result.about, &result.createdAt, &result.updatedAt)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		/* Success:

			{
			    "status": true,
			    "body": [
			        {
			            "id": 2,
			            "name": "Daniel Saputra",
			            "email": "danielwetan.io@gmail.com",
			            "profile_img": "profile.jpg",
			            "about": "Je parle français",
			            "created_at": "2020-07-26T08:11:14.000Z",
			            "updated_at": "2020-07-26T08:11:14.000Z",
			            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwibmFtZSI6IkRhbmllbCBTYXB1dHJhIiwiZW1haWwiOiJkYW5pZWx3ZXRhbi5pb0BnbWFpbC5jb20iLCJwcm9maWxlX2ltZyI6InByb2ZpbGUuanBnIiwiYWJvdXQiOiJKZSBwYXJsZSBmcmFuw6dhaXMiLCJjcmVhdGVkX2F0IjoiMjAyMC0wNy0yNlQwODoxMToxNC4wMDBaIiwidXBkYXRlZF9hdCI6IjIwMjAtMDctMjZUMDg6MTE6MTQuMDAwWiIsImlhdCI6MTU5ODg3NzYxNSwiZXhwIjoxNTk4ODgxMjE1fQ.lhqShOVzGL_8cIKhQiaVIid91xSs91nqTsJa11lNpmg"
			        }
			    ]
			}

		/* Failed:

			{
				"status": false,
				"body": "Username or password is wrong!""
			}

		*/

		type LoginSuccess struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Email      string `json:"email"`
			ProfileImg string `json:"profile_img"`
			About      string `json:"about"`
			CreatedAt  string `json:"created_at"`
			UpdatedAt  string `json:"updated_at"`
		}

		type LoginFailed struct {
			Message string `json:"msg"`
		}

		// Compare password dengan bcrypt
		match := CheckPasswordHash(r.FormValue("password"), result.password)
		// Jika true maka berikan response
		if match {

			type Res struct {
				Status       bool         `json:"status"`
				LoginSuccess LoginSuccess `json:"body"`
			}
			res := &Res{
				Status: true,
				LoginSuccess: LoginSuccess{
					ID:         result.id,
					Name:       result.name,
					Email:      result.email,
					ProfileImg: result.profileImg,
					About:      result.about,
					CreatedAt:  result.createdAt,
					UpdatedAt:  result.updatedAt,
				},
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
		} else {

			type Res struct {
				Status bool   `json:"status"`
				Body   string `json:"body"`
			}
			res := &Res{
				Status: false,
				Body:   "Username or password is wrong",
			}

			// Jika false berikan response gagal
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
		}
	} else {
		http.Error(w, "Something error", http.StatusBadRequest)
	}
}
