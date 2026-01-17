package util

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

type ExtImage struct {
	firefly.File
	firefly.Image
}

func NewExtImage(file firefly.File) ExtImage {
	return ExtImage{File: file, Image: file.Image()}
}

func (i ExtImage) GetColorAt(point firefly.Point) firefly.Color {
	if point.X < 0 || point.Y < 0 {
		return firefly.ColorNone
	}
	size := i.Size()
	if point.X >= size.W || point.Y >= size.H {
		return firefly.ColorNone
	}
	bpp := i.Raw[1]
	headerLen := 5 + (1 << (bpp - 1))
	body := i.Raw[headerLen:]

	pixelIndex := point.X + point.Y*size.W
	bodyIndex := pixelIndex * int(bpp) / 8
	pixelValue := body[bodyIndex]

	switch bpp {
	case 1:
		// the 'switch' can be reduced and just do a single equation
		// but I ain't got time for that right now
		byteOffset := 1 * (7 - pixelIndex%8)
		pixelValue = (pixelValue >> byte(byteOffset)) & 0b1
	case 2:
		byteOffset := 2 * (3 - pixelIndex%4)
		pixelValue = (pixelValue >> byte(byteOffset)) & 0b11
	case 4:
		byteOffset := 4 * (1 - pixelIndex%2)
		pixelValue = (pixelValue >> byte(byteOffset)) & 0b1111
	default:
		panic("invalid bpp")
	}

	return i.GetColor(pixelValue)
}
