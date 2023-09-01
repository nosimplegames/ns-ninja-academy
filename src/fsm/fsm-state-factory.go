package fsm

type FSMStateFactory[T any] struct {
	Name          string
	AllowedStates []string
}

func (factory FSMStateFactory[T]) Create() FSMState[T] {
	state := FSMState[T]{}
	factory.Init(&state)

	return state
}

func (factory FSMStateFactory[T]) Init(state *FSMState[T]) {
	state.name = factory.Name
	state.allowedStates = factory.AllowedStates
}
