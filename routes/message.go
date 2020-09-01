package routes

import (
	"net/http"

	"github.com/danielwetan/bonjour-go/controllers"
)

func Message() {
	http.HandleFunc("/msg", controllers.Message)
}
