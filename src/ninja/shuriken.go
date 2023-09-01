package ninja

import (
	"github.com/nosimplegames/ns-framework/hnbEntities"
	"github.com/nosimplegames/ns-framework/hnbMath"
)

type Shuriken struct {
	hnbEntities.Sprite

	MovingVector hnbMath.Vector
}

func (shuriken *Shuriken) UpdateFrame() {
	shuriken.Sprite.UpdateFrame()

	shuriken.Move(shuriken.MovingVector)
}
