package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type programSpec struct {
	name     string
	weight   int
	children []string
}

func FindRoot(input []programSpec) string {
	appearsAsChild := make(map[string]bool)
	for _, spec := range input {
		for _, child := range spec.children {
			appearsAsChild[child] = true
		}
	}
	for _, spec := range input {
		if !appearsAsChild[spec.name] {
			return spec.name
		}
	}
	panic("could not find")
}

type program struct {
	name           string
	weight         int
	towerWeight    int
	parent         *program
	children       []*program
}

func (p *program) isBalanced() bool {
	if p.children == nil {
		return true
	}
	childWeightCounts := make(map[int]int)
	for _, child := range p.children {
		childWeightCounts[child.towerWeight]++
	}
	return len(childWeightCounts) == 1
}

func FindProgramToAdjust(programs []programSpec) (string, int) {
	programsByName := make(map[string]programSpec)
	for _, program := range programs {
		programsByName[program.name] = program
	}

	root := FindRoot(programs)

	rootNode := buildNode(programsByName, root)

	nodesToProcess := []*program {rootNode}

	var targetNode *program
	for {
		n := len(nodesToProcess) - 1
		node := nodesToProcess[n]
		nodesToProcess = nodesToProcess[:n]

		childWeightCounts := make(map[int]int)
		for _, child := range node.children {
			childWeightCounts[child.towerWeight]++
		}

		if len(childWeightCounts) == 1 {
			targetNode = node
			break
		}
		for _, child := range node.children {
			if childWeightCounts[child.towerWeight] == 1 {
				nodesToProcess = append(nodesToProcess, child)
			}
		}
	}

	parentOfTarget := targetNode.parent
	var targetTowerWeight int
	for _, child := range parentOfTarget.children {
		if child.towerWeight != targetNode.towerWeight {
			targetTowerWeight = child.towerWeight
			break
		}
	}
	weightDiff := targetTowerWeight - targetNode.towerWeight

	return targetNode.name, targetNode.weight + weightDiff
}

func buildNode(programsByName map[string]programSpec, currentNode string) *program {
	p := &program{
		name: currentNode,
		weight: programsByName[currentNode].weight,
		towerWeight: programsByName[currentNode].weight,
	}
	for _, childName := range programsByName[currentNode].children {
		child := buildNode(programsByName, childName)
		child.parent = p
		p.children = append(p.children, child)
		p.towerWeight += child.towerWeight
	}
	return p
}

var lineRexexp = regexp.MustCompile(`([a-z]+) \(([\d]+)\)(?: -> (.*))?`)

func main() {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []programSpec
	for scanner.Scan() {
		line := scanner.Text()
		lineMatches := lineRexexp.FindStringSubmatch(line)
		name := lineMatches[1]
		weight, _ := strconv.Atoi(lineMatches[2])
		var children []string
		if lineMatches[3] != "" {
			children = strings.Split(lineMatches[3], ", ")
		}
		input = append(input, programSpec{name: name, weight: weight, children: children})
	}

	root := FindRoot(input)
	fmt.Printf("Part 1: %s\n", root)

	node, targetWeight := FindProgramToAdjust(input)
	fmt.Printf("Part 2: Adjust %s to weight %d\n", node, targetWeight)
}
