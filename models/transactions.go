package models

import (
	u "goodness/utils"
)

type RefreshToken struct {
	User uint64 `json:"user";sql:"user";gorm:"primary_key"`
	Url string `json:"url";sql:"url";gorm:"url"`
	Token string `json:"token";sql:"token";gorm:"token"`
}

func (token *RefreshToken) Save() (map[string]interface{}) {
	err := GetDB().Table("refresh_tokens").FirstOrCreate(&RefreshToken{User:token.User})
	err = GetDB().Table("refresh_tokens").Save(&token)
	if err != nil {
		return u.Message(false, "Connection error. Please retry")
	}
	return u.Message(true, "Saved")
}
