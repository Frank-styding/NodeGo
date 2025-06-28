package main

type Wire struct {
	value       int
	name        string
	inputNode   *Node
	outputsNode []*Node
}

type WireRegister struct {
	wires map[string]*Wire
}

func NewWireRegister() *WireRegister {
	return &WireRegister{
		wires: make(map[string]*Wire, 200),
	}
}

func (wr *WireRegister) add(names ...string) {
	for _, name := range names {
		if wr.wires[name] == nil {
			wr.wires[name] = &Wire{
				value: 0,
				name:  name,
			}
		}
	}
}

func (wr *WireRegister) set(name string, value int) {
	if wire, exists := wr.wires[name]; exists {
		wire.value = value
	}
}

func (wr *WireRegister) get(name string) int {
	if wire, exists := wr.wires[name]; exists {
		return wire.value
	}
	return 0
}
