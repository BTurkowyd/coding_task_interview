package api

import (
	"bike_renting_system/models"
	"bike_renting_system/queries"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRegister(c *gin.Context) {
	var registerRequest models.LoginRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := registerRequest.Name
	password := registerRequest.Password
	user, err := queries.CheckUser(name)

	if err != nil {
		queries.RegisterUser(name, password)
		c.JSON(http.StatusOK, gin.H{"message": "User registered. Please go to the login window and log in."})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Username %s is alrady taken", user.Name)})
	}
}
