package interpreter

type NodeInfo struct {
	Name        string
	Connections []GateConnection
	InputsName  []string
	OutputsName []string
}

func readNode(text string, i *int) (info NodeInfo, ok bool) {
	for get(text, i) == '\n' {
		advance(i)
	}
	if get(text, i) != '(' {
		ok = false
		return
	}
	properties := readArray(text, i, "()", "|")
	contain := readArea(text, i, "{}")
	k := 0
	connections, ok := readConenctions(contain, &k)
	if !ok {
		return
	}
	name := properties[0]
	j := 0
	inputsName := readArray(properties[1], &j, "[]", ",")
	j = 0
	outputsName := readArray(properties[2], &j, "[]", ",")

	info.Connections = connections
	info.Name = name + "_"
	info.InputsName = inputsName
	info.OutputsName = outputsName
	return
}
