package app

import (
	"github.com/dieg0code/go-microservices/controllers"
)

func mapUrls() {
	router.GET("/users:user_id", controllers.GetUser)
}
