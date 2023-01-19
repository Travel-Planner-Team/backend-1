package service
import(
    //"fmt"
   "travel-planner/backend"
	"travel-planner/model"
)

func getSitesList (vacationId string) ([]model.Site, error){
	var sites []model.Site
   sites, err := backend.DB.getSitesInVacation(vacationId)
   // getSitesInVacation (vacationId string)
   return sites, err

}
