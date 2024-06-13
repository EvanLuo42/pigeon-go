package systems

import (
	"pigeon/ecs"
	"pigeon/game/components"
	"pigeon/game/entities"
)

type MovementSystem struct {
	entities []entities.Player
}

func (m *MovementSystem) Add(entity *ecs.BasicEntity, position *components.PositionComponent,
	velocity *components.VelocityComponent) {
	m.entities = append(m.entities, entities.Player{
		BasicEntity:       entity,
		PositionComponent: position,
		VelocityComponent: velocity})
}

func (MovementSystem) Update(dt float64) {

}
