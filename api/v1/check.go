/**
  @author: qianyi  2021/12/28 20:34:00
  @note:
*/
package v1

import (
	"DuDao/dto"
	"DuDao/models/Res"
	"DuDao/models/common"
	v1 "DuDao/requests/v1"
	"DuDao/services/user"
	"github.com/gin-gonic/gin"
)

//LoginHandle 登陆
func LoginHandle(c *gin.Context) {
	request := v1.LoginRequest{}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(401, &Res.LoginResult{
			ReturnUrl: "null",
			Token:     "null",
			Authority: "null",
			Result:    "null",
			Message:   "Object reference not set to an instance of an object",
			Code:      500,
		})
		return
	}

	userDto := dto.UserDto{
		Account:  request.Account,
		Password: request.Password,
	}

	service := user.CheckService{}
	token, err, authority := service.Login(userDto)

	if err != nil {
		c.JSON(401, &Res.LoginResult{
			ReturnUrl: "null",
			Token:     "null",
			Authority: "",
			Result:    "null",
			Message:   "Object reference not set to an instance of an object",
			Code:      500,
		})
		return

	}

	c.SetCookie("X-Token", token, 14400, "/", "", false, false)

	c.JSON(200, Res.LoginResult{
		ReturnUrl: "",
		Token:     token,
		Authority: authority,
		Result:    "",
		Message:   "操作成功",
		Code:      200,
	})
}

// GetUserNameHandle 根据token获取用户名
func GetUserNameHandle(c *gin.Context) {
	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}
	service := user.CheckService{}

	result, err := service.GetUserName(token)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.JSON(200, Res.ResponseIn{
		Result:  result,
		Message: "操作成功",
		Code:    200,
	})
}


// 根据Token注销登录

func LogOutHandle(c *gin.Context) {

	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	service := user.CheckService{}
	result, err := service.LogOut(token)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.SetCookie("X-Token", token, -1, "/", "", false, false)
	c.JSON(200, Res.ResponseOut{
		Result:  result,
		Message: "操作成功",
		Code:    200,
	})
}

