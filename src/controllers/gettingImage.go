package controllers

import (
	"image"

	"github.com/kbinani/screenshot"
)

func gettingImage() (img *image.RGBA) {
	bounds := screenshot.GetDisplayBounds(0)
	img, _ = screenshot.CaptureRect(bounds)

	return
}
