package game

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"math/rand"
	"pigeon/pb/movement"
	"strconv"
	"time"
)

type (
	GameActor struct {
		Server *actor.PID
	}

	Update struct {
		deltaTime float64
	}
)

func (g *GameActor) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case *actor.Started:
		go func() {
			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()
			deltaTime := 0.0
			for {
				startTime := time.Now()
				select {
				case <-ticker.C:
					c.RequestFuture(c.Self(), Update{deltaTime: deltaTime}, 5*time.Second).Wait()
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

func (g *GameActor) Test(msg *movement.Test, c actor.Context) {
	c.Logger().Info(strconv.Itoa(int(msg.Test)))
}

func (g *GameActor) Update(msg Update, c actor.Context) {
	fmt.Println(msg.deltaTime)
	num := rand.Intn(4)
	time.Sleep(time.Duration(num) * time.Second)
	c.Respond(1)
}
