package main

import "fmt"

func StepsToPort(fromSquare int) int {
	if fromSquare == 1 {
		return 0
	}

	x := 1
	y := 0
	currentSquare := 2

	for squareSize := 3; currentSquare != fromSquare; squareSize += 2 {
		for yDiff := 1; yDiff < squareSize -1 && currentSquare != fromSquare; yDiff++ {
			y--
			currentSquare++
			//fmt.Printf("%d (%d,%d)\n", currentSquare, x,y)
		}
		for xDiff := 1; xDiff < squareSize && currentSquare != fromSquare; xDiff++ {
			x--
			currentSquare++
			//fmt.Printf("%d (%d,%d)\n", currentSquare, x,y)
		}
		for yDiff := 1; yDiff < squareSize && currentSquare != fromSquare; yDiff++ {
			y++
			currentSquare++
			//fmt.Printf("%d (%d,%d)\n", currentSquare, x,y)
		}
		for xDiff := 1; xDiff <= squareSize && currentSquare != fromSquare; xDiff++ {
			x++
			currentSquare++
			//fmt.Printf("%d (%d,%d)\n", currentSquare, x,y)
		}
	}

	return abs(x) + abs(y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	input := 277678
	fmt.Printf("Part 1: %d", StepsToPort(input))
}