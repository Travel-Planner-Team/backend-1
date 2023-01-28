package main

import (
	"fmt"
	// "net/http"
	// "io/ioutil"
	"travel-planner/constants"
	//"net/url"
	"net/url"
	//"path"
)


func tst() {
   url := getUrl("Intrepid Sea, Air & Space Museum")
   fmt.Println(url)
 
}

func getUrl(name string)(string){
key := constants.TRIPADVISOR_API_KEY

	//url := "https://api.content.tripadvisor.com/api/v1/location/search?key=62A808FFA5BB43458AA517B597F7C0E1&searchQuery=Intrepid%20Sea%2C%20Air%20%26%20Space%20Museum&language=en"
	//Intrepid Sea, Air & Space Museum
baseUrl := "https://api.content.tripadvisor.com/api/v1/location/search?key="
baseUrl += key
	chars := []rune(name)
	url, _ := url.Parse(baseUrl)
	var ret string
    for i := 0; i < len(chars); i++ {
        char := string(chars[i])
      
		if char == " " {
			ret += "%20"
			
		}else if char == "," {
			ret += "%2C"
		}else if char == "&" {
			ret += "%26"
		}else {
            ret += char
		}
    }
	ret += "&language=en"
	baseUrl += "&searchQuery="
	baseUrl +=ret;

	rel,_:= url.Parse(baseUrl)
	fmt.Println(rel)
	fmt.Println(rel.String())
	return rel.String()
}