/**
  @author: qianyi  2022/1/1 17:42:00
  @note:
*/
package user

import (
	"DuDao/dto"
	"DuDao/pkg/helper"
	"DuDao/repositories"
	"errors"
)

type UserService struct {
}

func (user UserService) ChangePassWd(token string,pwd string) error{
	claim,err := helper.Decode(token,helper.Key)

	if err != nil {
		return errors.New("解析Token失败")
	}

	userDto := dto.UserDto{
		Account: claim.Account,
		Password: pwd,
	}

	err = repositories.UpdateUser(userDto)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

