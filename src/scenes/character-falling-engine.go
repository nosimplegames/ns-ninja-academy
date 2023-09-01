package scenes

import "simple-games.com/ninja/src/movement"

func GetCharacterFallingEngine() movement.Engine {
	return movement.EngineFactory{
		MaxSpeed:          3,
		MinSpeed:          1,
		AccelerationForce: 120,
	}.Create()
}
