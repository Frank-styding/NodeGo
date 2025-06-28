package main

import (
	"fmt"
)

func main() {
	wireRegister := NewWireRegister()
	nodeRegister := NewNodeRegister()
	nodeRegister.processText(
		wireRegister,
		`[A,B]XOR0[C]`,
	)
	wireRegister.set("A", 0)
	wireRegister.set("B", 0)
	nodeRegister.exec(wireRegister)
	fmt.Printf("Result E: %v\n", wireRegister.get("C"))
}
