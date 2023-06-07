package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HandleLogout(c *gin.Context) {
	session := sessions.Default(c)
	authenticated := false
	session.Set("authenticated", authenticated)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "User logged out."})

}
