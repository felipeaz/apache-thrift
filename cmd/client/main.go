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
	// Get information about the RPC server running in a given address
	tSocket := thrift.NewTSocketConf("localhost:8080", nil)
	transportFactory := thrift.NewTTransportFactory()

	// Fetch the transport from the socket fetched above
	transport, err := transportFactory.GetTransport(tSocket)
	if err != nil {
		return
	}
	defer transport.Close()

	// Opens the transport for communication
	if err = transport.Open(); err != nil {
		log.Println("failed to ping with error:", err)
		return
	}

	// Initialize the protocol factory of type CompactProtocol.
	// More info about protocols here: https://thrift-tutorial.readthedocs.io/en/latest/thrift-stack.html
	protocolFactory := thrift.NewTCompactProtocolFactoryConf(&thrift.TConfiguration{})

	// Creates the RPC client using the transport + protocol factory
	cli := myapp.NewUserServiceClientFactory(transport, protocolFactory)

	// Calls RPC Server
	resp, err := cli.Ping(context.Background())
	if err != nil {
		log.Println("failed to ping with error:", err)
		return
	}
	log.Printf("Success: \nMessage:%s \nVersion:%d", resp.GetMessage(), resp.GetVersion())
}
