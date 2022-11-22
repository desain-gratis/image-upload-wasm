package lib

// Import the package to access the Wasm environment
import (
	"image"
)

func CropByCenterAndSpoke(rect image.Rectangle, centerX int, centerY int, ratioX int, ratioY, spoke int) image.Rectangle {

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

	return cropByCenterWidthHeight(centerX, centerY, width*2, height*2)
	// return cropFromPointByWidth(ratioX, ratioY, size, byHeight, centerX, centerY)
}

func cropByCenterWidthHeight(centerX int, centerY int, width int, height int) image.Rectangle {
	_minx := float64(centerX) - float64(width)/2
	_miny := float64(centerY) - float64(height)/2
	_maxx := float64(centerX) + float64(width)/2
	_maxy := float64(centerY) + float64(height)/2

	return image.Rect(int(_minx), int(_miny), int(_maxx), int(_maxy))
}

// resizeFromPoint
func cropFromPointByWidth(ratioX, ratioY int, size int, byHeight bool, centerX int, centerY int) image.Rectangle {
	if size <= 0 || ratioX <= 0 || ratioY <= 0 {
		return image.Rectangle{
			Min: image.Pt(centerX, centerY),
			Max: image.Pt(centerX, centerY),
		}
	}

	// resize by width.
	uW := float64(1)
	uH := float64(ratioY) / float64(ratioX)

	if byHeight {
		uW = float64(ratioX) / float64(ratioY)
		uH = float64(1)
	}

	_minx := float64(centerX) - float64(size)*uW/2
	_miny := float64(centerY) - float64(size)*uH/2
	_maxx := float64(centerX) + float64(size)*uW/2
	_maxy := float64(centerY) + float64(size)*uH/2

	return image.Rect(int(_minx), int(_miny), int(_maxx), int(_maxy))
}
