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
	Password string `json:"-"`
	Renting  bool   `json:"renting"`
	Bike_id  string `json:"bike_id"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Unauthorized struct {
	Message string `json:"message" example:"401 Unauthorized"`
}

type BadRequest struct {
	Message string `json:"message" example:"400 Bad Request"`
}

type NotFound struct {
	Message string `json:"message" example:"404 Not Found"`
}

type ServerError struct {
	Message string `json:"message" example:"500 Server Error"`
}
