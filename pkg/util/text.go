package util

import (
	"bytes"
	"strings"
)

func WordWrap(s string, maxWidth, charWidth int) string {
	if len(s)*charWidth <= maxWidth {
		return s
	}
	var buf [128]byte
	sb := bytes.NewBuffer(buf[:])
	sb.Reset()
	lineWidth := 0

	for word := range strings.FieldsSeq(s) {
		wordWidth := len(word) * charWidth
		if lineWidth == 0 {
			lineWidth += wordWidth
			sb.WriteString(word)
			continue
		}
		wordWidthPlusSpace := wordWidth + charWidth
		if lineWidth+wordWidthPlusSpace > maxWidth {
			sb.WriteByte('\n')
			lineWidth = wordWidth
		} else {
			sb.WriteByte(' ')
			lineWidth += wordWidthPlusSpace
		}
		sb.WriteString(word)
	}
	return sb.String()
}
