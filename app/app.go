package app

import (
	"net/http"

	"github.com/GaryLouisStewart/ms-mvc/app/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetElement)
}
