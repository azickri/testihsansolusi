package web

import (
	"net/http"
	"testihsansolusi/entity"
	"testihsansolusi/service"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	AccountService *service.AccountService
}

func (controller *Controller) register(context echo.Context) error {
	var body entity.BodyRegister
	if err := context.Bind(&body); err != nil {
		context.JSON(http.StatusInternalServerError, entity.ResponseError{Remark: err.Error()})
		return nil
	}

	if body.Name == "" || body.NIK == "" || body.Phone == "" {
		context.JSON(http.StatusBadRequest, entity.ResponseError{Remark: "nama, nik or no_hp should not be empty"})
		return nil
	}

	number, err := controller.AccountService.Register(body.Name, body.NIK, body.Phone)
	if err != nil {
		context.JSON(http.StatusBadRequest, entity.ResponseError{Remark: err.Error()})
		return nil
	}

	context.JSON(http.StatusOK, entity.ResponseRegister{Number: number})
	return nil
}

func (controller *Controller) deposit(context echo.Context) error {
	var body entity.BodyDepositOrWithdraw
	if err := context.Bind(&body); err != nil {
		context.JSON(http.StatusInternalServerError, entity.ResponseError{Remark: err.Error()})
		return nil
	}

	if body.Number == "" || body.Nominal == 0 {
		context.JSON(http.StatusBadRequest, entity.ResponseError{Remark: "no_rekening or nominal should not be empty"})
		return nil
	}

	balance, err := controller.AccountService.Deposit(body.Number, body.Nominal)
	if err != nil {
		context.JSON(http.StatusBadRequest, entity.ResponseError{Remark: err.Error()})
		return nil
	}

	context.JSON(http.StatusOK, entity.ResponseBalance{Balance: balance})
	return nil
}

func (controller *Controller) witdraw(context echo.Context) error {
	var body entity.BodyDepositOrWithdraw
	if err := context.Bind(&body); err != nil {
		context.JSON(http.StatusInternalServerError, entity.ResponseError{Remark: err.Error()})
		return nil
	}

	if body.Number == "" || body.Nominal == 0 {
		context.JSON(http.StatusBadRequest, entity.ResponseError{Remark: "no_rekening or nominal should not be empty"})
		return nil
	}

	balance, err := controller.AccountService.Withdraw(body.Number, body.Nominal)
	if err != nil {
		context.JSON(http.StatusBadRequest, entity.ResponseError{Remark: err.Error()})
		return nil
	}

	context.JSON(http.StatusOK, entity.ResponseBalance{Balance: balance})
	return nil
}

func (controller *Controller) balance(context echo.Context) error {
	var number = context.Param("no_rekening")
	if number == "" {
		context.JSON(http.StatusBadRequest, entity.ResponseError{Remark: "no_rekening should not be empty"})
		return nil
	}

	balance, err := controller.AccountService.Balance(number)
	if err != nil {
		context.JSON(http.StatusBadRequest, entity.ResponseError{Remark: err.Error()})
		return nil
	}

	context.JSON(http.StatusOK, entity.ResponseBalance{Balance: balance})
	return nil
}
