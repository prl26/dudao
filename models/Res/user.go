/**
  @author: qianyi  2022/1/1 19:53:00
  @note:
*/
package Res


type Response struct {
	Result string `json:"result"`
	Message string  `json:"message"`
	Code  int       `json:"code"`
}