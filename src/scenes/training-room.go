package scenes

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
	"simple-games.com/ninja/src/obstacles"
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
	tileMap.SetPosition(centerOfTheMap)

	player := PlayerFactory{}.Create()
	player.SetPosition(hnbMath.Vector{
		X: res.GameSize.X * 0.1,
		Y: floorPosition - 8,
	})

	log := obstacles.LogFactory{}.Create()
	log.SetPosition(hnbMath.Vector{
		X: res.GameSize.X * 0.8,
		Y: floorPosition - 8,
	})
	hnbPhysics.AddCollisionable(log)

	trap := obstacles.FloorTrapFactory{}.Create()
	trap.SetPosition(hnbMath.Vector{
		X: res.GameSize.X * 0.4,
		Y: floorPosition - 2.5,
	})
	hnbPhysics.AddCollisionable(trap)

	hnbCore.EntityAdder{
		Children: hnbCore.EntityChildren{
			tileMap,
			player,
			log,
			trap,
		},
		Parent: scene,
	}.Add()

	return scene
}
