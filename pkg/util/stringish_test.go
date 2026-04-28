package util

import "testing"

func TestTrimIndexRightByte(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		cutset byte
		want   string
	}{
		{name: "empty", input: "", cutset: 0, want: ""},
		{name: "no trimming", input: "hello", cutset: ' ', want: "hello"},
		{name: "trim once", input: " hello ", cutset: ' ', want: " hello"},
		{name: "trim multiple", input: "    hello    ", cutset: ' ', want: "    hello"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			index := TrimIndexRightByte(test.input, test.cutset)
			got := string(test.input[:index])
			if got != test.want {
				t.Errorf("wrong result\nwant: %q (len=%d)\ngot:  %q (len=%d)", test.want, len(test.want), got, len(got))
			}
		})
	}
}

func TestTrimIndexLeftByte(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		cutset byte
		want   string
	}{
		{name: "empty", input: "", cutset: 0, want: ""},
		{name: "no trimming", input: "hello", cutset: ' ', want: "hello"},
		{name: "trim once", input: " hello ", cutset: ' ', want: "hello "},
		{name: "trim multiple", input: "    hello    ", cutset: ' ', want: "hello    "},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			index := TrimIndexLeftByte(test.input, test.cutset)
			got := string(test.input[index:])
			if got != test.want {
				t.Errorf("wrong result\nwant: %q (len=%d)\ngot:  %q (len=%d)", test.want, len(test.want), got, len(got))
			}
		})
	}
}

func TestWriteWrapped(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		maxWidth  int
		charWidth int
		want      string
	}{
		{name: "empty", input: "", maxWidth: 40, charWidth: 4, want: ""},
		{
			name:      "no wrapping",
			input:     "hello world",
			maxWidth:  100,
			charWidth: 5,
			want:      "hello world",
		},
		{
			name:      "wrap once on space",
			input:     "hello world",
			maxWidth:  30,
			charWidth: 5,
			want:      "hello\nworld",
		},
		{
			name:      "exact sizing",
			input:     "hello world",
			maxWidth:  25,
			charWidth: 5,
			want:      "hello\nworld",
		},
		{
			name:      "word splitting",
			input:     "helloworld",
			maxWidth:  20,
			charWidth: 5,
			want:      "hell\nowor\nld",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf [1024]byte
			written := WriteWrapped(buf[:], test.input, test.maxWidth, test.charWidth)
			got := string(buf[:written])
			if got != test.want {
				t.Errorf("wrong result\nwant: %q (len=%d)\ngot:  %q (len=%d)", test.want, len(test.want), got, len(got))
			}
		})
	}
}

func TestConcatInto(t *testing.T) {
	tests := []struct {
		name    string
		buf     []byte
		strings []string
		want    string
	}{
		{
			name:    "no strings nil buffer",
			buf:     nil,
			strings: nil,
			want:    "",
		},
		{
			name:    "no strings empty buffer",
			buf:     make([]byte, 0),
			strings: nil,
			want:    "",
		},
		{
			name:    "just big enough buffer",
			buf:     make([]byte, 10),
			strings: []string{"hello", "world"},
			want:    "helloworld",
		},
		{
			name:    "bigger buffer",
			buf:     make([]byte, 14),
			strings: []string{"hello", "world"},
			want:    "helloworld",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			written := ConcatInto(test.buf, test.strings...)
			got := string(test.buf[:written])
			if got != test.want {
				t.Errorf("wrong result\nwant: %q (len=%d)\ngot:  %q (len=%d)", test.want, len(test.want), got, len(got))
			}
			t.Logf("buf: %q", test.buf)
		})
	}
}
