package backend

import (
	"fmt"
	"travel-planner/constants"
	"travel-planner/model"
	"travel-planner/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"errors"
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
	result := backend.db.Table("Users").Where("email = ?",userEmail).Find(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected != 0 {
		return &user, nil
	}
	
	return nil, errors.New("The email has not been registed before")
}
func (backend *MySQLBackend) SaveUser (user *model.User) (bool, error) {
	result := backend.db.Table("Users").Create(&user)
	if err := result.Error; err != nil{
		return false, err
	}
	fmt.Println("User saved in db")
	return true, nil
}

func (backend *MySQLBackend) ReadUserById (userId uint32)(*model.User, error){
	var user model.User
	result := backend.db.Table("Users").First(&user, userId)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}
// update interface has no return value in gorm?
func (backend *MySQLBackend) UpdateInfo (id uint32, password, username,gender string, age int64)(bool, error){
	var user model.User
	result := backend.db.Table("Users").First(&user, id)

	if result.Error != nil{
		fmt.Printf("error for update in db %v\n",result.Error)
		return false, result.Error
	}
	fmt.Printf("userID:%v\n", user.Id)
	fmt.Println(age)
    backend.db.Table("Users").Model(&user).Select("Password", "Username","Gender", "Age").
	Updates(model.User{Password: password, Username: username, Gender: gender, Age:age})
	fmt.Printf("usersAge:%v\n",user.Age)
    return true, nil
}


func (backend *MySQLBackend) GetSitesInVacation (vacationId uint32) ([]model.Site, error){
	var sites []model.Site
    result := backend.db.Table("Sites").Where("vacation_id = ?",vacationId).Find(&sites)
	if result.Error != nil{
		fmt.Println("Failed to get sites from db")
		return  nil, result.Error
	}
    if result.RowsAffected == 0{
		fmt.Printf("No sites record in vacation %v\n", vacationId)
      return nil, nil
	}
	return sites,nil
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

func (backend *MySQLBackend) SaveSites(sites []model.Site)(bool, error){
	var count = 0
	for _, item := range sites {
       result :=backend.db.Table("Sites").Create(&item)

	   if result.Error !=nil || result.RowsAffected == 0 {
		fmt.Printf("Faild to save site %v\n",item.Site_name)
	   }
	   count++
	}
	if count == 0{
		return false, errors.New("Failed to save all the sites")
	}
	return true, nil
}

func (backend *MySQLBackend) SaveSingleSite(site model.Site)(bool, error){

	
       result :=backend.db.Table("Sites").Create(&site)

	   if result.Error !=nil || result.RowsAffected == 0 {
		fmt.Printf("Faild to save site %v\n",site.Site_name)
	   }
	  
	return true, nil
}
