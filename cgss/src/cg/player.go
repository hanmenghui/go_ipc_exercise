package cg

import(
	"fmt"
)

type Player struct{
	Name string
	Level int
	Exp int
	Room int

	mq chan * string
}

func NewPlayer() *Player{
	m:=make(chan *string,1024)
	player:=&Player{"",0,0,0,m}
	go func(p *Player){
		for{
			message :=<-p.mq
			fmt.Println(p.Name,"receiced message:",message)
		}
	}(player)
	return player
}