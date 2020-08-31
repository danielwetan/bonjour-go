package routes

import (
	"net/http"

	"github.com/danielwetan/bonjour-go/controllers"
)

func Contact() {
	http.HandleFunc("/contact", controllers.Contact)
	// http.HandleFunc("/contact/", controllers.Contact)

}
