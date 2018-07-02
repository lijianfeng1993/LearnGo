package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method,params string) (resq *Response) {
	var body string
	if method == "get" {
		body = "ECHO:" + params
	}
	return &Response{"200", body}
}

func (server *EchoServer) Name() string {
	return "Echoserver"
}

func TestIpc(t *testing.T) {

	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("get", "From client1")
	resp2, _ := client2.Call("get", "From client2")

	if resp1.Body != "ECHO:From client1" || resp2.Body != "ECHO:From client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}

	client1.Close()
	client2.Close()
}
