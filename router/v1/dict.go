/**
  @author: qianyi  2022/3/1 19:08:00
  @note:
*/
package v1

import (
	v1 "DuDao/api/v1"
	"github.com/gin-gonic/gin"
)

func DictGroup(engine *gin.Engine)  {
	apiGroup := engine.Group("/api")
	dictGroup := apiGroup.Group("/Dict")

	dictGroup.POST("/GetChildren",v1.GetChildren)
}