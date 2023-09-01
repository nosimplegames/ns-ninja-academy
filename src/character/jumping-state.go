package character

import "github.com/nosimplegames/ns-framework/hnbEvents"

type JumpingState struct {
	CharacterState
}

func (state *JumpingState) OnStart(character ICharacter) {
	state.CharacterState.OnStart(character)
	character.AddEventListener(hnbEvents.EventListener{
		EventType: "jump-stop",
		Callback: func(_ hnbEvents.Event) {
			character.Fall()
		},
	})
}
