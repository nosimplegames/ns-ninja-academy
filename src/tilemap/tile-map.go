package tilemap

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbEntities"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type TileMap struct {
	hnbEntities.TileMap
}

type TileMapFactory struct {
	MapEntityCreators MapEntityCreators
	TileSet           hnbRender.Texture
	TileSize          hnbMath.Vector
	Layer             [][]int
}

func (factory TileMapFactory) Create() *TileMap {
	tileMap := factory.createTileMap()
	factory.createTileMapEntities(tileMap)

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

func (factory TileMapFactory) createTileMapEntities(tileMap *TileMap) []IMapEntity {
	mapEntities := []IMapEntity{}

	for rowIndex := 0; rowIndex < len(factory.Layer); rowIndex++ {
		row := factory.Layer[rowIndex]

		for columnIndex := 0; columnIndex < len(row); columnIndex++ {
			mapEntityKey := row[columnIndex]
			mapEntityCreator, doesMapEntityCreatorExist := factory.MapEntityCreators[mapEntityKey]

			if !doesMapEntityCreatorExist {
				continue
			}

			mapPosition := hnbMath.Vector{
				X: factory.TileSize.X * float64(columnIndex),
				Y: factory.TileSize.Y * float64(rowIndex),
			}.Add(factory.TileSize.By(0.5))
			mapEntity := mapEntityCreator.CreateMapEntity(mapPosition, factory.TileSize)

			hnbCore.EntityAdder{
				Parent: tileMap,
				Child:  mapEntity,
			}.Add()

			hnbPhysics.AddCollisionable(mapEntity)
		}
	}

	return mapEntities
}
