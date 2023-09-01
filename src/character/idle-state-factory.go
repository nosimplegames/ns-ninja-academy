package character

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"simple-games.com/ninja/src/fsm"
)

type IdleStateFactory struct {
	IdleAnimation hnbCore.IAnimation
}

func (factory IdleStateFactory) Create() *IdleState {
	state := &IdleState{}

	fsm.FSMStateFactory[ICharacter]{
		Name: "idle",
		AllowedStates: []string{
			"walking",
			"jumping",
		},
	}.Init(&state.FSMState)

	state.animation = factory.IdleAnimation

	return state
}
