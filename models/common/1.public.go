/**
  @author: qianyi  2022/3/17 15:43:00
  @note:
*/

package common

import (
	"DuDao/pkg/helper"

	"github.com/gin-gonic/gin"
)
func IsTokenValid(c *gin.Context) (string,error) {
	//token, err := c.Cookie("X-Token")
	//if err != nil {
	//	return "",errors.New("cookie取出失败，请重新登录")
	//}
	get := c.Request.Header.Get("X-Token")
	_ ,err := helper.Decode(get,helper.Key)
	if err !=nil{
		return "",err
	}
	return get,err
}
