package backend

import (
	"errors"
	"fmt"

	"travel-planner/model"
)

func (backend *MySQLBackend) FindUser(user *model.User) (bool, error) {
	result := backend.db.Table("Users").Select("email").Find(&user)
	fmt.Println(user, result)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected != 0 {
		return true, nil
	}
	return true, nil
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

	return nil, errors.New("The email has not been registed before.")
}

func (backend *MySQLBackend) ReadUserById(userId uint32) (*model.User, error) {
	var user model.User
	result := backend.db.Table("Users").First(&user, userId)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (backend *MySQLBackend) SaveUser(user *model.User) (bool, error) {
	fmt.Println(user)
	result := backend.db.Table("Users").Create(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
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
