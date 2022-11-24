package lib

// Import the package to access the Wasm environment
import (
	"image"
)

// scale (0 - 1.0]
func CropByCenterAndScale(rect image.Rectangle, centerX int, centerY int, ratioX int, ratioY int, scale float64) image.Rectangle {
	// We just need to find the max width here or use spoke
	x1 := centerX
	x2 := rect.Max.X - centerX
	y1 := centerY
	y2 := rect.Max.Y - centerY

	_ratioX := float64(ratioX)
	_ratioY := float64(ratioY)

	minX := x1
	if x2 < x1 {
		minX = x2
	}
	minY := y1
	if y2 < y1 {
		minY = y2
	}

	// how many height can be taken, given limited width (min width)
	// how many width can be taken, given limited height (min height)
	maxWidth := float64(minY) * _ratioX / _ratioY
	maxHeight := float64(minX) * _ratioY / _ratioX

	var width, height float64

	width = float64(minX)
	if float64(minX) > maxWidth {
		width = maxWidth
	}
	height = float64(minY)
	if float64(minY) > maxHeight {
		height = maxHeight
	}

	var boundedByWidth bool
	if width < height {
		boundedByWidth = true
	}

	if boundedByWidth {
		height = float64(width) * _ratioY / _ratioX
	} else {
		width = float64(height) * _ratioX / _ratioY
	}

	width = width * 2
	height = height * 2

	width = float64(width) * scale
	height = float64(height) * scale

	return cropByCenterWidthHeight(centerX, centerY, int(width), int(height))
}

func cropByCenterWidthHeight(centerX int, centerY int, width int, height int) image.Rectangle {
	_minx := float64(centerX) - float64(width)/2
	_miny := float64(centerY) - float64(height)/2
	_maxx := float64(centerX) + float64(width)/2
	_maxy := float64(centerY) + float64(height)/2

	return image.Rect(int(_minx), int(_miny), int(_maxx), int(_maxy))
}
