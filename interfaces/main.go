package main

import (
	"bytes"
	"fmt"
	"io"
)

/*
/ 1 byte = 8 bits of information, so a byte slice can represent a sequence of bytes,
/ which can be used to store data such as text, images, or other types of information.
/ The decode function is intended to process the byte slice
/ and extract meaningful information from it,
/ but in this implementation, it simply returns nil,
/ indicating that no error occurred during the decoding process.
*/
func decode(_ []byte) error {
	return nil
}

func DecodeReader(r io.Reader) error {
	// read all data from the reader and pass to decode
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return decode(b)
}

// entry point of the program
func main() {

	data := []byte("Hello, World!")

	DecodeReader(bytes.NewReader(data))
	// io.Reader
	// io.Writer

	DoAttack(StrongAttacker{})
}

// In Go, an interface is a custom type that defines a set of
// method signatures but does not provide their implementation
type Attacker interface {
	Attack() error
}

// StrongAttacker is a struct that represents a strong attacker
type StrongAttacker struct{}

// StrongAttacker implements the Attacker interface by providing
// an implementation of the Attack method
func (sa StrongAttacker) Attack() error {
	fmt.Println("What a strong attack!")

	return nil
}

// DoAttack is a function that takes an Attacker interface
// as an argument
func DoAttack(a Attacker) error {
	return a.Attack()
}
