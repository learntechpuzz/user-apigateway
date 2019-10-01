package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

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

	// Start server
	go func() {
		if err := e.Start(":" + viper.GetString(serverPort)); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
		log.Println("Close NATS server encoded connection.")
		nc.Close()
	}
	log.Println("Server exiting")

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
