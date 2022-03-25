/**
  @author: qianyi  2021/12/28 20:11:00
  @note:
*/
package hash

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
	Cost int
}

// make加密方法
func (b *Bcrypt) Make(password []byte) ([]byte,error){
	return bcrypt.GenerateFromPassword(password,b.Cost)
}

// check检查方法
func (b *Bcrypt) Check(hashedPassword, password []byte) error{
	return bcrypt.CompareHashAndPassword(hashedPassword,password)
}

func NewHash() *Bcrypt{
	return &Bcrypt{
		Cost: bcrypt.DefaultCost,
	}
}