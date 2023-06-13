package constants

import "fmt"

const (
	host     = "host.docker.internal"
	port     = 5432
	user     = "postgres"
	password = "0constans"
	dbname   = "bike_renting_system"
)

var Connection string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
