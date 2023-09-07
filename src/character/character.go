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
}

func (character Character) CanCollideWith(another string) bool {
	switch another {
	case "floor":
		return true

	case "floor-boundary":
		return character.isOverTheFloor

	case "wall":
		return true

	default:
		return character.Collisionable.CanCollideWith(another)
	}
}

func (character *Character) OnCollision(collision hnbPhysics.Collision) {
	switch collision.AnotherCollisionMask {
	case "floor-boundary":
		character.Fall()

	case "wall":
		xResolution := collision.CollisionResolverCalculator.CalculateXResolution()
		character.MoveX(xResolution)

	case "floor":
		anotherAABB := collision.AnotherAABB
		characterPrevAABB := character.GetPrevAABB()
		shouldResolveByX := character.IsMoving() && (characterPrevAABB.Right() <= anotherAABB.Left() || characterPrevAABB.Left() >= anotherAABB.Right())

		if !shouldResolveByX {
			characterAABB := character.GetAABB()
			yResolution := collision.CollisionResolverCalculator.CalculateYResolution()
			characterTopAfterYResolution := characterAABB.Top() + yResolution
			characterBottomAfterYResolution := characterAABB.Bottom() + yResolution

			shouldResolveByY := (character.IsJumping() && characterTopAfterYResolution >= anotherAABB.Bottom()) ||
				(character.IsFalling() && characterBottomAfterYResolution <= anotherAABB.Top())

			if shouldResolveByY {
				shouldNotResolve := (characterAABB.Right() == anotherAABB.Left() || characterAABB.Left() == anotherAABB.Right())
				if shouldNotResolve {
					return
				}
			} else {
				return
			}
		} else {
			characterAABB := character.GetAABB()
			shouldNotResolve := (characterAABB.Top() == anotherAABB.Bottom())
			if shouldNotResolve {
				return
			}
		}

		if shouldResolveByX {
			xResolution := collision.CollisionResolverCalculator.CalculateXResolution()
			character.MoveX(xResolution)
		} else {
			yResolution := collision.CollisionResolverCalculator.CalculateYResolution()
			character.MoveY(yResolution)

			if character.IsJumping() {
				character.Fall()
			} else if character.IsFalling() {
				character.StopFalling()
			}
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
		character.isOverTheFloor = false
	}
}

func (character *Character) StopFalling() {
	wasStateChanged := character.SetState("idle")

	if wasStateChanged {
		character.isOverTheFloor = true
		character.DynamicBody.StopFalling()
	}
}

func (character *Character) Jump() {
	wasStateChanged := character.SetState("jumping")

	if wasStateChanged {
		character.DynamicBody.Jump()
		character.isOverTheFloor = false
	}
}

func (character Character) GetLookingDirection() movement.Direction {
	return character.lookingDirection
}

func (character *Character) ResolveCollisionWithFloor(collision hnbPhysics.Collision) bool {
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
