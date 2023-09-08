package level

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbMath"
)

type MapEntityCreator struct {
	// CreateMapEntity(position hnbMath.Vector, tileSize hnbMath.Vector) IMapEntity
	TileId          int
	EntityName      string
	ShouldBeGrouped bool
	CreateEntity    func(position hnbMath.Vector, length int) hnbCore.IEntity
}

type MapEntityCreators = map[int]MapEntityCreator
