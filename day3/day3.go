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
		for yDiff := 1; yDiff < squareSize-1 && currentSquare != fromSquare; yDiff++ {
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

type position struct {
	x, y int
}

func (p position) neighbors() []position {
	var result []position
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			result = append(result, position{p.x + dx, p.y + dy})
		}
	}
	return result
}

type stressValues map[position]int

//147  142  133  122   59
//304    5    4    2   57
//330   10    1    1   54
//351   11   23   25   26
//362  747  806--->   ..

func (sv stressValues) updateValueAt(x, y int) int {
	position := position{x,y}
	value := 0
	for _, neighbor := range position.neighbors() {
		value += sv[neighbor]
	}
	sv[position] = value
	return value
}

func StressTest(limit int) int {
	values := stressValues{
		{0, 0}: 1,
		{1, 0}: 1,
	}
	x := 1
	y := 0

	for squareSize := 3; ; squareSize += 2 {
		for yDiff := 1; yDiff < squareSize-1; yDiff++ {
			y--
			newVal := values.updateValueAt(x,y)
			if newVal > limit {
				return newVal
			}
		}
		for xDiff := 1; xDiff < squareSize; xDiff++ {
			x--
			newVal := values.updateValueAt(x,y)
			if newVal > limit {
				return newVal
			}
		}
		for yDiff := 1; yDiff < squareSize; yDiff++ {
			y++
			newVal := values.updateValueAt(x,y)
			if newVal > limit {
				return newVal
			}
		}
		for xDiff := 1; xDiff <= squareSize; xDiff++ {
			x++
			newVal := values.updateValueAt(x,y)
			if newVal > limit {
				return newVal
			}
		}
	}

}

func main() {
	input := 277678
	fmt.Printf("Part 1: %d\n", StepsToPort(input))
	fmt.Printf("Part 2: %d\n", StressTest(input))
}
