package level

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbMath"
)

type WallFactory struct {
}

func (factory WallFactory) Create() hnbCore.IEntity {
	wall := &hnbCore.Entity{}

	wall.SetCanCollide(true)
	wall.SetCollisionMask("wall")
	wall.SetCollisioningMasks([]string{
		"character",
	})

	size := hnbMath.Vector{
		X: 16,
		Y: 16,
	}
	wall.SetSize(size)
	wall.SetOrigin(size.By(0.5))

	return wall
}
