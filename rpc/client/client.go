package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"

	rpcdemo "github.com/liubo0127/learn-go/rpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{
		A: 10,
		B: 3,
	}, &result)

	log.Printf("result is %v, err is %v", result, err)

	err = client.Call("DemoService.Div", rpcdemo.Args{
		A: 10,
		B: 0,
	}, &result)

	log.Printf("result is %v, err is %v", result, err)
}
