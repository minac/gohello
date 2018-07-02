package main

import (
//	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// can take any rotation, e.g. 13
func rot(orig int, rot int) int {
	rotated := 0
	switch {
	// Special characters
	case orig >= 32 && orig <= 63:
		rotated = orig
	// Simple rotation without wrap around (e.g. 90 - 78 + 1 = 13)
	case orig >= 65 && orig < (90-rot+1):
		rotated = orig + rot
	// Rotation of capitals with wrap around
	// 64 because first letter 'A' counts as well
	// 90 - orig is the difference to be rotated
	case orig >= 78 && orig <= 90:
		rotated = 64 + rot - (90 - orig)
	// Simple rotation without wrap around (e.g. 122 - 110 + 1 = 13)
	case orig >= 97 && orig < (122 - rot + 1):
		rotated = orig + rot
	case orig >= 110 && orig <= 122:
		// 96 because first letter 'a' counts as well
		rotated = 96 + rot - (122 - orig)
	default:
		rotated = 0
	}
	//fmt.Println(string(orig), ">", string(rotated))
	return rotated
}

func (rotr *rot13Reader) Read(b []byte) (int, error) {
	n, err := rotr.r.Read(b)
	var in int
	var out uint8
	for i := 0; i < n; i++ {
		in = int(b[i])
		out = uint8(rot(in, 13))
		b[i] = out
		//fmt.Println("in", in, "out", out)
	}
	return n, err
}

func main() {
	
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
