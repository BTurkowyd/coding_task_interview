package api

import (
	"bike_renting_system/models"
	"bike_renting_system/queries"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HandleLogin(c *gin.Context) {
	session := sessions.Default(c)
	authenticated := session.Get("authenticated")

	if authenticated != true {
		authenticated = false
	}

	if authenticated == true {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Another user is already logged in. Please logout first."})
		fmt.Println("User is already logged in")
	} else {
		var loginRequest models.LoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println("Wrong credentials")
			return
		}

		name := loginRequest.Name
		password := loginRequest.Password

		user, err := queries.CheckUser(name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Such user doesn't exist."})
			fmt.Println("User doesn't exist")
			panic(err)
		}

		if user.Password == password {
			authenticated := true
			session.Set("authenticated", authenticated)
			session.Set("user_name", user.Name)
			session.Save()

			c.Status(http.StatusOK)
			c.JSON(http.StatusOK, gin.H{"name": user.Name, "renting": user.Renting, "bike_id": user.Bike_id})
			fmt.Println("Login successful")
		} else {
			authenticated := false
			session.Set("authenticated", authenticated)
			session.Save()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed"})
			fmt.Println("Login failed")
		}
	}

}

func GetCookieSession(c *gin.Context) {
	session := sessions.Default(c)

	session.Set("username", "John Doe")
	session.Save()

}
