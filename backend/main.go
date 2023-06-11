package main

import (
	"bike_renting_system/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// @title Bike renting service API
// @version 1.0.0
// @host localhost:5000
// @BasePath /

func main() {

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization")
	config.AllowCredentials = true
	config.AllowAllOrigins = false
	config.AllowOrigins = []string{"http://localhost:3001"}
	config.AllowOriginFunc = func(origin string) bool {
		return true
	}

	router.Use(cors.New(config))

	store := cookie.NewStore([]byte("your-secret-key"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24})
	router.Use(sessions.Sessions("sessions", store))

	router.POST("/login", api.HandleLogin)
	router.GET("/logout", api.HandleLogout)
	router.POST("/register", api.HandleRegister)
	router.GET("/getCookie", api.GetCookieSession)
	router.GET("/fetchUserData", api.GetUserDataAPI)

	router.GET("/bikes", api.GetBikesAPI)
	router.GET("/bikes/:id", api.BikeByIdAPI)
	router.PATCH("/rent/:id", api.RentBikeAPI)
	router.PATCH("/return/:id", api.ReturnBikeAPI)
	router.Run("localhost:3000")
}
