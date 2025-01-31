package service

import (
	"database/sql"
	"errors"
	"fmt"
	"testihsansolusi/database"
	"testihsansolusi/entity/model"
	"testihsansolusi/helper"
	"time"
)

type AccountService struct {
	Database *database.PostgreSQL
}

func (accountService *AccountService) Register(name, nik, phone string) (string, error) {
	var account model.Account

	row := accountService.Database.Client.QueryRow("SELECT id FROM accounts WHERE nik = $1 OR phone = $2 LIMIT 1", nik, phone)
	err := row.Scan(&account.ID)
	if errors.Is(err, sql.ErrNoRows) {
		number := helper.RandomNumber(16)

		query := "INSERT INTO accounts (name, nik, phone, number, balance, created_at) VALUES ($1, $2, $3, $4, 0, $5)"
		accountService.Database.Client.QueryRow(query, name, nik, phone, number, time.Now())

		return number, nil
	}

	fmt.Println("INFO: user registered with existing nik or phone", "NIK:", nik, "PHONE:", phone)
	return "", errors.New("nik or phone has registered")
}

func (accountService *AccountService) Deposit(number string, nominal int64) (int64, error) {
	var account model.Account

	row := accountService.Database.Client.QueryRow("SELECT id, balance FROM accounts WHERE number = $1", number)
	err := row.Scan(&account.ID, &account.Balance)
	if err != nil {
		fmt.Println("ERROR: Client.QueryRow", err.Error())
		return 0, err
	}

	accountService.Database.Client.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", nominal, account.ID)
	go func() {
		query := "INSERT INTO account_histories (account_id, type, nominal, current_balance, new_balance, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
		accountService.Database.Client.Exec(query, account.ID, "deposit", nominal, account.Balance, account.Balance+nominal, time.Now())
	}()

	return account.Balance + nominal, nil
}

func (accountService *AccountService) Withdraw(number string, nominal int64) (int64, error) {
	var account model.Account

	row := accountService.Database.Client.QueryRow("SELECT id, balance FROM accounts WHERE number = $1", number)
	err := row.Scan(&account.ID, &account.Balance)
	if err != nil {
		fmt.Println("ERROR: Client.QueryRow", err.Error())
		return 0, err
	}

	if account.Balance < nominal {
		return 0, errors.New("nominal should be lower or same than balance")
	}

	accountService.Database.Client.Exec("UPDATE accounts SET balance = balance - $1 WHERE id = $2", nominal, account.ID)
	go func() {
		query := "INSERT INTO account_histories (account_id, type, nominal, current_balance, new_balance, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
		accountService.Database.Client.Exec(query, account.ID, "withdraw", nominal, account.Balance, account.Balance-nominal, time.Now())
	}()

	return account.Balance - nominal, nil
}

func (accountService *AccountService) Balance(number string) (int64, error) {
	var account model.Account

	row := accountService.Database.Client.QueryRow("SELECT id, balance FROM accounts WHERE number = $1", number)
	err := row.Scan(&account.ID, &account.Balance)
	if err != nil {
		fmt.Println("ERROR: Client.QueryRow", err.Error())
		return 0, err
	}

	return account.Balance, nil
}
