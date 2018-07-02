package ipc

import (
	"encoding/json"
	"fmt"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{c}
}

func (client *IpcClient)Close() {
	client.conn <- "CLOSE"
}

func (client *IpcClient)Call(method, params string) (resq *Response, err error) {
	req := &Request{method, params}

	var b []byte
	b, err = json.Marshal(req)
	if err != nil {
		return
	}

	client.conn <- string(b)

	res := <- client.conn

	var resq1 Response
	err2 := json.Unmarshal([]byte(res), &resq1)
	if err2 != nil {
		fmt.Println("Invalid response format:", res)
		return
	}

	resq = &resq1
	return resq, nil
}

