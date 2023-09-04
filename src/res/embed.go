package res

import (
	_ "embed"
)

var (
	//go:embed tileset.png
	tileset []byte

	//go:embed ninja-idle-animation.png
	ninjaIdleSpriteSheet []byte
	//go:embed ninja-walking-animation.png
	ninjaWalkingSpriteSheet []byte

	//go:embed shuriken-animation.png
	shurikenSpriteSheet []byte

	//go:embed log.png
	logTexture []byte
)
