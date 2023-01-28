package service

import (
	"errors"
	"travel-planner/backend"
	"travel-planner/model"
)

func GetVacationsInfo() ([]model.Vacation, error) {
	vacations, err := backend.DB.GetVacations()
	if err != nil {
		return nil, err
	}

	if vacations == nil || len(vacations) == 0 {
		return nil, errors.New("empty or invalid vacations, check the Database")
	}
	return vacations, nil
}

func AddVacation(vacation *model.Vacation) (bool, error) {
	success, err := backend.DB.SaveVacation(vacation)
	return success, err
}

func GetActivitiesInfoFromPlanId(plan_id uint32) ([]model.Activity, error) {
	activities, err := backend.DB.GetActivityFromPlanId(plan_id)
	if err != nil {
		return nil, err
	}

	if activities == nil || len(activities) == 0 {
		return nil, errors.New("empty or invalid vacations, check the Database")
	}
	return activities, nil
}

func SaveVacationPlan(plan model.Plan) (error) {
	err := backend.DB.SaveVacationPlanToSQL(plan)
	return err
}
