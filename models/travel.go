package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Destination struct {
	gorm.Model
	User uint64 `json:"user";sql:"user"`
  TimeStamp time.Time  `json:"time";sql:"time"`
  Destination  string `json:"destination";sql:"destination"`
}
