/**
  @author: qianyi  2022/3/1 19:07:00
  @note:
*/
package v1

import (
	v1 "DuDao/api/v1"
	"github.com/gin-gonic/gin"
)

func CountGroup(engine *gin.Engine)  {

	apiGroup := engine.Group("/api")
	countGroup := apiGroup.Group("/Count")


	countGroup.GET("/GetCount",v1.GetCountHandle)
	countGroup.GET("/GetNewCount",v1.GetNewCountHandle)
}
