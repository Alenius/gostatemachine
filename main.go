package main

type Edge struct {
	origin string
	dest   string
}

type Node struct {
	name  string
	edges []Edge
	// when the state machine reaches this node it will
	// perform the action and then respond with the current edge
	action func() Edge
}

type StateMachine struct {
	name  string
	nodes map[string]Node
}

func (sm *StateMachine) AddNodes(nodes ...Node) {
	for _, n := range nodes {
		nodename := n.name

		if sm.nodes == nil {
			sm.nodes = map[string]Node{}
		}

		sm.nodes[nodename] = n
	}
}

func InitStateMachine(name string) StateMachine {
	return StateMachine{
		name:  name,
		nodes: map[string]Node{},
	}
}

func main() {
	sm := InitStateMachine("machine")
	node1 := Node{
		name: "a",
	}
	node2 := Node{
		name: "b",
	}
	node3 := Node{
		name: "c",
	}

	sm.AddNodes(node1, node2, node3)

}
