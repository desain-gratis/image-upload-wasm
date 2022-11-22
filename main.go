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
	js.Global().Set("CenterCrop", CenterCrop())
	select {}
}

func GetSize() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 2 {
			fmt.Println("Ohnono, args <= 0")
			return nil
		}

		size := args[0].Int()
		data := args[1]

		if size > (10 * (1 << 20)) {
			fmt.Println("Ohnono, size too large")
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

// CenterCrop returns a JavaScript function
func CenterCrop() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 2 {
			fmt.Println("Ohnono, args <= 0")
			return nil
		}

		size := args[0].Int()
		data := args[1]
		centerX := args[2].Int()
		centerY := args[3].Int()
		ratioX := args[4].Int()
		ratioY := args[5].Int()
		spoke := args[6].Int()

		// validate size too large

		if size > (10 * (1 << 20)) {
			fmt.Println("Ohnono, size too large")
			return nil
		}

		buffer := make([]byte, size)

		bytesRead := js.CopyBytesToGo(buffer, data)

		// in case we use fixed sized buffer
		buffer = buffer[:bytesRead]

		r := bytes.NewReader(buffer)
		imc, format, err := image.DecodeConfig(r)
		fmt.Printf("Size: %+v x %+v; format: %v, err: %v\n", imc.Width, imc.Height, format, err)
		fmt.Printf("Format: %v\n", format)
		fmt.Printf("Err: %v\n", err)

		r.Reset(buffer)
		img, format, err := image.Decode(r)

		// fmt.Println("%+v\n", img.)
		fmt.Printf("Format is: %v\n", format)
		fmt.Printf("Err: %v\n", err)

		// centerX := img.Bounds().Min.X + img.Bounds().Max.X/2
		// centerY := img.Bounds().Min.Y + img.Bounds().Max.Y/2

		// cropByCenterAndSpoke(img.Bounds(), centerX, centerY, 3, 1, 20000)
		crop := lib.CropByCenterAndSpoke(img.Bounds(), image.Pt(centerX, centerY), ratioX, ratioY, spoke)

		fmt.Println(crop.Min.X, crop.Min.Y, crop.Max.X, crop.Max.Y)
		fmt.Println(crop)

		// todo can use sub image
		cropped := image.NewRGBA(image.Rect(0, 0, crop.Max.X-crop.Min.X, crop.Max.Y-crop.Min.Y))

		draw.Draw(cropped, cropped.Rect, img, image.Pt(0, 0), draw.Src)

		//Global().Get("Uint8Array").New(len(bufff.Bytes()))
		copied := js.CopyBytesToJS(dst, bufff.Bytes())
		fmt.Println("COPIED", copied)

		return map[string]interface{}{
			"ORIGINAL": nil,
			"1x1":      nil,
			"1x2":      nil,
			"1x3":      nil,
			"1x4":      nil,
			"2x1":      nil,
			"2x3":      nil,
			"3x1":      nil,
			"3x2":      nil,
			"3x4":      nil,
			"4x1":      nil,
			"4x2":      nil,
			"4x3":      nil,
			"9x16":     nil,
			"16x9":     nil,
			"test":     dst,
		}
	})
}
