/**
  @author: qianyi  2022/1/2 16:44:00
  @note:
*/
package Res

import (
	"time"
)

// 加载听课记录及其列表
type ResponseLoad struct {
	Code          int        `json:"code"`
	Message       string     `json:"message"`
	Count         int64      `json:"count"`
	ColumnHeaders []CHeaders `json:"columnHeaders"`
	Data          []LoadData `json:"data"`
}

type CHeaders struct {
	Key         string `json:"key"`
	Description string `json:"description"`
}

type LoadData struct {
	Type             string    `json:"type"`
	TName            string    `json:"tName" gorm:"column:tName"`
	TNameInSchool    string    `json:"tNameInSchool" gorm:"column:tNameInSchool"`
	CourseName       string    `json:"courseName" gorm:"column:courseName"`
	CourseType       string    `json:"courseType" gorm:"column:courseType"`
	ClassName        string    `json:"className" gorm:"column:className"`
	StuInSchool      string    `json:"stuInSchool" gorm:"column:stuInSchool"`
	AllNum           string    `json:"allNum" gorm:"column:allNum"`
	FactNum          string    `json:"factNum" gorm:"column:factNum"`
	Attendance       string    `json:"attendance" gorm:"column:attendance"`
	Evaluation       string    `json:"evaluation" gorm:"column:evaluation"`
	EvaluationDetail string    `json:"evaluationDetail" gorm:"column:evaluationDetail"`
	Supervisor       string    `json:"supervisor" gorm:"column:supervisor"`
	ListenAdd        string    `json:"listenAdd" gorm:"column:listenAdd"`
	Status           string    `json:"status" gorm:"column:status"`
	Semester         string    `json:"semester" gorm:"column:semester"`
	Id               string    `json:"id" gorm:"column:id"`
	ListenDay        time.Time `json:"listenDay"gorm:"column:listenDate"`
	//ListenDayString  string	    `json:"listenDay" gorm:"-"`
	ListenTimeArrayt string    `json:"-" gorm:"column:listenTime"`
	ListenTimeArray  []string  `json:"listenTimeArray" gorm:"-"`
	//Notions          string    `json:"notions"`
	//Targets          []string  `json:"targets"`
}

type LoadData1 struct {
	Type             string    `json:"type"`
	TName            string    `json:"tName" gorm:"column:tName"`
	TNameInSchool    string    `json:"tNameInSchool" gorm:"column:tNameInSchool"`
	CourseName       string    `json:"courseName" gorm:"column:courseName"`
	CourseType       string    `json:"courseType" gorm:"column:courseType"`
	ClassName        string    `json:"className" gorm:"column:className"`
	StuInSchool      string    `json:"stuInSchool" gorm:"column:stuInSchool"`
	AllNum           string    `json:"allNum" gorm:"column:allNum"`
	FactNum          string    `json:"factNum" gorm:"column:factNum"`
	Attendance       string    `json:"attendance" gorm:"column:attendance"`
	Evaluation       string    `json:"evaluation" gorm:"column:evaluation"`
	EvaluationDetail string    `json:"evaluationDetail" gorm:"column:evaluationDetail"`
	Supervisor       string    `json:"supervisor" gorm:"column:supervisor"`
	ListenAdd        string    `json:"listenAdd" gorm:"column:listenAdd"`
	Status           int       `json:"status" gorm:"column:status"`
	Semester         string    `json:"semester" gorm:"column:semester"`
	Id               string    `json:"id" gorm:"column:id"`
	ListenDay        time.Time `json:"listenDay" gorm:"column:listenDate"`
	ListenTimeArray  string    `json:"listenTimeArray" gorm:"column:listenTime"`
	//Notions          string    `json:"notions"`
	//Targets          []string  `json:"targets"`
}

// 根据ID 逻辑删除data
type ResponseDeleteData struct {
	Message string
	Code    int
}

// 查看打分详情
type ShowDetails struct {
	Result  LoadData     `json:"result"`
	Notions Notion       `json:"notions"`
	Targets []TargetData `json:"targets"`
	Message string       `json:"message"`
	Code    int          `json:"code"`
}

type TargetData struct {
	Id                 string   `json:"id" gorm:"column:id"`
	DataId             string   `json:"dataId" gorm:"column:dataID"`
	TargetName         string   `json:"targetName" gorm:"column:targetName"`
	TargetSort         int      `json:"targetSort" gorm:"column:targetSort"`
	GroupName          string   `json:"groupName" gorm:"column:groupName"`
	GroupSort          int      `json:"groupSort" gorm:"column:groupSort"`
	AddAt              string   `json:"addAt" gorm:"column:addAt"`
	AddBy              string   `json:"addBy" gorm:"column:addBy"`
	AddIp              string   `json:"addIp" gorm:"column:addIp"`
	AddAgent           string   `json:"addAgent" gorm:"column:addAgent"`
	IsGroupNameDisplay int      `json:"isGroupNameDisplay" gorm:"column:isGroupNameDisplay"`
	Content            string   `json:"content" gorm:"content"`
	TargetContent      string   `json:"targetContent" gorm:"column:targetContent"`
	ContentArray       []string `json:"contentArray" gorm:"-"`
}

