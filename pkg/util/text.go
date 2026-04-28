package util

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

func DrawTextRightAligned(font firefly.Font, text string, right firefly.Point, color firefly.Color) {
	width := font.CharWidth() * len(text)
	font.Draw(text, right.Add(firefly.P(-width, 0)), color)
}

func DrawTextCentered(font firefly.Font, text string, center firefly.Point, color firefly.Color) {
	width := font.CharWidth() * len(text)
	font.Draw(text, center.Add(firefly.P(-width/2, 0)), color)
}

func FormatIntInto(buf []byte, num int) int {
	if num < 0 {
		buf[0] = '-'
		return FormatIntInto(buf[1:], -num) + 1
	}
	if num < 10 {
		buf[0] = '0' + byte(num)
		return 1
	}
	size := numberOfDigits(num)
	index := size - 1
	for num > 0 && index >= 0 {
		buf[index] = '0' + byte(num%10)
		num /= 10
		index--
	}
	return size
}

func numberOfDigits(num int) int {
	switch {
	case num < 0:
		return numberOfDigits(-num) + 1
	case num < 1e1:
		return 1
	case num < 1e2:
		return 2
	case num < 1e3:
		return 3
	case num < 1e4:
		return 4
	case num < 1e5:
		return 5
	case num < 1e6:
		return 6
	case num < 1e7:
		return 7
	case num < 1e8:
		return 8
	case num < 1e9:
		return 9
	default:
		panic("number is too big")
	}
}
