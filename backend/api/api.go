package api

import (
	"bike_renting_system/funcs"
	"bike_renting_system/models"
	"bike_renting_system/queries"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ReturnBikesAPI ... User returns the bike.
// @Summary User returns the bike.
// @Description User returns the bike.
// @Accept text/plain
// @Param id path string true "bike_id in models.Bike"
// @Tags Bikes
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /return/:id [patch]
func ReturnBikeAPI(c *gin.Context) {
	session := sessions.Default(c)

	if auth, _ := session.Get("authenticated").(bool); !auth {
		c.JSON(http.StatusUnauthorized, models.Response{Code: http.StatusUnauthorized, Message: "Unauthorized"})
		return
	}

	id := c.Param("id")
	bike, err := queries.FetchBikebyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Response{Code: http.StatusNotFound, Message: "Bike not found"})
		return
	} else {
		if bike.User_name == session.Get("user_name").(string) {
			if !bike.Rented {
				c.IndentedJSON(http.StatusBadRequest, models.Response{Code: http.StatusBadRequest, Message: "This bike can't be returned"})
			} else {
				queries.ReturnBike(session.Get("user_name").(string), bike)
				c.IndentedJSON(http.StatusOK, models.Response{Code: http.StatusOK, Message: fmt.Sprintf("Bike %s was returned by %s", bike.Name, session.Get("user_name").(string))})
			}
		} else {
			c.IndentedJSON(http.StatusBadRequest, models.Response{Code: http.StatusBadRequest, Message: "Logged in user can't return this bike"})
		}
	}
}

// RentBikeAPI ... User rents the bike.
// @Summary User rents the bike.
// @Description User rents the bike.
// @Accept text/plain
// @Param id path string true "bike_id in models.Bike"
// @Tags Bikes
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /rent/:id [patch]
func RentBikeAPI(c *gin.Context) {
	session := sessions.Default(c)
	user, err := queries.CheckUser(session.Get("user_name").(string))
	funcs.CheckError(err)

	if auth, _ := session.Get("authenticated").(bool); !auth {
		c.JSON(http.StatusUnauthorized, models.Response{Code: http.StatusUnauthorized, Message: "Unauthorized"})
		return
	}

	id := c.Param("id")
	bike, err := queries.FetchBikebyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Response{Code: http.StatusNotFound, Message: "Bike not found"})
		return
	} else {
		if user.Renting {
			c.IndentedJSON(http.StatusBadRequest, models.Response{Code: http.StatusBadRequest, Message: "Logged in user already rents another bike"})

		} else {
			if bike.Rented {
				c.IndentedJSON(http.StatusBadRequest, models.Response{Code: http.StatusBadRequest, Message: "UThis bike is rented by another user"})
			} else {
				queries.RentBike(session.Get("user_name").(string), bike)
				c.IndentedJSON(http.StatusOK, models.Response{Code: http.StatusOK, Message: fmt.Sprintf("Bike %s is rented by %s", bike.Name, session.Get("user_name").(string))})
			}
		}
	}
}

// BikeByIdAPI ... Fetches a bike with a user-provided ID.
// @Summary Fetches a bike with a user-provided ID.
// @Description Fetches a bike with a user-provided ID.
// @Accept text/plain
// @Param id path string true "bike_id in models.Bike"
// @Tags Bikes
// @Success 200 {object} models.BikeResponseStruct
// @Failure 401 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /bikes/:id [get]
func BikeByIdAPI(c *gin.Context) {
	session := sessions.Default(c)

	if auth, _ := session.Get("authenticated").(bool); !auth {
		c.JSON(http.StatusUnauthorized, models.Response{Code: http.StatusUnauthorized, Message: "Unauthorized"})
		return
	}

	id := c.Param("id")
	bike, err := queries.FetchBikebyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Response{Code: http.StatusNotFound, Message: "Bike not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, models.BikeResponseStruct{Bike_id: bike.Bike_id, Name: bike.Name, Rented: bike.Rented, Latitude: bike.Latitude, Longtitude: bike.Longtitude})
}

// GetBikesAPI ... Fetches all bikes from the database.
// @Summary Fetches all bikes from the database.
// @Description Fetches all bikes from the database.
// @Tags Bikes
// @Success 200 {array} models.Bike
// @Failure 401 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /bikes [get]
func GetBikesAPI(c *gin.Context) {
	session := sessions.Default(c)

	if auth, _ := session.Get("authenticated").(bool); !auth {
		c.JSON(http.StatusUnauthorized, models.Response{Code: http.StatusUnauthorized, Message: "Unauthorized"})
		return
	}

	bikes := queries.FetchBikes()
	c.IndentedJSON(http.StatusOK, bikes)
}

// GetUserDataAPI ... Fetches the data of the logged in user.
// @Summary Fetches the data of the logged in user.
// @Description Fetches the data of the logged in user.
// @Tags Users
// @Success 200 {object} models.User
// @Failure 401 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /fetchUserData [get]
func GetUserDataAPI(c *gin.Context) {
	session := sessions.Default(c)

	if auth, _ := session.Get("authenticated").(bool); !auth {
		fmt.Println("User is not authorized")
		c.JSON(http.StatusUnauthorized, models.Response{Code: http.StatusUnauthorized, Message: "Unauthorized"})
		return
	}

	user, err := queries.CheckUser(session.Get("user_name").(string))
	funcs.CheckError(err)

	c.IndentedJSON(http.StatusOK, user)
}

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
