package main

import (
	"fmt"
	"main/compiler/controllers"
)

func main() {

	text :=
		`(ADDER1 | [a,b,Cn] | [s,Cout]){
			[a,b] XOR_0 [x0]
			[x0,Cn] XOR_1 [s]
			[x0,Cn] AND_0 [x1]
			[a,b] AND_1 [x2]
			[x1,x2] OR_0 [Cout]
		}
		(ADDER8 | [a, b, Cn] | [c, Cout]){
			[a] 1N_0 [ a0, a1, a2, a3, a4, a5, a6, a7]
			[b] 1N_1 [ b0, b1, b2, b3, b4, b5, b6, b7]
			[a0, b0, Cn   ] ADDER1_0 [c0, Cout0]
			[a1, b1, Cout0] ADDER1_1 [c1, Cout1]
			[a2, b2, Cout1] ADDER1_2 [c2, Cout2]
			[a3, b3, Cout2] ADDER1_3 [c3, Cout3]
			[a4, b4, Cout3] ADDER1_4 [c4, Cout4]
			[a5, b5, Cout4] ADDER1_5 [c5, Cout5]
			[a6, b6, Cout5] ADDER1_6 [c6, Cout6]
			[a7, b7, Cout6] ADDER1_7 [c7, Cout]
			[c0,c1,c2,c3,c4,c5,c6,c7] N1_0 [c]
		}

		[a,b,c] ADDER8_0 [s,Cout]
		`

	nodeController := controllers.NewGateController()
	nodeController.ProcessText(text)
	nodeController.Set("a", 13)
	nodeController.Set("b", 1)
	nodeController.Set("c", 1)

	nodeController.Exec()
	fmt.Println(nodeController.Get("s"))
}
