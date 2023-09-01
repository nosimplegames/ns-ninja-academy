package tilemap

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
)

type IMapEntity interface {
	hnbCore.IEntity
	hnbPhysics.ICollisionable
}
