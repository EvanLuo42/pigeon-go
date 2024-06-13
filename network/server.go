package network

import (
	"github.com/asynkron/protoactor-go/actor"
	"log"
	"net"
	"pigeon/game"
	"strconv"
)

type (
	ServerActor struct {
		Port     int
		sessions map[net.Addr]*actor.PID
		game     *actor.PID
	}

	RemoveSession struct {
		addr net.Addr
	}

	AddSession struct {
		conn net.Conn
	}
)

func (s *ServerActor) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case *actor.Started:
		listener, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(s.Port))
		if err != nil {
			log.Fatal(err)
		}
		props := actor.PropsFromProducer(func() actor.Actor {
			return &ListenerActor{listener: listener, server: c.Self()}
		})
		c.Spawn(props)
		props = actor.PropsFromProducer(func() actor.Actor {
			return &game.WorldActor{Server: c.Self()}
		})
		pid := c.Spawn(props)
		s.game = pid
	case RemoveSession:
		s.RemoveSession(msg.addr, c)
	case AddSession:
		s.AddSession(msg.conn, c)
	}
}

func (s *ServerActor) RemoveSession(addr net.Addr, c actor.Context) {
	c.Stop(s.sessions[addr])
	delete(s.sessions, addr)
	c.Logger().Info("Destroy session", "addr", addr)
}

func (s *ServerActor) AddSession(conn net.Conn, c actor.Context) {
	props := actor.PropsFromProducer(func() actor.Actor {
		return &SessionActor{conn: conn, server: c.Self(), game: s.game}
	})
	pid := c.Spawn(props)
	if s.sessions == nil {
		s.sessions = make(map[net.Addr]*actor.PID)
	}
	s.sessions[conn.RemoteAddr()] = pid
	c.Logger().Info("Create new session", "addr", conn.RemoteAddr().String())
}
