/**
  @author: qianyi  2022/3/17 17:04:00
  @note:
*/
package repositories

import (
	"DuDao/models"
	"DuDao/models/Res"
	v1 "DuDao/requests/v1"
)

func GetCount(req v1.GetCountReq) (*Res.GetCountRes,error) {
	var tcount []Res.CountResult
	var res Res.GetCountRes

	tcName := req.TcName
	tcSemester := req.TcSemester

	err := models.DB.Table("tcount").Where("tc_name = ? AND tc_semester = ?", tcName, tcSemester).Find(&tcount).Error
	if err != nil {
		return nil, err
	}

	res.Result = tcount

	return &res,nil

}

// 0 填写，1提交
func GetNewCount(name ,authority,semester string) (Res.NewCount,error) {
	var haveDone int64
	var listening int64
	var listened  int64
	var newCount Res.NewCount


	if authority == "督导员" {
		models.DB.Table("tdata").Where("supervisor = ? AND status = 0 AND semester = ?",name,semester).Count(&listening)
		models.DB.Table("tdata").Where("supervisor = ? AND status = 1 AND semester = ?",name,semester).Count(&listened)
		models.DB.Table("tdata").Where("supervisor = ? AND status != 2 AND semester = ?",name,semester).Count(&haveDone)
	}

	if authority == "管理员" {
		models.DB.Table("tdata").Where("status = 0 AND  semester = ?",semester).Count(&listening)
		models.DB.Table("tdata").Where("status = 1 AND semester = ?",semester).Count(&listened)
		models.DB.Table("tdata").Where("status != 2 AND semester = ?",semester).Count(&haveDone)
	}



	newCount.HaveDone = int(haveDone)
	newCount.Listening = int(listening)
	newCount.Listened = int(listened)


	return newCount, nil
}

