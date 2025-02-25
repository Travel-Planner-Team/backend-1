package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"travel-planner/model"
	"travel-planner/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	vacationID := mux.Vars(r)["vacation_id"]
	fmt.Println("vacationID:", vacationID)
	w.Header().Set("Content_Type", "application/json")
	// get plans
	intId, _ := strconv.ParseInt(vacationID, 0, 64)
	parsedVacationId := uint32(intId)
	plans, err := service.GetPlanInfoFromVactionId(parsedVacationId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var plansInfo []model.PlansInfo
	fmt.Println("plan leng: ", len(plans))

	for planIdx := 0; planIdx < len(plans); planIdx++ {
		plan := &plans[planIdx]
		parsedPlanId := plan.Id
		fmt.Println("planId: ", parsedPlanId)

		activities, err := service.GetActivitiesInfoFromPlanId(parsedPlanId)

		fmt.Println("act leng: ", len(activities))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var totalActList []model.ActivitiesList

		// process each plan
		for activityIdx := 0; activityIdx < len(activities); activityIdx++ {
			activity := &activities[activityIdx]
			site, err := service.GetSiteFromSiteId(activity.SiteId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			activityList := model.ActivitiesList{
				ActivityID:            activity.Id,
				ActivityName:          site.SiteName,
				ActivityType:          "type",
				ActivityDescription:   site.Description,
				ActivityAddress:       site.Address,
				ActivityPhone:         site.PhoneNumber,
				ActivityWebsite:       site.SiteUrl,
				ActivityImage:         site.ImageUrl,
				ActivityStartDatetime: activity.StartTime,
				ActivityEndDatetime:   activity.EndTime,
				ActivityDate:          activity.Date,
				ActivityDuration:      activity.DurationHrs,
			}

			totalActList = append(totalActList, activityList)
		}

		// get []transportations
		transportations, err := service.GetTransportationFromPlanId(parsedPlanId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var transportationsList []model.TransportationList
		for _, trans := range transportations {
			transResponse := model.TransportationList{
				TransportationId:          trans.Id,
				TransportationType:        trans.Type,
				TransportationStartTime:   trans.StartTime,
				TransportationEndTime:     trans.EndTime,
				TransportationDate:        trans.Date,
				TransportationDurationHrs: trans.DurationHrs,
				TransportationPlanId:      trans.PlanId,
			}
			transportationsList = append(transportationsList, transResponse)
		}

		dayIdx := 1
		transportationIdx := 0
		nActs := len(totalActList)
		nTrans := len(transportationsList)
		var currentActlist []model.ActivitiesList
		var currentTransportationList []model.TransportationList
		var daysInfoList []model.DayInfo

		for idx, activity := range totalActList {
			if (idx > 0) && (activity.ActivityDate.After(totalActList[idx-1].ActivityDate)) {
				dayInfo := model.DayInfo{DayIDX: dayIdx, Act: currentActlist, Trans: currentTransportationList}
				daysInfoList = append(daysInfoList, dayInfo)
				dayIdx++
				currentActlist = make([]model.ActivitiesList, 0)
				currentTransportationList = make([]model.TransportationList, 0)
				fmt.Println(daysInfoList)
			}
			fmt.Println(idx)
			currentActlist = append(currentActlist, activity)
			if (transportationIdx < nTrans) && (idx < nActs-1) &&
				(transportationsList[transportationIdx].TransportationDate.Equal(activity.ActivityDate)) {
				transportationsList[transportationIdx].TransportationStartAddress = activity.ActivityName
				transportationsList[transportationIdx].TransportationEndAddress = totalActList[idx+1].ActivityName

				currentTransportationList = append(currentTransportationList, transportationsList[transportationIdx])
				transportationIdx++
			}
		}

		dayInfo := model.DayInfo{dayIdx, currentActlist, currentTransportationList}
		daysInfoList = append(daysInfoList, dayInfo)

		plansInfo = append(plansInfo, model.PlansInfo{parsedPlanId, daysInfoList})
	}

	planDetail := model.PlanDetail{parsedVacationId, plansInfo}
	jsonData, err := json.Marshal(planDetail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func SavePlanInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /vacation/{vacation_id}/plan/{plan_id}/save")

	vacationID := mux.Vars(r)["vacation_id"]
	intVacation, _ := strconv.ParseInt(vacationID, 0, 64)
	parsedVacationId := uint32(intVacation)
	planID := mux.Vars(r)["plan_id"]
	intPlan, _ := strconv.ParseInt(planID, 0, 64)
	parsedPlanId := uint32(intPlan)

	fmt.Println(parsedVacationId, parsedPlanId)

	var planInfo model.SavePlanRequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&planInfo)
	if err != nil {
		fmt.Println(err)
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

	vacationID := mux.Vars(r)["vacation_id"]
	intVacation, _ := strconv.ParseInt(vacationID, 0, 64)
	parsedVacationId := uint32(intVacation)

	var newPlan model.Plan

	err := json.NewDecoder(r.Body).Decode(&newPlan)
	if err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// newPlan.Id = uuid.New()
	newPlan.Id = uuid.New().ID()
	newPlan.VacationId = parsedVacationId

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

// type Schedule struct {
// 	Plan_idx       int32                 `json:"plan_idx"`
// 	Activities     []model.Activity      `json:"activity_info_list"`
// 	Transportation []model.Transportaion `json:"transportation_info_list"`
// }

// func GetRouteForVacation(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Received request: /vacation/{vacation_id}/plan/routes")
// 	planIdx, activities, transportations := service.GetRoutesFromSites(nil)
// 	var route Schedule
// 	route.Plan_idx = planIdx
// 	route.Activities = activities
// 	route.Transportation = transportations
// 	js, err := json.Marshal(route)
// 	if err != nil {
// 		http.Error(w, "Fail to save vacation into DB", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Write(js)
// 	w.Write([]byte("Potential Routes Sent"))

// }
