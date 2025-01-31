package model

import "time"

type Account struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	NIK       string    `json:"nik" db:"nik"`
	Phone     string    `json:"phone" db:"phone"`
	Number    string    `json:"number" db:"number"`
	Balance   int64     `json:"balance" db:"balance"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
