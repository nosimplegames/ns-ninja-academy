package tilemap

import "github.com/nosimplegames/ns-framework/hnbMath"

type MapEntityCreator interface {
	CreateMapEntity(position hnbMath.Vector, tileSize hnbMath.Vector) IMapEntity
}

type MapEntityCreators = map[int]MapEntityCreator
