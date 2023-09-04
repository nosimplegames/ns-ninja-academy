package obstacles

import (
	"github.com/nosimplegames/ns-framework/hnbEntities"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
)

type Log struct {
	hnbEntities.Sprite
}

func (log *Log) OnCollision(_ hnbPhysics.Collision) {
	log.Die()
}
