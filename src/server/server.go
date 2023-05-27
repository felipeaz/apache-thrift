package server

import (
	handler "apache-thrift/src/handler"
	"apache-thrift/src/thrift-rpc/myapp"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
)

func RunServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	usrHandler := handler.NewUserHandler()
	processor := myapp.NewUserServiceProcessor(usrHandler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	log.Printf("Starting simple thrift server on: %s\n", addr)
	return server.Serve()
}
