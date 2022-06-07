/**
* @Author: 云坠
* @Date: 2022/6/7 15:28
**/
package helper

import "time"

var timeTemplates = []string {
	"2006-01-02 15:04:05", //常规类型
	"2006/01/02 15:04:05",
	"2006-01-02",
	"2006/01/02",
	"15:04:05",
}

func TimeStringToGoTime(tm string) time.Time {
	for i := range timeTemplates {
		t, err := time.ParseInLocation(timeTemplates[i], tm, time.Local)
		if nil == err && !t.IsZero() { return t }
	}
	return time.Time{}
}


