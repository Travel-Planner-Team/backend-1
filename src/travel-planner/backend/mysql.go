package backend

import (
	"fmt"
	"travel-planner/constants"
	"travel-planner/model"
	"travel-planner/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *MySQLBackend

type MySQLBackend struct {
	db *gorm.DB
}

func InitMySQLBackend(config *util.MySQLInfo) {
	endpoint, username, password := config.Endpoint, config.Password, config.Username
	dsn := username + ":" + password + "@tcp(" + endpoint + ")/" + constants.MYSQL_DBNAME + "?" + constants.MYSQL_CONFIG
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = &MySQLBackend{db}
}

func (backend *MySQLBackend) exampleQueryFunc() error {
	var users []model.User
	result := backend.db.Table("Users").Find(&users)
	fmt.Println(users)
	fmt.Println(result.RowsAffected)
	return result.Error
}
