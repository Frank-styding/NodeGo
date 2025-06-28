package main

import (
	"fmt"
	"main/controllers"
)

func main() {
	wireRegister := controllers.NewWireController()
	nodeRegister := controllers.NewGateController()

	nodeRegister.ProcessText(
		wireRegister,
		`[A]1N_0[B,C]
		[B,C]N1_0[D]`,
	)
	wireRegister.Set("A", 2)
	nodeRegister.Exec(wireRegister)
	fmt.Println(
		wireRegister.Get("D"),
	)
}
