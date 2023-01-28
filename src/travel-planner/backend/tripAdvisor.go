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
	//"travel-planner/backend"
	//"errors"
)



func SearchDetailFromTrip(sites []model.Site) {
     for key, item := range sites {
		fmt.Printf("Sitename:%v\n",item.Site_name)
        location := GetSearchTripAdvisor(item.Site_name)
	    if location == nil {
			continue
		}
		var location_id string
		location_id = location.Location_id
    
        
		res := GetDetailsWithLocationId(location_id)

		if res == ""{
			continue
		}
	
		resBytes := []byte(res) // Converting the string "res" into byte array
		var jsonRes map[string]interface{} // declaring a map for key names as string and values as interface 
        _ = json.Unmarshal(resBytes, &jsonRes) // Unmarshalling

        item.Description = jsonRes["description"].(string)
		if jsonRes["phone"] != nil{
          item.Phone_number = jsonRes["phone"].(string)
		}
		
		if jsonRes["rating"] != nil{
		item.Rating =jsonRes["rating"].(string)
		}
		
		if jsonRes["address_obj"] != nil{
		details_Address := jsonRes["address_obj"].(map[string]interface{})
		item.Address = details_Address["address_string"].(string)
		}
		
		fmt.Println(item)
		DB.SaveSingleSite(item)
		sites[key] = item

	 }
}
func GetSearchTripAdvisor(name string)(*model.TripSite){
	//key := constants.TRIPADVISOR_API_KEY
	//ttps://api.content.tripadvisor.com/api/v1/location/search
	//url := "https://api.content.tripadvisor.com/api/v1/location/search?key=62A808FFA5BB43458AA517B597F7C0E1&searchQuery=hi&language=en"
	// params := url.Values{}
	// params.Add("key", key)
	// params.Add("searchQuery", name)
	// url := fmt.Sprintf("https://api.content.tripadvisor.com/api/v1/location/search?" +params.Encode()+"&language=en")
	url := getUrl(name)
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

    data := jsonRes["data"].([]interface{})
	if data == nil {
		return nil;
	}
	firstData := data[0]
	if firstData == nil {
		return nil
	}
	fmt.Print(firstData)
	firstDataJson, _ := json.Marshal(firstData)
   var tripSites model.TripSite
   json.Unmarshal(firstDataJson, &tripSites)

   return &tripSites
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
	//location_id = "17553630"
	//url := "https://api.content.tripadvisor.com/api/v1/location/105125/details?key=62A808FFA5BB43458AA517B597F7C0E1&language=en&currency=USD"
    key := constants.TRIPADVISOR_API_KEY
	url := fmt.Sprintf("https://api.content.tripadvisor.com/api/v1/location/%s/details?language=en&currency=USD&key=%s", location_id, key)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)
    if res.StatusCode == 404 {
		return ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	return string(body)

}

func getUrl(name string)(string){
key := constants.TRIPADVISOR_API_KEY
baseUrl := "https://api.content.tripadvisor.com/api/v1/location/search?"
    baseUrl += "key=" + key
	baseUrl += "&searchQuery=" + url.QueryEscape(name)
	baseUrl += "&language=en"
	return baseUrl
	//url := "https://api.content.tripadvisor.com/api/v1/location/search?key=62A808FFA5BB43458AA517B597F7C0E1&searchQuery=Intrepid%20Sea%2C%20Air%20%26%20Space%20Museum&language=en"
	//Intrepid Sea, Air & Space Museum
//baseUrl := "https://api.content.tripadvisor.com/api/v1/location/search?key="
// baseUrl += key
// 	chars := []rune(name)
// 	url, _ := url.Parse(baseUrl)
// 	var ret string
//     for i := 0; i < len(chars); i++ {
//         char := string(chars[i])
      
// 		if char == " " {
// 			ret += "%20"
			
// 		}else if char == "," {
// 			ret += "%2C"
// 		}else if char == "&" {
// 			ret += "%26"
// 		}else {
//             ret += char
// 		}
//     }
// 	ret += "&language=en"
// 	baseUrl += "&searchQuery="
// 	baseUrl +=ret;

// 	rel,_:= url.Parse(baseUrl)
// 	fmt.Println(rel)
// 	fmt.Println(rel.String())
// 	return rel.String()
    
}
