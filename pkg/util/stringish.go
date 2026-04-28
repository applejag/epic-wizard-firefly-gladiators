package util

type Stringish interface {
	~string | ~[]byte
}

func TrimRightByte[S Stringish](s S, c byte) S {
	for len(s) > 0 && s[len(s)-1] == c {
		s = s[:len(s)-1]
	}
	return s
}

func TrimLeftByte[S Stringish](s S, c byte) S {
	for len(s) > 0 && s[0] == c {
		s = s[1:]
	}
	return s
}

func LastIndexByte[S Stringish](s S, c byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func WriteWrapped[S Stringish](buf []byte, s S, maxWidth, charWidth int) int {
	written := 0
	rest := TrimRightByte(s, ' ')
	for len(rest) != 0 {
		rest = TrimLeftByte(rest, ' ')

		if len(rest)*charWidth <= maxWidth {
			written += copy(buf[written:], rest)
			break
		}

		maxLen := maxWidth / charWidth
		splitAt := LastIndexByte(rest[:maxLen], ' ')
		if splitAt == -1 {
			written += copy(buf[written:], rest[:maxLen])
			buf[written] = '\n'
			written += 1
			rest = rest[maxLen:]
			continue
		}
		written += copy(buf[written:], rest[:splitAt])
		buf[written] = '\n'
		written += 1
		rest = rest[maxLen:]
	}
	return written
}

// ConcatInto writes the given strings into a buffer with zero allocations.
//
// Returns the number of bytes written. Intended usage:
//
//	var buf [10]
//	written := ConcatInto(buf[:], "hello", "world")
//	foobar(string(buf[:written]))
//
// Panics if "buf" is too small.
func ConcatInto[S Stringish](buf []byte, parts ...S) int {
	written := 0
	for _, s := range parts {
		if len(s) > len(buf) {
			panic("ConcatBuf: buffer is not big enough")
		}
		written += copy(buf, s)
		buf = buf[len(s):]
	}
	return written
}
