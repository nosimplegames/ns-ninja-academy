package character

import (
	"github.com/nosimplegames/ns-framework/hnbAnimations"
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbEvents"
)

type ICharacter interface {
	hnbEvents.IEventTarget
	hnbAnimations.ISprite
	hnbCore.IEntity

	Fall()
	StopFalling()
	Jump()
}
