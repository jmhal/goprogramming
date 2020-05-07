package main

import (
	"crypto/sha256"
	"fmt"
)

const (
	first = 1 << iota
	second
	third
	fourth
	fifth
	sixth
	seventh
	eighth
)

var bits = [8]uint8{first, second, third, fourth, fifth, sixth, seventh, eighth}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%d\n", bitCount(c1, c2))
}

func bitCount(h1 [32]uint8, h2 [32]uint8) uint32 {
	var count uint32
	for i := 0; i < 32; i++ {
		a := h1[i]
		b := h2[i]
		if a == b {
			continue
		} else {
			x := a ^ b
			for _, bit := range bits {
				if x&bit == bit {
					count++
				}
			}
		}
	}
	return count
}
