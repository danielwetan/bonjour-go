package routes

import (
	"net/http"

	"github.com/danielwetan/bonjour-go/controllers"
)

func Auth() {
	http.HandleFunc("/auth/register", controllers.Register)
	http.HandleFunc("/auth/login", controllers.Login)
}
