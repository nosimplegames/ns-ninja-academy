package fsm

type IFSMState[T any] interface {
	OnStart(T)
	OnEnd(T)
	GetName() string
	CanChangeToState(stateName string) bool
	UpdateFrame(T)
}
