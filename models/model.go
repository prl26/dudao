/**
  @author: qianyi  2022/1/1 17:36:00
  @note:
*/
package models

import (
	"time"
)

type User struct {
	ID        string `gorm:"column:id"`
	Account   string `gorm:"column:mobilePhone"`
	Password  string `gorm:"column:password"`
	Name      string `gorm:"column:name"`
	Status    int    `gorm:"column:status"`
	Authority string `gorm:"column:authority"`
}

type Dict struct {
	ID       string `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Code     string `gorm:"column:code"`
	ParentID string `gorm:"column:parentid"`
	Value    string `gorm:"column:value"`
	Sort     int    `gorm:"column:sort"`
}

type TData struct {
	ID               string    `gorm:"column:id"`
	Type             string    `gorm:"column:type"`
	TName            string    `gorm:"column:tName"`
	TNameSchool      string    `gorm:"column:tNameInSchool"`
	CourseName       string    `gorm:"column:courseName"`
	CourseType       string    `gorm:"column:courseType"`
	ClassName        string    `gorm:"column:className"`
	StuInSchool      string    `gorm:"column:stuInSchool"`
	AllNum           string    `gorm:"column:allNum"`
	FactNum          string    `gorm:"column:factNum"`
	Attendance       string    `gorm:"column:attendance"`
	Evaluation       string    `gorm:"column:evaluation"`
	EvaluationDetail string    `gorm:"column:evaluationDetail"`
	Supervisor       string    `gorm:"column:supervisor"`
	ListenDate       time.Time `gorm:"column:listenDate"`
	ListenTime       string    `gorm:"column:listenTime"`
	ListenAdd        string    `gorm:"column:listenAdd"`
	Status           int       `gorm:"column:status"`
	AddAt            time.Time `gorm:"column:addAt"`
	AddBy            string    `gorm:"column:addBy"`
	AddIP            string    `gorm:"column:addIP"`
	AddAgent         string    `gorm:"column:addAgent"`
	Semester         string    `gorm:"column:semester"`
}

type TdataResult struct {
	ID                 string    `json:"id" gorm:"column:id"`
	DataID             string    `json:"dataID" gorm:"column:dataID"`
	TargetName         string    `json:"targetName" gorm:"column:targetName"`
	TargetSort         uint64    `json:"targetSort" gorm:"column:targetSort"`
	GroupName          string    `json:"groupName" gorm:"column:groupName"`
	GroupSort          uint64    `json:"groupSort" gorm:"column:groupSort"`
	TargetContent      string    `json:"targetContent" gorm:"column:targetContent"`
	AddAt              time.Time `json:"addAt" gorm:"column:addAt"`
	AddBy              string    `json:"addBy" gorm:"column:addBy"`
	AddIP              string    `json:"addIp" gorm:"column:addIP"`
	AddAgent           string    `json:"addAgent" gorm:"column:addAgent"`
	IsGroupNameDisplay int       `json:"isGroupNameDisplay" gorm:"column:isGroupNameDisplay"`
	Content            string    `json:"content" gorm:"column:content"`
}
type Ttarget struct {
	ID                 string `json:"id" gorm:"column:id"`
	Type               string `gorm:"column:type"`
	Name               string `gorm:"column:name"`
	Content            string `gorm:"column:content"`
	Sort               uint64 `gorm:"column:sort"`
	GroupName          string `gorm:"column:groupName"`
	GroupSort          uint64 `gorm:"column:groupSort"`
	Memo               string `gorm:"memo"`
	IsGroupNameDisplay int    `gorm:"column:isGroupNameDisplay"`
	IsImportant        int    `gorm:"column:isImportant"`
}

