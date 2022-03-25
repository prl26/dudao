/**
  @author: qianyi  2022/3/17 16:45:00
  @note:
*/
package v1

import (
	"DuDao/models/common"
	v1 "DuDao/requests/v1"
	"DuDao/services/user"
	"github.com/gin-gonic/gin"
)

func GetCountHandle(c *gin.Context) {

	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401, err.Error())
	}

	tcName := c.Query("tc_name")
	tcSemester := c.Query("tc_semester")

	var query = v1.GetCountReq{
		tcName,
		tcSemester,
	}

	var service user.CountService

	CountRes, err := service.GetCount(token, query)
	if err != nil {
		c.String(401, err.Error())
		return
	}

	c.JSON(200, CountRes)

}

func GetNewCountHandle(c *gin.Context) {
	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401, err.Error())
	}

	var service user.CountService
	query := c.Query("semester")
	Res, err := service.GetNewCount(token,query)
	if err != nil {
		c.String(401, err.Error())
	}
	c.JSON(200, Res)
}
