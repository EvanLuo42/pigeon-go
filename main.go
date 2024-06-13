package main

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"pigeon/network"
	"time"
)

func main() {
	system := actor.NewActorSystem(actor.WithLoggerFactory(coloredConsoleLogging))
	props := actor.PropsFromProducer(func() actor.Actor {
		return &network.ServerActor{Port: 8080}
	})
	system.Root.Spawn(props)
	select {}
}

func coloredConsoleLogging(system *actor.ActorSystem) *slog.Logger {
	return slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.RFC3339,
		AddSource:  true,
	})).With("lib", "Proto.Actor").
		With("system", system.ID)
}
