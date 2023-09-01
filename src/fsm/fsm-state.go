package fsm

type FSMState[T any] struct {
	name          string
	allowedStates []string
}

func (state FSMState[T]) GetName() string {
	return state.name
}

func (state FSMState[T]) CanChangeToState(stateName string) bool {
	for _, allowedState := range state.allowedStates {
		foundState := allowedState == stateName

		if foundState {
			return true
		}
	}

	return false
}
