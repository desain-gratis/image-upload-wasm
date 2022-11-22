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
	maxWidth := minY * ratioX / ratioY
	maxHeight := minX * ratioY / ratioX

	// Check between the two, whos limiting who
	boundedByWidth := false
	if maxWidth < maxHeight {
		boundedByWidth = true
	}

	width := maxHeight * ratioX / ratioY
	height := maxHeight

	if boundedByWidth {
		width = maxWidth
		height = maxWidth * ratioY / ratioX
	}

	width = width * 2
	height = height * 2

	width = int(float64(width) * scale)
	height = int(float64(height) * scale)

	return cropByCenterWidthHeight(centerX, centerY, width, height)
}

func cropByCenterWidthHeight(centerX int, centerY int, width int, height int) image.Rectangle {
	_minx := float64(centerX) - float64(width)/2
	_miny := float64(centerY) - float64(height)/2
	_maxx := float64(centerX) + float64(width)/2
	_maxy := float64(centerY) + float64(height)/2

	return image.Rect(int(_minx), int(_miny), int(_maxx), int(_maxy))
}
