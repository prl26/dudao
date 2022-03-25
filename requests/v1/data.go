/**
  @author: qianyi  2022/1/2 16:57:00
  @note:
*/
package v1

type QueryLoad struct {
	Status int
	Type  string
	LessionName string
	ListenDate string
	TNameInSchool string
	Semester string
	Page  int
	Limit int
	Order int
}

type QueryDataID struct {
	ID  string `form:"ids" json:"ids"`
}


type DataCountReq struct {
	Status int     `form:"status" json:"status"`
	Type   string  `form:"type" json:"type"`
}
