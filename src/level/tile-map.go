package level

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
	TileSet        hnbRender.Texture
	TileSize       hnbMath.Vector
	Data           [][]int
	EntityCreators MapEntityCreators
}

func (factory TileMapFactory) Create() *TileMap {
	tileMap := factory.createTileMap()
	factory.createEntities(tileMap)

	return tileMap
}

func (factory TileMapFactory) createTileMap() *TileMap {
	tileMap := &TileMap{}

	tileMapLayer := factory.createTileMapLayer()

	hnbEntities.TileMapFactory{
		TileSet:  factory.TileSet,
		TileSize: factory.TileSize,
		Layer:    tileMapLayer,
	}.Init(&tileMap.TileMap)

	return tileMap
}

func (factory TileMapFactory) createTileMapLayer() [][]int {
	tileMapLayer := [][]int{}

	for _, row := range factory.Data {
		tileMapRow := []int{}
		for _, entityCreatorId := range row {
			entityCreator, hasEntityCreator := factory.EntityCreators[entityCreatorId]

			if hasEntityCreator {
				tileMapRow = append(tileMapRow, entityCreator.TileId)
			} else {
				tileMapRow = append(tileMapRow, entityCreatorId)
			}
		}

		tileMapLayer = append(tileMapLayer, tileMapRow)
	}

	return tileMapLayer
}

func (factory TileMapFactory) createEntities(parent hnbCore.IEntity) {
	for rowIndex := 0; rowIndex < len(factory.Data); rowIndex++ {
		row := factory.Data[rowIndex]

		isTrackingEntity := false
		startingColumn := -1
		trackingEntityLength := 0
		var trackingEntityCreator *MapEntityCreator = nil

		createEntity := func() {
			size := hnbMath.Vector{
				X: float64(trackingEntityLength * 16),
				Y: 16,
			}
			position := hnbMath.Vector{
				X: float64(startingColumn*16) + size.X*0.5,
				Y: float64(rowIndex*16) + 8.0,
			}
			entity := trackingEntityCreator.CreateEntity(
				position,
				trackingEntityLength,
			)

			hnbCore.EntityAdder{
				Child:  entity,
				Parent: parent,
			}.Add()

			startingColumn = -1
			trackingEntityLength = 0
			trackingEntityCreator = nil
			isTrackingEntity = false
		}

		for columnIndex, entityCreatorId := range row {
			entityCreator, hasEntityCreator := factory.EntityCreators[entityCreatorId]

			if !hasEntityCreator {
				if isTrackingEntity {
					createEntity()
				}

				continue
			}

			shouldStartTrackingEntity := entityCreator.ShouldBeGrouped

			if isTrackingEntity {
				isSameEntityAsTracking := trackingEntityCreator.EntityName == entityCreator.EntityName

				if isSameEntityAsTracking {
					trackingEntityLength++
					continue
				} else {
					createEntity()

					shouldStartTrackingEntity = true
				}
			}

			startingColumn = columnIndex
			trackingEntityLength = 1

			if shouldStartTrackingEntity {
				isTrackingEntity = true
				trackingEntityCreator = &entityCreator
			} else {
				trackingEntityCreator = &entityCreator
				createEntity()
			}
		}

		if isTrackingEntity {
			createEntity()
		}
	}
}
