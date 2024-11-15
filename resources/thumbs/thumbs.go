package thumbs

import (
	"fmt"
	"image"

	assets "github.com/xlcaptain/go-captcha-staic/bindata/thumbs"
	"github.com/xlcaptain/go-captcha-staic/helper"
)

func GetThumbs() ([]image.Image, error) {
	var images []image.Image
	var err error
	for i := 1; i <= 5; i++ {
		var img image.Image
		var asset []byte
		asset, err = assets.Asset(fmt.Sprintf("sourcedata/thumbs/thumb-%d.png", i))
		if err != nil {
			return images, err
		}

		img, err = helper.DecodeByteToPng(asset)
		if err != nil {
			return images, err
		}

		images = append(images, img)
	}

	return images, nil
}
