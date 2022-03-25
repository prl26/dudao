package Res

type GetCountRes struct {
	Result  []CountResult `json:"result"`
	Message string        `json:"message"`
	Code    int           `json:"code"`
}

type CountResult struct {
	TcName      string `json:"tc_name" gorm:"column:tc_name"`
	TcTotal     int    `json:"tc_total" gorm:"column:tc_total"`
	TcIng       int    `json:"tc_ing" gorm:"column:tc_ing"`
	TcRemainder int    `json:"tc_remainder" gorm:"column:tc_remainder"`
	TcSemester  string `json:"tc_semester" gorm:"column:tc_semester"`
	Id          string `json:"id" gorm:"column:id"`
}

type ResCount struct {
	Result  NewCount `json:"result"`
	Message string   `json:"message"`
	Code    int      `json:"code"`
}

type NewCount struct {
	HaveDone  int `json:"haveDone"`
	Listening int `json:"listening"`
	Listened  int `json:"listened"`
}
