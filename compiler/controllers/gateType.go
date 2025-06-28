package controllers

import "main/utils"

type AndGate struct{ Gate }
type NotGate struct{ Gate }
type OrGate struct{ Gate }
type XorGate struct{ Gate }
type Node1N struct{ Gate }
type NodeN1 struct{ Gate }

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

func (p *Node1N) Exec() {
	if len(p.inputs) < 1 || len(p.outputs) < 1 {
		return
	}
	for i := range p.outputs {
		p.outputs[i].Value = utils.GetBit(p.inputs[0].Value, i)
	}
}

func (p *NodeN1) Exec() {
	if len(p.inputs) < 1 || len(p.outputs) < 1 {
		return
	}
	output := &p.outputs[0].Value
	for i := range p.inputs {
		utils.SetBit(output, i, p.inputs[i].Value)
	}
}
