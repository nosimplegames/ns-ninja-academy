package ninja

import (
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
	"simple-games.com/ninja/src/character"
	"simple-games.com/ninja/src/res"
)

type NinjaFactory struct {
}

func (factory NinjaFactory) Create() *Ninja {
	ninja := &Ninja{}
	factory.Init(ninja)

	return ninja
}

func (factory NinjaFactory) Init(ninja *Ninja) {
	animations := res.GetAnimations()

	character.CharacterFactory{
		FallingAnimation: animations.NinjaFallingAnimation,
		IdleAnimation:    animations.NinjaIdleAnimation,
		WalkingAnimation: animations.NinjaWalkingAnimation,
		JumpingAnimation: animations.NinjaJumpingAnimation,
	}.Init(&ninja.Character)

	ninja.SetCollisioningMasks([]string{
		"map-floor",
		"floor-trap",
	})
	ninja.SetSize(hnbMath.Vector{
		X: res.NinjaFrameSize.X * 0.6,
		Y: res.NinjaFrameSize.Y,
	})
	ninja.SetOrigin(res.NinjaFrameSize.By(0.5))

	hnbPhysics.AddCollisionable(ninja)
}
