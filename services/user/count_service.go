/**
  @author: qianyi  2022/3/17 16:50:00
  @note:
*/
package user

import (
	"DuDao/models/Res"
	"DuDao/pkg/helper"
	user2 "DuDao/repositories"
	v1 "DuDao/requests/v1"
	"errors"
	"log"
)

type CountService struct {
}

func (Cs CountService) GetCount(token string, req v1.GetCountReq) (*Res.GetCountRes, error) {

	_, err := helper.Decode(token, helper.Key)
	if err != nil {
		log.Println("解析token失败")
		return nil, errors.New("登录已失效，请重新登录")
	}

	count, err := user2.GetCount(req)
	if err != nil {
		return nil, errors.New("未找到相关信息")
	}

	count.Message = "操作成功"
	count.Code = 200

	return count, nil
}

// 获取最新学期的count
func (Cs CountService) GetNewCount(token string,semester string) (*Res.ResCount, error) {
	claims, err := helper.Decode(token, helper.Key)
	if err != nil {
		return nil, errors.New("登录已失效")
	}
	getNewCount, err := user2.GetNewCount(claims.UserName,claims.Authority,semester)
	if err != nil {
		return nil, errors.New("数据库操作失败")
	}
	var resCount = &Res.ResCount{
		Result:  getNewCount,
		Message: "操作成功",
		Code:    200,
	}

	return resCount, nil
}
