package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Calulator struct{}

type Args struct {
	A, B int
}

func (c *Calulator) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func (c *Calulator) Sub(args *Args, reply *int) error {
	*reply = args.A - args.B
	return nil
}

func main() {
	calulator := new(Calulator)
	rpc.Register(calulator)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go jsonrpc.ServeConn(conn)
		fmt.Println("Server is running on port 8080")
	}

}
