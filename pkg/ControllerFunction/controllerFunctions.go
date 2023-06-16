package ControllerFunction

import (
	"Generating/pkg/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
	"text/template"
)

func GenMain() {
	os.Mkdir("AsanTest", os.ModePerm)
	type DummyModel struct {
		Greeting string
	}

	dummyObj := DummyModel{
		Greeting: "Hello welcome",
	}

	//templating

	tmp1, err := template.ParseFiles("codesource/genMain.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)

	err = tmp1.Execute(buf, dummyObj)
	buff := []byte(buf.String())

	// saving to go file

	err = os.WriteFile("AsanTest/main.go", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}

}

func GenDatabaseConnection() {

	//templating
	tmp1, err := template.ParseFiles("codesource/genDatabaseConnection.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)

	err = tmp1.Execute(buf, "")
	buff := []byte(buf.String())

	// saving to go file

	os.Mkdir("AsanTest/connections/", os.ModePerm)                               //to make a directory
	err = os.WriteFile("AsanTest/connections/databaseConncetion.go", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}

}

func GenPaths(c echo.Context) error {
	//templating
	data := &models.Data{}
	if err := c.Bind(data); err != nil {
		return err
	}
	obj := models.Data{
		ID:           data.ID,
		Database_URI: data.Database_URI,
		DatabaseName: data.DatabaseName,
	}
	tmp1, err := template.ParseFiles("codesource/genPaths.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)

	err = tmp1.Execute(buf, obj)
	buff := []byte(buf.String())

	// saving to go file
	err = os.WriteFile("AsanTest/connections/paths.go", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}
	return err

}

func Api_connections() {

	//templating

	tmp1, err := template.ParseFiles("codesource/genapi_connection.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)
	obj := ""
	err = tmp1.Execute(buf, obj)
	buff := []byte(buf.String())
	//make directory
	os.Mkdir("AsanTest/pkg", os.ModePerm)
	os.Mkdir("AsanTest/pkg/apis", os.ModePerm)
	// saving to go file
	err = os.WriteFile("AsanTest/pkg/apis/api_connection.go", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}

}

func Api_createCustomer() {

	//templating

	tmp1, err := template.ParseFiles("codesource/genapi_createCustomer.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)
	obj := ""
	err = tmp1.Execute(buf, obj)
	buff := []byte(buf.String())
	// saving to go file
	err = os.WriteFile("AsanTest/pkg/apis/api_createCustomer.go", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}

}

func Api_deleteCustomer() {

	//templating

	tmp1, err := template.ParseFiles("codesource/genapi_deleteCustomer.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)
	obj := ""
	err = tmp1.Execute(buf, obj)
	buff := []byte(buf.String())
	// saving to go file
	err = os.WriteFile("AsanTest/pkg/apis/api_deleteCustomer.go", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}

}

func Api_findCustomer() {

	//templating

	tmp1, err := template.ParseFiles("codesource/genapi_findCustomer.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)
	obj := ""
	err = tmp1.Execute(buf, obj)
	buff := []byte(buf.String())
	// saving to go file
	err = os.WriteFile("AsanTest/pkg/apis/api_findCustomer.go", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}

}

func Api_updateCustomer() {

	//templating

	tmp1, err := template.ParseFiles("codesource/genapi_updateCustomer.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)
	obj := ""
	err = tmp1.Execute(buf, obj)
	buff := []byte(buf.String())
	// saving to go file
	err = os.WriteFile("AsanTest/pkg/apis/api_updateCustomer.go", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}

}

func ApiController() {

	//templating

	tmp1, err := template.ParseFiles("codesource/genapiController.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)
	obj := ""
	err = tmp1.Execute(buf, obj)
	buff := []byte(buf.String())
	// saving to go file
	err = os.WriteFile("AsanTest/pkg/apis/apiController.go", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}

}

func CreateCustomer() {

	//templating

	tmp1, err := template.ParseFiles("codesource/genCreateCustomer.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)
	obj := ""
	err = tmp1.Execute(buf, obj)
	buff := []byte(buf.String())
	os.Mkdir("AsanTest/pkg/databaseOps", os.ModePerm)
	// saving to go file
	err = os.WriteFile("AsanTest/pkg/databaseOps/CreateCustomer.go", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}

}

func DeleteCustomer() {

	//templating

	tmp1, err := template.ParseFiles("codesource/gendeleteCustomer.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)
	obj := ""
	err = tmp1.Execute(buf, obj)
	buff := []byte(buf.String())
	// saving to go file
	err = os.WriteFile("AsanTest/pkg/databaseOps/deleteCustomer.go", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}

}

func FindCustomer() {

	//templating

	tmp1, err := template.ParseFiles("codesource/genfindCustomer.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)
	obj := ""
	err = tmp1.Execute(buf, obj)
	buff := []byte(buf.String())
	// saving to go file
	err = os.WriteFile("AsanTest/pkg/databaseOps/findCustomer.go", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}

}

func CreateModel() {

	//templating

	tmp1, err := template.ParseFiles("codesource/genCustomerModel.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)
	obj := ""
	err = tmp1.Execute(buf, obj)
	buff := []byte(buf.String())
	os.Mkdir("AsanTest/pkg/models", os.ModePerm)
	// saving to go file
	err = os.WriteFile("AsanTest/pkg/models/CustomerModel.go", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}

}

func CreateGo() {

	//templating

	tmp1, err := template.ParseFiles("codesource/genGo.tmp")
	if err != nil {
		fmt.Println(err)
	}
	buf := new(strings.Builder)
	obj := ""
	err = tmp1.Execute(buf, obj)
	buff := []byte(buf.String())
	// saving to go file
	err = os.WriteFile("AsanTest/go.mod", buff, 0644) //
	if err != nil {
		fmt.Println(err)
	}

}

func genFunction() {

}
