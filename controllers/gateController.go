package controllers

import (
	"main/structs"
	"maps"
	"sort"
)

type GateController struct {
	nodes          map[string]structs.IGate
	compliledNodes []structs.IGate
}

func NewGateController() *GateController {
	return &GateController{
		nodes: make(map[string]structs.IGate, 200),
	}
}

func (nr *GateController) add(name string, node structs.IGate) {
	if _, exists := nr.nodes[name]; exists {
		return
	}

	// Asignar el nombre directamente usando la interfaz
	if setNameNode, ok := node.(interface{ SetName(string) }); ok {
		setNameNode.SetName(name)
	}

	nr.nodes[name] = node
}

func (nr *GateController) getGateByTag(name string) structs.IGate {
	switch name {
	case "NOT":
		return &structs.NotGate{}
	case "AND":
		return &structs.AndGate{}
	case "OR":
		return &structs.OrGate{}
	case "XOR":
		return &structs.XorGate{}
	case "1N_":
		return &structs.Node1N{}
	case "N1_":
		return &structs.NodeN1{}
	}
	return &structs.Gate{}
}

func (nr *GateController) connect(wireRegister *WireController, name string, inputs []string, outputs []string) {
	node, exists := nr.nodes[name]

	if !exists {
		nodeName := name[0 : len(name)-1]
		node = nr.getGateByTag(nodeName)
		nr.add(name, node)
	}

	node.ConnectInputs(wireRegister, inputs...)
	node.ConnectOutputs(wireRegister, outputs...)
}

func (nr *GateController) ProcessText(wireRegister *WireController, text string) {
	connections := processText(text)

	for _, conn := range connections {
		wireRegister.Add(conn.Inputs...)
		wireRegister.Add(conn.Outputs...)
		nr.connect(wireRegister, conn.NodeName, conn.Inputs, conn.Outputs)
	}
}

// calcDistance calcula la distancia máxima desde un nodo hasta los nodos de entrada
func calcDistance(node structs.IGate, depth int, visited map[string]bool) int {
	if visited == nil {
		visited = make(map[string]bool)
	}

	nodeName := node.GetName()
	if visited[nodeName] {
		return 0 // Evitar ciclos
	}
	visited[nodeName] = true

	maxDepth := depth
	inputs := node.GetInputs()

	for _, wire := range inputs {
		if wire.InputNode != nil {
			newVisited := make(map[string]bool)
			maps.Copy(newVisited, visited)
			wireDepth := calcDistance(wire.InputNode, depth+1, newVisited)
			if wireDepth > maxDepth {
				maxDepth = wireDepth
			}
		}
	}

	return maxDepth
}

func (nr *GateController) Exec(wireRegister *WireController) {
	if len(nr.compliledNodes) > 0 {
		for _, node := range nr.compliledNodes {
			node.Exec()
		}
		return
	}
	// Calcular distancias de todos los nodos
	nodeDistances := make(map[string]int, len(nr.nodes))
	for nodeName, node := range nr.nodes {
		nodeDistances[nodeName] = calcDistance(node, 0, nil)
	}

	// Crear una lista de nodos ordenados por distancia (de mayor a menor)
	type nodeWithDistance struct {
		name     string
		node     structs.IGate
		distance int
	}

	nodeList := make([]nodeWithDistance, 0, len(nr.nodes))
	for nodeName, node := range nr.nodes {
		nodeList = append(nodeList, nodeWithDistance{
			name:     nodeName,
			node:     node,
			distance: nodeDistances[nodeName],
		})
	}

	// Ordenar por distancia de mayor a menor (nodos más profundos primero)
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].distance < nodeList[j].distance
	})

	for _, nodeWithDist := range nodeList {
		nr.compliledNodes = append(nr.compliledNodes, nodeWithDist.node)
	}

	//fmt.Println(nodeList[0])
	// Ejecutar los nodos en el orden ordenado
	for _, nodeWithDist := range nodeList {
		nodeWithDist.node.Exec()
	}
}
