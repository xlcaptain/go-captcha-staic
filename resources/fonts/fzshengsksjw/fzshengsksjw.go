package fzshengsksjw

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	assets "github.com/xlcaptain/go-captcha-staic/bindata/fonts/fzshengsksjw"
)

func GetFont() (font *truetype.Font, err error) {
	asset, err := assets.Asset("sourcedata/fonts/fzshengsksjw_cu.ttf")
	if err != nil {
		return font, err
	}

	font, err = freetype.ParseFont(asset)
	if err != nil {
		return nil, err
	}

	return font, nil
}
