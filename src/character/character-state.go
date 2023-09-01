package character

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"simple-games.com/ninja/src/fsm"
)

type CharacterState struct {
	fsm.FSMState[ICharacter]
	animation hnbCore.IAnimation
}

func (state *CharacterState) OnStart(character ICharacter) {
	animation := state.animation.Copy(character)
	hnbCore.AddAnimation(animation)
	state.animation = animation
	state.animation.Apply()
}

func (state *CharacterState) OnEnd(character ICharacter) {
	state.animation.Stop()
}

func (state CharacterState) UpdateFrame(_ ICharacter) {
}
