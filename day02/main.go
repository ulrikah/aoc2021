package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func LinesToCommands(file *os.File) (commands []Command, err error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineArray := strings.Split(line, " ")
		command := lineArray[0]
		magnitude, err := strconv.Atoi(lineArray[1])
		if err != nil {
			return nil, err
		}
		commands = append(commands, Command{command, magnitude})
	}
	return
}

type Position struct {
	vertical   int
	horizontal int
}

type Command struct {
	direction string
	magnitude int
}

func calculatePosition(commands []Command) Position {
	position := Position{
		vertical:   0,
		horizontal: 0,
	}
	for _, command := range commands {
		switch command.direction {
		case "forward":
			position.horizontal += command.magnitude
		case "up":
			position.vertical -= command.magnitude
		case "down":
			position.vertical += command.magnitude
		}
	}
	return position
}
func calculatePositionWithAim(commands []Command) Position {
	aim := 0
	position := Position{
		vertical:   0,
		horizontal: 0,
	}
	for _, command := range commands {
		switch command.direction {
		case "forward":
			position.horizontal += command.magnitude
			position.vertical += aim * command.magnitude
		case "up":
			aim -= command.magnitude
		case "down":
			aim += command.magnitude
		}
	}
	return position
}

func main() {
	// test
	testData := []Command{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}

	fmt.Println(calculatePosition(testData))
	fmt.Println(calculatePositionWithAim(testData))

	f, err := os.Open("day02/input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	commands, err := LinesToCommands(f)

	// 1.1
	position := calculatePosition(commands)
	fmt.Println(position.horizontal * position.vertical)
	// 1.2
	positionWithAim := calculatePositionWithAim(commands)
	fmt.Println(positionWithAim.horizontal * positionWithAim.vertical)

}
