package web

import (
	"os"
	"testihsansolusi/database"
	"testihsansolusi/service"

	"github.com/labstack/echo/v4"
)

func InitRoute(postgtresql *database.PostgreSQL) {
	router := echo.New()
	controller := Controller{
		AccountService: &service.AccountService{Database: postgtresql},
	}

	router.POST("/daftar", controller.register)
	router.POST("/tabung", controller.deposit)
	router.POST("/tarik", controller.witdraw)
	router.GET("/saldo/:no_rekening", controller.balance)

	router.Start(":" + os.Getenv("APPLICATION_PORT"))
}
