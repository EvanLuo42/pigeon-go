package main

import (
	"github.com/asynkron/protoactor-go/actor"
	"pigeon/network"
)

func main() {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(func() actor.Actor {
		return &network.ServerActor{Port: 8080}
	})
	system.Root.Spawn(props)
	select {}
}
