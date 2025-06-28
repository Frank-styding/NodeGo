package main

type Node struct {
	inputs  []*Wire
	outputs []*Wire
	Name    string
}

type AndNode struct{ Node }
type NotNode struct{ Node }
type OrNode struct{ Node }
type XorNode struct{ Node }

type INode interface {
	exec()
	connectOutputs(register WireRegister, wires ...string)
	connectInputs(register WireRegister, wires ...string)
	getName() string
	getInputs() []*Wire
	setName(name string)
}

func (p *Node) connectInputs(register WireRegister, wires ...string) {
	p.inputs = make([]*Wire, len(wires))
	for i := range wires {
		p.inputs[i] = register.wires[wires[i]]
		p.inputs[i].outputsNode = append(p.inputs[i].outputsNode, p)
	}
}

func (p *Node) connectOutputs(register WireRegister, wires ...string) {
	p.outputs = make([]*Wire, len(wires))
	for i := range wires {
		p.outputs[i] = register.wires[wires[i]]
		p.outputs[i].inputNode = p
	}
}

func (p *Node) exec() {

}

func (p *AndNode) exec() {
	if len(p.inputs) < 2 || len(p.outputs) < 1 {
		return
	}
	p.outputs[0].value = p.inputs[0].value & p.inputs[1].value
}

func (p *OrNode) exec() {
	if len(p.inputs) < 2 || len(p.outputs) < 1 {
		return
	}
	p.outputs[0].value = p.inputs[0].value | p.inputs[1].value
}

func (p *NotNode) exec() {
	if len(p.inputs) < 1 || len(p.outputs) < 1 {
		return
	}
	p.outputs[0].value = ^p.inputs[0].value
}

func (p *XorNode) exec() {
	if len(p.inputs) < 2 || len(p.outputs) < 1 {
		return
	}
	p.outputs[0].value = (^p.inputs[0].value & p.inputs[1].value) |
		(p.inputs[0].value & ^p.inputs[1].value)
}

func (p *Node) getName() string {
	return p.Name
}

func (p *Node) getInputs() []*Wire {
	return p.inputs
}

func (p *Node) setName(name string) {
	p.Name = name
}
