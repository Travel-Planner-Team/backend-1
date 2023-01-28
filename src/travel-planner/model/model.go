package model

import (
	// "fmt"
	"time"
	// "gorm.io/gorm"
)

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

type Vacation struct {
	Id           uint32    `json:"id"`
	Destination  string    `json:"destination"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	DurationDays int64     `json:"duration_days"`
	UserId       uint32    `json:"user_id"`
}

//	type Model struct {
//		ID uint `jason:"id"` // `gorm:"primary_key jason:"id"`
//		CreatedAt   time.Time  `json:"created_at"`
//		UpdatedAt   time.Time  `json:"updated_at"`
//		DeletedAt   *time.Time `json:"deleted_at"`
//	}
type User struct {
	// gorm.Model
	Id       uint32 `gorm:"primaryKey;autoIncrement:true" jason:"id"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Username string `gorm:"unique" json:"username"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
}

// func (user *User) Validate() *errors.RestErr {
// 	user.Username = strings.TrimSpace(user.Username)
// 	user.Password = strings.TrimSpace(user.Password)
// 	user.Email = strings.TrimSpace(user.Email)
// 	if user.Email == "" {
// 		return errors.NewBadRequestError("Invalid email address")
// 	}
// 	if user.Username == "" || regexp.MustCompile(`^[a-z0-9]$`).MatchString(user.Username) {
//     return errors.NewBadRequestError("Invalid username")
//   }
// 	if user.Password == "" {
//     return errors.NewBadRequestError("Invalid password")
//   }
// 	return nil
// }
