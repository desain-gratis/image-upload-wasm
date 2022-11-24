package main

import (
	"bytes"
	"fmt"
	"image"
	"syscall/js"

	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/id-auction/image-upload-wasm/lib"

	_ "golang.org/x/image/webp"
)

func main() {
	js.Global().Set("GetSize", GetSize())
	js.Global().Set("Crop", Crop())
	js.Global().Set("GetRGBA", GetRGBA())
	js.Global().Set("CropRGBA", CropRGBA())
	select {}
}

func CropRGBA() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 4 {
			fmt.Println("Ohnono, args < 5")
			return nil
		}

		size := args[0].Int()
		data := args[1]
		width := args[2].Int()
		height := args[3].Int()
		stride := args[4].Int()

		centerX := args[5].Int()
		centerY := args[6].Int()
		ratioX := args[7].Int()
		ratioY := args[8].Int()
		scale := args[9].Float()

		if size > (100 * (1 << 20)) {
			fmt.Println("Ohnono, size too large (greater than 100 Mb)")
			return nil
		}

		buffer := make([]byte, size)
		bytesRead := js.CopyBytesToGo(buffer, data)
		// in case we use fixed sized buffer
		buffer = buffer[:bytesRead]

		originalCopy := image.NewRGBA(image.Rect(0, 0, width, height))
		originalCopy.Pix = buffer
		originalCopy.Stride = stride

		crop := lib.CropByCenterAndScale(originalCopy.Bounds(), centerX, centerY, ratioX, ratioY, scale)

		// todo can use sub image
		cropped := image.NewRGBA(image.Rect(0, 0, crop.Max.X-crop.Min.X, crop.Max.Y-crop.Min.Y))
		draw.Draw(cropped, cropped.Rect, originalCopy, crop.Min, draw.Src)

		dst := js.Global().Get("Uint8Array").New(len(cropped.Pix))
		_ = js.CopyBytesToJS(dst, cropped.Pix)

		return map[string]interface{}{
			"success": map[string]interface{}{
				"data":   dst,
				"width":  cropped.Rect.Dx(),
				"height": cropped.Rect.Dy(),
				"stride": cropped.Stride,
			},
		}
	})
}

// GetRGBA without using Canvas :(
func GetRGBA() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 2 {
			fmt.Println("Ohnono, args < 2")
			return nil
		}

		size := args[0].Int()
		data := args[1]

		if size > (10 * (1 << 20)) {
			fmt.Println("Ohnono, size too large (greater than 10 Mb)")
			return nil
		}

		buffer := make([]byte, size)
		bytesRead := js.CopyBytesToGo(buffer, data)
		// in case we use fixed sized buffer
		buffer = buffer[:bytesRead]

		r := bytes.NewReader(buffer)
		// imc, format, err := image.DecodeConfig(r)
		// fmt.Printf("Size: %+v x %+v; format: %v, err: %v\n", imc.Width, imc.Height, format, err)
		// fmt.Printf("Format: %v\n", format)
		// fmt.Printf("Err: %v\n", err)
		// r.Reset(buffer)

		img, _, err := image.Decode(r)
		if err != nil {
			return map[string]interface{}{
				"failed": err.Error(),
			}
		}

		originalCopy := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))

		draw.Draw(originalCopy, originalCopy.Rect, img, image.Pt(0, 0), draw.Src)

		dst := js.Global().Get("Uint8Array").New(len(originalCopy.Pix))
		_ = js.CopyBytesToJS(dst, originalCopy.Pix)

		return map[string]interface{}{
			"success": map[string]interface{}{
				"data":   dst,
				"width":  originalCopy.Rect.Dx(),
				"height": originalCopy.Rect.Dy(),
				"stride": originalCopy.Stride,
			},
		}

	})
}

func GetSize() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 2 {
			fmt.Println("Ohnono, args < 2")
			return nil
		}

		size := args[0].Int()
		data := args[1]

		if size > (10 * (1 << 20)) {
			fmt.Println("Ohnono, size too large (greater than 10 Mb)")
			return nil
		}

		buffer := make([]byte, size)
		bytesRead := js.CopyBytesToGo(buffer, data)
		// in case we use fixed sized buffer
		buffer = buffer[:bytesRead]

		r := bytes.NewReader(buffer)
		imc, _, _ := image.DecodeConfig(r)

		return map[string]interface{}{
			"width":  imc.Width,
			"height": imc.Height,
		}

	})
}

// Crop returns a JavaScript function
func Crop() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 2 {
			return map[string]interface{}{
				"failed": "Ohnono, args <= 0",
			}
		}

		size := args[0].Int()
		data := args[1]
		centerX := args[2].Int()
		centerY := args[3].Int()
		ratioX := args[4].Int()
		ratioY := args[5].Int()
		scale := args[6].Float()

		// validate size too large

		if size > (10 * (1 << 20)) {
			return map[string]interface{}{
				"failed": "Ohnono, size too large",
			}
		}

		buffer := make([]byte, size)

		bytesRead := js.CopyBytesToGo(buffer, data)

		// in case we use fixed sized buffer
		buffer = buffer[:bytesRead]

		r := bytes.NewReader(buffer)
		// imc, format, err := image.DecodeConfig(r)
		// fmt.Printf("Size: %+v x %+v; format: %v, err: %v\n", imc.Width, imc.Height, format, err)
		// fmt.Printf("Format: %v\n", format)
		// fmt.Printf("Err: %v\n", err)
		// r.Reset(buffer)

		img, _, err := image.Decode(r)
		if err != nil {
			return map[string]interface{}{
				"failed": err.Error(),
			}
		}

		// fmt.Println("%+v\n", img.)
		// fmt.Printf("Format is: %v\n", format)
		// fmt.Printf("Err: %v\n", err)

		// centerX := img.Bounds().Min.X + img.Bounds().Max.X/2
		// centerY := img.Bounds().Min.Y + img.Bounds().Max.Y/2

		// cropByCenterAndSpoke(img.Bounds(), centerX, centerY, 3, 1, 20000)
		crop := lib.CropByCenterAndScale(img.Bounds(), centerX, centerY, ratioX, ratioY, scale)

		// fmt.Println(crop.Min.X, crop.Min.Y, crop.Max.X, crop.Max.Y)
		// fmt.Println(crop)

		// todo can use sub image
		cropped := image.NewRGBA(image.Rect(0, 0, crop.Max.X-crop.Min.X, crop.Max.Y-crop.Min.Y))
		draw.Draw(cropped, cropped.Rect, img, image.Pt(0, 0), draw.Src)

		dst := js.Global().Get("Uint8Array").New(len(cropped.Pix))
		_ = js.CopyBytesToJS(dst, cropped.Pix)

		return map[string]interface{}{
			"success": dst,
		}
	})
}
