/**
  @author: qianyi  2022/3/1 19:07:00
  @note:
*/
package v1

import (
	v1 "DuDao/api/v1"
	"github.com/gin-gonic/gin"
)

//func apiGroup(e *gin.Engine) *gin.Engine{
//	e.Group("/api")
//	return e
//}

func CheckRouters(engine *gin.Engine) {
	apiGroup := engine.Group("/api")
	// check及登录相关接口
	CheckGroup := apiGroup.Group("/Check")

	CheckGroup.POST("/Login",v1.LoginHandle)
	CheckGroup.GET("/GetUserName",v1.GetUserNameHandle)
	CheckGroup.POST("/Logout",v1.LogOutHandle)
}
