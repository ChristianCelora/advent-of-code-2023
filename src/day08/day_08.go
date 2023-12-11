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
	is_dead   bool // left and right are this node
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
	left := strings.TrimSpace(strings.Replace(info_left_right[0], "(", "", -1))
	right := strings.TrimSpace(strings.Replace(info_left_right[1], ")", "", -1))

	return Node{
		label: node_label,
		left:  left,
		right: right,
		// step 2
		is_start:  node_label[len(node_label)-1] == 'A',
		is_finish: node_label[len(node_label)-1] == 'Z',
		is_dead:   left == node_label && right == node_label,
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

func GetEndNodes(node_map map[string]Node, l_r_instr []byte, start string) []int {
	var steps int
	var node Node
	var end_node_steps []int
	node = node_map[start]
	first_run := true
	node_visited_map := make(map[string]bool)
	node_visited_map[node.label] = true

	// finish when find a dead node or we are back at the starting node
	for (first_run || node.label != start) && !node.is_dead {
		first_run = false
		for i := 0; i < len(l_r_instr); i++ {
			instruction := l_r_instr[i]
			steps++
			if instruction == 'L' {
				node = node_map[node.left]
			} else if instruction == 'R' {
				node = node_map[node.right]
			} else {
				panic("instruction not recognized")
			}
			if node.is_finish && !node_visited_map[node.label] {
				end_node_steps = append(end_node_steps, steps)
			}
			node_visited_map[node.label] = true

			if node.label == start || node.is_dead {
				fmt.Printf("break")
				break
			}
		}
		// horrible solution to break graph cycles
		if steps > 100000 {
			break
		}
	}
	return end_node_steps
}

func CalcStepsCombinationLCM(slices_steps [][]int) []int {
	var steps_combinations []int
	var steps_combinations_new []int
	var steps_lcm int

	steps_combinations = slices_steps[0]
	for _, node_ending_steps := range slices_steps[1:] {
		steps_combinations_new = make([]int, 0)
		for i := 0; i < len(node_ending_steps); i++ {
			steps := node_ending_steps[i]
			for j := 0; j < len(steps_combinations); j++ {
				v := steps_combinations[j]
				steps_lcm = LCM(v, steps)
				steps_combinations_new = append(steps_combinations_new, steps_lcm)
			}
		}
		steps_combinations = steps_combinations_new
	}

	return steps_combinations
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	lines := reader.ReadLines("./day08/data/input1_3.txt")
	nodes_map := make(map[string]Node)

	instructions := GetInstructionsFromLine(lines[0])

	var starting_nodes []Node
	var finish_nodes []Node
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}
		node := GetNodeFromLine(line)
		nodes_map[node.label] = node
		if node.is_start {
			starting_nodes = append(starting_nodes, node)
		}
		if node.is_finish {
			finish_nodes = append(finish_nodes, node)
		}
	}

	// step 1
	steps := WalkNodes(nodes_map, instructions, "AAA", "ZZZ")
	fmt.Printf("It takes %d steps to reach %s from %s\n", steps, "AAA", "ZZZ")

	// step 2
	// step 2. solution 1. too slow for input
	// steps_ghost := WalkNodesAsGhost(nodes_map, instructions, starting_nodes)
	// fmt.Printf("It takes %d steps to reach the destination as ghost\n", steps_ghost)

	// step 2. solution 2
	var ending_node_steps [][]int
	for _, start := range starting_nodes {
		finish_nodes_steps := GetEndNodes(nodes_map, instructions, start.label)
		ending_node_steps = append(ending_node_steps, finish_nodes_steps)
	}

	min_steps_to_reach_finish_nodes := int(^uint(0) >> 1)
	steps_combinations := CalcStepsCombinationLCM(ending_node_steps)
	for _, steps := range steps_combinations {
		if steps < min_steps_to_reach_finish_nodes {
			min_steps_to_reach_finish_nodes = steps
		}
	}
	fmt.Printf("It takes %d steps to reach the destinations as ghost\n", min_steps_to_reach_finish_nodes)
}
