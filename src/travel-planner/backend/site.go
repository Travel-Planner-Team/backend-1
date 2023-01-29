package backend

import (
	"errors"
	"fmt"
	"travel-planner/model"
)

func (backend *MySQLBackend) GetSitesInVacation(vacationId uint32) ([]model.Site, error) {
	var sites []model.Site
	result := backend.db.Table("Sites").Where("vacation_id = ?", vacationId).Find(&sites)
	if result.Error != nil {
		fmt.Println("Failed to get sites from db")
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		fmt.Printf("No sites record in vacation %v\n", vacationId)
		return nil, nil
	}
	return sites, nil
}

func (backend *MySQLBackend) SaveSites(sites []model.Site) (bool, error) {
	var count = 0
	for _, item := range sites {
		result := backend.db.Table("Sites").Create(&item)

		if result.Error != nil || result.RowsAffected == 0 {
			fmt.Printf("Faild to save site %v\n", item.SiteName)
		}
		count++
	}
	if count == 0 {
		return false, errors.New("Failed to save all the sites")
	}
	return true, nil
}

func (backend *MySQLBackend) SaveSingleSite(site model.Site) (bool, error) {

	result := backend.db.Table("Sites").Omit("vacation_id").Create(&site)

	if result.Error != nil || result.RowsAffected == 0 {
		fmt.Printf("Faild to save site %v\n", site.SiteName)
	}

	return true, nil
}

func (backend *MySQLBackend) AddVacationIdToSite(siteID uint32, vacationID uint32) (bool, error) {
	var site model.Site
	result := backend.db.Table("Sites").First(&site, siteID)

	if result.Error != nil {
		fmt.Printf("error for update in db %v\n", result.Error)
		return false, result.Error
	}
	fmt.Printf("siteID:%v\n", siteID)
	fmt.Printf("vacationID:%v\n", vacationID)
	backend.db.Table("Sites").Model(&site).Select("vacation_id").Updates(model.Site{VacationId: vacationID})

	return true, nil
}
