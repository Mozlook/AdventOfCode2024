package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isSafe(report []int) bool {
	if len(report) < 2 {
		return false
	}

	d0 := report[1] - report[0]
	if d0 == 0 {
		return false
	}

	dir := 1
	if d0 < 0 {
		dir = -1
	}

	for i := 0; i < len(report)-1; i++ {
		d := report[i+1] - report[i]

		if dir == 1 && d <= 0 {
			return false
		}
		if dir == -1 && d >= 0 {
			return false
		}

		if a := abs(d); a < 1 || a > 3 {
			return false
		}
	}
	return true
}

func main() {
	rawInput, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("input error: %v", err)
		return
	}
	defer rawInput.Close()

	scanner := bufio.NewScanner(rawInput)

	safe := 0
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		fields := strings.Fields(scanner.Text())

		report := make([]int, len(fields))
		for i, s := range fields {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Fprintf(os.Stderr, "parsing error (line %d, token %q): %v\n", lineNo, s, err)
				return
			}
			report[i] = n
		}

		if isSafe(report) {
			safe++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "scan error: %v\n", err)
		return
	}
	fmt.Println(safe)
}
