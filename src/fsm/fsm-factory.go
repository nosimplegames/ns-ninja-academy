package fsm

type FSMFactory[T any] struct {
	Target           T
	States           FSMStates[T]
	InitialStateName string
}

func (factory FSMFactory[T]) Create() FSM[T] {
	fsm := FSM[T]{}
	factory.Init(&fsm)

	return fsm
}

func (factory FSMFactory[T]) Init(fsm *FSM[T]) {
	fsm.target = factory.Target
	fsm.states = factory.States
	fsm.SetState(factory.InitialStateName)
}
