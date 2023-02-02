package model

import (
	// "fmt"
	"regexp"
	"strings"
	"time"
	"travel-planner/util/errors"
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
	Id           string    `json:"id"`
	Destination  string    `json:"destination"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	DurationDays int64     `json:"duration_days"`
	UserId       uint32    `json:"user_id"`
}

type User struct {
	Id       uint32 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
}

type Site struct {
	Id          uint32  `json:"id"`
	SiteName    string  `json:"site_name"`
	Rating      string  `json:"rating"`
	PhoneNumber string  `json:"phone_number"`
	VacationId  uint32  `json:"vacation_id"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	Url         string  `json:"url"`
}

type TripSite struct {
	LocationId string     `json:"location_id"`
	Name       string     `json:"name"`
	Address    AddressObj `json:"address_obj"`
}

type AddressObj struct {
	Street1       string `json:"street1"`
	Street2       string `json:"street2"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	Postalcode    string `json:"postalcode"`
	AddressString string `json:"address_string"`
}

type TripDetails struct {
	LocationId    string `json:"location_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	WebUrl        string `json:"web_url"`
	AddressString string `json:"address_string"`
	Rating        string `json:"rating"`
	Phone         string `json:"phone"`
}

type Plan struct {
	Id         uint32    `json:"id"`
	StartDate  time.Time `json:"start_date"`
	Duration   int64     `json:"duration"`
	VacationId uint32    `json:"vacation_id"`
}

type Activity struct {
	Id        uint32    `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Date      time.Time `json:"date"`
	Duration  int64     `json:"duration"`
	SiteId    uint32    `json:"site_id"`
	Plan_id   uint32    `json:"plan_id"`
}

type Transportaion struct {
	Id        uint32 `json:"id"`
	Type      int    `json:"type"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Date      string `json:"date"`
	Plan_id   uint32 `json:"plan_id"`
}

type SavePlanRequestBody struct {
	ActivityInfoList       []Activity      `json:"activity_info_list"`
	TransportationInfoList []Transportaion `json:"transportation_info_list"`
}

type ShowPlan struct {
	Plan_id                uint32          `json:"planid"`
	ActivityInfoList       []Activity      `json:"activity_info_list"`
	TransportationInfoList []Transportaion `json:"transportation_info_list"`
}

type ListOfShowPlan struct {
	ShowPlans   []ShowPlan `json:"show_plan_list"`
	Vacation_id string     `json:"vacation_id"`
}

type ActivitiesList struct {
	ActivityID            int       `json:"activity_id"`
	ActivityName          string    `json:"activity_name"`
	ActivityType          string    `json:"activity_type"`
	ActivityDescription   string    `json:"activity_description"`
	ActivityAddress       string    `json:"activity_address"`
	ActivityPhone         string    `json:"activity_phone"`
	ActivityWebsite       string    `json:"activity_website"`
	ActivityImage         string    `json:"activity_image"`
	ActivityStartDatetime time.Time `json:"activity_start_datetime"`
	ActivityEndDatetime   time.Time `json:"activity_end_datetime"`
	ActivityDate          time.Time `json:"activity_date"`
	ActivityDuration      int64     `json:"activity_duration"`
}

type DaysInfo struct {
	DayIDX int              `json:"day_idx"`
	Act    []ActivitiesList `json:"activities"`
	Trans  []Transportaion  `json:"transportation"`
}

type PlansInfo struct {
	PlanIDX int      `json:"plan_idx"`
	Days    DaysInfo `json:"days"`
}

type PlanDetail struct {
	VacationID int         `json:"vacation_id"`
	Plans      []PlansInfo `json:"plans"`
}

func (user *User) Validate() *errors.RestErr {
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	if user.Username == "" || regexp.MustCompile(`^[a-z0-9]$`).MatchString(user.Username) {
		return errors.NewBadRequestError("Invalid username")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}
	return nil
}
