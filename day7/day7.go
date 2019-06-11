package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type programSpec struct {
	name     string
	children []string
}

//type program struct {
//	name     string
//	parent   *program
//	children []*program
//}

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

var lineRexexp = regexp.MustCompile(`([a-z]+) \([\d]+\)(?: -> (.*))?`)

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
		var children []string
		if lineMatches[2] != "" {
			children = strings.Split(lineMatches[2],", ")
		}
		input = append(input, programSpec{name:name, children:children})
	}

	root := FindRoot(input)
	fmt.Printf("Part 1: %s", root)
}
