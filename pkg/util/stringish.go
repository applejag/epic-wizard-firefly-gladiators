package util

type Stringish interface {
	~string | ~[]byte
}

// TrimIndexRightByte is equivalent to [strings.TrimRight],
// but does not return a slice to not trigger heap escape analysis.
//
// TinyGo's escape analyzer cannot guarantee that a slice can stay on the stack
// when it is given to- and then returned from a function.
// As soon as a function takes in a slice and then returns a modified version
// of that slice (such as with [strings.TrimRight]), then TinyGo will fall back
// to escaping it to the heap.
// But if you let the caller do the slice manipulation, then TinyGo is able
// to guarantee that we're talking about the same slice, and is then able to
// keep it on the stack.
//
// Intended usage:
//
//	trimmed := myString[:util.TrimIndexRightByte(myString, ' ')]
func TrimIndexRightByte[S Stringish](s S, c byte) int {
	i := len(s)
	for i > 0 && s[i-1] == c {
		i--
	}
	return i
}

// TrimIndexRightByte is equivalent to [strings.TrimLeft],
// but does not return a slice to not trigger heap escape analysis.
//
// Same explanation as with [TrimIndexRightByte].
//
// Intended usage:
//
//	trimmed := myString[util.TrimIndexLeftByte(myString, ' '):]
func TrimIndexLeftByte[S Stringish](s S, c byte) int {
	i := 0
	for i < len(s) && s[i] == c {
		i++
	}
	return i
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
	rest := s[:TrimIndexRightByte(s, ' ')]
	for len(rest) != 0 {
		rest = rest[TrimIndexLeftByte(rest, ' '):]

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
