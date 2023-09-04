package main

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbEntities"
	"simple-games.com/ninja/src/res"
	"simple-games.com/ninja/src/scenes"
)

func main() {
	game := hnbCore.Game{
		Size:          res.GameSize,
		WindowSize:    res.WindowSize,
		MustDrawWorld: true,
	}

	scene := scenes.TrainingRoomFactory{}.Create()
	game.PushScene(scene)

	camera := &hnbEntities.Camera{
		RenderingBox: game.GetDefaultRenderingBox(),
	}
	game.AddCamera(camera)
	camera.SetScale(res.WindowSize.DivideV(res.GameSize))

	game.Run()
}
