package network

import (
	"github.com/asynkron/protoactor-go/actor"
	"net"
)

type (
	SessionActor struct {
		conn   net.Conn
		server *actor.PID
	}
)

func (s *SessionActor) Receive(c actor.Context) {
	switch c.Message().(type) {
	case *actor.Started:
		for {

			c.Send(s.server, RemoveSession{addr: s.conn.RemoteAddr()})
		}
	}
}
