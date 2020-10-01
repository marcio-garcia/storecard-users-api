package app

import "github.com/marcio-garcia/storecard-users-api/controllers"

func mapRoutes() {
	router.GET("/ping", controllers.Ping)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:user_id", controllers.GetUser)
}
