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

func safeWithSkip(a []int, skip int) bool {
	prev := -1
	dir := 0

	for i := range len(a) {
		if i == skip {
			continue
		}
		if prev == -1 {
			prev = i
			continue
		}
		d := a[i] - a[prev]

		if d == 0 {
			return false
		}

		if dir == 0 {
			if d > 0 {
				dir = 1
			} else {
				dir = -1
			}
		}

		if dir == 1 && d <= 0 {
			return false
		}
		if dir == -1 && d >= 0 {
			return false
		}

		ad := abs(d)
		if ad < 1 || ad > 3 {
			return false
		}

		prev = i
	}
	return prev != -1
}

func isSafe(report []int) bool {
	if safeWithSkip(report, -1) {
		return true
	}
	for k := range report {
		if safeWithSkip(report, k) {
			return true
		}
	}
	return false
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
