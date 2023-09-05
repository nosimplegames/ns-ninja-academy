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
	isOverTheFloor   bool
}

func (character *Character) UpdateFrame() {
	character.FSM.UpdateFrame()
	character.DynamicBody.UpdateFrame()
	character.Move(character.Speed)

	mustFall := !character.isOverTheFloor && !character.IsJumping()
	if mustFall {
		character.Fall()
	} else {
		character.isOverTheFloor = false
	}
}

func (character *Character) OnCollision(collision hnbPhysics.Collision) {
	stateName, hasState := character.GetCurrentStateName()

	if !hasState {
		return
	}

	isCollidingWithFloor := collision.AnotherCollisionMask == "map-floor"

	switch stateName {
	case "falling":
		if isCollidingWithFloor {
			wasResolvedByY := character.ResolveCollisionWithFloor(collision)

			if wasResolvedByY {
				character.StopFalling()
				character.SetState("idle")
			}
		}

	case "jumping":
		if isCollidingWithFloor {
			wasResolvedByY := character.ResolveCollisionWithFloor(collision)

			if wasResolvedByY {
				character.Fall()
			}
		}
		characterYResolution := collision.CollisionResolverCalculator.CalculateYResolution()
		character.Move(hnbMath.Vector{
			Y: characterYResolution,
		})

	case "walking":
		isCollidingWithTheFloor := collision.AnotherCollisionMask == "map-floor"

		if isCollidingWithTheFloor {
			character.isOverTheFloor = true
		}
	}
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

func (character *Character) StopFalling() {
	character.isOverTheFloor = true
	character.DynamicBody.StopFalling()
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

func (character *Character) ResolveCollisionWithFloor(collision hnbPhysics.Collision) bool {
	// stateName, _ := character.GetCurrentStateName()
	shouldResolveByX := character.IsMoving() && !character.isOverTheFloor

	if shouldResolveByX {
		xResolution := collision.CollisionResolverCalculator.CalculateXResolution()
		yResolution := collision.CollisionResolverCalculator.CalculateYResolution()
		mustResolveByX := xResolution < yResolution

		if mustResolveByX {
			character.Move(hnbMath.Vector{
				X: xResolution,
			})
			return false
		}
	}

	character.Move(hnbMath.Vector{
		Y: collision.CollisionResolverCalculator.CalculateYResolution(),
	})
	return true
}
