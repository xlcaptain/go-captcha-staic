package images

import (
	"fmt"
	"image"

	assets "github.com/xlcaptain/go-captcha-staic/bindata/images"
	"github.com/xlcaptain/go-captcha-staic/helper"
)

func GetImages() ([]image.Image, error) {
	var images []image.Image
	var err error
	for i := 1; i <= 5; i++ {
		var img image.Image
		var asset []byte
		asset, err = assets.Asset(fmt.Sprintf("sourcedata/images/image-%d.jpg", i))
		if err != nil {
			return images, err
		}

		img, err = helper.DecodeByteToJpeg(asset)
		if err != nil {
			return images, err
		}

		images = append(images, img)
	}

	return images, nil
}
