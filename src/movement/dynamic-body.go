package movement

import (
	"github.com/nosimplegames/ns-framework/hnbEvents"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbUtils"
)

type DynamicBodyState int

const (
	DynamicBodyFalling DynamicBodyState = iota
	DynamicBodyOnFloor
	DynamicBodyJumping
)

type DynamicBody struct {
	Speed     hnbMath.Vector
	State     DynamicBodyState
	Direction Direction

	hnbEvents.EventTarget
}

func (body *DynamicBody) UpdateFrame() {
	switch body.State {
	case DynamicBodyFalling:
		body.Speed.Y += 9.8 * hnbUtils.FrameTime
	case DynamicBodyJumping:
		body.Speed.Y += 9.8 * hnbUtils.FrameTime

		if body.HasStopJumping() {
			body.Speed.Y = 0
			body.Fire("jump-stop")
		}
	}

	switch body.Direction {
	case NoDirection:
		body.Speed.X = 0
	case DirectionLeft:
		body.Speed.X = -69.9 * hnbUtils.FrameTime
	case DirectionRight:
		body.Speed.X = 69.9 * hnbUtils.FrameTime
	}
}

func (body *DynamicBody) StopFalling() {
	body.State = DynamicBodyOnFloor
	body.Speed.Y = 0
}

func (body *DynamicBody) Fall() {
	body.State = DynamicBodyFalling
	body.Speed.Y = 0
}

func (body DynamicBody) IsFalling() bool {
	return body.State == DynamicBodyFalling
}

func (body *DynamicBody) Jump() {
	body.State = DynamicBodyJumping
	body.Speed.Y -= 4
}

func (body DynamicBody) HasStopJumping() bool {
	isJumping := body.State == DynamicBodyJumping

	if !isJumping {
		return true
	}

	isJumping = body.Speed.Y >= 0.0

	return isJumping
}

func (body *DynamicBody) SetMovingDirection(direction Direction) {
	body.Direction = direction
}
