package main

import (
	"apache-thrift/src/thrift-rpc/myapp"
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
)

func main() {
	Ping()
}

func Ping() {
	tSocket := thrift.NewTSocketConf("localhost:8080", nil)
	transportFactory := thrift.NewTTransportFactory()

	transport, err := transportFactory.GetTransport(tSocket)
	if err != nil {
		return
	}
	defer transport.Close()

	if err = transport.Open(); err != nil {
		log.Println("failed to ping with error:", err)
		return
	}

	protocolFactory := thrift.NewTCompactProtocolFactoryConf(&thrift.TConfiguration{})
	cli := myapp.NewUserServiceClientFactory(transport, protocolFactory)

	resp, err := cli.Ping(context.Background())
	if err != nil {
		log.Println("failed to ping with error:", err)
		return
	}

	log.Printf("Success: \nMessage:%s \nVersion:%d", resp.GetMessage(), resp.GetVersion())
}
