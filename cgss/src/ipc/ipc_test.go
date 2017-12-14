package ipc

import (
	"fmt"
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, param string) *Response {
	return &Response{"OK", "ECHO: " + method + " ~ " + param}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("foo", "From client1")
	resp2, _ := client2.Call("foo", "From client2")

	fmt.Println(resp1.Code + ":" + resp1.Body)
	fmt.Println(resp2.Code + ":" + resp2.Body)

}
