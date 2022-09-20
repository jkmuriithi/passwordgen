package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const printable = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

// Takes in a positive integer length as a command line argument and returns a
// randomly generated password of that length.
func main() {
	// Read command-line args
	length, err := strconv.Atoi(os.Args[1])
	if err != nil  {
		panic("Usage: ./passwordgen password_length")
	}
	if length <= 0 {
		panic("Error: password_length must be a positive integer")
	}

	// Build password
	pw := strings.Builder{}

	chars := []rune(printable)
	n := len(chars)
	for i := 0; i < length; i++ {
		bigIdx, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
		if err != nil {
			panic(err)
		}

		pw.WriteRune(chars[bigIdx.Uint64()])
	}

	fmt.Println(pw.String())
}
