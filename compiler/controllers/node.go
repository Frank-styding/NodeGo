package controllers

type NodeGate struct {
	Gate
	Name           string
	InputsName     []string
	OutputsName    []string
	gateController *GateController
}

func (p *NodeGate) NewNodeGate() NodeGate {
	return NodeGate{
		gateController: NewGateController(),
	}
}
func (p *NodeGate) Exec() {
	for i, wire := range p.GetInputs() {
		if i < len(p.InputsName) {
			p.gateController.Set(p.InputsName[i], wire.Value)
		}
	}

	p.gateController.Exec()

	for i, wire := range p.GetOutputs() {
		if i < len(p.OutputsName) {
			wire.Value = p.gateController.Get(p.OutputsName[i])
		}
	}
}
