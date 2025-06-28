package interpreter

type GateConnection struct {
	Inputs, Outputs []string
	NodeName        string
}

func readConenctions(text string, i *int) (connections []GateConnection, ok bool) {

	if get(text, i) != '[' {
		ok = false
		return
	}

	for {
		if endLine(text, i) {
			break
		}
		inputs := readArray(text, i, "[]", ",")
		nodeName := readString(text, i)
		outputs := readArray(text, i, "[]", ",")
		if endLine(text, i) {
			connections = append(connections,
				GateConnection{
					Inputs:   inputs,
					Outputs:  outputs,
					NodeName: nodeName,
				})
			break
		}
		if get(text, i) == '\n' {
			connections = append(connections,
				GateConnection{
					Inputs:   inputs,
					Outputs:  outputs,
					NodeName: nodeName,
				})
			advance(i)
		}
		advance(i)
	}

	ok = true
	return
}
