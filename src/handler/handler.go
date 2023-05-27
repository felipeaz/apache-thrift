package handler

import (
	"apache-thrift/src/thrift-rpc/myapp"
	"context"
	"log"
)

type UserHandler struct{}

var ApiVersion myapp.Int

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (p *UserHandler) Ping(ctx context.Context) (*myapp.PingResponse, error) {
	message := "ping"
	log.Println(message)

	PingResponse := myapp.NewPingResponse()
	PingResponse.Message = &message
	PingResponse.Version = &ApiVersion

	return PingResponse, nil
}
