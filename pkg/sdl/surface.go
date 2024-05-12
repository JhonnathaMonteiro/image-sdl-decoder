package display

import (
	"image"
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
)

type Surface sdl.Surface

// ColorModel returns the color model of the surface.
func (s *Surface) ColorModel() color.Model {
	switch s.Format.Format {
	case sdl.PIXELFORMAT_ARGB8888, sdl.PIXELFORMAT_ABGR8888:
		return color.RGBAModel
	case sdl.PIXELFORMAT_RGB888:
		return color.RGBAModel
	default:
		panic("unsupported pixel format")
	}
}

func (s *Surface) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(s.W), int(s.H))
}

func (s *Surface) At(x, y int) color.Color {
	pix := (*sdl.Surface)(s).Pixels()
	i := int32(y)*s.Pitch + int32(x)*int32(s.Format.BytesPerPixel)
	switch s.Format.Format {
	case sdl.PIXELFORMAT_RGB888:
		return color.RGBA{pix[i], pix[i+1], pix[i+2], 0xff}
	default:
		panic("not implemented")
	}
}

func (s *Surface) Set(x, y int, c color.Color) {
	pix := (*sdl.Surface)(s).Pixels()
	i := int32(y)*s.Pitch + int32(x)*int32(s.Format.BytesPerPixel)
	switch s.Format.Format {
	case sdl.PIXELFORMAT_ARGB8888:
		col := s.ColorModel().Convert(c).(color.RGBA)
		pix[i+0] = col.R
		pix[i+1] = col.G
		pix[i+2] = col.B
		pix[i+3] = col.A
	case sdl.PIXELFORMAT_ABGR8888:
		col := s.ColorModel().Convert(c).(color.RGBA)
		pix[i+3] = col.R
		pix[i+2] = col.G
		pix[i+1] = col.B
		pix[i+0] = col.A
	case sdl.PIXELFORMAT_RGB24, sdl.PIXELFORMAT_RGB888:
		col := s.ColorModel().Convert(c).(color.RGBA)
		pix[i+0] = col.B
		pix[i+1] = col.G
		pix[i+2] = col.R
	case sdl.PIXELFORMAT_BGR24, sdl.PIXELFORMAT_BGR888:
		col := s.ColorModel().Convert(c).(color.RGBA)
		pix[i+2] = col.R
		pix[i+1] = col.G
		pix[i+0] = col.B

	default:
		panic("Unknown pixel format!")
	}
}
