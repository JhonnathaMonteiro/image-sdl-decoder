package data

type BmpFileHeader struct {
	Magic    uint16 // BM (0x424D) or MB (if using little endian)
	Size     uint32 // File size in bytes
	Reserved uint32 // unused (0x0)
	Offset   uint32 // Offset from the beginning of the file to the beginning of the bitmap data
}

type BpmInfoHeader struct {
	Size            uint32 // Size of this header (in bytes)
	Width           uint32 // Width of bitmap in pixels
	Heigth          uint32 // Heigth of bitmap in pixels
	Planes          uint16 // No. of planes for the target device, this is always 1
	BitsPerPixel    uint16 // No. of bits per pixel
	Compression     uint32 
	ImageSize       uint32
	XpixelsPerM     uint32
	YpixelsPerM     uint32
	ColorsUsed      uint32
	ImportantColors uint32
}
