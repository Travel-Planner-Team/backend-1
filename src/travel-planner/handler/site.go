package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"travel-planner/backend"
	"travel-planner/model"
	"travel-planner/service"

	"github.com/gorilla/mux"
)

func GetSitesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a get sites request in the vacation")
	w.Header().Set("Content-Type", "application/json")

	//line 66 is hardcode for test, we cannot get info from http yet, we should use line65
	//vacationId := mux.Vars(r)["vacationid"]
	var vacationId uint32 = 1
	var sites []model.Site
	var err error
	sites, err = service.GetSitesList(vacationId)

	if err != nil || sites == nil {
		http.Error(w, "Failed to get sites from bd", http.StatusInternalServerError)
		return
	}

	// change to json
	js, err := json.Marshal(sites)
	if err != nil {
		http.Error(w, "Failed to parse sites to JSON format", http.StatusInternalServerError)
	}

	w.Write(js)
}

// Search sites be send on query keywords in current vacation
func SearchSitesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a search sites request in vacation")
	w.Header().Set("Content-Type", "application/json")

	//line 93 is hardcode for test, we cannot get info from http yet, we should use line65
	// vacationId := mux.Vars(r)["vacationid"]
	city := r.URL.Query().Get("city")
	interest := r.URL.Query().Get("interest")

	var sites []model.Site
	sites, err := service.SearchSites(interest, city)

	if err != nil {
		http.Error(w, "Failed to search sites", http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(sites)

	if err != nil {
		http.Error(w, "Failed to parse sites into JSON format", http.StatusInternalServerError)
		return
	}
	//返回
	w.Write(js)
}

func addSiteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one checkout request")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		return
	}

	siteId := mux.Vars(r)["id"]
	vacationId := mux.Vars(r)["vacation_id"]
	fmt.Printf("siteid: %v\n", siteId)
	fmt.Printf("vacationid: %v\n", vacationId)

	siteIdInt, _ := strconv.ParseInt(siteId, 0, 64)
	parsedSiteId := uint32(siteIdInt)
	vacationIdInt, _ := strconv.ParseInt(vacationId, 0, 64)
	parsedVacationId := uint32(vacationIdInt)

	success, err := backend.DB.AddVacationIdToSite(parsedSiteId, parsedVacationId)
	if err != nil {
		fmt.Println("Add LocationId to site failed.")
		w.Write([]byte(err.Error()))
		return
	}

	if !success {
		http.Error(w, "Failed to update LocationId to site", http.StatusInternalServerError)
		fmt.Printf("Failed to update LocationId to site %v\n ", err)
	}

	fmt.Println("LocationId updated successfully")
	fmt.Fprintf(w, "Update request received %s\n", siteId)
}
