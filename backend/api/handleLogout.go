package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// HandleLogout ... Allows the user to log out from the service.
// @Summary Allows the user to log out from the service.
// @Description Allows the user to log out from the service.
// @Tags Logout
// @Success 200 {application/json}
// @Failure 500 {object} object
// @Router /logout [get]
func HandleLogout(c *gin.Context) {
	host := c.Request.Header.Get("Origin")
	c.Writer.Header().Set("Access-Control-Allow-Origin", host)
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "User logged out."})
	c.Redirect(301, "/login")

}
