package entities

import (
	"pigeon/ecs"
	"pigeon/game/components"
)

type Player struct {
	*ecs.BasicEntity
	*components.PositionComponent
	*components.VelocityComponent
}
