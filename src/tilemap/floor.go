package tilemap

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
)

type Floor struct {
	hnbCore.Entity
}

func (floor Floor) CanCollide() bool {
	return true
}

func (floor Floor) CanCollideWith(_ string) bool {
	return true
}

func (floor Floor) GetCollisionMask() string {
	return "map-floor"
}

func (floor Floor) OnCollision(collision hnbPhysics.Collision) {
}

type FloorFactory struct {
}

func (factory FloorFactory) Create() *Floor {
	floor := &Floor{}

	return floor
}

func (factory FloorFactory) CreateMapEntity(
	position hnbMath.Vector,
	tileSize hnbMath.Vector,
) IMapEntity {
	mapLimit := factory.Create()

	mapLimit.SetSize(tileSize)
	mapLimit.SetOrigin(tileSize.By(0.5))
	mapLimit.SetPosition(position)

	return mapLimit
}
