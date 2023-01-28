package model

import (
	"time"

	"gorm.io/gorm"
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

type User struct {
	gorm.Model
	Id       uint
	Email    string
	Password string
}

type Plan struct {
	Id          uint32    `json:"id"`
	Start_date  time.Time `json:"start_date"`
	Duration    int64     `json:"duration"`
	Vacation_id uint32    `json:"vacation_id"`
}

type Activity struct {
	Id        uint32    `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Date      time.Time `json:"date"`
	Duration  int64		`json:"duration"`	
	Site_id   uint32	`json:"site_id"`
}

type Transportaion struct {
	Id        uint32    `json:"id"`
	Type      string    `json:"type"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Date      time.Time `json:"date"`
}
