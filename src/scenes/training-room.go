package scenes

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbEntities"
	"github.com/nosimplegames/ns-framework/hnbMath"
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

	centerOfTheMap := res.GameSize.By(0.5)
	floorPosition := res.GameSize.Y - 16

	tileMap := tilemap.TileMapFactory{
		TileSet:  res.GetTextures().Tileset,
		TileSize: res.TileSize,
		Data: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 3, 0, 5, 0, 0, 0, 4},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 4},
			{0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 5, 4},
			{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 1, 1, 0, 4},
			{4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 4},
			{4, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		},
		EntityCreators: tilemap.GetDefaultEntityCreators(),
	}.Create()
	tileMap.SetPosition(centerOfTheMap)

	player := PlayerFactory{}.Create()
	player.SetPosition(hnbMath.Vector{
		X: 16 * 2.5,
		Y: floorPosition - 8,
	})

	text := hnbEntities.TextFactory{
		Text:     "Destroy all logs!",
		FontFace: res.GetFonts().NormalText,
	}.Create()
	text.SetPosition(hnbMath.Vector{
		X: centerOfTheMap.X,
		Y: 16,
	})

	// log := obstacles.LogFactory{}.Create()
	// log.SetPosition(hnbMath.Vector{
	// 	X: res.GameSize.X * 0.8,
	// 	Y: floorPosition - 8,
	// })
	// hnbPhysics.AddCollisionable(log)

	// trap := obstacles.FloorTrapFactory{}.Create()
	// trap.SetPosition(hnbMath.Vector{
	// 	X: res.GameSize.X * 0.4,
	// 	Y: floorPosition - 2.5,
	// })
	// hnbPhysics.AddCollisionable(trap)

	hnbCore.EntityAdder{
		Children: hnbCore.EntityChildren{
			tileMap,
			player,
			text,
			// log,
			// trap,
		},
		Parent: scene,
	}.Add()

	return scene
}
