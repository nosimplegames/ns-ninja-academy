package ninja

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbEntities"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"simple-games.com/ninja/src/movement"
	"simple-games.com/ninja/src/res"
)

type ShurikenFactory struct {
	MovingDirection movement.Direction
}

func (factory ShurikenFactory) Create() *Shuriken {
	shuriken := &Shuriken{}

	shuriken.MovingVector = factory.getMovementVector()
	shuriken.SetLifeSpan(5)

	hnbEntities.SpriteFactory{}.Init(&shuriken.Sprite)
	animation := res.GetAnimations().ShurikenAnimation.Copy(shuriken)
	shuriken.SetSize(res.ShurikenFrameSize)
	shuriken.SetOrigin(res.ShurikenFrameSize.By(0.5))
	shuriken.OnDie.AddCallback(animation.Stop)
	hnbCore.AddAnimation(animation)

	return shuriken
}

func (factory ShurikenFactory) getMovementVector() hnbMath.Vector {
	shurikenSpeed := 6.0
	shouldMoveToLeft := factory.MovingDirection == movement.DirectionLeft

	if shouldMoveToLeft {
		shurikenSpeed *= -1.0
	}

	return hnbMath.Vector{
		X: shurikenSpeed,
		Y: 0,
	}
}
