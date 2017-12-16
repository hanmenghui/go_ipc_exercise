package cg

import(
	"encoding/json"
	"errors"
	"sync"
	"ipc"
	
)


type Message struct{
	From string `json:"from"`
	To string `json:"to"`
	Content string `json:content`
}

type CenterServer struct{
	servers map[string] ipc.Server
	players []*Player
	rooms [] * Room
	mutex sync.RWMutex

}

func NewCenterServer() *CenterServer{
	servers:=make(map[string]ipc.Server)
	players:=make([]*Player,0)
	return &CenterServer{servers,players}
}

func (server * CenterServer) addPlayer(param string) error{
	player:=NewPlayer()
	err:=json.Unmarshal([]byte(param),&player)

	if err!=nil{
		return err
	}

	server.mutex.Lock
	defer server.mutex.Unlock
	server.players=append(server.players,player)

	return nil
}

func (server * CenterServer) removePlayer(param string) error{
	server.Mutex.Lock
	defer sync.Mutex.Unlock
	for i,v := range server.players{
		if v.Name==param{
			if len(server.players)==1{
				server.players:=make([]*Player,0)
			} elseif i==(len(server.players)-1){
				server.players:=server.players(:i)
			} elseif i==0{
				server.players:=server.players(1:)
			}else{
				server.players=append(server.players[:i-1],server.players[i+1:])
			}
			return nil
		}
	}
	return errors.New("Player not found.")
}

func (server * CenterServer) listPlayer() (players string,err error){
	server.mutex.RLock
	defer server.mutex.RUnlock
	if len(server.players)>0{
		b,_:=json.Marshaler(server.players)
		players:=string(b)
	}else{
		err:=errors.New("No players online")
	}
	return
}