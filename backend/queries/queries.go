package queries

import (
	"bike_renting_system/constants"
	"bike_renting_system/funcs"
	"bike_renting_system/models"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

func FetchBikes() []models.Bike {
	db, err := sql.Open("postgres", constants.Connection)
	funcs.CheckError(err)
	rows, err := db.Query(`SELECT * FROM "bike_renting_system"`)
	funcs.CheckError(err)
	defer db.Close()

	bikes := make([]models.Bike, 0)

	for rows.Next() {
		var bike_id string
		var name string
		var latitude float32
		var longtitude float32
		var rented bool
		var user_name string

		err = rows.Scan(&bike_id, &name, &latitude, &longtitude, &rented, &user_name)
		funcs.CheckError(err)

		bike := models.Bike{Bike_id: bike_id, Name: name, Latitude: latitude, Longtitude: longtitude, Rented: rented, User_name: user_name}

		bikes = append(bikes, bike)
	}
	return bikes
}

func FetchBikebyID(bike_id string) (*models.Bike, error) {
	db, err := sql.Open("postgres", constants.Connection)
	funcs.CheckError(err)
	query := `SELECT * FROM "bike_renting_system" WHERE bike_id = '` + bike_id + `'`
	rows, err2 := db.Query(query)
	funcs.CheckError(err2)
	defer db.Close()

	for rows.Next() {
		var bike_id string
		var name string
		var latitude float32
		var longtitude float32
		var rented bool
		var user_name string

		err := rows.Scan(&bike_id, &name, &latitude, &longtitude, &rented, &user_name)
		funcs.CheckError(err)
		bike := models.Bike{Bike_id: bike_id, Name: name, Latitude: latitude, Longtitude: longtitude, Rented: rented, User_name: user_name}
		return &bike, nil
	}
	return nil, errors.New("bike not found")
}

func RentBike(user_name string, bike *models.Bike) {
	db, err := sql.Open("postgres", constants.Connection)
	funcs.CheckError(err)

	if !bike.Rented {
		query := `UPDATE bike_renting_system SET rented = $2, user_name = $3 WHERE bike_id = $1`
		_, err := db.Exec(query, bike.Bike_id, true, user_name)
		funcs.CheckError(err)

		query = `UPDATE users SET renting = $1, bike_id = $2 WHERE name = $3`
		_, err = db.Exec(query, true, bike.Bike_id, user_name)
		funcs.CheckError(err)
	}
}

func ReturnBike(user_name string, bike *models.Bike) {
	db, err := sql.Open("postgres", constants.Connection)
	funcs.CheckError(err)

	if user_name == bike.User_name {
		if bike.Rented {
			query := `UPDATE bike_renting_system SET rented = $2, user_name = '0' WHERE bike_id = $1`
			_, err := db.Exec(query, bike.Bike_id, false)
			funcs.CheckError(err)

			query = `UPDATE users SET renting = $1, bike_id = '0' WHERE name = $2`
			_, err = db.Exec(query, false, user_name)
			funcs.CheckError(err)
		}
	}
}

func CheckUser(name string) (*models.User, error) {
	db, err := sql.Open("postgres", constants.Connection)
	funcs.CheckError(err)
	query := `SELECT * FROM "users" WHERE name = '` + name + `'`
	rows, err := db.Query(query)
	funcs.CheckError(err)
	defer db.Close()

	for rows.Next() {
		var user_id string
		var name string
		var password []byte
		var renting bool
		var bike_id string

		err := rows.Scan(&user_id, &name, &password, &renting, &bike_id)
		funcs.CheckError(err)
		user := models.User{User_id: user_id, Name: name, Password: password, Renting: renting, Bike_id: bike_id}
		return &user, nil
	}
	return nil, errors.New("user not found")
}

func RegisterUser(name string, password []byte) {
	db, err := sql.Open("postgres", constants.Connection)
	funcs.CheckError(err)
	uuid_gen, err := uuid.NewRandom()
	funcs.CheckError(err)
	query := `INSERT INTO users (user_id, name, password, renting, bike_id) VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(query, uuid_gen.String(), name, password, false, "0")
	funcs.CheckError(err)
	defer db.Close()
}
