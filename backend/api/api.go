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
