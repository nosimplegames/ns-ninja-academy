package character

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"simple-games.com/ninja/src/fsm"
	"simple-games.com/ninja/src/movement"
)

type CharacterFactory struct {
	FallingAnimation hnbCore.IAnimation
	IdleAnimation    hnbCore.IAnimation
	WalkingAnimation hnbCore.IAnimation
	JumpingAnimation hnbCore.IAnimation
}

func (factory CharacterFactory) Create() ICharacter {
	character := &Character{}
	factory.Init(character)

	return character
}

func (factory CharacterFactory) Init(character *Character) {
	movement.DynamicBodyFactory{}.Init(&character.DynamicBody)

	fsm.FSMFactory[ICharacter]{
		Target:           character,
		InitialStateName: "falling",
		States: fsm.FSMStates[ICharacter]{
			FallingStateFactory{
				FallingAnimation: factory.FallingAnimation,
			}.Create(),
			IdleStateFactory{
				IdleAnimation: factory.IdleAnimation,
			}.Create(),
			WalkingStateFactory{
				WalkingAnimation: factory.WalkingAnimation,
			}.Create(),
			JumpingStateFactory{
				JumpingAnimation: factory.JumpingAnimation,
			}.Create(),
		},
	}.Init(&character.FSM)
}
