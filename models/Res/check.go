/**
  @author: qianyi  2022/1/1 19:53:00
  @note:
*/
package Res


type LoginResult struct {
	ReturnUrl	string 	 `json:"returnUrl"`
	Token	string       `json:"token"`
	Authority	string   `json:"authority"`
	Result	string       `json:"result"`
	Message	string       `json:"message"`
	Code	int  		 `json:"code"`
}

type ResponseIn struct {
	Result []string `json:"result"`
	Message string  `json:"message"`
	Code  int       `json:"code"`
}

type ResponseOut struct {
	Result bool `json:"result"`
	Message string  `json:"message"`
	Code  int       `json:"code"`
}