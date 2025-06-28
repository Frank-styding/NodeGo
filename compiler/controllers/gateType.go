package controllers

import "main/utils"

type AndGate struct{ Gate }
type NotGate struct{ Gate }
type OrGate struct{ Gate }
type XorGate struct{ Gate }
type Node1N struct{ Gate }
type NodeN1 struct{ Gate }

func (p *AndGate) Exec() {
	if len(p.Inputs) < 2 || len(p.Outputs) < 1 {
		return
	}
	p.Outputs[0].Value = p.Inputs[0].Value & p.Inputs[1].Value
}

func (p *OrGate) Exec() {
	if len(p.Inputs) < 2 || len(p.Outputs) < 1 {
		return
	}
	p.Outputs[0].Value = p.Inputs[0].Value | p.Inputs[1].Value
}

func (p *NotGate) Exec() {
	if len(p.Inputs) < 1 || len(p.Outputs) < 1 {
		return
	}
	p.Outputs[0].Value = ^p.Inputs[0].Value
}

func (p *XorGate) Exec() {
	if len(p.Inputs) < 2 || len(p.Outputs) < 1 {
		return
	}
	p.Outputs[0].Value = (^p.Inputs[0].Value & p.Inputs[1].Value) |
		(p.Inputs[0].Value & ^p.Inputs[1].Value)
}

func (p *Node1N) Exec() {
	if len(p.Inputs) < 1 || len(p.Outputs) < 1 {
		return
	}
	for i := len(p.Outputs) - 1; i >= 0; i-- {
		p.Outputs[i].Value = utils.GetBit(p.Inputs[0].Value, i)
	}
}

func (p *NodeN1) Exec() {
	if len(p.Inputs) < 1 || len(p.Outputs) < 1 {
		return
	}

	output := &p.Outputs[0].Value

	for i := len(p.Inputs) - 1; i >= 0; i-- {
		utils.SetBit(output, i, p.Inputs[i].Value)
	}
}
