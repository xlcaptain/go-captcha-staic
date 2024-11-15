package shapes

import (
	"fmt"
	"image"

	assets "github.com/xlcaptain/go-captcha-staic/bindata/shapes"
	"github.com/xlcaptain/go-captcha-staic/helper"
)

func GetShapes() (map[string]image.Image, error) {
	var imgMaps = make(map[string]image.Image)
	var err error
	for i := 1; i <= 13; i++ {
		var img image.Image
		var asset []byte
		asset, err = assets.Asset(fmt.Sprintf("sourcedata/shapes/shape-%d.png", i))
		if err != nil {
			return imgMaps, err
		}

		img, err = helper.DecodeByteToPng(asset)
		if err != nil {
			return imgMaps, err
		}

		key := fmt.Sprintf("shape-%d", i)
		imgMaps[key] = img
	}

	return imgMaps, nil
}
