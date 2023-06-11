package api

import (
	"bike_renting_system/funcs"
	"bike_renting_system/queries"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ReturnBikeAPI(c *gin.Context) {
	session := sessions.Default(c)

	if auth, _ := session.Get("authenticated").(bool); !auth {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	id := c.Param("id")
	bike, err := queries.FetchBikebyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Bike not found"})
		return
	} else {
		if bike.User_name == session.Get("user_name").(string) {
			if !bike.Rented {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "This bike can't be returned"})
			} else {
				queries.ReturnBike(session.Get("user_name").(string), bike)
				c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Bike %s was returned by %s", bike.Name, session.Get("user_name").(string))})
			}
		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "This user can't return this bike"})
		}
	}
}

func RentBikeAPI(c *gin.Context) {
	session := sessions.Default(c)
	user, err := queries.CheckUser(session.Get("user_name").(string))
	funcs.CheckError(err)

	if auth, _ := session.Get("authenticated").(bool); !auth {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	id := c.Param("id")
	bike, err := queries.FetchBikebyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Bike not found"})
		return
	} else {
		if user.Renting {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "This user already rents another bike."})

		} else {
			if bike.Rented {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "This bike can't be rented"})
			} else {
				queries.RentBike(session.Get("user_name").(string), bike)
				c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Bike %s is rented to %s", bike.Name, session.Get("user_name").(string))})
			}
		}
	}
}

func BikeByIdAPI(c *gin.Context) {
	session := sessions.Default(c)

	if auth, _ := session.Get("authenticated").(bool); !auth {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	id := c.Param("id")
	bike, err := queries.FetchBikebyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Bike not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"bike_id": bike.Bike_id, "name": bike.Name, "rented": bike.Rented, "latitude": bike.Latitude, "longtitude": bike.Longtitude})
}

func GetBikesAPI(c *gin.Context) {
	session := sessions.Default(c)

	if auth, _ := session.Get("authenticated").(bool); !auth {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	bikes := queries.FetchBikes()
	c.IndentedJSON(http.StatusOK, bikes)
}

func GetUserDataAPI(c *gin.Context) {
	session := sessions.Default(c)

	if auth, _ := session.Get("authenticated").(bool); !auth {
		fmt.Println("User is not authorized")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	user, err := queries.CheckUser(session.Get("user_name").(string))
	funcs.CheckError(err)

	c.IndentedJSON(http.StatusOK, user)
}
