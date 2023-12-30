package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func calculateFib(c echo.Context) error {
	inputStr := c.Param("n")
	n, err := strconv.Atoi(inputStr)
	if err != nil {
		panic(err)
	}
	return c.String(http.StatusOK, fmt.Sprintf("hello, my name is sina sharaiti and result of fib(%s) = %v", inputStr, fib(n)))
}

func main() {
	e := echo.New()
	e.GET("/fib/:n", calculateFib)
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
