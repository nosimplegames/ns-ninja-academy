package level

import "github.com/nosimplegames/ns-framework/hnbCore"

type EntityCreatorFunc func() hnbCore.IEntity

type EntitySpawner struct {
	hnbCore.Entity

	CreateEntity EntityCreatorFunc

	isEntityAlive bool
}

func (spawner *EntitySpawner) UpdateFrame() {
	spawner.Entity.UpdateFrame()

	if !spawner.isEntityAlive {
		spawner.SpawnEntity()
	}
}

func (spawner *EntitySpawner) SpawnEntity() {
	canSpawn := spawner.CreateEntity != nil

	if !canSpawn {
		return
	}

	entity := spawner.CreateEntity()
	entity.On("die", func() {
		spawner.isEntityAlive = false
	})
	entity.MoveY(8)

	spawner.isEntityAlive = true
	hnbCore.EntityAdder{
		Child:  entity,
		Parent: spawner,
	}.Add()
}
