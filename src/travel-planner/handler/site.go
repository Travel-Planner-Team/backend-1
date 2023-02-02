package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"travel-planner/backend"
	"travel-planner/model"
	"travel-planner/service"

	//"github.com/form3tech-oss/jwt-go"
	"github.com/google/uuid"
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

func AddSiteInVacationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one add site request")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
       return
   	}
   	var site model.Site

	// siteID := r.FormValue("siteId")
  	// intId ,_:=strconv.ParseInt(siteID, 0, 64)
  	// fmt.Printf("intId : %v\n", intId)
  	// parsedId := uint32(intId)

	vacationID := r.FormValue("vacationId")
  	fmt.Printf("vacationid: %v\n", vacationID)
    intVacation ,_:=strconv.ParseInt(vacationID, 0, 64)
  	parsedVacation := uint32(intVacation)

	latitude := r.FormValue("latitude")
  	fmt.Printf("latitude: %v\n", latitude)
    floatLatitude ,_:=strconv.ParseFloat(latitude, 32)
  	parsedLatitude := float32(floatLatitude)
	
	longitude := r.FormValue("longitude")
  	fmt.Printf("vacationid: %v\n", longitude)
    floatLongitude ,_:=strconv.ParseFloat(longitude, 32)
  	parsedLongitude := float32(floatLongitude)
   
	

	
	site = model.Site{
	  SiteName:    r.FormValue("siteName"),
	  Rating:      r.FormValue("rating"),
	  PhoneNumber: r.FormValue("phoneNumber"),
	  VacationId:  parsedVacation,
	  Description: r.FormValue("description"),
	  Address:     r.FormValue("address"),
	  Latitude:    parsedLatitude,
	  Longitude:   parsedLongitude,
	}
	site.Id = uuid.New().ID()
   

	// vacationID := mux.Vars(r)["vacation_id"]
  	// fmt.Printf("vacationid: %v\n", vacationID)
    // intVacation ,_:=strconv.ParseInt(vacationID, 0, 64)
  
  	// pasedVacation := uint32(intVacation)

 	// // get site id
  	// siteID := mux.Vars(r)["id"]
  	// intId ,_:=strconv.ParseInt(siteID, 0, 64)
  	// fmt.Printf("intId : %v\n", intId)
  	// parsedId := uint32(intId)

  	// siteF := r.Context().Value("site")
  	// fmt.Println(siteF)
  	// claims := siteF.(*jwt.Token).Claims
  	// //siteId := claim.(jwt.MapClaims)["site_name"].(string)
  	// siteName := claims.(jwt.MapClaims)["site_name"].(string)
  	// siteRating := claims.(jwt.MapClaims)["rating"].(string)
  	// sitePhone := claims.(jwt.MapClaims)["phone_numbner"].(string)
  	// siteDescription := claims.(jwt.MapClaims)["description"].(string)
  	// siteAddress := claims.(jwt.MapClaims)["address"].(string)
  	// lati :=claims.(jwt.MapClaims)["latitude"].(string)
  	// value, _ := strconv.ParseFloat(lati, 32)
  	// siteLatitude := float32(value)
  	// longi :=claims.(jwt.MapClaims)["latitude"].(string)
  	// value1, _ := strconv.ParseFloat(longi, 32)
  	// siteLongitude := float32(value1)


  	// site.Id = parsedId
  	// site.SiteName = siteName
  	// site.Rating = siteRating
  	// site.PhoneNumber = sitePhone
  	// site.VacationId = pasedVacation
  	// site.Description = siteDescription
  	// site.Address = siteAddress
  	// site.Latitude =siteLatitude
  	// site.Longitude = siteLongitude

  
  
	success, err := backend.DB.SaveSingleSite(site)

  	 if !success {
		http.Error(w, "Failed to save site",http.StatusInternalServerError)
		fmt.Printf("Failed to save site %v\n ", err)
	}

	fmt.Println("sites added successfully")
	fmt.Fprintf(w, "Site saved %s\n", site.Id)
}