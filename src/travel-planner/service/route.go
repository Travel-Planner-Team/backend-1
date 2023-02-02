package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"travel-planner/backend"
	"travel-planner/model"

	"github.com/google/uuid"
	"googlemaps.github.io/maps"
)


func ShowRoute(vacationId uint32) (*model.ListOfShowPlan, error) {
	// Step 1
	// get all the site list from the backend using GetSitesInVacation
	var sites []model.Site
	sites, err := backend.DB.GetSitesInVacation(vacationId)
	// if cannot get site, then just return empty plan
	if err != nil {
		return nil, nil
	}
	fmt.Println("Successfully get sites")
	fmt.Println(sites)

	// Step 2
	// get the corresponding vacation by vacationId, because we need
	// the vacation Start Date, and End Date for later use to create Activity, and Transportation
	vacation, err := backend.DB.GetSingleVacation(vacationId)
	if err != nil {
		return nil, nil
	}
	fmt.Println("Successfully get the corresponding vacation")
	startDate := vacation.StartDate
	endDate := vacation.EndDate
	fmt.Println(startDate, endDate)

	// Step 4
	// get list of ShowPlan from DB
	return CreatePlans(sites, vacation), err
}


func CreatePlans(si []model.Site, vacation *model.Vacation) *model.ListOfShowPlan {
	var showplanlist model.ListOfShowPlan

	// for planId
	var planId uint32

	vaID := vacation.Id

	// we only generate for 3 different ShowPlan
	for i := 0; i < 3; i++ {
		// deepCopy the site array and make a new planId
		cpy := make([]model.Site, len(si))
		copy(cpy, si)
		planId = uuid.New().ID()
		startDate := vacation.StartDate
		endDate := vacation.EndDate

		plan := &model.Plan{
			Id:         planId,
			StartDate:  startDate,
			Duration:   vacation.DurationDays,
			VacationId: vaID, // convert string to uint32
		}
		backend.DB.SaveVacationPlanToSQL(*plan)

		activities, transportations := GenerateActivityAndTransportation(cpy, startDate, endDate, planId)

		// Save show plan
		// add it to showPlanList

		sp := &model.ShowPlan{
			Plan_id:                planId,
			ActivityInfoList:       activities,
			TransportationInfoList: transportations,
		}
		showplanlist.ShowPlans = append(showplanlist.ShowPlans, *sp)
	}
	return &showplanlist
}
func GenerateActivityAndTransportation(site []model.Site, startDate time.Time, endDate time.Time, planId uint32) ([]model.Activity, []model.Transportaion) {

	// need to make the randomIndex real random
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(site))
	var activitylist []model.Activity
	var transportaionlist []model.Transportaion

	temp := site[randomIndex]
	site[randomIndex] = site[0]
	site[0] = temp
	// initialize index
	i := 0
	nextDateSetter := startDate.AddDate(0, 0, 1)
	tempDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 21, startDate.Minute(), 0, 0, startDate.UTC().Location())

	// must be site.length - 1, because we do not want nil pointer when we call google api in last one element
	for i < len(site)-1 {
		// check valid address, check date
		if startDate.After(endDate) {
			break
		}
		if site[i].Address == "" {
			i++
			continue
		}

		a := &model.Activity{
			Id:        uuid.New().ID(),
			StartTime: startDate,
			EndTime:   startDate.Add(time.Hour * 2),
			Date:      startDate,
			Duration:  2,
			SiteId:    site[i].Id,
			Plan_id:   planId,
		}
		// update time after visiting the site
		startDate = startDate.Add(time.Hour * 2)
		activitylist = append(activitylist, *a)
		// save activity to db
		backend.DB.SaveActivity(a)

		if startDate.After(endDate) {
			break
		}

		// get distance matrix from google api
		googleResponse, err := backend.GetDistanceMatrix(site, i)
		if err != nil {
			fmt.Println("cannot get google api matrix")
			return nil, nil
		}

		duration, index, err := FindTheClosestSite(googleResponse)
		if err != nil {
			i++
			continue
		}
		// use while loop would be easier to control the index
		i++
		temp = site[index+i]
		site[index+i] = site[i]
		site[i] = temp

		t := &model.Transportaion{
			Id:        uuid.New().ID(),
			Type:      0,
			StartTime: startDate.String(),
			EndTime:   startDate.Add(duration).String(), // need to add processed duration
			Date:      startDate.String(),
			// need to add duration
			Plan_id: planId,
		}

		// Again, update the startDate time
		startDate = startDate.Add(duration)
		transportaionlist = append(transportaionlist, *t)
		backend.DB.SaveTransportation(t)

		if startDate.After(endDate) {
			break
		}
		// change to next date
		if startDate.After(tempDate) {
			startDate = nextDateSetter
			nextDateSetter = startDate.AddDate(0, 0, 1)
			tempDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 21, startDate.Minute(), 0, 0, startDate.UTC().Location())
		}
	}
	if startDate.Before(endDate) {
		a := &model.Activity{
			Id:        uuid.New().ID(),
			StartTime: startDate,
			EndTime:   startDate.Add(time.Hour * 2),
			Date:      startDate,
			Duration:  2,
			SiteId:    site[i].Id,
			Plan_id:   planId,
		}
		// update time after visiting the site
		startDate = startDate.Add(time.Hour * 2)
		activitylist = append(activitylist, *a)
		// save activity to db
		backend.DB.SaveActivity(a)
	}

	return activitylist, transportaionlist
}

func FindTheClosestSite(res *maps.DistanceMatrixResponse) (time.Duration, int, error) {
	element := res.Rows[0].Elements
	// initialize globalMin
	globalMinDuration := element[0].Duration
	var index int
	for i := 0; i < len(element); i++ {
		if res.OriginAddresses[0] == "" {
			return element[0].Duration, -1, errors.New("originalAddress error")
		}
		if res.DestinationAddresses[i] == "" {
			continue
		}
		du := element[i].Duration.Minutes()
		if du <= globalMinDuration.Minutes() {
			globalMinDuration = element[i].Duration
			index = i
		}
	}
	return globalMinDuration, index, nil
}
