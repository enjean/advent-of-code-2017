package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Checksum(spreadsheet [][]int) int {
	return doChecksum(spreadsheet, rowChecksum)
}

func Checksum2(spreadsheet [][]int) int {
	return doChecksum(spreadsheet, rowChecksum2)
}

func doChecksum(spreadsheet [][]int, rowCheckF func([]int) int) int {
	rowCh := make(chan []int)
	rowChecksumCh := make(chan int)
	go func() {
		for _, row := range spreadsheet {
			rowCh <- row
		}
		close(rowCh)
	}()

	go func() {
		for row := range rowCh {
			rowChecksumCh <- rowCheckF(row)
		}
		close(rowChecksumCh)
	}()

	checksum := 0
	for rowChecksumResult := range rowChecksumCh {
		checksum += rowChecksumResult
	}
	return checksum
}

func rowChecksum(row []int) int {
	min := row[0]
	max := row[0]
	for _, val := range row {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return max - min
}

func rowChecksum2(row []int) int {
	for i := 0; i < len(row); i++ {
		for j:=0; j <len(row); j++ {
			if i == j {
				continue
			}
			if row[i] % row[j] == 0 {
				return row[i] / row[j]
			}
		}
	}
	panic("no match found")
}

func main() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var spreadsheet[][]int
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, val := range strings.Split(line, "\t") {
			intVal, _ := strconv.Atoi(val)
			row = append(row, intVal)
		}
		spreadsheet = append(spreadsheet, row)
	}

	fmt.Printf("Part 1 = %d\n", Checksum(spreadsheet))
	fmt.Printf("Part 2 = %d\n", Checksum2(spreadsheet))
}
