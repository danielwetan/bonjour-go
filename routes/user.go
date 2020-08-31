package routes

import (
	"net/http"

	"github.com/danielwetan/bonjour-go/controllers"
)

func User() {
	http.HandleFunc("/u", controllers.User)

}
