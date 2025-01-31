package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name    string `json:"name" gorm:"type:text;column:name"`
	NIK     string `json:"nik" gorm:"type:text;unique;column:nik"`
	Phone   string `json:"phone" gorm:"type:text;unique;column:phone"`
	Number  string `json:"number" gorm:"type:text;unique;column:number"`
	Balance int64  `json:"balance" gorm:"type:bigint;column:balance"`
}
