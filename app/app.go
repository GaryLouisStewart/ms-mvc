package app

import (
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.getElement)
}
