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
