package routers

import (
	"log"
	"market_apis/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	e *echo.Echo = echo.New()
)

// GetRouters ..
func GetRouters() *echo.Echo {
	return e
}

func init() {
	log.Println("Initializing Router for Market Service")

	var (
		productionController *controllers.ProductionController = controllers.NewProductionController()
	)

	echoGroup := e.Group("/api")
	echoGroup.Use(middleware.Logger())  // Logger
	echoGroup.Use(middleware.Recover()) // Recover

	// CORS
	echoGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	productRoute := echoGroup.Group("/product")
	productRoute.POST("/upload", productionController.UploadProduct)
	productRoute.GET("/list", productionController.GetProduct)

}
