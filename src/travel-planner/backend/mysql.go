package backend

import (
	"fmt"
	"travel-planner/constants"
	"travel-planner/model"
	"travel-planner/util"

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

func (backend *MySQLBackend) GetVacations() ([]model.Vacation, error) {
	var vacations []model.Vacation
	result := backend.db.Table("Vacations").Find(&vacations)
	fmt.Println(vacations, result)
	if result.Error != nil {
		return nil, result.Error
	}
	return vacations, nil
}

func (backend *MySQLBackend) SaveVacation(vacation *model.Vacation) (bool, error) {
	result := backend.db.Table("Vacations").Create(&vacation)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
