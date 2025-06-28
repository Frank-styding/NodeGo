package controllers

import (
	"main/structs"
)

type WireController struct {
	Wires map[string]*structs.Wire
}

func NewWireController() *WireController {
	return &WireController{
		Wires: make(map[string]*structs.Wire, 200),
	}
}

func (wr *WireController) Add(names ...string) {
	for _, name := range names {
		if wr.Wires[name] == nil {
			wr.Wires[name] = &structs.Wire{
				Value: 0,
				Name:  name,
			}
		}
	}
}

func (wr *WireController) Set(name string, value int) {
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

func (wr *WireController) GetWire(name string) *structs.Wire {
	return wr.Wires[name]
}
