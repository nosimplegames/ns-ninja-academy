package fsm

type FSMStates[T any] []IFSMState[T]

func (states FSMStates[T]) GetByName(stateName string) (IFSMState[T], bool) {
	for _, state := range states {
		foundState := state.GetName() == stateName

		if foundState {
			return state, true
		}
	}

	return nil, false
}
