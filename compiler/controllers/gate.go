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
	Inputs  []*Wire
	Outputs []*Wire
	Name    string
}

func (p *Gate) Exec() {

}

func (p *Gate) GetName() string {
	return p.Name
}

func (p *Gate) GetInputs() []*Wire {
	return p.Inputs
}

func (p *Gate) GetOutputs() []*Wire {
	return p.Outputs
}

func (p *Gate) SetName(name string) {
	p.Name = name
}

type IWireController interface {
	GetWire(name string) *Wire
}

func (p *Gate) ConnectInputs(register IWireController, wires ...string) {
	p.Inputs = make([]*Wire, len(wires))
	for i := range wires {
		p.Inputs[i] = register.GetWire(wires[i])
		p.Inputs[i].OutputsNode = append(p.Inputs[i].OutputsNode, p)
	}
}

func (p *Gate) ConnectOutputs(register IWireController, wires ...string) {
	p.Outputs = make([]*Wire, len(wires))
	for i := range wires {
		p.Outputs[i] = register.GetWire(wires[i])
		p.Outputs[i].InputNode = p
	}
}
