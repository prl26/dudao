/**
  @author: qianyi  2022/1/2 15:28:00
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

// GetChildren 获取子级数据
func GetChildren(c *gin.Context)  {

	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	request := v1.ParentID{}
	err = c.ShouldBind(&request)
	if err != nil {
		c.Status(401)
		c.Writer.Write([]byte("请求参数出错"))
		return
	}

	service := user.DictService{}

	result, err := service.GetCollegeChildren(token,request.ParentID)
	if err!=nil {
		c.Writer.Write([]byte(err.Error()))
	}
	c.JSON(200,Res.ResponseDict{
		Result: result,
		Message: "操作成功",
		Code: 200,
	})
}
