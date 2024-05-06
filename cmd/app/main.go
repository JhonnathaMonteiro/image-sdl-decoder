package main

import (
	"app-sdl/internals/parsers"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	imageData, err := parsers.BMP("images/bmp/bmp_24.bmp")
	if err != nil {
		panic(err)
	}

	pixelData := imageData.PixelData

	pitch := ((imageData.InfoHeader.Width*3 + 3) >> 2) << 2 // padding to 4 bytes (+3 helps to round up to the nearest multiple of 4)
	surface, err := sdl.CreateRGBSurfaceFrom(
		unsafe.Pointer(&pixelData),             // pixels
		int32(imageData.InfoHeader.Width),      // width
		int32(imageData.InfoHeader.Heigth),     // height
		int(imageData.InfoHeader.BitsPerPixel), // depth
		int(pitch),                             // pitch
		0,                                      // Rmask
		0,                                      // Gmask
		0,                                      // Bmask
		0,                                      // Amask
	)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}

	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	renderer.Clear()

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}

	renderer.SetRenderTarget(texture)
	renderer.Copy(texture, nil, nil)
	renderer.Present()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent: // NOTE: Please use `*sdl.QuitEvent` for `v0.4.x` (current version).
				println("Quit")
				running = false
				break
			}
		}

		sdl.Delay(33)
	}

	println("DONE!")
}

// ➜  app-sdl git:(main) ✗ go run cmd/app/main.go
// [120000/120000]0xc00012a000
// panic: runtime error: cgo argument has Go pointer to unpinned Go pointer
//
// goroutine 1 [running, locked to thread]:
// github.com/veandco/go-sdl2/sdl.CreateRGBSurfaceFrom.func1(0xc00014a000, 0xc8, 0xc8, 0x18, 0x258, 0x0, 0x0, 0x0, 0x0)
//         /home/jhonnatha-ubt-note2/go/pkg/mod/github.com/veandco/go-sdl2@v0.4.38/sdl/surface.go:202 +0x53
// github.com/veandco/go-sdl2/sdl.CreateRGBSurfaceFrom(0xc00012a000?, 0x1d4c0?, 0x0?, 0x1d4c0?, 0x25800000320?, 0x4?, 0xc0?, 0x418930?, 0x0?)
//         /home/jhonnatha-ubt-note2/go/pkg/mod/github.com/veandco/go-sdl2@v0.4.38/sdl/surface.go:202 +0x13
// main.main()
//         /home/jhonnatha-ubt-note2/Documents/sandbox/app-sdl/cmd/app/main.go:32 +0x1a5
// exit status 2
