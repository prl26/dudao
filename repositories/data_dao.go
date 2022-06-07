/**
  @author: qianyi  2022/1/2 17:00:00
  @note:
*/
package repositories

import (
	"DuDao/models"
	"DuDao/models/Res"
	"DuDao/pkg/helper"
	v1 "DuDao/requests/v1"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"time"
)

// LoadRecord 加载听课记录Load
func LoadRecord(load v1.QueryLoad, name string, authority string) ([]Res.LoadData, int64, error) {

	var data []Res.LoadData
	var count int64
	var pageindex, pagesize int
	// 分页功能
	if load.Limit <= 0 {
		load.Limit = 10
		pagesize = load.Limit
	} else {
		pagesize = load.Limit
	}

	if load.Page <= 0 {
		pageindex = 1
	} else {
		pageindex = load.Page
	}
	db := models.DB

	if authority == "督导员" {
		db = db.Table("tdata").Where("supervisor = ?", name)
	}

	if authority == "管理员" {
		db = db.Table("tdata")
	}

	if load.LessionName != "" {
		db.Where("courseName = ?", load.LessionName)
	}

	if load.ListenDate != "" {
		db.Where("listenDate = ?", load.ListenDate)
	}

	if load.TNameInSchool != "" {
		db.Where("TNameInSchool = ?", load.TNameInSchool)
	}

	if load.Semester != "" {
		db.Where("semester = ?", load.Semester)
	}

	if load.Order == 0 {
		db.Order("rand()")
	}

	if load.Order == 1 {
		db.Order("addAt ASC")

	}

	if load.Order == 2 {
		db.Order("addAt DESC")
	}

	if load.Status == 3 {
		db.Where("status != 2").Find(&data)
	} else {
		db.Where(map[string]interface{}{"status": load.Status}).Find(&data)
	}

	db.Count(&count).Find(&data)
	db.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&data)

	n := len(data)
	for i := 0; i < n; i++ {
		spilt := strings.Split(data[i].ListenTimeArrayt, "|")
		data[i].ListenTimeArray = spilt
		//format := data[i].ListenDay.Format("2006-01-02")
		//data[i].ListenDayString = format
	}

	return data, count, nil
}

// DeleteDataById 根据id 逻辑删除data
func DeleteDataById(Qd v1.QueryDataID) error {

	err := models.DB.Table("tdata").Where("id = ?", Qd.ID).Update("status", "2").Error

	return err

}

// GetDetails 显示打分详情
func GetDetails(dataID string, authority string, userName string) (*Res.ShowDetails, error) {

	var data Res.LoadData
	var target []Res.TargetData
	var s Res.Notion

	err := models.DB.Table("tnotion").Where("id = 1").Find(&s).Error
	if authority == "管理员" {
		err = models.DB.Table("tdata").Where("id = ?", dataID).Find(&data).Error
		if err != nil {
			return nil, err
		}
		err = models.DB.Table("tdataResult").Where("dataID = ?", dataID).Order("targetSort ASC").Find(&target).Error
		if err != nil {
			return nil, err
		}
	}

	if authority == "督导员" {

		addBy := GetAddbyByDataID(dataID)
		serchUser := GetUserByAddby(addBy)
		fmt.Println(dataID)
		fmt.Println(addBy)
		fmt.Println(serchUser.Name)
		fmt.Println(userName)
		if serchUser.Name != userName {
			log.Println("权限不足")
			return nil, err
		}
		if serchUser.Name == userName {
			err = models.DB.Table("tdata").Where("id = ?", dataID).Find(&data).Error
			if err != nil {
				return nil, err
			}
			err = models.DB.Table("tdataResult").Where("dataID = ?", dataID).Order("targetSort ASC").Find(&target).Error
			if err != nil {
				return nil, err
			}
		}
	}

	n := len(target)
	for i := 0; i < n; i++ {
		scta := strings.Split(target[i].TargetContent, "|")
		target[i].ContentArray = scta
	}

	//splitData := getSplitData(data)

	slta := data.ListenTimeArrayt
	split := strings.Split(slta, "|")
	data.ListenTimeArray = split

	var res = Res.ShowDetails{
		Result:  data,
		Notions: s,
		Targets: target,
	}

	return &res, nil
}