// Notion type Target struct {
//Id                 string      `json:"id" gorm:"column:id" `
//DataId             string      `json:"dataId"`
//TargetName         string      `json:"targetName"`
//TargetSort         int         `json:"targetSort"`
//GroupName          string      `json:"groupName"`
//GroupSort          int         `json:"groupSort"`
//AddAt              time.Time   `json:"addAt"`
//AddBy              string      `json:"addBy"`
//AddIp              interface{} `json:"addIp"`
//AddAgent           interface{} `json:"addAgent"`
//IsGroupNameDisplay int         `json:"isGroupNameDisplay"`
//Content            interface{} `json:"content"`
//TargetContent      string      `json:"targetContent"`
//ContentArray       []string    `json:"contentArray"`
//}
type Notion struct {
	Memo string `gorm:"column:memo"`
}

// 创建听课记录并返回指标信息
type GetAdd struct {
	Result  Result `json:"result"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Result struct {
	Type             string       `json:"type"`
	TName            string       `json:"tName" gorm:"column:tName"`
	TNameInSchool    string       `json:"tNameInSchool" gorm:"column:tNameInSchool"`
	CourseName       string       `json:"courseName" gorm:"column:courseName"`
	CourseType       string       `json:"courseType" gorm:"column:courseType"`
	ClassName        string       `json:"className" gorm:"column:className"`
	StuInSchool      string       `json:"stuInSchool" gorm:"column:stuInSchool"`
	AllNum           string       `json:"allNum" gorm:"column:allNum"`
	FactNum          string       `json:"factNum" gorm:"column:factNum"`
	Attendance       string       `json:"attendance" gorm:"column:attendance"`
	Evaluation       string       `json:"evaluation" gorm:"column:evaluation"`
	EvaluationDetail string       `json:"evaluationDetail" gorm:"column:evaluationDetail"`
	Supervisor       string       `json:"supervisor" gorm:"column:supervisor"`
	ListenAdd        string       `json:"listenAdd" gorm:"column:listenAdd"`
	Status           int          `json:"status" gorm:"column:status"`
	Semester         string       `json:"semester" gorm:"column:semester"`
	Id               string       `json:"id" gorm:"column:id"`
	ListenDay        time.Time    `json:"listenDay" gorm:"column:listenDate"`
	ListenTime       string       `json:"listenTimeArray" gorm:"column:listenTime"`
	Notions          string       `json:"notions"`
	Targets          []TargetData `json:"targets" gorm:"-"`
}

type ResFault struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// 返回督导员列表
type ResponseSupervisor struct {
	Result  []Supervisors `json:"result"`
	Message string        `json:"message"`
	Code    int           `json:"code"`
}

type Supervisors struct {
	MobilePhone string `json:"mobilePhone" gorm:"column:mobilePhone"`
	Name        string `json:"name" gorm:"column:name"`
	Status      int    `json:"status" gorm:"column:status"`
	Password    string `json:"password" gorm:"column:password"`
	Authority   string `json:"authority" gorm:"column:authority"`
	//Token     string `json:"token" gorm:"column:token"`
}

// 获取听课的课程及其次数
type ResponseDataCount struct {
	Result  DataCount `json:"result"`
	Message string    `json:"message"`
	Code    int       `json:"code"`
}

type DataCount struct {
	TotalCount int64       `json:"totalCount"`
	TypeCount  []TypeCount `json:"typeCount"`
}

type TypeCount struct {
	Type  string `json:"type" gorm:"column:type"`
	Count int    `json:"count" gorm:"column:count"`
}

type UpData struct {
	Id               string    `json:"id" gorm:"column:id"`
	TName            string    `json:"tName" gorm:"column:tName"`
	TNameInSchool    string    `json:"tNameInSchool" gorm:"column:tNameInSchool"`
	CourseName       string    `json:"courseName" gorm:"column:className"`
	CourseType       string    `json:"courseType" gorm:"column:courseType"`
	ClassName        string    `json:"className" gorm:"column:className"`
	StuInSchool      string    `json:"stuInSchool" gorm:"column:stuInSchool"`
	AllNum           string    `json:"allNum" gorm:"column:allNum"`
	FactNum          string    `json:"factNum" gorm:"column:factNum"`
	Attendance       string    `json:"attendance" gorm:"column:attendance"`
	Evaluation       string    `json:"evaluation" gorm:"column:evaluation"`
	EvaluationDetail string    `json:"evaluationDetail" gorm:"column:evaluationDetail"`
	Supervisor       string    `json:"supervisor" gorm:"column:supervisor"`
	ListenDate       time.Time `json:"listenDate" gorm:"column:listenDate"`
	ListenTime       string    `gorm:"column:listenTime"`
	ListenTimeArray  []int     `json:"listenTimeArray" gorm:"listenTimeArray"`
	ListenAdd        string    `json:"listenAdd" gorm:"column:listenAdd"`
	Status           int       `json:"status" gorm:"column:status"`
	Semester         string    `json:"semester" gorm:"column:semester"`
	Targets          []Targets `json:"targets" gorm:"column:targets"`
}

type Targets struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}
