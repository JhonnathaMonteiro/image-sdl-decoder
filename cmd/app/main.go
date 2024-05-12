package main

import (
	_ "app-sdl/pkg/bmp"
	"fmt"
	"image"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	file, err := os.Open("images/bmp/lena.bmp")
	// file, err := os.Open("images/bmp/bmp_24.bmp")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, formatName, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		fmt.Sprintf("Rendering: %s", formatName),
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(img.Bounds().Dx()),
		int32(img.Bounds().Dy()),
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Convert image to SDL surface
	surface := imageToSurface(img)
	defer surface.Free()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	texture, err := renderer.CreateTextureFromSurface(imageToSurface(img))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return
	}
	defer texture.Destroy()

	// Clear renderer
	renderer.Clear()

	// Copy surface to renderer
	renderer.Copy(texture, nil, nil)

	// Present renderer
	renderer.Present()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
	}
}

// Convert Go image to SDL surface
func imageToSurface(img image.Image) *sdl.Surface {
	// Get image dimensions
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create SDL surface
	surface, err := sdl.CreateRGBSurface(
		0,
		int32(width),
		int32(height),
    32, // 32 bits per pixel in the surface R: 8, G: 8, B: 8, A: 8
    0, 
    0, 
    0,
    0,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create surface: %s\n", err)
		return nil
	}

	// Set pixel colors
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			surface.Set(x, y, img.At(x, y))
		}
	}

	return surface
}
