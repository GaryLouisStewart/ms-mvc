package app

import (
	"github.com/GaryLouisStewart/ms-mvc/controllers"
)

func mapUrls() {
	router.GET("/elements/:element_id", controllers.GetElement)
}
