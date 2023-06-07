package api

import (
	"bike_renting_system/funcs"
	"bike_renting_system/models"
	"bike_renting_system/queries"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HandleLogin(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("authenticated").(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Another user is already logged in. Please logout first."})
	} else {

		var loginRequest models.LoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		name := loginRequest.Name
		password := loginRequest.Password

		user, err := queries.CheckUser(name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Such user doesn't exist."})
			funcs.CheckError(err)
		}

		if user.Password == password {
			authenticated := true
			session.Set("authenticated", authenticated)
			session.Set("user_name", user.Name)
			session.Save()
			c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user.Name, "password": user.Password})
		} else {
			authenticated := false
			session.Set("authenticated", authenticated)
			session.Save()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed"})
		}
	}

}
