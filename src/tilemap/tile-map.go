package tilemap

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbEntities"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type TileMap struct {
	hnbEntities.TileMap
}

type TileMapFactory struct {
	TileSet  hnbRender.Texture
	TileSize hnbMath.Vector
	Layer    [][]int

	FloorLayer [][]int
}

func (factory TileMapFactory) Create() *TileMap {
	tileMap := factory.createTileMap()
	factory.createCollisionables(tileMap)

	return tileMap
}

func (factory TileMapFactory) createTileMap() *TileMap {
	tileMap := &TileMap{}

	hnbEntities.TileMapFactory{
		TileSet:  factory.TileSet,
		TileSize: factory.TileSize,
		Layer:    factory.Layer,
	}.Init(&tileMap.TileMap)

	return tileMap
}

func (factory TileMapFactory) createCollisionables(parent hnbCore.IEntity) {
	for rowIndex := 0; rowIndex < len(factory.FloorLayer); rowIndex++ {
		row := factory.Layer[rowIndex]

		floorLength := 0
		startingColumn := -1

		for columnIndex := 0; columnIndex < len(row); columnIndex++ {
			cell := row[columnIndex]
			isFloor := cell != 0

			if isFloor {
				floorLength++

				mustSetStartingColumn := startingColumn == -1

				if mustSetStartingColumn {
					startingColumn = columnIndex
				}
			} else {
				floor := FloorFactory{
					FloorLength:    floorLength,
					Row:            rowIndex,
					StartingColumn: startingColumn,
				}.Create()
				hnbCore.EntityAdder{
					Parent: parent,
					Child:  floor,
				}.Add()

				floorLength = 0
				startingColumn = -1
			}
		}

		floor := FloorFactory{
			FloorLength:    floorLength,
			Row:            rowIndex,
			StartingColumn: startingColumn,
		}.Create()
		hnbCore.EntityAdder{
			Parent: parent,
			Child:  floor,
		}.Add()
	}
}
