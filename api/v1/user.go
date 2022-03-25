/**
  @author: qianyi  2022/1/1 17:34:00
  @note:
*/
package v1

import (
	"DuDao/models/Res"
	"DuDao/models/common"
	v1 "DuDao/requests/v1"
	"DuDao/services/user"
	"github.com/gin-gonic/gin"
)

// ChangePassWdHandle 用户修改密码
func ChangePassWdHandle(c *gin.Context) {

	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	request := v1.ChangePwdRep{}
	err = c.ShouldBind(&request)
	if err != nil {
		c.Status(401)
		c.Writer.Write([]byte("参数绑定出错"))
		return
	}
	pwd := request.Password

	userservice := user.UserService{}

	err = userservice.ChangePassWd(token, pwd)

	if err != nil {
		c.Status(401)
		c.Writer.Write([]byte("修改密码失败"))
		return
	}

	c.SetCookie("X-Token", token, -1, "/", "", false, false)
	c.JSON(200, &Res.Response{
		Result:  "",
		Message: "操作成功",
		Code:    200,
	})
}
