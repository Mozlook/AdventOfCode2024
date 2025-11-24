package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawInput, err := os.Open("../01/input.txt")
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

	frequency := make(map[int]int)

	for _, number := range secondArray {
		frequency[number]++
	}

	sum := 0
	for _, number := range firstArray {
		freq := frequency[number]
		multi := number * freq
		sum = sum + multi
	}
	fmt.Println(sum)
}
