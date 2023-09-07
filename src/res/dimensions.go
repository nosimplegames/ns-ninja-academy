package res

import "github.com/nosimplegames/ns-framework/hnbMath"

var (
	GameSize   = hnbMath.Vector{X: 320, Y: 16 * 13}
	WindowSize = GameSize.By(2)
	TileSize   = hnbMath.Vector{X: 16, Y: 16}

	NinjaFrameSize = hnbMath.Vector{X: 16, Y: 16}

	ShurikenFrameSize = hnbMath.Vector{X: 5, Y: 5}

	LogFrameSize       = hnbMath.Vector{X: 11, Y: 16}
	FloorTrapFrameSize = hnbMath.Vector{X: 16, Y: 5}
)
