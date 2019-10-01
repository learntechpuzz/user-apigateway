package main

import (
	"log"

	"user-apigateway/app/config"
	"user-apigateway/app/handler"
	"user-apigateway/app/platform/nats"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	envFlag    = "env"
	defaultEnv = "dev"
	serverPort = "server.port"
	natsServer = "nats.server"
)

func main() {

	// Get environment flag
	env := pflag.String(envFlag, defaultEnv, "environment config value to use")
	pflag.Parse()

	if err := config.LoadConfiguration(*env); err != nil {
		checkErr(err)
	}

	// Create new NATS server connection
	nc, err := natsclient.NewNATSServerConnection(viper.GetString(natsServer))
	checkErr(err)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Handlers
	handler.NewUserHandler(e, nc)

	// Server
	e.Logger.Fatal(e.Start(":" + viper.GetString(serverPort)))

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
