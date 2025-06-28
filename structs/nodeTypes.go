package structs

type Node1N struct {
	Gate
}

func (p *Node1N) Exec() {
	inputs := p.GetInputs()
	outputs := p.GetOutputs()
	if len(inputs) < 1 || len(outputs) < 1 {
		return
	}
	for i := range outputs {
		outputs[i].Value = getBit(inputs[0].Value, i)
	}
}

type NodeN1 struct {
	Gate
}

func (p *NodeN1) Exec() {
	inputs := p.GetInputs()
	outputs := p.GetOutputs()
	if len(inputs) < 1 || len(outputs) < 1 {
		return
	}
	output := &outputs[0].Value
	for i := range inputs {
		setBit(output, i, inputs[i].Value)
	}
}

// Funciones auxiliares para manipulación de bits
func getBit(value int, bitIndex int) int {
	return (value >> bitIndex) & 1
}

func setBit(value *int, bitIndex int, bitValue int) {
	if bitValue == 1 {
		*value |= (1 << bitIndex)
	} else {
		*value &= ^(1 << bitIndex)
	}
}
