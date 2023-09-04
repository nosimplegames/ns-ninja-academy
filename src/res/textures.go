package res

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/nosimplegames/ns-framework/hnbAssets"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type Textures struct {
	Tileset hnbRender.Texture

	NinjaFallingTexture     hnbRender.Texture
	NinjaIdleTexture        hnbRender.Texture
	NinjaIdleSpriteSheet    hnbRender.Texture
	NinjaWalkingSpriteSheet hnbRender.Texture

	ShurikenSpriteSheet hnbRender.Texture
	Log                 hnbRender.Texture
	FloorTrap           hnbRender.Texture
}

var textures *Textures = nil

func GetTextures() *Textures {
	areTexturesInited := textures != nil

	if !areTexturesInited {
		textures = loadTextures()
	}

	return textures
}

func loadTextures() *Textures {
	textures := &Textures{}

	textures.Tileset = hnbAssets.LoadTexture(tileset)

	ninjaIdleSpriteSheet := hnbAssets.LoadTexture(ninjaIdleSpriteSheet)
	textures.NinjaFallingTexture = ninjaIdleSpriteSheet.SubImage(
		image.Rect(
			0,
			0,
			int(NinjaFrameSize.X),
			int(NinjaFrameSize.Y),
		),
	).(*ebiten.Image)
	textures.NinjaIdleTexture = ninjaIdleSpriteSheet.SubImage(
		image.Rect(
			0,
			0,
			int(NinjaFrameSize.X),
			int(NinjaFrameSize.Y),
		),
	).(*ebiten.Image)
	textures.NinjaIdleSpriteSheet = ninjaIdleSpriteSheet
	textures.NinjaWalkingSpriteSheet = hnbAssets.LoadTexture(ninjaWalkingSpriteSheet)

	textures.ShurikenSpriteSheet = hnbAssets.LoadTexture(shurikenSpriteSheet)

	textures.Log = hnbAssets.LoadTexture(logTexture)
	textures.FloorTrap = hnbAssets.LoadTexture(floorTrapTexture)

	return textures
}
