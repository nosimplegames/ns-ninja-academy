package ninja

import (
	"github.com/nosimplegames/ns-framework/hnbEntities"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
)

type Shuriken struct {
	hnbEntities.Sprite

	MovingVector hnbMath.Vector
}

func (shuriken *Shuriken) UpdateFrame() {
	shuriken.Sprite.UpdateFrame()

	shuriken.Move(shuriken.MovingVector)
}

func (shuriken *Shuriken) OnCollision(_ hnbPhysics.Collision) {
	shuriken.Die()
}
