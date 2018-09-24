package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i, c := range b {
		if (c >= 'A' && c < 'N') || (c >= 'a' && c < 'n') {
			b[i] += 13
		} else if (c > 'M' && c <= 'Z') || (c > 'm' && c <= 'z') {
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
