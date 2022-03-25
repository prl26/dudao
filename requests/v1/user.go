/**
  @author: qianyi  2021/12/28 20:34:00
  @note:
*/
package v1

////RegisterRequest 注册入参
//type RegisterRequest struct {
//	Account string `form:"account" binding:"required" json:"account"`
//	Password string `form:"password" binding:"required" json:"password"`
//}

//LoginRequest 登陆入参
type LoginRequest struct {
	Account string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// ChangePwdRep 修改密码入参
type  ChangePwdRep struct {
	Password string `form:"passwd" json:"passwd" binding:"required"`
}