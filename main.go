package main

import "main/compiler/controllers"

func main() {

	nodeController := controllers.GateController{}
	/* 	i := 0

	   	text := `(hola | [a,b,c] ){123}`
	   	text = strings.ReplaceAll(text, " ", "")
	   	interpreter.ReadNode(text, &i) */
	/* 	wireRegister := controllers.NewWireController()
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
	   	) */
}
