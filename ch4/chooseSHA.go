package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		if os.Args[2] == "384" {
			fmt.Printf("%x\n", sha512.Sum384([]byte(os.Args[1])))
		} else if os.Args[2] == "512" {
			fmt.Printf("%x\n", sha512.Sum512([]byte(os.Args[1])))
		}
	} else {
		fmt.Printf("%x\n", sha256.Sum256([]byte(os.Args[1])))
	}
}
