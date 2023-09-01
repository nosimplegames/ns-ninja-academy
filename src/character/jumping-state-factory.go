package character

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"simple-games.com/ninja/src/fsm"
)

type JumpingStateFactory struct {
	JumpingAnimation hnbCore.IAnimation
}

func (factory JumpingStateFactory) Create() *JumpingState {
	state := &JumpingState{}

	fsm.FSMStateFactory[ICharacter]{
		Name: "jumping",
		AllowedStates: []string{
			"falling",
		},
	}.Init(&state.FSMState)

	state.animation = factory.JumpingAnimation

	return state
}
