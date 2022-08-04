package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	fmt.Println("user list")
	return nil
}

func Create(c echo.Context) error {
	fmt.Println("user Create")
	return nil
}

func Update(c echo.Context) error {
	fmt.Println("user Update")
	return nil
}

func Delete(c echo.Context) error {
	fmt.Println("user Delete")
	return nil
}
