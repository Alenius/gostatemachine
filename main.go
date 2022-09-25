package main

import "fmt"

type Node struct {
	name string
	// list of node names which the state machine can jump to after this node
	// n.b. can be self-referencing
	edges []string
	// when the state machine reaches this node it will
	// perform the action and then respond with the current edge
	action func() (string, error)
}

type StateMachine struct {
	name        string
	currentNode *Node
	nodes       map[string]Node
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

func (sm *StateMachine) Run() error {
	if sm.currentNode == nil {
		return fmt.Errorf("start node must be set and it cannot be nil")
	}

	for {
		fmt.Printf("running action for node: %v \n", sm.currentNode.name)
		nextnodename, err := sm.currentNode.action()
		if err != nil {
			return fmt.Errorf("node action failed with error: %v \n", err)
		}

		nextnode, ok := sm.nodes[nextnodename]
		if !ok {
			return fmt.Errorf("could not find any node with name: %v in the list of nodes \n", nextnodename)
		}

		// if next node does not have a list of edges it means it's a terminal
		// TODO: run this action anyway?
		if len(nextnode.edges) == 0 {
			fmt.Println("state machine is finished")
			return nil
		}

		sm.currentNode = &nextnode
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
		name:  "a",
		edges: []string{"a", "b", "c"},
		action: func() (string, error) {
			return "b", nil
		},
	}
	node2 := Node{
		name:  "b",
		edges: []string{"b", "c"},
		action: func() (string, error) {
			return "c", nil
		},
	}
	node3 := Node{
		name: "c",
	}

	sm.AddNodes(node1, node2, node3)

	// set start node
	sm.currentNode = &node1

	err := sm.Run()
	if err != nil {
		fmt.Printf("state machine encountered an error: %v \n", err)
	}
}