// GetAdd 根据类型创建听课记录并返回指标信息
func GetAdd(id string, name string) (models.TData, []models.TdataResult, error) {
	typeid := id
	var UtdataResult = make([]models.TdataResult, 14)
	var Utdict models.Dict
	models.DB.Where("id = ?", typeid).Find(&Utdict)
	var Uttarget = make([]models.Ttarget, 14)
	models.DB.Where("type =?", Utdict.Name).Find(&Uttarget)
	uuid1 := helper.GetUuid()
	var user models.User
	models.DB.Table("tuser").Where("name = ?", name).Find(&user)
	Utdata := models.TData{
		ID:               uuid1,
		Type:             Utdict.Name,
		TName:            "",
		TNameSchool:      "",
		CourseName:       "",
		CourseType:       "",
		ClassName:        "",
		StuInSchool:      "",
		AllNum:           "",
		FactNum:          "",
		Attendance:       "",
		Evaluation:       "",
		EvaluationDetail: "",
		Supervisor:       name,
		ListenDate:       time.Now(),
		ListenTime:       "",
		ListenAdd:        "",
		Status:           0,
		AddAt:            time.Now(),
		AddBy:            user.ID,
		AddIP:            "",
		AddAgent:         "",
		Semester:         "",
	}
	models.DB.Create(Utdata)
	for i := 0; i < len(Uttarget); i++ {
		b := &Uttarget[i]
		valueofb := reflect.ValueOf(b).Elem()
		TdataResult := models.TdataResult{
			ID:                 helper.GetUuid(),
			DataID:             uuid1,
			TargetName:         valueofb.FieldByName("Name").String(),
			TargetSort:         valueofb.FieldByName("Sort").Uint(),
			GroupName:          valueofb.FieldByName("GroupName").String(),
			GroupSort:          valueofb.FieldByName("GroupSort").Uint(),
			TargetContent:      valueofb.FieldByName("Content").String(),
			AddAt:              time.Now(),
			AddBy:              user.ID,
			AddIP:              "",
			AddAgent:           "",
			IsGroupNameDisplay: 0,
			Content:            "",
		}
		UtdataResult[i] = TdataResult
		models.DB.Create(&TdataResult)
	}
	return Utdata, UtdataResult, nil

}

// GetSupervisors 返回督导员列表
func GetSupervisors(authority string) ([]Res.Supervisors, error) {
	var supervisors []Res.Supervisors

	if authority == "管理员" {
		err := models.DB.Table("tuser").Find(&supervisors).Error
		if err != nil {
			return supervisors, errors.New("数据库操作失败")
		}
	} else {
		return supervisors, errors.New("权限不足")
	}

	return supervisors, nil
}

// GetCourseCount 获取听课的总次数 以及听课的课程
func GetCourseCount(name string, request v1.DataCountReq) Res.DataCount {

	status := request.Status
	typeCourse := request.Type

	var totalCount int64
	var typeCount []Res.TypeCount

	if isExistCourse(typeCourse) {
		models.DB.Raw("SELECT count(*) AS count FROM tdata WHERE supervisor = ? AND status = ? AND  type = ?", name, status, typeCourse).Scan(&typeCount)
		models.DB.Table("tdata").Where("supervisor = ? AND status = ?", name, status).Count(&totalCount)
		typeCount[0].Type = typeCourse
	} else {
		models.DB.Raw("SELECT type,count(*) AS count FROM tdata WHERE supervisor = ? AND status = ? GROUP BY type", name, status).Scan(&typeCount)
		models.DB.Table("tdata").Where("supervisor = ? AND status = ?", name, status).Count(&totalCount)
	}



	dataCount := Res.DataCount{
		TotalCount: totalCount,
		TypeCount:  typeCount,
	}

	return dataCount
}

func isExistCourse(typeCourse string) bool {
	courseArray := []string{"理论课", "外语课", "实验课", "研究生课", "体育课"}
	for i := 0; i < len(courseArray); i++ {
		if courseArray[i] == typeCourse {
			return true
		}
	}
	return false
}
