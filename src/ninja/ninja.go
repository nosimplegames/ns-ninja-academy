package ninja

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
	"simple-games.com/ninja/src/character"
)

type Ninja struct {
	character.Character
}

func (ninja Ninja) ThrowShuriken() {
	shuriken := ShurikenFactory{
		MovingDirection: ninja.GetLookingDirection(),
	}.Create()
	shuriken.SetPosition(ninja.GetPosition())
	hnbCore.AddChildToRoot(shuriken)
	hnbPhysics.AddCollisionable(shuriken)
}
