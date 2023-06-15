package models

type Bike struct {
	Bike_id    string  `json:"bike_id"`
	Name       string  `json:"name"`
	Latitude   float32 `json:"latitude"`
	Longtitude float32 `json:"longtitude"`
	Rented     bool    `json:"rented"`
	User_name  string  `json:"user_id"`
}

type User struct {
	User_id  string `json:"user_id"`
	Name     string `json:"name"`
	Password []byte `json:"-"`
	Renting  bool   `json:"renting"`
	Bike_id  string `json:"bike_id"`
}

type UserResponseStruct struct {
	Name    string `json:"name"`
	Renting bool   `json:"renting"`
	Bike_id string `json:"bike_id"`
}

type BikeResponseStruct struct {
	Bike_id    string  `json:"bike_id"`
	Name       string  `json:"name"`
	Latitude   float32 `json:"latitude"`
	Longtitude float32 `json:"longtitude"`
	Rented     bool    `json:"rented"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Response struct {
	Code    int    `json:"code" example:"418"`
	Message string `json:"message" example:"A simple message briefly explaining the issue"`
}
