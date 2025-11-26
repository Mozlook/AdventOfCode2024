package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	matches := re.FindAllStringSubmatch(string(data), -1)
	sum := 0
	for _, m := range matches {
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
	print(sum)
}
