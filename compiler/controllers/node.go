package controllers

import "slices"

type NodeGate struct {
	Gate
	wiresValid     []string
	connections    string
	wireController *WireController
	gateController *GateController
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