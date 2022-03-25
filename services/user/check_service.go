package user

import (
	"DuDao/dto"
	"DuDao/pkg/helper"
	user2 "DuDao/repositories"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
)

type CheckService struct {
}

//Login 登陆
func (user CheckService) Login(userDto dto.UserDto) (string, error, string) {
	model := user2.GetUserByAccount(userDto.Account)
	if model.Account == "" {
		return "", errors.New("账号不存在"), ""
	}

	if model.Password != userDto.Password {
		return "", errors.New("密码错误"), ""
	}

	//整合jwt
	claims := helper.Claims{
		Account:        model.Account,
		Password:       model.Password,
		Authority:      model.Authority,
		UserName:       model.Name,
		StandardClaims: jwt.StandardClaims{},
	}

	//生成token
	token, err := helper.Encode(claims, helper.Key)
	if err != nil {
		return "", errors.New("token生成失败"), ""
	}

	return token, nil, model.Authority
}

// GetUserName 根据Token获取用户名称
func (user CheckService) GetUserName(token string) ( []string, error) {
	// 解析token
	claim, err := helper.Decode(token, helper.Key)
	if err != nil {
		log.Println("解析token失败")
		return nil, errors.New("解析失败，请重新登录")
	}

	model := user2.GetUserByAccount(claim.Account)

	var result []string
	result = append(result, model.Name, model.Authority)

	return result, nil
}

// LogOut 根据Token注销登录
func (user CheckService) LogOut(token string) (bool, error) {
	// 解析token
	_, err := helper.Decode(token, helper.Key)
	if err != nil {
		log.Println("解析token失败")
		return false, errors.New("登录已失效，请重新登录")
	}
	return true, nil
}
