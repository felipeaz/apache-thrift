package main

import (
	"flag"
	"log"

	"apache-thrift/src/handler"
	"apache-thrift/src/server"
	"apache-thrift/src/thrift-rpc/myapp"
)

var (
	addr = flag.String("addr", "localhost:8080", "Thrift Address to listen on")
)

var ApiVersion = 1

func main() {
	flag.Parse()

	handler.ApiVersion = myapp.Int(ApiVersion)

	if err := server.RunServer(*addr, false); err != nil {
		log.Println("error running thrift server: ", err)
	}
}
