package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"simple-games.com/ninja/src/character"
	"simple-games.com/ninja/src/movement"
)

type Player struct {
	character.Character
}

func (player *Player) HandleInput() {
	isMoving := false

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		player.SetMovingDirection(movement.DirectionRight)
		isMoving = true
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		player.SetMovingDirection(movement.DirectionLeft)
		isMoving = true
	}

	if !isMoving {
		player.SetMovingDirection(movement.NoDirection)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		player.Jump()
	}
}
