package handler

import (
	"encoding/json"
	"fmt"

	"net/http"
	//"regexp"
	//"strconv"
	"travel-planner/model"

	//"time"

	"travel-planner/backend"
	"travel-planner/service"
	//"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	//"github.com/pborman/uuid"
	"strconv"
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
	//vacationId := mux.Vars(r)["vacationid"]
	city := r.URL.Query().Get("city")
	interest := r.URL.Query().Get("interest")

	// interest := "Museum"
	// city := "New York"

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

   siteID := mux.Vars(r)["id"]
   vacationID := mux.Vars(r)["vacation_id"]
  fmt.Printf("siteid: %v\n", siteID)
  fmt.Printf("vacationid: %v\n", vacationID)

   //userFound, _ := backend.DB.ReadUserByEmail(user.Email)
   
  intId ,_:=strconv.ParseInt(siteID, 0, 64)
  fmt.Printf("intId : %v\n", intId)
  pasedId := uint32(intId)
  
   success, err := backend.DB.AddVacationIdToSite(pasedId, vacationID)
   if err != nil {
       fmt.Println("Add LocationId to site failed.")
       w.Write([]byte(err.Error()))
       return
   }

   if !success {
		http.Error(w, "Failed to update LocationId to site",http.StatusInternalServerError)
		fmt.Printf("Failed to update LocationId to site %v\n ", err)
	}

	fmt.Println("LocationId updated successfully")
	fmt.Fprintf(w, "Update request received %s\n", siteID)
}
