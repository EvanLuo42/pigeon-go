package network

import (
	"github.com/asynkron/protoactor-go/actor"
	"net"
)

type (
	ListenerActor struct {
		listener net.Listener
		server   *actor.PID
	}
)

func (l *ListenerActor) Receive(c actor.Context) {
	switch c.Message().(type) {
	case *actor.Started:
		for {
			conn, err := l.listener.Accept()
			if err != nil {
				c.Logger().Error(err.Error())
			}
			c.Send(l.server, AddSession{conn: conn})
		}
	}
}
