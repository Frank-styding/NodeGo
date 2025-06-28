package controllers

type IGate interface {
	Exec()
	ConnectOutputs(register IWireController, wires ...string)
	ConnectInputs(register IWireController, wires ...string)
	GetName() string
	GetInputs() []*Wire
	GetOutputs() []*Wire
	SetName(name string)
}

type Gate struct {
	inputs      []*Wire
	outputs     []*Wire
	Name        string
	inputsNode  []IGate
	outputsNode []IGate
}

func (p *Gate) Exec() {

}

func (p *Gate) GetName() string {
	return p.Name
}

func (p *Gate) GetInputs() []*Wire {
	return p.inputs
}

func (p *Gate) GetOutputs() []*Wire {
	return p.outputs
}

func (p *Gate) SetName(name string) {
	p.Name = name
}

type IWireController interface {
	GetWire(name string) *Wire
}

func (p *Gate) ConnectInputs(register IWireController, wires ...string) {
	p.inputs = make([]*Wire, len(wires))
	for i := range wires {
		p.inputs[i] = register.GetWire(wires[i])
		p.inputs[i].OutputsNode = append(p.inputs[i].OutputsNode, p)
	}
}

func (p *Gate) ConnectOutputs(register IWireController, wires ...string) {
	p.outputs = make([]*Wire, len(wires))
	for i := range wires {
		p.outputs[i] = register.GetWire(wires[i])
		p.outputs[i].InputNode = p
	}
}