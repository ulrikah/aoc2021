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
	prev := -1
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

func windowedSums(numbers []int, windowSize int) (sums []int) {
	for idx := 0; idx < len(numbers)-(windowSize-1); idx++ {
		sum := 0
		for i := idx; i < idx+windowSize; i++ {
			sum += numbers[i]
		}
		sums = append(sums, sum)
	}
	return
}

func main() {
	// test
	testData := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	fmt.Println(countIncrements(testData))
	fmt.Println(windowedSums(testData, 3))
	fmt.Println(countIncrements(windowedSums(testData, 3)))

	f, err := os.Open("day01/input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	numbers, err := LinesToInts(f)

	// 1.1
	fmt.Println(countIncrements(numbers))
	// 1.2
	fmt.Println(countIncrements(windowedSums(numbers, 3)))

}
