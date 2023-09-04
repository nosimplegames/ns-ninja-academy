package tilemap

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
)

type Floor struct {
	hnbCore.Entity
}

type FloorFactory struct {
}

func (factory FloorFactory) Create() *Floor {
	floor := &Floor{}

	hnbCore.EntityFactory{
		CollisionableFactory: hnbPhysics.CollisionableFactory{
			CanCollide:    true,
			CollisionMask: "map-floor",
			CollisioningMasks: []string{
				"character",
			},
		},
	}.Init(&floor.Entity)

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
