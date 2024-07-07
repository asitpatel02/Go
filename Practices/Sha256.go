package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Print("Enter a string: ")
	var input string
	fmt.Scanln(&input)

	hash := sha256.Sum256([]byte(input))
	fmt.Printf("SHA256 hash of '%s': %x\n", input, hash)
}
