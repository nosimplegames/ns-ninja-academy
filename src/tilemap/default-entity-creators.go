package tilemap

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"simple-games.com/ninja/src/obstacles"
)

func GetDefaultEntityCreators() MapEntityCreators {
	creators := MapEntityCreators{
		1: MapEntityCreator{
			TileId:          1,
			EntityName:      "floor",
			ShouldBeGrouped: true,
			CreateEntity: func(position hnbMath.Vector, length int) hnbCore.IEntity {
				floor := FloorFactory{
					FloorLength: length,
					Position:    position,
				}.Create()

				return floor
			},
		},
		2: MapEntityCreator{
			TileId:     0,
			EntityName: "log",
			CreateEntity: func(position hnbMath.Vector, _ int) hnbCore.IEntity {
				log := obstacles.LogFactory{}.Create()
				log.SetPosition(position)

				return log
			},
		},
		3: MapEntityCreator{
			TileId:     0,
			EntityName: "floor-trap",
			CreateEntity: func(position hnbMath.Vector, _ int) hnbCore.IEntity {
				trap := obstacles.FloorTrapFactory{}.Create()
				trapPosition := hnbMath.Vector{
					X: position.X,
					Y: position.Y + 8.0 - trap.GetSize().Y*0.5,
				}
				trap.SetPosition(trapPosition)

				return trap
			},
		},
		4: MapEntityCreator{
			TileId:     0,
			EntityName: "wall",
			CreateEntity: func(position hnbMath.Vector, _ int) hnbCore.IEntity {
				wall := WallFactory{}.Create()
				wall.SetPosition(position)

				return wall
			},
		},
		5: MapEntityCreator{
			TileId:     0,
			EntityName: "floor-boundary",
			CreateEntity: func(position hnbMath.Vector, length int) hnbCore.IEntity {
				entity := &hnbCore.Entity{}
				entity.SetCanCollide(true)
				entity.SetCollisionMask("floor-boundary")
				entity.SetCollisioningMasks([]string{
					"character",
				})
				entity.SetSize(hnbMath.Vector{
					X: 16,
					Y: 16,
				})
				entity.SetPosition(position)
				entity.SetOrigin(entity.GetSize().By(0.5))

				return entity
			},
		},
	}

	return creators
}
