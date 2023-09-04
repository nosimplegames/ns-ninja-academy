package obstacles

import "simple-games.com/ninja/src/res"

type LogFactory struct {
}

func (factory LogFactory) Create() *Log {
	log := &Log{}

	log.SetTexture(res.GetTextures().Log)
	log.SetSize(res.LogFrameSize)
	log.SetOriginCenter()
	log.SetCanCollide(true)
	log.SetCollisioningMasks([]string{
		"shuriken",
	})
	log.SetCollisionMask("log")

	return log
}
