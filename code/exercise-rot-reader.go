package main

import (
	"strings"
	"io"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func rot13byte(b byte) byte {
	var beg byte

	if b >= 'A' && b <= 'Z' {
		beg = 'A'
	} else if b >= 'a' && b <= 'z' {
		beg = 'a'
	} else {
		return b
	}

	return (((b - beg) + 13) % 26) + beg
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	size, err := rot13.r.Read(b)

	for i, v := range b {
		b[i] = rot13byte(v)
	}

	return size, err
}

func exerciseRotReader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}