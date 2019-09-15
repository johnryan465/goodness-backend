package models

import (
	"github.com/jinzhu/gorm"
	u "goodness/utils"
	"time"
	"fmt"
)

func Scores(user uint64) (map[string] interface{}, bool) {
	var  temp []Score
	err := GetDB().Table("scores").Where("user = ?", user).Find(&temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	resp := u.Message(false, "Success")
	fmt.Println(temp)
	resp["scores"] = temp
	return resp, true
}

type Score struct {
	gorm.Model
	User int `json:"user";sql:"user"`
  TimeStamp time.Time  `json:"time";sql:"time"`
  Value  float64 `json:"score";sql:"score"`
}
