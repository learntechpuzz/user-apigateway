package handler

import (
	"fmt"
	"net/http"
	"sync"

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
	var wg sync.WaitGroup
	wg.Add(1)

	user := new(model.User)


	ch := make(chan *model.User)

	// BindSendChan() allows binding of a Go channel to a nats
	// subject for publish operations. The Encoder attached to the
	// EncodedConn will be used for marshaling.	
	handler.nc.BindSendChan("user.create", ch)
	ch <- user

	// BindRecvChan() allows binding of a Go channel to a nats
	// subject for subscribe operations. The Encoder attached to the
	// EncodedConn will be used for un-marshaling.	
	handler.nc.BindRecvChan("user.create.completed", ch)

	u := <-ch

	fmt.Printf("%+v\n", u)

	return c.JSON(http.StatusCreated, u)
}



func publishUserCreate(){

}
