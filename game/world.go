package game

import (
	"fmt"
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
			freq := time.Second
			timeout := 5 * time.Second
			deltaTime := 0.0
			timer := time.NewTimer(freq)
			for {
				startTime := time.Now()
				err := c.RequestFuture(c.Self(), Update{deltaTime: deltaTime}, timeout).Wait()
				if err != nil {
					c.Logger().Error(err.Error())
				}
				endTime := time.Now()
				deltaTime = endTime.Sub(startTime).Seconds()
				if deltaTime > freq.Seconds() {
					continue
				}
				<-timer.C
				deltaTime = time.Second.Seconds()
				timer.Reset(time.Second)
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
	fmt.Println(msg.deltaTime)
	g.world.Update(msg.deltaTime)
	c.Respond(true)
}
