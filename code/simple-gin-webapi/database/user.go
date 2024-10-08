package database

import (
	"fmt"
	"webapi/models"
)

func GetUser() ([]models.User, error) {
	users := make([]models.User, 0)
	err := GetDB().Table("sys_user").Where("status = 0").Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}

func GetUserDetail(id int32) (*models.User, error) {
	var user models.User
	err := GetDB().Table("sys_user").Where("id = ? and status = 0", id).First(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func CreateUser(user models.User) (*models.User, error) {
	err := GetDB().Table("sys_user").Select("name", "password").Create(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id int32, user models.User) (*models.User, error) {
	err := GetDB().Table("sys_user").Where("id = ? and status = 0", id).Updates(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func DeleteUser(id int32) (*models.User, error) {
	var user models.User
	err := GetDB().Table("sys_user").Where("id = ? and status = 0", id).Delete(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}
