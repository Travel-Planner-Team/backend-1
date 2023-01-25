package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	//"strings"
	"travel-planner/constants"
	"travel-planner/model"
	
)



func SearchDetailFromTrip(sites []model.Site) {
     for _, item := range sites {
        location := GetSearchTripAdvisor(item.Site_name)
		
		var location_id string
		location_id = location.Location_id

		// tripDetails := GetDetailTripAdvisor(location_id)
		// // if err != nil {
        // //     return false,err
		// // }
		// item.Rating = tripDetails.Rating
		// item.Phone_number = tripDetails.Phone
		// item.Description = tripDetails.Description
		// item.Address = tripDetails.Address_string
		// fmt.Printf(item.Site_name)

		res := GetDetailsWithLocationId(location_id)
		resBytes := []byte(res) // Converting the string "res" into byte array
		var jsonRes map[string]interface{} // declaring a map for key names as string and values as interface 
        _ = json.Unmarshal(resBytes, &jsonRes) // Unmarshalling
		item.Rating = jsonRes["rating"].(string)
		item.Phone_number = jsonRes["phone_number"].(string)
		item.Description = jsonRes["description"].(string)
		item.Address = jsonRes["address_string"].(string);
	 }
}
func GetSearchTripAdvisor(name string)(model.TripSite){
	key := constants.TRIPADVISOR_API_KEY
	//ttps://api.content.tripadvisor.com/api/v1/location/search
	//url := "https://api.content.tripadvisor.com/api/v1/location/search?key=62A808FFA5BB43458AA517B597F7C0E1&searchQuery=hi&language=en"
	params := url.Values{}
	params.Add("key", key)
	params.Add("searchQuery", name)
	url := fmt.Sprintf("https://api.content.tripadvisor.com/api/v1/location/search?" +params.Encode()+"&language=en")
	fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
  //res := MakeRequest(api_url) // Making the request
	res, _ := http.DefaultClient.Do(req)

	//defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	//fmt.Println(string(body))
     
	resBytes := []byte(body) // Converting the string "res" into byte array
    var jsonRes map[string]interface{} // declaring a map for key names as string and values as interface 
     _ = json.Unmarshal(resBytes, &jsonRes) // Unmarshalling
    data := jsonRes["data"].(string)

	fmt.Print(data)

	//convert response string to json 
	var tripSites []model.TripSite
	json.Unmarshal([]byte(data), &tripSites)
	


//    if tripSites == nil {
// 	 return nil
//    }
   return tripSites[0]
}

func GetDetailTripAdvisor(location_id string) (model.TripDetails) {
	key := constants.TRIPADVISOR_API_KEY
	url := fmt.Sprintf("https://api.content.tripadvisor.com/api/v1/location/%s/details?language=en&currency=USD&key=%s", location_id, key)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
    
	fmt.Println(res)
	fmt.Println(string(body))

	var tripDetails model.TripDetails
	json.Unmarshal([]byte(body), &tripDetails)
    
    return tripDetails
}

func GetDetailsWithLocationId(location_id string)( string){
	//url := "https://api.content.tripadvisor.com/api/v1/location/105125/details?key=62A808FFA5BB43458AA517B597F7C0E1&language=en&currency=USD"
    key := constants.TRIPADVISOR_API_KEY
	url := fmt.Sprintf("https://api.content.tripadvisor.com/api/v1/location/%s/details?language=en&currency=USD&key=%s", location_id, key)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	return string(body)

}
