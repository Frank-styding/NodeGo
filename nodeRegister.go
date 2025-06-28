package main

import (
	"maps"
	"sort"
)

type NodeRegister struct {
	nodes map[string]INode
}

func NewNodeRegister() *NodeRegister {
	return &NodeRegister{
		nodes: make(map[string]INode, 200),
	}
}

func (nr *NodeRegister) add(name string, node INode) {
	if _, exists := nr.nodes[name]; exists {
		return
	}

	// Asignar el nombre directamente usando la interfaz
	if setNameNode, ok := node.(interface{ setName(string) }); ok {
		setNameNode.setName(name)
	}

	nr.nodes[name] = node
}

func (nr *NodeRegister) getNodeByTag(name string) INode {
	switch name {
	case "NOT":
		return &NotNode{}
	case "AND":
		return &AndNode{}
	case "OR":
		return &OrNode{}
	case "XOR":
		return &XorNode{}
	}
	return &Node{}
}
func (nr *NodeRegister) connect(wireRegister *WireRegister, name string, inputs []string, outputs []string) {
	node, exists := nr.nodes[name]

	if !exists {
		nodeName := name[0 : len(name)-1]
		node = nr.getNodeByTag(nodeName)
		nr.add(name, node)
	}

	node.connectInputs(*wireRegister, inputs...)
	node.connectOutputs(*wireRegister, outputs...)
}

func (nr *NodeRegister) processText(wireRegister *WireRegister, text string) {
	connections := processText(text)

	for _, conn := range connections {
		wireRegister.add(conn.inputs...)
		wireRegister.add(conn.outputs...)
		nr.connect(wireRegister, conn.nodeName, conn.inputs, conn.outputs)
	}
}

// calcDistance calcula la distancia máxima desde un nodo hasta los nodos de entrada
func calcDistance(node INode, depth int, visited map[string]bool) int {
	if visited == nil {
		visited = make(map[string]bool)
	}

	nodeName := node.getName()
	if visited[nodeName] {
		return 0 // Evitar ciclos
	}
	visited[nodeName] = true

	maxDepth := depth
	inputs := node.getInputs()

	for _, wire := range inputs {
		if wire.inputNode != nil {
			newVisited := make(map[string]bool)
			maps.Copy(newVisited, visited)
			wireDepth := calcDistance(wire.inputNode, depth+1, newVisited)
			if wireDepth > maxDepth {
				maxDepth = wireDepth
			}
		}
	}

	return maxDepth
}

func (nr *NodeRegister) exec(wireRegister *WireRegister) {
	// Calcular distancias de todos los nodos
	nodeDistances := make(map[string]int, len(nr.nodes))
	for nodeName, node := range nr.nodes {
		nodeDistances[nodeName] = calcDistance(node, 0, nil)
	}

	// Crear una lista de nodos ordenados por distancia (de mayor a menor)
	type nodeWithDistance struct {
		name     string
		node     INode
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

	// Ejecutar los nodos en el orden ordenado
	for _, nodeWithDist := range nodeList {
		nodeWithDist.node.exec()
	}
}
