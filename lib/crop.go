package lib

// Import the package to access the Wasm environment
import (
	"image"
)

func CropByCenterAndSpoke(rect image.Rectangle, center image.Point, ratioX int, ratioY, spoke int) image.Rectangle {

	// We just need to find the max width here or use spoke
	x1 := center.X
	x2 := rect.Max.X - center.X
	y1 := center.Y
	y2 := rect.Max.Y - center.Y

	minX := x1
	if x2 < x1 {
		minX = x1
	}
	minY := y1
	if y2 < y1 {
		minY = y2
	}

	// if bounded by width, we don't do anything
	byHeight := false
	size := minX * 2
	if minY < minX {
		byHeight = true
		size = minY * 2
	}

	if spoke < size {
		size = spoke
	}

	return cropFromPointByWidth(ratioX, ratioY, size, byHeight, center)
}

// resizeFromPoint
func cropFromPointByWidth(ratioX, ratioY int, size int, byHeight bool, center image.Point) image.Rectangle {
	if size <= 0 || ratioX <= 0 || ratioY <= 0 {
		return image.Rectangle{
			Min: center,
			Max: center,
		}
	}

	// resize by width.
	uW := float64(1)
	uH := float64(ratioY) / float64(ratioX)

	if byHeight {
		uW = float64(ratioX) / float64(ratioY)
		uH = float64(1)
	}

	_minx := float64(center.X) - float64(size)*uW/2
	_miny := float64(center.Y) - float64(size)*uH/2
	_maxx := float64(center.X) + float64(size)*uW/2
	_maxy := float64(center.Y) + float64(size)*uH/2

	return image.Rect(int(_minx), int(_miny), int(_maxx), int(_maxy))
}
