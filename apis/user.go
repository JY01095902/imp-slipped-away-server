package apis

import (
	"net/http"
	"github.com/labstack/echo"
	services "imp-slipped-away-server/services"
	entities "imp-slipped-away-server/entities"
)

func InitUserAPIS(e *echo.Echo){
	e.POST("/users", saveUser)
}

func saveUser(c echo.Context) error {
	user := new(entities.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	err := services.SaveUser(user)
	if (err == nil) {
		return c.JSON(http.StatusCreated, user)
	}else {
		return c.JSON(http.StatusInternalServerError, err)
	}
}