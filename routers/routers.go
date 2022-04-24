package routers

import (
	"log"
	"market_apis/controllers"

	"market_apis/middlewares"

	"github.com/labstack/echo-contrib/prometheus"
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
		p                 = prometheus.NewPrometheus("myapp", nil)
		metricsMiddleware = middlewares.NewMetricsMiddleware()

		productionController *controllers.ProductionController = controllers.NewProductionController()
	)

	p.Use(e)

	echoGroup := e.Group("/api")
	echoGroup.Use(middleware.Logger())  // Logger
	echoGroup.Use(middleware.Recover()) // Recover
	echoGroup.Use(metricsMiddleware.Metrics)

	// CORS
	echoGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	productRoute := echoGroup.Group("/product")
	productRoute.POST("/upload", productionController.UploadProduct)
	productRoute.GET("/list/:id", productionController.GetProduct)

}
