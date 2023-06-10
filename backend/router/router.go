package router

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("your-secret-key"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24})

	router.Use(sessions.Sessions("sessions", store))
	router.Use(cors.Default())

	router.Use(func(c *gin.Context) {
		host := c.Request.Header.Get("Origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", host)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// fmt.Println(store)

	return router
}
