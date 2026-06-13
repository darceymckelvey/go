package main

import (
	"fmt"
	"io"
)

// decode is a function that takes a byte slice as an argument
// and returns an error
func decode(b []byte) error {
	return nil
}

func DecodeReader(r io.Reader) error {
	return nil
}

// entry point of the program
func main() {
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
