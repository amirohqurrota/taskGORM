package controllers

import (
	"GORM-2/lib/database"
	"GORM-2/models"
	"net/http"

	"github.com/labstack/echo"
)

func ReqUserBookController(echoContext echo.Context) error {
	// var user models.User
	users, err := database.ReqUserBook()

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "not found",
			"data":   err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "successful operation",
		"data":   users,
	})

}

func ReqAddNewUserController(echoContext echo.Context) error {
	var user models.User
	echoContext.Bind(&user)

	result, err := database.ReqAddNewUser(user)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "error",
			"data":   err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "successful operation",
		"data":   result,
	})
}

func ReqDeleteUserController(echoContext echo.Context) error {
	var user models.User
	users, err := database.ReqDeleteUser(int(user.ID))

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "not found",
			"data":   err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "successful deleted",
		"data":   users,
	})

}
