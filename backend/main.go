package main

import (
	"bike_renting_system/api"
	"bike_renting_system/router"

	_ "github.com/lib/pq"
)

func main() {

	router := router.SetupRouter()

	router.POST("/login", api.HandleLogin)
	router.GET("/logout", api.HandleLogout)
	router.POST("/register", api.HandleRegister)
	router.GET("/getCookie", api.GetCookieSession)

	router.GET("/bikes", api.GetBikesAPI)
	router.GET("/bikes/:id", api.BikeByIdAPI)
	router.PATCH("/rent/:id", api.RentBikeAPI)
	router.PATCH("/return/:id", api.ReturnBikeAPI)
	router.Run("localhost:3000")
}
