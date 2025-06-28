package interpreter

type NodeInfo struct {
	name        string
	connections []GateConnection
	wires       []string
}

func readNode(text string, i *int) (info NodeInfo, ok bool) {
	if get(text, i) != '(' {
		ok = false
		return
	}
	properties := readArray(text, i, "()", "|")
	contain := readArea(text, i, "{}")
	k := 0
	connections, ok := readConenctions(contain, &k)
	if ok {
		return
	}
	name := properties[0]
	j := 0
	wires := readArray(properties[1], &j, "[]", ",")
	info.connections = connections
	info.name = name
	info.wires = wires
	return
}
