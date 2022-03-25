/**
  @author: qianyi  2021/12/30 22:08:00
  @note:
*/
package main

import (
	"DuDao/models"
	v1 "DuDao/router/v1"
)

func main() {

	models.Init()
	v1.Include(v1.CheckRouters,v1.UserRouter,v1.DataGroup,v1.CountGroup,v1.DictGroup)
	engine := v1.RInit()
	engine.Run()

}


