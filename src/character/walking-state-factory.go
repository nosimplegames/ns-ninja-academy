package character

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"simple-games.com/ninja/src/fsm"
)

type WalkingStateFactory struct {
	WalkingAnimation hnbCore.IAnimation
}

func (factory WalkingStateFactory) Create() *WalkingState {
	state := &WalkingState{}

	fsm.FSMStateFactory[ICharacter]{
		Name: "walking",
		AllowedStates: []string{
			"idle",
			"jumping",
		},
	}.Init(&state.FSMState)

	state.animation = factory.WalkingAnimation

	return state
}
