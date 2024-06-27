package Data

import (
	"math/rand"
	"sync"
)

type EntityType string
type MovingMode string
type Direction string

const (
	Human  EntityType = "_human"
	Zombie EntityType = "_Zombie"
)

const (
	Slow  MovingMode = "_slow"
	Panic MovingMode = "_panic"
)

const (
	NORTH Direction = "_north"
	SOUTH Direction = "_south"
	EAST  Direction = "_east"
	WEST  Direction = "_west"
)

type Entity struct {
	world         *World
	entityType    EntityType
	x, y          int
	activityLevel int
	direction     Direction
	mutex         sync.Mutex
}

func newEntity(world *World) *Entity {
	entity := &Entity{}

	var isFound bool
	for !isFound {
		entity.x = rand.Intn(world.width)
		entity.y = rand.Intn(world.height)

		isFound = world.maze[entity.x][entity.y]
	}

	return entity
}

func (e *Entity) Interact(other *Entity) {
	if other.entityType == Zombie || e.entityType == Zombie {
		e.entityType = Zombie
		other.entityType = Zombie
	}
}

func (e *Entity) Draw() (x, y int) {
	return e.x, e.y
}

func movingHuman(e *Entity) {
	//directions := []Direction{NORTH, SOUTH, EAST, WEST}

	//switch this.direction {
	//case this.world.DIRECTIONS.NORTH:
	//	dy--
	//	break
	//case this.world.DIRECTIONS.EAST:
	//	dx++
	//	break
	//case this.world.DIRECTIONS.SOUTH:
	//	dy++
	//	break
	//case this.world.DIRECTIONS.WEST:
	//	dx--
	//	break
	//}
}
