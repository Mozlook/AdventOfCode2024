package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\)`)

	matches := re.FindAllStringSubmatch(string(data), -1)
	sum := 0
	enabled := true
	for _, m := range matches {
		token := m[0]

		switch {
		case strings.HasPrefix(token, "mul("):
			if enabled {
				intA, err := strconv.Atoi(m[1])
				if err != nil {
					log.Fatal(err)
				}
				intB, err := strconv.Atoi(m[2])
				if err != nil {
					log.Fatal(err)
				}
				product := intA * intB
				sum = sum + product
			}

		case token == "do()":
			enabled = true

		case token == "don't()":
			enabled = false
		}
	}
	fmt.Println(sum)
}
