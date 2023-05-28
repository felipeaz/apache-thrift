package server

import (
	handler "apache-thrift/src/handler"
	"apache-thrift/src/thrift-rpc/myapp"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
)

func RunServer(addr string) error {
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	processor := buildProcessor()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactoryConf(&thrift.TConfiguration{})
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	log.Printf("Starting simple thrift server on: %s\n", addr)
	return server.Serve()
}

func buildProcessor() *myapp.UserServiceProcessor {
	usrHandler := handler.NewUserHandler()
	return myapp.NewUserServiceProcessor(usrHandler)
}
