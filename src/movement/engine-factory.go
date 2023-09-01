package movement

type EngineFactory struct {
	MaxSpeed          float64
	MinSpeed          float64
	AccelerationForce float64
}

func (factory EngineFactory) Init(engine *Engine) {
	engine.accelerationForce = factory.AccelerationForce
	engine.minSpeed = factory.MinSpeed
	engine.maxSpeed = factory.MaxSpeed
	engine.currentSpeed = factory.MinSpeed
}

func (factory EngineFactory) Create() Engine {
	engine := Engine{}

	factory.Init(&engine)

	return engine
}
