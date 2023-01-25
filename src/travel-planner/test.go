package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func test() {

	url := "https://api.content.tripadvisor.com/api/v1/location/search?key=62A808FFA5BB43458AA517B597F7C0E1&searchQuery=The%20Metropolitan%20Museum%20of%20Art&language=en"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
