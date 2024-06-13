package game

import (
	"github.com/asynkron/protoactor-go/actor"
	"pigeon/pb/movement"
	"strconv"
)

type (
	GameActor struct {
		Server *actor.PID
	}
)

func (g *GameActor) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case *movement.Test:
		g.Test(msg, c)
	}
}

func (g *GameActor) Test(msg *movement.Test, c actor.Context) {
	c.Logger().Info(strconv.Itoa(int(msg.Test)))
}
