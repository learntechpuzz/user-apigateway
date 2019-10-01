package handler

import (
	"fmt"
	"net/http"

	"user-apigateway/app/model"

	"github.com/labstack/echo/v4"
	nats "github.com/nats-io/nats.go"
)

// NewUserHandler creates a new user handler with NATS server encoded connection
func NewUserHandler(e *echo.Echo, nc *nats.EncodedConn) {
	handler := UserHandler{nc: nc}
	e.POST("/v1/api/users", handler.CreateUser)
}

// UserHandler contains NATS server encoded connection
type UserHandler struct {
	nc *nats.EncodedConn
}

// CreateUser
func (handler *UserHandler) CreateUser(c echo.Context) error {
	user := new(model.User)

	handler.nc.Subscribe("user.create", func(u *model.User) {
		fmt.Printf("\nReceived a user: %+v", u)
	})

	handler.nc.Publish("user.create", user)

	return c.JSON(http.StatusCreated, user)
}
