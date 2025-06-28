package controllers

type WireController struct {
	Wires map[string]*Wire
}

func NewWireController() *WireController {
	return &WireController{
		Wires: make(map[string]*Wire, 200),
	}
}

func (wr *WireController) Add(names ...string) {
	for _, name := range names {
		if wr.Wires[name] == nil {
			wr.Wires[name] = &Wire{
				Value: 0,
				Name:  name,
			}
		}
	}
}

func (wr *WireController) Set(name string, value int) {
	wr.Add(name)
	if wire, exists := wr.Wires[name]; exists {
		wire.Value = value
	}
}

func (wr *WireController) Get(name string) int {
	if wire, exists := wr.Wires[name]; exists {
		return wire.Value
	}
	return 0
}

func (wr *WireController) GetWire(name string) *Wire {
	return wr.Wires[name]
}
