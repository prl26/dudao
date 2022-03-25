
package user

import (
	"DuDao/models/Res"
	"DuDao/pkg/helper"
	"DuDao/repositories"
	"errors"
)

type DictService struct {
	
}

func (service DictService) GetCollegeChildren(token ,parentID string) (result []Res.ResultDict, err error){

	_, err = helper.Decode(token, helper.Key)
	if err!=nil {
		return nil,errors.New("解析token失败")
	}

	result = repositories.GetChildren(parentID)

	return result,nil

}
