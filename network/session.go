package network

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/golang/protobuf/proto"
	"net"
	"pigeon/pb/movement"
	"pigeon/utils"
)

type (
	SessionActor struct {
		conn   net.Conn
		server *actor.PID
		game   *actor.PID
	}
)

var packets = map[int32]proto.Message{
	1: &movement.Test{},
}

func (s *SessionActor) Receive(c actor.Context) {
	switch c.Message().(type) {
	case *actor.Started:
		for {
			buffer := make([]byte, 1024)
			_, err1 := s.conn.Read(buffer)
			if err1 != nil {
				s.HandleError(c, err1)
				return
			}
			tag, data, err2 := utils.Decode(buffer)
			if err2 != nil {
				s.HandleError(c, err2)
				return
			}
			packet := proto.Clone(packets[tag])
			err3 := ReadProto(data, &packet)
			if err3 != nil {
				s.HandleError(c, err3)
				return
			}

			c.Send(s.game, packet)
		}
	}
}

func (s *SessionActor) HandleError(c actor.Context, err error) {
	c.Logger().Error(err.Error())
	c.Send(s.server, RemoveSession{addr: s.conn.RemoteAddr()})
	err1 := s.conn.Close()
	if err1 != nil {
		c.Logger().Error(err1.Error())
	}
}

func ReadProto[T *proto.Message](data []byte, packet T) error {
	err := proto.Unmarshal(data, *packet)
	return err
}
