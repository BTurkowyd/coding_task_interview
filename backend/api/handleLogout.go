package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HandleLogout(c *gin.Context) {
	host := c.Request.Header.Get("Origin")
	c.Writer.Header().Set("Access-Control-Allow-Origin", host)
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "User logged out."})
	c.Redirect(301, "/login")

}
