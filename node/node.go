package node

import (
	"main/structs"
	"slices"
)

type NodeGate struct {
	structs.Gate
	wiresValid     []string
	connections    string
	wireController *WireController
	gateController *GateController
}

// WireController y GateController se definen localmente para evitar el ciclo
type WireController struct {
	// Implementación básica - puedes expandir según necesites
	wires map[string]int
}

func NewWireController() *WireController {
	return &WireController{
		wires: make(map[string]int),
	}
}

func (wc *WireController) Add(names ...string) {
	for _, name := range names {
		if _, exists := wc.wires[name]; !exists {
			wc.wires[name] = 0
		}
	}
}

func (wc *WireController) Set(name string, value int) {
	wc.wires[name] = value
}

func (wc *WireController) Get(name string) int {
	return wc.wires[name]
}

type GateController struct {
	nodes map[string]structs.IGate
}

func NewGateController() *GateController {
	return &GateController{
		nodes: make(map[string]structs.IGate),
	}
}

func (gc *GateController) ProcessText(wireController *WireController, text string) {
	// Implementación básica - puedes expandir según necesites
	// Por ahora, solo un placeholder
}

func (p *NodeGate) NewNodeGate() NodeGate {
	return NodeGate{
		wireController: NewWireController(),
		gateController: NewGateController(),
	}
}

func (p *NodeGate) Exec() {
	for _, wire := range p.GetInputs() {
		if slices.Contains(p.wiresValid, wire.Name) {
			p.wireController.Add("E_" + wire.Name)
			p.wireController.Set("E_"+wire.Name, wire.Value)
		}
	}

	p.gateController.ProcessText(p.wireController, p.connections)

	for _, wire := range p.GetOutputs() {
		if slices.Contains(p.wiresValid, wire.Name) {
			wire.Value = p.wireController.Get("E_" + wire.Name)
		}
	}
}
