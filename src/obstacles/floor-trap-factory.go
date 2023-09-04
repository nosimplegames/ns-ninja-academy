package obstacles

import "simple-games.com/ninja/src/res"

type FloorTrapFactory struct {
}

func (factory FloorTrapFactory) Create() *FloorTrap {
	trap := &FloorTrap{}

	trap.SetTexture(res.GetTextures().FloorTrap)
	trap.SetSize(res.FloorTrapFrameSize)
	trap.SetOriginCenter()
	trap.SetCanCollide(true)
	trap.SetCollisioningMasks([]string{
		"character",
	})
	trap.SetCollisionMask("floor-trap")

	return trap
}
