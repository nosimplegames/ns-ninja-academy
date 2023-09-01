package scenes

import (
	"github.com/nosimplegames/ns-framework/hnbPhysics"
	"simple-games.com/ninja/src/character"
	"simple-games.com/ninja/src/res"
)

type PlayerFactory struct {
}

func (factory PlayerFactory) Create() *Player {
	player := &Player{}
	player.SetSize(res.NinjaFrameSize)
	player.SetOrigin(res.NinjaFrameSize.By(0.5))
	animations := res.GetAnimations()

	character.CharacterFactory{
		FallingAnimation: animations.NinjaFallingAnimation,
		IdleAnimation:    animations.NinjaIdleAnimation,
		WalkingAnimation: animations.NinjaWalkingAnimation,
		JumpingAnimation: animations.NinjaJumpingAnimation,
	}.Init(&player.Character)

	hnbPhysics.AddCollisionable(player)

	return player
}
