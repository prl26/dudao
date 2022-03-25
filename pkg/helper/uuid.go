/**
  @author: qianyi  2022/1/8 15:52:00
  @note:
*/
package helper

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func GetUuid() string{
	id := uuid.NewV4()
	ids := id.String()
	s := strings.Replace(ids,"-","",-1)
	return s
}
