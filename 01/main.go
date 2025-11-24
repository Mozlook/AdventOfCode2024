package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rawInput, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("input error: %v", err)
		return
	}
	defer rawInput.Close()

	scanner := bufio.NewScanner(rawInput)

	var firstArray []int
	var secondArray []int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		firstNumber, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Printf("parsing error: %v", err)
			return
		}
		secondNumber, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Printf("parsing error: %v", err)
			return
		}

		firstArray = append(firstArray, firstNumber)
		secondArray = append(secondArray, secondNumber)

	}

	slices.Sort(firstArray)
	slices.Sort(secondArray)

	var sum int
	for i := range 1000 {
		diff := firstArray[i] - secondArray[i]
		if diff < 0 {
			diff = -diff
		}
		sum = sum + diff
	}
	fmt.Println(sum)
}
