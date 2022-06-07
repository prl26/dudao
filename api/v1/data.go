/**
  @author: qianyi  2022/1/2 20:14:00
  @note:
*/
package v1

import (
	"DuDao/models/Res"
	"DuDao/models/common"
	v1 "DuDao/requests/v1"
	"DuDao/services/user"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetListenRecordHandle Load听课记录
func GetListenRecordHandle(c *gin.Context) {

	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	status := c.Query("status")
	i, _ := strconv.Atoi(status)
	semester := c.Query("semester")
	types := c.Query("type")
	page := c.Query("page")
	p, _ := strconv.Atoi(page)
	limit := c.Query("limit")
	l, _ := strconv.Atoi(limit)
	lessionName := c.Query("lessionName")
	listenDate := c.Query("listenDate")
	tNameInSchool := c.Query("tNameInSchool")
	od := c.Query("order")
	order, _ := strconv.Atoi(od)

	var Load = v1.QueryLoad{
		Status:        i,
		Semester:      semester,
		ListenDate:    listenDate,
		LessionName:   lessionName,
		TNameInSchool: tNameInSchool,
		Type:          types,
		Page:          int(p),
		Limit:         int(l),
		Order:         order,
	}

	service := user.DataService{}

	loadRecord, err := service.LoadRecord(token, Load)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	c.JSON(200, loadRecord)
}

// DeleteDataByIDHandle 根据data ID逻辑删除data
func DeleteDataByIDHandle(c *gin.Context) {

	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	var request v1.QueryDataID
	err = c.ShouldBind(&request)
	if err != nil {
		c.Writer.Write([]byte("参数绑定出错"))
		return
	}

	service := user.DataService{}

	byID, err := service.DeleteDataByID(token, request)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	c.JSON(200, byID)
}

// GetDetailsHandle 显示详情
func GetDetailsHandle(c *gin.Context) {

	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	dataID := c.Query("id")
	if dataID == "" {
		c.JSON(401, gin.H{
			"result":  nil,
			"message": "Non-static method requires a target.",
			"code":    500,
		})
		return
	}

	service := user.DataService{}

	res, err := service.GetDetails(token, dataID)
	if err != nil {
		c.String(401,err.Error())
		return
	}
	if res==nil {
		s := "权限不足"
		c.String(401,s)
		return
	}

	res.Message = "操作成功"
	res.Code = 200

	c.JSON(200, res)
}

// GetAddHandle 根据类型创建听课记录并返回指标信息
func GetAddHandle(c *gin.Context) {

	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	typeid := c.Query("typeid")

	service := user.DataService{}

	add1,add2, err := service.GetAdd(token, typeid)
	if err != nil {
		c.JSON(401, &Res.ResFault{
			Message: "未找到此类型相关信息！",
			Code:    500,
		})
		return
	}

	c.JSON(200, gin.H{
		"result":add1,
		"target":add2,
		"message":"操作成功",
		"code":200,
	})

}

// GetSupervisorsHandle 获取督导员列表
func GetSupervisorsHandle(c *gin.Context) {
	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	service := user.DataService{}

	supervisors, err := service.GetSupervisors(token)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	c.JSON(200, Res.ResponseSupervisor{
		Result:  supervisors,
		Message: "操作成功",
		Code:    200,
	})
}

// GetCourseCountHandle 获取听课的数目 以及听课的课程
func GetCourseCountHandle(c *gin.Context) {

	token, err := common.IsTokenValid(c)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	request := v1.DataCountReq{}
	c.ShouldBind(&request)
	if err != nil {
		c.Status(401)
		c.Writer.Write([]byte("传入参数出错"))
		return
	}


	service := user.DataService{}

	dataCount, err := service.GetCourseCount(token, request)
	if err != nil {
		c.String(401,err.Error())
		return
	}

	c.JSON(200, Res.ResponseDataCount{
		Result:  dataCount,
		Message: "操作成功",
		Code:    200,
	})
}

// UpdateHandle 更新或编辑data
func UpdateHandle(c *gin.Context)  {

	token, err := common.IsTokenValid(c)   // cookie里面的token
	if err != nil {
		c.String(401,err.Error())
		return
	}
	request := Res.UpData{}
	err = c.ShouldBind(&request)

	if err != nil {
		err := "绑定参数出错"
		c.String(401,err)
		return
	}

	var service user.DataService
	err = service.Update(token, request)
	if err != nil {
		c.JSON(401, gin.H{
			"message": err,
			"code":    400,
		})
	}else{
		c.JSON(200,gin.H{
			"message":"操作成功",
			"code":200,
		})
	}


}
