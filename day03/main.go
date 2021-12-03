package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func transposeArray(array []string) []string {
	transposedArray := make([]string, len(array[0]))
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			transposedArray[j] += string(array[i][j])
		}
	}
	return transposedArray
}

func decFromBinary(bin string) int {
	dec, err := strconv.ParseInt(bin, 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	return int(dec)
}

func findGammaAndEpsilon(binaryStrings []string) (int, int) {
	transposedArray := transposeArray(binaryStrings)
	gammaRate := ""
	for _, binaryString := range transposedArray {
		ones := strings.Count(binaryString, "1")
		if ones > (len(binaryString) - ones) {
			gammaRate += "1"
		} else {
			gammaRate += "0"
		}
	}
	epsilonRate := invertBits(gammaRate)
	return decFromBinary(gammaRate), decFromBinary(epsilonRate)
}

func getSubsets(binStrings []string, position int) (oxygenSubset []string, co2Subset []string) {
	ones, zeros := 0, 0
	for _, binaryString := range binStrings {
		if string(binaryString[position]) == "1" {
			ones++
		} else {
			zeros++
		}
	}
	for _, binaryString := range binStrings {
		if string(binaryString[position]) == "1" {
			if ones > zeros {
				oxygenSubset = append(oxygenSubset, binaryString)
			} else if ones == zeros {
				oxygenSubset = append(oxygenSubset, binaryString)
			} else {
				co2Subset = append(co2Subset, binaryString)
			}
		} else {
			if ones < zeros {
				oxygenSubset = append(oxygenSubset, binaryString)
			} else if ones == zeros {
				oxygenSubset = append(oxygenSubset, binaryString)
			} else {
				co2Subset = append(co2Subset, binaryString)
			}
		}
	}
	return
}

func findOxygenAndCO2(binaryStrings []string) ([]string, []string) {
	totalSubsets := binaryStrings
	oxygenSubset, co2Subset := []string{}, []string{}
	for i := 0; i < len(binaryStrings[0]); i++ {
		oxygenSubset, co2Subset = getSubsets(totalSubsets, i)
		fmt.Println("LENGTH", len(oxygenSubset), len(co2Subset))
		totalSubsets = append(oxygenSubset, co2Subset...)
	}
	return oxygenSubset, co2Subset
}

func invertBits(bin string) (inv string) {
	for _, bit := range bin {
		if bit == '0' {
			inv += "1"
		} else {
			inv += "0"
		}
	}
	return
}

func LinesFromFile(file *os.File) (lines []string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return
}

func main() {
	testData := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	fmt.Println(findGammaAndEpsilon((testData)))

	f, err := os.Open("day03/input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	lines := LinesFromFile(f)
	// 1
	gamma, epsilon := findGammaAndEpsilon((lines))
	fmt.Println(gamma * epsilon)

	// 2
	panic("Not implemented")
	oxygenSubset, co2Subset := findOxygenAndCO2((lines))
	fmt.Println(len(oxygenSubset), len(co2Subset))
}
