package app

import (
	"net/http"

	"github.com/GaryLouisStewart/ms-mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/elements", controllers.GetElement)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
