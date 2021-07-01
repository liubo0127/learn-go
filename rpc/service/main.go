package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	rpcdemo "github.com/liubo0127/learn-go/rpc"
)

func main() {
	err := rpc.Register(rpcdemo.DemoService{})
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("error occured: %v", err)
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
