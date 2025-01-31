package model

import "time"

type AccountHistoryModel struct {
	ID             int       `json:"id" db:"id"`
	AccountID      int       `json:"account_id" db:"account_id"`
	Type           string    `json:"type" db:"type"`
	Nominal        int       `json:"nominal" db:"nominal"`
	CurrentBalance int       `json:"current_balance" db:"current_balance"`
	NewBalance     int       `json:"new_balance" bson:"new_balance"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
