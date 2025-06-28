package interpreter

import (
	"strings"
)

type Interpreter struct {
	NodesInfo   []NodeInfo
	Connections []GateConnection
}

func (nr *Interpreter) ProcessText(text string) {
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ReplaceAll(text, "\t", "")
	text = deleteComments(text)

	for i := 0; i < len(text); i++ {
		info, ok := readNode(text, &i)
		if ok {
			nr.NodesInfo = append(nr.NodesInfo, info)
		}
		connections, ok := readConenctions(text, &i)
		if ok {
			nr.Connections = append(nr.Connections, connections...)
		}
	}
}
