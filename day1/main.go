package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func LinesToInts(file *os.File) (ints []int, err error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		integer, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		ints = append(ints, integer)
	}
	return
}

func countIncrements(numbers []int) int {
	count := 0
	prev := 0
	for idx, next := range numbers {
		if idx == 0 {
			continue
		}
		if next > prev {
			count += 1
		}
		prev = next
	}
	return count
}

func main() {
	fmt.Println("Hello, World!")
	f, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	numbers, err := LinesToInts(f)
	fmt.Println(countIncrements(numbers))
}
