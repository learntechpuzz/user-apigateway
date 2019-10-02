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

	if err := c.Bind(&user); err != nil {
		fmt.Printf("Failed to Bind JSON: %v", err.Error())
		return err
	}	

	// Publish to user.create via channel
	handler.nc.Publish("user.create", user)

	// Subscribe to user.create.completed via channel
	ch := make(chan *model.User)
	handler.nc.BindRecvChan("user.create.completed", ch)

	u := <-ch
	fmt.Printf("Received a User: %+v\n", u)


	return c.JSON(http.StatusCreated, u)
}



func publishUserCreate(){

}
