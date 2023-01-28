package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"travel-planner/model"
	"travel-planner/service"

	"github.com/google/uuid"
)

func GetVacationsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation")
	w.Header().Set("Content-Type", "application/json")

	vacations, err := service.GetVacationsInfo()
	if err != nil {
		http.Error(w, "Fail to read vacation info from backend", http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(vacations)
	if err != nil {
		http.Error(w, "Fail to parse vacations list into JSON", http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func SaveVacationsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/init")
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var vacation model.Vacation

	if err := decoder.Decode(&vacation); err != nil {
		fmt.Println(err)
		http.Error(w, "Cannot decode vacation input", http.StatusBadRequest)
		return
	}

	vacation.Id = uuid.New().ID()
	success, err := service.AddVacation(&vacation)
	if err != nil || !success {
		fmt.Println(err)
		http.Error(w, "Unable to save", http.StatusInternalServerError)
	}

	js, err := json.Marshal(vacation)
	if err != nil {
		http.Error(w, "Fail to save vacation into DB", http.StatusInternalServerError)
		return
	}
	// w.Write([]byte("Vacation saved: " + fmt.Sprint(vacation.Id)))
	w.Write(js)
}

func GetVacationPlanHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/{vacation_id}/plan")
	vacationID := r.Context().Value("vacation_id")
	fmt.Printf("vacationID: %v\n", vacationID)
	w.Header().Set("Content_Type", "application/json")
	// Create a slice of activities
	activities := []model.Activity{
		{Id: 1, StartTime: time.Now(), EndTime: time.Now().Add(time.Hour), Date: time.Now(), Duration: 3600, Site_id: 100},
		{Id: 2, StartTime: time.Now().Add(time.Hour * 2), EndTime: time.Now().Add(time.Hour * 3), Date: time.Now(), Duration: 3600, Site_id: 200},
		{Id: 3, StartTime: time.Now().Add(time.Hour * 4), EndTime: time.Now().Add(time.Hour * 5), Date: time.Now(), Duration: 3600, Site_id: 300},
	}

	// Marshal the activities to JSON
	jsonData, err := json.Marshal(activities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON data to the response
	w.Write(jsonData)
}

func SaveActivitiesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/{vacation_id}/plan/{plan_id}/save")
	vacationId := r.Context().Value("vacation_id")
	plan_id := r.Context().Value("plan_id")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Activities saved: " + fmt.Sprint(vacationId) + fmt.Sprint(plan_id)))
}

func InitPlanHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Plan had been init"))
}

func MakeRouteForVacation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/{vacation_id}/plan/routes")
	w.Write([]byte("Potential Routes Sent"))
}
