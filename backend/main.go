package main

import (
	"bike_renting_system/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	_ "bike_renting_system/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Bike renting service API
// @version 1.0.0
// @host localhost:3000
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
	router.GET("/fetchUserData", api.GetUserDataAPI)

	router.GET("/bikes", api.GetBikesAPI)
	router.GET("/bikes/:id", api.BikeByIdAPI)
	router.PATCH("/rent/:id", api.RentBikeAPI)
	router.PATCH("/return/:id", api.ReturnBikeAPI)
	router.GET("/swaggerAPI/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:3000")

}
