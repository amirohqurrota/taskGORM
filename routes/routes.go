package routes

import (
	"GORM-2/controllers"

	"github.com/labstack/echo"
)

func NewRoutes() *echo.Echo {

	echoInit := echo.New()

	echoInit.GET("/user", controllers.ReqUserBookController)
	echoInit.POST("/user", controllers.ReqAddNewUserController)
	echoInit.PUT("/article", controllers.ReqUpdateUserController)
	echoInit.DELETE("/user", controllers.ReqDeleteUserController)

	return echoInit

}
