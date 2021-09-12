package routers

import (
	"github.com/labstack/echo/v4"
)

var (
	e *echo.Echo = echo.New()
)

// GetRouters ..
func GetRouters() *echo.Echo {
	return e
}

func init(){
	
}
