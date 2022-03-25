/**
  @author: qianyi  2022/1/2 14:44:00
  @note:
*/
package repositories

import (
	"DuDao/models"
	"DuDao/models/Res"
)

// GetChildren 根据父id获取子级数据
func GetChildren(parentID string) []Res.ResultDict {
	var result []Res.ResultDict
	models.DB.Table("tdict").Where("parentid = ?", parentID).Find(&result)
	return result
}

// GetTypeNameByID 根据id获取类型名字
func GetTypeNameByID(id string) (string, error) {
	var name Res.GetTypeNameByID
	err := models.DB.Table("tdict").Where("id = ?", id).Find(&name).Error
	if err != nil {
		return "", err
	}
	typeName := name.Name

	return typeName, nil
}
