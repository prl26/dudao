/**
  @author: qianyi  2021/12/28 20:14:00
  @note:
*/
package repositories

import (
	"DuDao/dto"
	"DuDao/models"
	"errors"
)


type Addby struct {
	Addby string `gorm:"column:addBy"`
}

// GetUserByAccount  通过账户查询用户
func GetUserByAccount(account string) models.User {
	user := models.User{}
	models.DB.Find(&user, models.DB.Where("MobilePhone = ?",account))

	return user
}



//UpdateUser 修改用户密码
func UpdateUser(userDto dto.UserDto) error {
	user := models.User{}

	if userDto.Password=="" {
		return errors.New("密码不能为空")
	}

	err := models.DB.Model(&user).Where("mobilePhone = ?",userDto.Account).Update("password",userDto.Password).Error

	return err
}

func GetUserByAddby(addby string) models.User {
	user := models.User{}
	models.DB.Find(&user, models.DB.Where("id = ?", addby))
	return user
}

func GetAddbyByDataID(dataID string)  string {
	var addby Addby
	models.DB.Table("tdata").Find(&addby, models.DB.Where("id = ?",dataID))
	return addby.Addby
}












////CreateUser 创建用户
//func CreateUser(dto dto.UserDto) error {
//	user := models.User{}
//	user.Account = dto.Account
//	user.Password = dto.Password
//
//	err := models.DB.Create(&user).Error
//
//	return err
//}
//
//// 修改密码
//func ChangePassword( rep v1.ChangePwdRep) error {
//	user := models.User{}
//
//	if rep.Password != "" {
//		user.Password = rep.Password
//	}
//
//	err := models.DB.Model(&user).Updates(&user).Error
//
//	return err
//}





