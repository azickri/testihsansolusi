package service

import (
	"errors"
	"fmt"
	"testihsansolusi/database"
	"testihsansolusi/entity/model"
	"testihsansolusi/helper"

	"gorm.io/gorm"
)

type AccountService struct {
	Database *database.PostgreSQL
}

func (accountService *AccountService) Register(name, nik, phone string) (string, error) {
	var account model.Account

	err := accountService.Database.Client.Where("nik = ? OR phone = ?", nik, phone).First(&account).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("ERROR: Register Client.Where.First", err.Error())
		return "", err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		number := helper.RandomNumber(16)
		newAccount := model.Account{
			Name:    name,
			NIK:     nik,
			Phone:   phone,
			Number:  number,
			Balance: 0,
		}

		err := accountService.Database.Client.Create(&newAccount).Error
		if err != nil {
			fmt.Println("ERROR: Client.Create", err.Error())
			return "", err
		}

		return number, nil
	}

	fmt.Println("INFO: user registered with existing nik or phone", "NIK:", nik, "PHONE:", phone)
	return "", errors.New("nik or no_hp has registered")
}

func (accountService *AccountService) Deposit(number string, nominal int64) (int64, error) {
	var account model.Account

	err := accountService.Database.Client.Where("number = ?", number).First(&account).Error
	if err != nil {
		fmt.Println("ERROR: Deposit Client.Where", err.Error())
		return 0, err
	}

	currentBalance := account.Balance
	newBalance := account.Balance + nominal

	newAccountHistory := model.AccountHistory{
		AccountID:      account.ID,
		Type:           "deposit",
		Nominal:        nominal,
		CurrentBalance: currentBalance,
		NewBalance:     newBalance,
	}

	err = accountService.Database.Client.Create(&newAccountHistory).Error
	if err != nil {
		fmt.Println("ERROR: Deposit Client.Create", err.Error())
		return 0, err
	}

	account.Balance = newBalance
	accountService.Database.Client.Save(&account)

	return newBalance, nil
}

func (accountService *AccountService) Withdraw(number string, nominal int64) (int64, error) {
	var account model.Account

	err := accountService.Database.Client.Where("number = ?", number).First(&account).Error
	if err != nil {
		fmt.Println("ERROR: Withdraw Client.Where", err.Error())
		return 0, err
	}

	if account.Balance < nominal {
		return 0, errors.New("nominal should be lower or same than balance")
	}

	currentBalance := account.Balance
	newBalance := account.Balance - nominal

	newAccountHistory := model.AccountHistory{
		AccountID:      account.ID,
		Type:           "withdraw",
		Nominal:        nominal,
		CurrentBalance: currentBalance,
		NewBalance:     newBalance,
	}

	err = accountService.Database.Client.Create(&newAccountHistory).Error
	if err != nil {
		fmt.Println("ERROR: Withdraw Client.Create", err.Error())
		return 0, err
	}

	account.Balance = newBalance
	accountService.Database.Client.Save(&account)

	return newBalance, nil
}

func (accountService *AccountService) Balance(number string) (int64, error) {
	var account model.Account

	err := accountService.Database.Client.Where("number = ?", number).First(&account).Error
	if err != nil {
		fmt.Println("ERROR: Balance Client.Where", err.Error())
		return 0, err
	}

	return account.Balance, nil
}
