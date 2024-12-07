package helloworld

import (
	"fmt"
	"github/chaiyaphat20/go-basic/helloworld/calculator"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Helloworld() {
	fmt.Println("Hello World")
	fmt.Println(calculator.Add(1, 2))

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
