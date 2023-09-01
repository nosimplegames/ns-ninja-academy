package character

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"simple-games.com/ninja/src/fsm"
)

type FallingStateFactory struct {
	FallingAnimation hnbCore.IAnimation
}

func (factory FallingStateFactory) Create() *FallingState {
	state := &FallingState{}

	fsm.FSMStateFactory[ICharacter]{
		Name: "falling",
		AllowedStates: []string{
			"idle",
		},
	}.Init(&state.FSMState)

	state.animation = factory.FallingAnimation

	return state
}
