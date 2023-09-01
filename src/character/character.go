package character

import (
	"github.com/nosimplegames/ns-framework/hnbEntities"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
	"simple-games.com/ninja/src/fsm"
	"simple-games.com/ninja/src/movement"
)

type Character struct {
	hnbEntities.Sprite

	fsm.FSM[ICharacter]
	movement.DynamicBody

	lookingDirection movement.Direction
}

func (character *Character) UpdateFrame() {
	character.FSM.UpdateFrame()
	character.DynamicBody.UpdateFrame()
	character.Move(character.Speed)
}

func (character Character) CanCollide() bool {
	return true
}

func (character Character) CanCollideWith(collisionMask string) bool {
	switch collisionMask {
	case "map-floor":
		return character.IsFalling()
	default:
		return true
	}
}

func (character Character) GetCollisionMask() string {
	return "character"
}

func (character *Character) OnCollision(collision hnbPhysics.Collision) {
	characterYResolution := collision.CollisionResolverCalculator.CalculateYResolution()
	character.Move(hnbMath.Vector{
		Y: characterYResolution,
	})
	character.StopFalling()
	character.SetState("idle")
}

func (character *Character) SetMovingDirection(direction movement.Direction) {
	character.DynamicBody.SetMovingDirection(direction)

	switch direction {
	case movement.NoDirection:
		if !character.IsFalling() {
			character.SetState("idle")
		}

	case movement.DirectionRight:
		character.SetScale(hnbMath.Vector{X: 1, Y: 1})
		character.SetState("walking")
		character.lookingDirection = direction

	case movement.DirectionLeft:
		character.SetScale(hnbMath.Vector{X: -1, Y: 1})
		character.SetState("walking")
		character.lookingDirection = direction
	}
}

func (character *Character) Fall() {
	wasStateChanged := character.SetState("falling")

	if wasStateChanged {
		character.DynamicBody.Fall()
	}
}

func (character *Character) Jump() {
	wasStateChanged := character.SetState("jumping")

	if wasStateChanged {
		character.DynamicBody.Jump()
	}
}

func (character Character) GetLookingDirection() movement.Direction {
	return character.lookingDirection
}
