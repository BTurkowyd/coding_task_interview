package api

import (
	"bike_renting_system/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// HandleLogout ... Allows the user to log out from the service.
// @Summary Allows the user to log out from the service.
// @Description Allows the user to log out from the service.
// @Tags Authorization
// @Failure 500 {object} models.Response
// @Router /logout [get]
func HandleLogout(c *gin.Context) {
	host := c.Request.Header.Get("Origin")
	c.Writer.Header().Set("Access-Control-Allow-Origin", host)
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, Message: "User logged out"})
	c.Redirect(301, "/login")

}
