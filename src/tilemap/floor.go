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
	FloorLength    int
	Row            int
	StartingColumn int
}

func (factory FloorFactory) Create() *Floor {
	needToCreateFloor := factory.FloorLength > 0

	if !needToCreateFloor {
		return nil
	}

	floor := &Floor{}

	floor.SetCanCollide(true)
	floor.SetCollisionMask("map-floor")
	floor.SetCollisioningMasks([]string{
		"character",
	})

	size := hnbMath.Vector{
		X: float64(factory.FloorLength) * 16,
		Y: 16,
	}
	floor.SetSize(size)
	floor.SetOrigin(size.By(0.5))
	position := hnbMath.Vector{
		X: float64(factory.StartingColumn*16) + size.X*0.5,
		Y: float64(factory.Row*16) + size.Y*0.5,
	}
	floor.SetPosition(position)

	hnbPhysics.AddCollisionable(floor)

	return floor
}
