package model

import (
	"gorm.io/gorm"
)

type AccountHistory struct {
	gorm.Model
	AccountID      uint    `json:"account_id" gorm:"not null;column:account_id"`
	Type           string  `json:"type" gorm:"type:text;column:type"`
	Nominal        int64   `json:"nominal" gorm:"type:bigint;column:nominal"`
	CurrentBalance int64   `json:"current_balance" gorm:"type:bigint;column:current_balance"`
	NewBalance     int64   `json:"new_balance" bson:"type:bigint;column:new_balance"`
	Account        Account `gorm:"foreignKey:AccountID;references:ID;onDelete:CASCADE"`
}
