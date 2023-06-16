package apis

import (
	"Generating/pkg/ControllerFunction"
	"Generating/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	users = map[int]*models.Data{}
	seq   = 1
)

func createData(c echo.Context) error {

	data := &models.Data{
		ID: seq,
	}
	if err := c.Bind(data); err != nil {
		return err
	}
	users[data.ID] = data
	seq++
	return c.JSON(http.StatusCreated, data)

}

func getData(c echo.Context) error {
	data := &models.Data{}
	if err := c.Bind(data); err != nil {
		return err
	}

	ControllerFunction.GenPaths(c)

	return c.JSON(http.StatusOK, data)
}

func triggerAPI(c echo.Context) error {

	ControllerFunction.GenMain()
	ControllerFunction.GenDatabaseConnection()
	ControllerFunction.GenPaths(c)
	ControllerFunction.Api_connections()
	ControllerFunction.Api_createCustomer()
	ControllerFunction.Api_deleteCustomer()
	ControllerFunction.Api_findCustomer()
	ControllerFunction.Api_updateCustomer()
	ControllerFunction.ApiController()
	ControllerFunction.CreateCustomer()
	ControllerFunction.CreateModel()
	ControllerFunction.CreateGo()
	ControllerFunction.FindCustomer()
	ControllerFunction.DeleteCustomer()
	return c.String(http.StatusOK, "API triggered successfully")

}
