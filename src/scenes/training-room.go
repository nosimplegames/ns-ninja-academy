package scenes

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"simple-games.com/ninja/src/res"
	"simple-games.com/ninja/src/tilemap"
)

type TrainingRoom struct {
	hnbCore.Scene
}

type TrainingRoomFactory struct {
}

func (factory TrainingRoomFactory) Create() hnbCore.IScene {
	scene := &TrainingRoom{}

	tileMap := tilemap.TileMapFactory{
		MapEntityCreators: tilemap.GetDefaultMapEntityCreators(),
		TileSet:           res.GetTextures().Tileset,
		TileSize:          res.TileSize,
		Layer: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}.Create()

	player := PlayerFactory{}.Create()
	player.SetPosition(res.GameSize.By(0.5))
	tileMap.SetPosition(res.GameSize.By(0.5))

	hnbCore.EntityAdder{
		Children: hnbCore.EntityChildren{
			tileMap,
			player,
		},
		Parent: scene,
	}.Add()

	return scene
}
