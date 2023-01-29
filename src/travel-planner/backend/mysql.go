package backend

import (
	"fmt"
	"travel-planner/constants"
	"travel-planner/model"
	"travel-planner/util"

	//"travel_planner/handler"

	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *MySQLBackend
)

type MySQLBackend struct {
	db *gorm.DB
}

func InitMySQLBackend(config *util.MySQLInfo) {

	endpoint, username, password := config.Endpoint, config.Username, config.Password

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		username, password, endpoint, constants.MYSQL_DBNAME, constants.MYSQL_CONFIG)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = &MySQLBackend{db}
}

func (backend *MySQLBackend) ExampleQueryFunc() error {
	var users []model.User
	result := backend.db.Table("Users").Find(&users)
	fmt.Println(users)
	fmt.Println(result.RowsAffected)
	return result.Error

}

func (backend *MySQLBackend) ReadUserByEmail(userEmail string) (*model.User, error) {
	var user model.User
	result := backend.db.Table("Users").Where("email = ?", userEmail).Find(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected != 0 {
		return &user, nil
	}

	return nil, errors.New("the email has not been registed before")
}

func (backend *MySQLBackend) ReadUserById(userId uint32) (*model.User, error) {
	var user model.User
	result := backend.db.Table("Users").First(&user, userId)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// update interface has no return value in gorm?
func (backend *MySQLBackend) UpdateInfo(id uint32, password, username, gender string, age int64) (bool, error) {
	var user model.User
	result := backend.db.Table("Users").First(&user, id)

	if result.Error != nil {
		fmt.Printf("error for update in db %v\n", result.Error)
		return false, result.Error
	}
	fmt.Printf("userID:%v\n", user.Id)
	fmt.Println(age)
	backend.db.Table("Users").Model(&user).Select("Password", "Username", "Gender", "Age").
		Updates(model.User{Password: password, Username: username, Gender: gender, Age: age})
	fmt.Printf("usersAge:%v\n", user.Age)
	return true, nil
}

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

func (backend *MySQLBackend) GetVacations() ([]model.Vacation, error) {
	var vacations []model.Vacation
	result := backend.db.Table("Vacations").Find(&vacations)
	fmt.Println(vacations, result)
	if result.Error != nil {
		return nil, result.Error
	}
	return vacations, nil
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

func (backend *MySQLBackend) SaveVacation(vacation *model.Vacation) (bool, error) {
	result := backend.db.Table("Vacations").Create(&vacation)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (backend *MySQLBackend) GetActivityFromPlanId(plan_id uint32) ([]model.Activity, error) {
	var activities []model.Activity
	result := backend.db.Table("Activity").Find(&activities)
	fmt.Print(activities, result)
	if result.Error != nil {
		return nil, result.Error
	}
	return activities, nil
}

func (backend *MySQLBackend) SaveSites(sites []model.Site) (bool, error) {
	var count = 0
	for _, item := range sites {
		result := backend.db.Table("Sites").Create(&item)

		if result.Error != nil || result.RowsAffected == 0 {
			fmt.Printf("Faild to save site %v\n", item.Site_name)
		}
		count++
	}
	if count == 0 {
		return false, errors.New("failed to save all the sites")
	}
	return true, nil
}

func (backend *MySQLBackend) SaveSingleSite(site model.Site) (bool, error) {

	result := backend.db.Table("Sites").Create(&site)

	if result.Error != nil || result.RowsAffected == 0 {
		fmt.Printf("Faild to save site %v\n", site.Site_name)
	}

	return true, nil
}

func (backend *MySQLBackend) SaveVacationPlanToSQL(plan model.Plan) error {
	fmt.Println("Saving new plan to SQL")
	result := backend.db.Table("Plans").Create(&plan)
	if result.Error != nil || result.RowsAffected == 0 {
		fmt.Printf("Faild to save plan %v\n", plan.Id)
	}
	return nil
}

func (backend *MySQLBackend) SavePlanInfoToSQL(planInfo model.SavePlanRequestBody) error {
	var count = 0
	for _, activity := range planInfo.ActivityInfoList {
		result := backend.db.Table("Activities").Create(&activity)
		if result.Error != nil || result.RowsAffected == 0 {
			fmt.Printf("Faild to save activities %v\n", activity.Id)
		}
		count++
	}
	if count == 0 {
		return errors.New("failed to save all the activities info")
	}

	for _, transportaion := range planInfo.TransportationInfoList {
		result := backend.db.Table("Transportations")
		if result.Error != nil || result.RowsAffected == 0 {
			fmt.Printf("Faild to save activities %v\n", transportaion.Id)
		}
		count++
	}

	if count == 0 {
		return errors.New("failed to save all the activities info")
	}
	return nil
}

func (backend *MySQLBackend) AddVacationIdToSite(siteID uint32, vacationID string) (bool, error) {
	var site model.Site
	result := backend.db.Table("Sites").First(&site, siteID)

	if result.Error != nil {
		fmt.Printf("error for update in db %v\n", result.Error)
		return false, result.Error
	}
	fmt.Printf("siteID:%v\n", siteID)
	fmt.Printf("vacationID:%v\n", vacationID)
	backend.db.Table("Sites").Model(&site).Select("vacation_id").Updates(model.Site{Vacation_id: vacationID})

	return true, nil
}

func (backend *MySQLBackend) ReadFromDB(user *model.User) (bool, error) {
	result := backend.db.Table("User").Select("email").Find(&user)
	fmt.Println(user, result)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected != 0 {
		return true, nil
	}
	return true, nil
}

func (backend *MySQLBackend) SaveUser(user *model.User) (bool, error) {
	fmt.Println(user)
	result := backend.db.Table("Users").Create(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
