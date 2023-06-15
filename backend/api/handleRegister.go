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
		c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, Message: "User registered. Please go to the login window and log in."})

	} else {
		c.JSON(http.StatusBadRequest, models.Response{Code: http.StatusBadRequest, Message: fmt.Sprintf("Username %s is alrady taken", user.Name)})
		fmt.Println("This username is already taken")
	}
}
