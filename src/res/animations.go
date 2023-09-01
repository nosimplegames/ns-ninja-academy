package res

import (
	"github.com/nosimplegames/ns-framework/hnbAnimations"
	"github.com/nosimplegames/ns-framework/hnbCore"
)

type Animations struct {
	NinjaIdleAnimation    hnbCore.IAnimation
	NinjaFallingAnimation hnbCore.IAnimation
	NinjaWalkingAnimation hnbCore.IAnimation
	NinjaJumpingAnimation hnbCore.IAnimation
}

var animations *Animations = nil

func GetAnimations() *Animations {
	areTexturesInited := animations != nil

	if !areTexturesInited {
		animations = loadAnimations()
	}

	return animations
}

func loadAnimations() *Animations {
	animations := &Animations{}
	textures := GetTextures()

	animations.NinjaIdleAnimation = hnbAnimations.SpriteAnimationFactory{
		Texture:          textures.NinjaIdleSpriteSheet,
		FrameSize:        NinjaFrameSize,
		FrameDuration:    0.05,
		LoopCount:        hnbAnimations.AnimationInfiniteLoop,
		TimeBetweenLoops: 3,
	}.Create(nil)
	animations.NinjaFallingAnimation = hnbAnimations.SpriteAnimationFactory{
		Texture:       textures.NinjaFallingTexture,
		FrameSize:     NinjaFrameSize,
		FrameDuration: 1,
		LoopCount:     hnbAnimations.AnimationInfiniteLoop,
	}.Create(nil)
	animations.NinjaWalkingAnimation = hnbAnimations.SpriteAnimationFactory{
		Texture:       textures.NinjaWalkingSpriteSheet,
		FrameSize:     NinjaFrameSize,
		FrameDuration: 0.075,
		LoopCount:     hnbAnimations.AnimationInfiniteLoop,
	}.Create(nil)
	animations.NinjaJumpingAnimation = animations.NinjaFallingAnimation

	return animations
}
