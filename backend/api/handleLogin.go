package api

import (
	"bike_renting_system/models"
	"bike_renting_system/queries"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// HandleLogin ... Handles the login procedure to the service.
// @Summary Handles the login procedure to the service.
// @Description Handles the login procedure to the service.
// @Accept application/json
// @Param credentials body models.LoginRequest true "name and password values in models.User"
// @Tags Authorization
// @Success 200 {object} models.UserResponseStruct
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /login [post]
func HandleLogin(c *gin.Context) {
	session := sessions.Default(c)
	authenticated := session.Get("authenticated")

	if authenticated != true {
		authenticated = false
	}

	if authenticated == true {
		c.JSON(http.StatusBadRequest, models.Response{Code: http.StatusBadRequest, Message: "Another user is already logged in. Please log out first"})
	} else {
		var loginRequest models.LoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			fmt.Println("User is already logged in")
			c.JSON(http.StatusBadRequest, models.Response{Code: http.StatusBadRequest, Message: err.Error()})
			return
		}

		name := loginRequest.Name
		password := loginRequest.Password

		user, err := queries.CheckUser(name)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Code: http.StatusBadRequest, Message: "Such user doesn't exist"})
			panic(err)
		}

		if user.Password == password {
			authenticated := true
			session.Set("authenticated", authenticated)
			session.Set("user_name", user.Name)
			session.Save()
			c.Status(http.StatusOK)
			c.JSON(http.StatusOK, models.UserResponseStruct{Name: user.Name, Renting: user.Renting, Bike_id: user.Bike_id})
		} else {
			authenticated := false
			session.Set("authenticated", authenticated)
			session.Save()
			c.JSON(http.StatusUnauthorized, models.Response{Code: http.StatusUnauthorized, Message: "Login failed"})
		}
	}

}
