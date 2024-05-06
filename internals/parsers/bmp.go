package parsers

import (
	"app-sdl/internals/data"
	"encoding/binary"
	"os"
)

func BMP(filepath string) (*data.BmpFile, error) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	var fileHeader data.BmpFileHeader
	err = binary.Read(file, binary.LittleEndian, &fileHeader)
	if err != nil {
		panic(err)
	}

	if fileHeader.Magic != 0x4d42 && fileHeader.Magic != 0x424d {
		panic("invalid BMP file format")
	}

	var infoHeader data.BpmInfoHeader
	err = binary.Read(file, binary.LittleEndian, &infoHeader)
	if err != nil {
		panic(err)
	}

	imagePPB := infoHeader.BitsPerPixel >> 3
	rowSize := ((infoHeader.Width*uint32(imagePPB) + 3) >> 2) << 2 // padding to 4 bytes (+3 helps to round up to the nearest multiple of 4)

	pixelData := make([]uint8, rowSize*infoHeader.Heigth)
	file.ReadAt(pixelData, int64(fileHeader.Offset))

  return &data.BmpFile{
    FileHeader: fileHeader,
    InfoHeader: infoHeader,
    PixelData:  pixelData,
  }, nil

}
