/**
  @author: qianyi  2021/12/28 20:36:00
  @note:
*/
package v1

import (
	v1 "DuDao/api/v1"
	"github.com/gin-gonic/gin"
)

func UserRouter(engine *gin.Engine)  {

	apiGroup := engine.Group("/api")

	// 用户操作
	userRouter := apiGroup.Group("/Users")
	userRouter.POST("/ChangePassWd",v1.ChangePassWdHandle)



	//// check及登录相关接口
	//CheckGroup := apiGroup.Group("/Check")
	//CheckGroup.POST("/Login",v1.LoginHandle)
	//CheckGroup.GET("/GetUserName",v1.GetUserNameHandle)
	//CheckGroup.POST("/Logout",v1.LogOutHandle)
	//
	////听课次数
	//CountGroup := apiGroup.Group("/Count")
	//CountGroup.GET("GetCount",)
	//
	////Data
	//DataGroup := apiGroup.Group("/Data")
	//DataGroup.GET("/Load",v1.GetListenRecordHandle)			// 加载听课记录列表
	//DataGroup.POST("/Delete",v1.DeleteDataByIDHandle)			// 逻辑删除该条记录
	//DataGroup.GET("Details",v1.GetDetailsHandle)				// 显示详情
	//DataGroup.GET("/Add",v1.GetAddHandle)						// 添加听课记录并返回指标信息
	//DataGroup.GET("/GetSupervisors",v1.GetSupervisorsHandle)	// 返回督导员列表
	//DataGroup.POST("/Count",v1.GetCourseCountHandle)			// 获取听课的课程数目

	//// 字典操作，获取学院数据
	//DictGroup := apiGroup.Group("/Dict")
	//DictGroup.POST("/GetChildren",v1.GetChildren)
	//
	//
	//
	//
	//
	////使用JwtCheck中间件测试jwt
	//engine.GET("/test/jwt", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{"message": "jwt检查通过"})
	//}).Use(middlewares.JwtCheck())
}


