package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"travel-planner/model"
	"travel-planner/service"

	//"github.com/pborman/uuid"
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
	fmt.Println(r.Body)
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
	//w.Write([]byte("Vacation saved: " + fmt.Sprint(vacation.Id)))
	
	w.Write(js)
}

func GetVacationPlanHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/{vacation_id}/plan")
	vacationID := r.Context().Value("vacation_id")
	fmt.Printf("vacationID: %v\n", vacationID)
	w.Header().Set("Content_Type", "application/json")
}

func SavePlanInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/{vacation_id}/plan/{plan_id}/save")
	// vacationId := r.Context().Value("vacation_id")
	// plan_id := r.Context().Value("plan_id")

	var planInfo model.SavePlanRequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&planInfo)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(planInfo)

	err = service.SavePlanInfo(planInfo)
	if err != nil {
		http.Error(w, "Failed to save plan info", http.StatusInternalServerError)
	}
}

func InitVacationPlanHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/{vacation_id}/plan/init")

	var newPlan model.Plan
	err := json.NewDecoder(r.Body).Decode(&newPlan)
	if err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// newPlan.Id = uuid.New()
	newPlan.Id = newPlan.VacationId

	// Write the JSON data to the response
	jsonData, err := json.Marshal(newPlan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)

	// Save the plan to the database
	err = service.SaveVacationPlan(newPlan)
	if err != nil {
		http.Error(w, "Error saving plan to database", http.StatusInternalServerError)
		return
	}
}

type Schedule struct {
	Plan_idx       int32                 `json:"plan_idx"`
	Activities     []model.Activity      `json:"activity_info_list"`
	Transportation []model.Transportaion `json:"transportation_info_list"`
}

func GetRouteForVacation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/{vacation_id}/plan/routes")
	planIdx, activities, transportations := service.GetRoutesFromSites(nil)
	var route Schedule
	route.Plan_idx = planIdx
	route.Activities = activities
	route.Transportation = transportations
	js, err := json.Marshal(route)
	if err != nil {
		http.Error(w, "Fail to save vacation into DB", http.StatusInternalServerError)
		return
	}
	w.Write(js)
	w.Write([]byte("Potential Routes Sent"))

}
