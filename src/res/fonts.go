package res

import (
	"github.com/nosimplegames/ns-framework/hnbAssets"
	"golang.org/x/image/font"
)

type Fonts struct {
	NormalText font.Face
}

var fonts *Fonts = nil

func GetFonts() *Fonts {
	areFontsInited := fonts != nil

	if !areFontsInited {
		fonts = loadFonts()
	}

	return fonts
}

func loadFonts() *Fonts {
	fonts := &Fonts{}

	fonts.NormalText = hnbAssets.FontFaceFactory{
		FontData: hnbAssets.FontData{
			Bytes: mainFont,
			Name:  "main",
		},
		DPI:  72,
		Size: 20,
	}.Create()

	return fonts
}
