package structs

type IWireController interface {
	GetWire(name string) *Wire
}

type Gate struct {
	inputs  []*Wire
	outputs []*Wire
	Name    string
}

type AndGate struct{ Gate }
type NotGate struct{ Gate }
type OrGate struct{ Gate }
type XorGate struct{ Gate }

type IGate interface {
	Exec()
	ConnectOutputs(register IWireController, wires ...string)
	ConnectInputs(register IWireController, wires ...string)
	GetName() string
	GetInputs() []*Wire
	GetOutputs() []*Wire
	SetName(name string)
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

func (p *Gate) Exec() {

}

func (p *AndGate) Exec() {
	if len(p.inputs) < 2 || len(p.outputs) < 1 {
		return
	}
	p.outputs[0].Value = p.inputs[0].Value & p.inputs[1].Value
}

func (p *OrGate) Exec() {
	if len(p.inputs) < 2 || len(p.outputs) < 1 {
		return
	}
	p.outputs[0].Value = p.inputs[0].Value | p.inputs[1].Value
}

func (p *NotGate) Exec() {
	if len(p.inputs) < 1 || len(p.outputs) < 1 {
		return
	}
	p.outputs[0].Value = ^p.inputs[0].Value
}

func (p *XorGate) Exec() {
	if len(p.inputs) < 2 || len(p.outputs) < 1 {
		return
	}
	p.outputs[0].Value = (^p.inputs[0].Value & p.inputs[1].Value) |
		(p.inputs[0].Value & ^p.inputs[1].Value)
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
