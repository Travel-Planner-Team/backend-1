package backend

import (
	"context"
	"log"
	"travel-planner/constants"
	"travel-planner/model"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func GetDistanceMatrix(sites []model.Site, index int) (*maps.DistanceMatrixResponse, error){
	c, err := maps.NewClient(maps.WithAPIKey(constants.GOOGLE_API_KEY))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return nil, nil
	}
	var destination []string
	for i := index + 1; i < len(sites); i++ {
		s := sites[i].Address
		destination = append(destination, s)
	}

	r := &maps.DistanceMatrixRequest{
		Origins:      []string{sites[index].Address},
		Destinations: destination,
		Mode: "TravelModeDriving",
	}
	route, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return nil, nil
	}
	pretty.Println(route)

	// after this we will get distancematrix response, 
	// we wanna pick the next closest site as our next point.
	// once we choose it, we can create a transportation object.
	// once create transportation, and return back to service level

	return route, err

}
