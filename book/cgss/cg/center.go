package cg

import (
	"encoding/json"
	"errors"
	"sync"
	"LearnGo/book/cgss/ipc"


)


var _, ipc.Server = &CenterServer{}


type Message struct {
	From string "from"
	To string "to"
	Content string "content"
}

type CenterServer struct {
	servers map[string] ipc.Server
	players []*Player
	rooms []*Room
	mutex sync.RWMutex
}

func NewCenterServer() *CenterServer {
	servers := make(map[string] ipc.Server)
	players := make([]*Player, 0)
	return CenterServer{servers:servers, players:players}
}

func (server *CenterServer)addPlayer(params string) error {
	player := &NewPlayer()

	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	server.players = append(server.players, player)
	return nil
}

func (server *CenterServer)remotePlayer(params string) error {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	for i,v := range server.players {
		if v.Name == params {
			if len(server.players) == 1 {
				server.players = make([]*Player, 0)
			} else if i == len(server.players) - 1 {
				server.players = server.players[:i-1]
			} else if i == 0 {
				server.players = server.players[1:]
			} else {
				server.players = append(server.players[:i-1], server.players[i+1:]...)
			}
		}
		return nil
	}
	return error.New("Player not found.")
}

func (server *CenterServer)listPlayer(params string) (players string, err error) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	if len(server.players) > 0 {
		b,_ := json.marshal(server.players)
		players = string(b)
	} else {
		errors.New("No player online.")
	}
	return
}

func (server *CenterServer)broadcast(params string) error {
	var message Message
	err := json.Unmarshal([]byte(params), &message)
	if err != nil {
		return err
	}

	if len(server.players) > 0 {
		for _, player := range server.players {
			player.mq <- &message
		}
	} else {
		errors.New("No player online.")
	}

	return nil
}

func (server *CenterServer)Name() string {
	return "CenterServer"
}

func (server *CenterServer)Handle(method, params string) *ipc.Response{
	switch method {
	case "addplayer":
		err := server.addPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error(), Body: "addplayer failed."}
		}
	case "removeplayer":
		err := server.remotePlayer(params)
		if err != nil {
			return &ipc.Response{Code:err.Error(), Body: "removeplayer failed."}
		}
	case "listplayer":

	}


}