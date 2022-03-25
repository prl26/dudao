/**
  @author: qianyi  2022/1/2 20:05:00
  @note:
*/
package user

import (
	"DuDao/models"
	"DuDao/models/Res"
	"DuDao/pkg/helper"
	"DuDao/repositories"
	v1 "DuDao/requests/v1"
	"errors"
	"strconv"
	"strings"
	"time"
)

type DataService struct {
}

// LoadRecord Load听课记录
func (ds DataService) LoadRecord(token string,load v1.QueryLoad) (Res.ResponseLoad,error){
	claim, err := helper.Decode(token, helper.Key)
	if err != nil {
		return Res.ResponseLoad{}, err
	}
	name := claim.UserName
	authority := claim.Authority
	record, i, err := repositories.LoadRecord(load,name,authority)
	if err != nil {
		return Res.ResponseLoad{}, err
	}

	loadData := Res.ResponseLoad{Code:200,Message:"加载成功",Count: i,Data: record}

	return loadData,nil
}

// DeleteDataByID 逻辑删除一条数据
func (ds DataService) DeleteDataByID (token string,qd v1.QueryDataID) (Res.ResponseDeleteData,error) {
	_, err := helper.Decode(token, helper.Key)
	if err != nil {
		return Res.ResponseDeleteData{}, err
	}
	err = repositories.DeleteDataById(qd)

	if err != nil {
		return Res.ResponseDeleteData{}, err
	}

	var res  = Res.ResponseDeleteData{
		Message: "操作成功",
		Code: 200,
	}

	return res,nil
}

// GetDetails 显示打分详情
func (ds DataService) GetDetails(token string,dataId string) (*Res.ShowDetails,error){
	claim , err := helper.Decode(token, helper.Key)
	if err != nil {
		return nil, err
	}
	details, err := repositories.GetDetails(dataId,claim.Authority,claim.UserName)
	if err != nil {
		return nil, errors.New("权限不足")
	}

	return details,nil
}

// GetAdd 根据id创建记录并返回指标信息
func (ds DataService) GetAdd(token string,id string) (models.TData,[]models.TdataResult,error){
	claim, err := helper.Decode(token, helper.Key)
	if err != nil {
		return models.TData{},[]models.TdataResult{}, err
	}
	user := repositories.GetUserByAccount(claim.Account)
	supervisor := user.Name
	add1,add2, err := repositories.GetAdd(id, supervisor)
	if err != nil {
		return models.TData{},[]models.TdataResult{}, err
	}
	return add1,add2,nil
}


// GetSupervisors 返回督导员列表
func (ds DataService) GetSupervisors (token string) ([]Res.Supervisors,error){
	claim, err := helper.Decode(token, helper.Key)
	if err != nil {
		return nil, err
	}
	supervisors, err := repositories.GetSupervisors(claim.Authority)
	if err != nil {
		return nil, err
	}

	return supervisors,nil
}




// GetCourseCount 获取听课的总次数 以及听课的课程
func (ds DataService) GetCourseCount(token string, request v1.DataCountReq) (Res.DataCount,error) {
	claims, err := helper.Decode(token, helper.Key)
	if err!=nil {
		return Res.DataCount{}, err
	}

	// 根据账号找到用户姓名
	user := repositories.GetUserByAccount(claims.Account)
	supervisor := user.Name

	// 返回typeCount
	dataCount := repositories.GetCourseCount(supervisor, request)

	return dataCount,nil
}

func (ds DataService) Update(token string, nu Res.UpData) error{
	timeArray := nu.ListenTimeArray
	var listenArray []string
	for i:=1;i<len(timeArray);i++{
		if strconv.Itoa(timeArray[i]) == ","{
			continue
		}
		if i == len(timeArray)-1{
			listenArray = append(listenArray,strconv.Itoa(timeArray[i]))
		}else{
			listenArray = append(listenArray,strconv.Itoa(timeArray[i]),"|")
		}
	}
	Array := strings.Join(listenArray,"")
		tar := nu.Targets
		var ou models.TData
		err := models.DB.Model(&ou).Where("ID = ?", nu.Id).Updates(models.TData{
			ID:               nu.Id,
			Type:             "",
			TName:            nu.TName,
			TNameSchool:      nu.TNameInSchool,
			CourseName:       nu.CourseName,
			CourseType:       nu.CourseType,
			ClassName:        nu.ClassName,
			StuInSchool:      nu.StuInSchool,
			AllNum:           nu.AllNum,
			FactNum:          nu.FactNum,
			Attendance:       nu.Attendance,
			Evaluation:       nu.Evaluation,
			EvaluationDetail: nu.EvaluationDetail,
			Supervisor:       nu.Supervisor,
			ListenDate:       nu.ListenDate,
			ListenTime:       Array,
			ListenAdd:        nu.ListenAdd,
			Status:           nu.Status,
			AddAt:            time.Time{},
			AddBy:            "",
			AddIP:            "",
			AddAgent:         "",
			Semester:         nu.Semester,
		}).Error
		if err != nil {
			return err
		}
		var our models.TdataResult
		for k,v := range tar {
			models.DB.Model(&our).Where("id= ?",tar[k].Id).Update("content",v.Content)
		}
	return nil
}
