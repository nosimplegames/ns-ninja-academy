package level

import (
	"simple-games.com/ninja/src/ninja"
)

type PlayerFactory struct {
}

func (factory PlayerFactory) Create() *Player {
	player := &Player{}
	ninja.NinjaFactory{}.Init(&player.Ninja)

	return player
}
