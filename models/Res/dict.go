/**
  @author: qianyi  2022/1/2 14:33:00
  @note:
*/
package Res


type ResponseDict struct {
	Result   []ResultDict `json:"result"`
	Message  string       `json:"message"`
	Code     int      `json:"code"`
}

type ResultDict struct {
	ID string     `json:"id" gorm:"column:id"`
	Name string   `json:"name" gorm:"column:name"`
}

type GetTypeNameByID struct {
	Name string `gorm:"column:name"`
}
