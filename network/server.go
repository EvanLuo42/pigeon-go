package network

import (
	"github.com/asynkron/protoactor-go/actor"
	"log"
	"net"
	"strconv"
)

type (
	ServerActor struct {
		Port    int
		session map[net.Addr]*actor.PID
	}
)

func (s ServerActor) Receive(c actor.Context) {
	switch c.Message().(type) {
	case *actor.Started:
		s.session = make(map[net.Addr]*actor.PID)
		s.Listen(s.Port, c)
	}
}

func (s ServerActor) Listen(port int, c actor.Context) {
	listen, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		props := actor.PropsFromProducer(func() actor.Actor {
			return &SessionActor{}
		})
		pid := c.Spawn(props)
		s.session[conn.RemoteAddr()] = pid
		c.Logger().Info("Create new session", "ip", conn.RemoteAddr().String())
	}
}
