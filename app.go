package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"market_apis/configs"
	"market_apis/routers"

	"github.com/labstack/echo/v4"
	gomLog "github.com/labstack/gommon/log"
)

var (
	configuation = configs.GetConfig()
)

func main() {

	var (
		e *echo.Echo = routers.GetRouters()

		appName     string = configs.GetConfig().AppName
		appVersion  string = configs.GetConfig().AppVersion
		environment string = configs.GetConfig().Enironment
		host        string = configs.GetConfig().APIHost
		port        string = configs.GetConfig().APIPort
	)

	go func() {
		e.HideBanner = true

		switch configuation.Enironment {
		case "local":
			e.Logger.SetLevel(gomLog.DEBUG)
		case "staging":
			e.Logger.SetLevel(gomLog.INFO)
		case "production":
			e.Logger.SetLevel(gomLog.ERROR)
		}

		e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", host, port)))
	}()

	log.Printf(`
	-----------------------------------------------------
	App name: %s
	Version: %s
	Listening Port: %v
	Environment: %s
	-----------------------------------------------------
	`, appName, appVersion, port, environment)

	// trap sigterm or interupt and gracefully shutdown the server.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 20 seconds for current operations to complete
	ctx, cancelFunction := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunction()

	e.Shutdown(ctx)
}
