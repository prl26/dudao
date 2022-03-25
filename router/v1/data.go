/**
  @author: qianyi  2022/3/1 19:08:00
  @note:
*/
package v1

import (
	v1 "DuDao/api/v1"
	"github.com/gin-gonic/gin"
)

func DataGroup(engine *gin.Engine)  {

	apiGroup := engine.Group("/api")
	dataGroup := apiGroup.Group("/Data")

	dataGroup.GET("/Load",v1.GetListenRecordHandle)          // 加载听课记录列表
	dataGroup.POST("/Delete",v1.DeleteDataByIDHandle)        // 逻辑删除该条记录
	dataGroup.GET("/Details",v1.GetDetailsHandle)             // 显示详情
	dataGroup.POST("/Update",v1.UpdateHandle)					// 更新或编辑
	dataGroup.GET("/Add",v1.GetAddHandle)                    // 添加听课记录并返回指标信息
	dataGroup.GET("/GetSupervisors",v1.GetSupervisorsHandle) // 返回督导员列表
	dataGroup.POST("/Count",v1.GetCourseCountHandle)         // 获取听课的课程数目

}

