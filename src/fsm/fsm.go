package fsm

type FSM[T any] struct {
	target       T
	currentState IFSMState[T]
	states       FSMStates[T]
}

func (fsm *FSM[T]) SetState(stateName string) bool {
	canChangeToState := true
	hasCurrentState := fsm.HasCurrentState()

	if hasCurrentState {
		canChangeToState = fsm.currentState.CanChangeToState(stateName)
	}

	if !canChangeToState {
		return false
	}

	if hasCurrentState {
		fsm.currentState.OnEnd(fsm.target)
	}

	newState, doesNewStateExist := fsm.states.GetByName(stateName)
	fsm.currentState = newState

	if doesNewStateExist {
		fsm.currentState.OnStart(fsm.target)
		return true
	}

	return false
}

func (fsm FSM[T]) UpdateFrame() {
	if !fsm.HasCurrentState() {
		return
	}

	fsm.currentState.UpdateFrame(fsm.target)
}

func (fsm FSM[T]) HasCurrentState() bool {
	return fsm.currentState != nil
}

func (fsm FSM[T]) IsCurrentState(stateName string) bool {
	return fsm.currentState != nil && fsm.currentState.GetName() == stateName
}

func (fsm FSM[T]) GetCurrentStateName() (string, bool) {
	if !fsm.HasCurrentState() {
		return "", false
	}

	return fsm.currentState.GetName(), true
}
