package model

type AppStub struct {
	Id          string `json:"id"`
	User        string `json:"user"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Url         string `json:"url"`
	ProductID   string `json:"product_id"`
	PriceID     string `json:"price_id"`
}

type UserStub struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
}

type User struct {
	Id       string    `json:"id"`
	Email    string  `json:"email"`
	Password string   `json:"password"`
	Username string `json:"username"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
}
type Vacation struct {
	Id       string    `json:"id"`
	Destination_city    string  `json:"destication_city"`
	State_date string   `json:"state_date"`
	End_date string `json:"end_date"`
	Duration      int64  `json:"duration"`
	User_id   string `json:"user_id"`
}

type Site struct {
	Id       string    `json:"id"`
	Site_name    string  `json:"destication_city"`
	Rating string   `json:"rating"`
	Phone_number string `json:"phone_number"`
	Vacation_id   string `json:"vacation_id"`
	Description string `json:"description"`
	Address string `json:"address"`
}

   type TripSite struct{
      Location_id string
      Name string
	  Address_obj Address_obj

    }
	type Address_obj struct{
        Street1 string
        Street2 string
        City string
        State string
        Country string
        Postalcode string
        Address_string string
	}

	type TripDetails struct{
        Location_id string
		Name string
		Description string
		Web_url string
		Address_string string
        Rating string
		Phone string
	}

