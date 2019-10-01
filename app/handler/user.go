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

	// Publish to user.create via channel
	pch := make(chan *model.User)
	handler.nc.BindSendChan("user.create", pch)
	pch <- user

	// Subscribe to user.create.completed via channel
	sch := make(chan *model.User)
	handler.nc.BindRecvChan("user.create.completed", sch)

	u := <-sch
	fmt.Printf("Received a User: %+v\n", u)


	return c.JSON(http.StatusCreated, u)
}



func publishUserCreate(){

}
