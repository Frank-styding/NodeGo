package controllers

import (
	"main/compiler/interpreter"
	"maps"
	"sort"
)

type GateController struct {
	gates          map[string]IGate
	nodesInfo      map[string]interpreter.NodeInfo
	compliledNodes []IGate
	wireController WireController
}

func NewGateController() *GateController {
	return &GateController{
		gates:          make(map[string]IGate, 200),
		nodesInfo:      make(map[string]interpreter.NodeInfo, 200),
		wireController: *NewWireController(),
	}
}

func (nr *GateController) Set(name string, value int) {
	nr.wireController.Set(name, value)
}

func (nr *GateController) Get(name string) int {
	return nr.wireController.Get(name)
}

func (nr *GateController) Add(name string) {
	nr.wireController.Add(name)
}

func (nr *GateController) addGate(name string, node IGate) {
	if _, exists := nr.gates[name]; exists {
		return
	}

	// Asignar el nombre directamente usando la interfaz
	if setNameNode, ok := node.(interface{ SetName(string) }); ok {
		setNameNode.SetName(name)
	}

	nr.gates[name] = node
}

func (nr *GateController) getGateByTag(name string) IGate {
	switch name {
	case "NOT_":
		return &NotGate{}
	case "AND_":
		return &AndGate{}
	case "OR_":
		return &OrGate{}
	case "XOR_":
		return &XorGate{}
	case "1N_":
		return &Node1N{}
	case "N1_":
		return &NodeN1{}
	default:
		node, ok := nr.nodesInfo[name]
		if ok {
			aux := &NodeGate{
				Name:           node.Name,
				InputsName:     node.InputsName,
				OutputsName:    node.OutputsName,
				gateController: NewGateController(),
			}
			for _, i := range nr.nodesInfo {
				if node.Name != i.Name {
					aux.gateController.nodesInfo[i.Name] = i
				}
			}
			aux.gateController.CreateConnections(node.Connections...)
			return aux
		}
	}
	return &Gate{}
}

func (nr *GateController) connect(name string, inputs []string, outputs []string) {
	node, exists := nr.gates[name]
	if !exists {
		nodeName := name[0 : len(name)-1]
		node = nr.getGateByTag(nodeName)
		nr.addGate(name, node)
	}
	node.ConnectInputs(&nr.wireController, inputs...)
	node.ConnectOutputs(&nr.wireController, outputs...)
}

func (nr *GateController) createNodes(infos ...interpreter.NodeInfo) {
	for _, info := range infos {
		if _, ok := nr.nodesInfo[info.Name]; !ok {
			nr.nodesInfo[info.Name] = info
		}
	}
}

func (nr *GateController) CreateConnections(conections ...interpreter.GateConnection) {
	for _, conn := range conections {
		nr.wireController.Add(conn.Inputs...)
		nr.wireController.Add(conn.Outputs...)
		nr.connect(conn.NodeName, conn.Inputs, conn.Outputs)
	}
}

func (nr *GateController) ProcessText(text string) {
	inter := interpreter.Interpreter{}
	inter.ProcessText(text)
	nr.createNodes(inter.NodesInfo...)
	nr.CreateConnections(inter.Connections...)
}

// calcDistance calcula la distancia máxima desde un nodo hasta los nodos de entrada
func calcDistance(node IGate, depth int, visited map[string]bool) int {
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

func (nr *GateController) Exec() {
	if len(nr.compliledNodes) > 0 {
		for _, node := range nr.compliledNodes {
			node.Exec()
		}
		return
	}
	// Calcular distancias de todos los nodos
	nodeDistances := make(map[string]int, len(nr.gates))
	for nodeName, node := range nr.gates {
		nodeDistances[nodeName] = calcDistance(node, 0, nil)
	}

	// Crear una lista de nodos ordenados por distancia (de mayor a menor)
	type nodeWithDistance struct {
		name     string
		node     IGate
		distance int
	}

	nodeList := make([]nodeWithDistance, 0, len(nr.gates))
	for nodeName, node := range nr.gates {
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

	// Ejecutar los nodos en el orden ordenado
	for _, nodeWithDist := range nodeList {
		nodeWithDist.node.Exec()
	}
}
