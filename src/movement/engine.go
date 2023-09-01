package movement

import (
	"math"

	"github.com/nosimplegames/ns-framework/hnbUtils"
)

type Engine struct {
	maxSpeed          float64
	minSpeed          float64
	currentSpeed      float64
	isAccelerating    bool
	accelerationForce float64
}

func (engine *Engine) FrameUpdateFrame() {
	if !engine.isAccelerating {
		return
	}

	engine.currentSpeed += engine.accelerationForce * hnbUtils.FrameTime
	engine.currentSpeed = math.Min(engine.currentSpeed, engine.maxSpeed)
	engine.currentSpeed = math.Max(engine.currentSpeed, engine.minSpeed)

	engine.setIsAccelerating(false)
}

func (engine *Engine) Accelerate() {
	engine.setIsAccelerating(true)
}

func (engine *Engine) setIsAccelerating(isAccelerating bool) {
	engine.isAccelerating = isAccelerating
}

func (engine Engine) GetCurrentSpeed() float64 {
	return engine.currentSpeed
}

func (engine *Engine) Reset() {
	engine.currentSpeed = engine.minSpeed
	engine.isAccelerating = false
}
