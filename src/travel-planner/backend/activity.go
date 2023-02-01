package backend

import (
	"fmt"
	"travel-planner/model"
)

func (backend *MySQLBackend) GetActivityFromPlanId(plan_id uint32) ([]model.Activity, error) {
	var activities []model.Activity
	result := backend.db.Table("Activity").Find(&activities)
	fmt.Print(activities, result)
	if result.Error != nil {
		return nil, result.Error
	}
	return activities, nil
}

func (backend *MySQLBackend) GetRoutes(sites []uint32) (int32, []model.Activity, []model.Transportaion) {
	var activities []model.Activity
	result := backend.db.Table("Activities").Find(&activities)
	fmt.Println(activities, result)
	if result.Error != nil {
		return -1, nil, nil
	}
	var tranportations []model.Transportaion
	result = backend.db.Table("Transportations").Find(&tranportations)
	fmt.Println(activities, result)
	if result.Error != nil {
		return -1, activities, nil
	}
	return 8, activities, tranportations
}
