package level

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbMath"
)

type Floor struct {
	hnbCore.Entity
}

type FloorFactory struct {
	FloorLength int
	Position    hnbMath.Vector
}

func (factory FloorFactory) Create() *Floor {
	needToCreateFloor := factory.FloorLength > 0

	if !needToCreateFloor {
		return nil
	}

	floor := &Floor{}

	floor.SetCanCollide(true)
	floor.SetCollisionMask("floor")
	floor.SetCollisioningMasks([]string{
		"character",
		"shuriken",
	})

	size := hnbMath.Vector{
		X: float64(factory.FloorLength) * 16,
		Y: 16,
	}
	floor.SetSize(size)
	floor.SetOrigin(size.By(0.5))
	floor.SetPosition(factory.Position)

	// beginningBoundary, endingBoundary := factory.createBoundaries(size)
	// hnbCore.EntityAdder{
	// 	Children: hnbCore.EntityChildren{beginningBoundary, endingBoundary},
	// 	Parent:   floor,
	// }.Add()

	return floor
}

// func (factory FloorFactory) createBoundaries(floorSize hnbMath.Vector) (hnbCore.IEntity, hnbCore.IEntity) {
// 	beginningPosition := hnbMath.Vector{
// 		X: -floorSize.X*0.5 - 8,
// 		// Y: -16,
// 	}
// 	endingPosition := hnbMath.Vector{
// 		X: floorSize.X*0.5 + 8,
// 		// Y: -16,
// 	}
// 	beginning := factory.createBoundary(beginningPosition)
// 	ending := factory.createBoundary(endingPosition)

// 	return beginning, ending
// }

// func (factory FloorFactory) createBoundary(position hnbMath.Vector) hnbCore.IEntity {
// 	boundary := &hnbCore.Entity{}

// 	boundarySize := hnbMath.Vector{
// 		X: 16,
// 		Y: 16,
// 	}
// 	boundary.SetCanCollide(true)
// 	boundary.SetCollisionMask("floor-boundary")
// 	boundary.SetCollisioningMasks([]string{
// 		"character",
// 	})
// 	boundary.SetSize(boundarySize)
// 	boundary.SetOrigin(boundarySize.By(0.5))
// 	boundary.SetPosition(position)

// 	return boundary
// }
