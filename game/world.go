package game

import (
	"github.com/asynkron/protoactor-go/actor"
	"pigeon/ecs"
	"pigeon/pb/movement"
	"strconv"
	"time"
)

type (
	WorldActor struct {
		Server *actor.PID
		world  ecs.World
	}

	Update struct {
		deltaTime float64
	}
)

func (g *WorldActor) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case *actor.Started:
		g.world = ecs.World{}
		go func() {
			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()
			deltaTime := 0.0
			for {
				startTime := time.Now()
				select {
				case <-ticker.C:
					err := c.RequestFuture(c.Self(), Update{deltaTime: deltaTime}, 5*time.Second).Wait()
					if err != nil {
						c.Logger().Error(err.Error())
					}
					endTime := time.Now()
					deltaTime = endTime.Sub(startTime).Seconds()
				}
			}
		}()
	case Update:
		g.Update(msg, c)
	case *movement.Test:
		g.Test(msg, c)
	}
}

func (g *WorldActor) Test(msg *movement.Test, c actor.Context) {
	c.Logger().Info(strconv.Itoa(int(msg.Test)))
}

func (g *WorldActor) Update(msg Update, c actor.Context) {
	g.world.Update(msg.deltaTime)
	c.Respond(true)
}
