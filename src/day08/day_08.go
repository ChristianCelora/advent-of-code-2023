package main

import (
	"adventcode/reader"
	"fmt"
	"strings"
)

type Node struct {
	label string
	left  string
	right string
	// step 2
	is_start  bool
	is_finish bool
}

func GetInstructionsFromLine(line string) []byte {
	var instructions []byte
	for i := 0; i < len(line); i++ {
		instructions = append(instructions, line[i])
	}
	return instructions
}

func GetNodeFromLine(line string) Node {
	info := strings.Split(line, "=")
	info_left_right := strings.Split(info[1], ",")
	node_label := strings.TrimSpace(info[0])

	return Node{
		label: node_label,
		left:  strings.TrimSpace(strings.Replace(info_left_right[0], "(", "", -1)),
		right: strings.TrimSpace(strings.Replace(info_left_right[1], ")", "", -1)),
		// step 2
		is_start:  node_label[len(node_label)-1] == 'A',
		is_finish: node_label[len(node_label)-1] == 'Z',
	}
}

// step 1
func WalkNodes(node_map map[string]Node, l_r_instr []byte, start string, finish string) int {
	var steps int
	var node Node
	node = node_map[start]

	for node.label != finish {
		for _, instruction := range l_r_instr {
			steps++
			if instruction == 'L' {
				node = node_map[node.left]
			} else if instruction == 'R' {
				node = node_map[node.right]
			} else {
				panic("instruction not recognized")
			}
			if node.label == finish {
				break
			}
		}
	}

	return steps
}

// step 2
func WalkNodesAsGhost(node_map map[string]Node, l_r_instr []byte, starting_nodes []Node) int {
	var steps int
	var all_nodes_are_finished bool
	nodes := starting_nodes

	for !all_nodes_are_finished {
		for _, instruction := range l_r_instr {
			steps++
			fmt.Printf("%d\n", steps)
			for i, node := range nodes {
				if instruction == 'L' {
					nodes[i] = node_map[node.left]
				} else if instruction == 'R' {
					nodes[i] = node_map[node.right]
				} else {
					panic("instruction not recognized")
				}
			}
			if all_nodes_are_finished = areEndNodes(nodes); all_nodes_are_finished {
				break
			}
		}
	}

	return steps
}

func areEndNodes(nodes []Node) bool {
	for _, node := range nodes {
		if !node.is_finish {
			return false
		}
	}
	return true
}

func main() {
	lines := reader.ReadLines("./day08/data/input1_3.txt")
	nodes_map := make(map[string]Node)

	instructions := GetInstructionsFromLine(lines[0])

	var starting_nodes []Node
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}
		node := GetNodeFromLine(line)
		nodes_map[node.label] = node
		if node.is_start {
			starting_nodes = append(starting_nodes, node)
		}
	}

	// step 1
	//steps := WalkNodes(nodes_map, instructions, "AAA", "ZZZ")
	//fmt.Printf("It takes %d steps to reach %s from %s\n", steps, "AAA", "ZZZ")

	// step 2
	fmt.Printf("starting nodes %v", starting_nodes)
	steps_ghost := WalkNodesAsGhost(nodes_map, instructions, starting_nodes)
	fmt.Printf("It takes %d steps to reach the destination as ghost\n", steps_ghost)
}
