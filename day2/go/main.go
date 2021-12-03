package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(data []string) {
	depth := 0
	horizontal := 0
	for _, s := range data {
		splitted := strings.Split(s, " ")
		value, err := strconv.Atoi(splitted[1])
		if err != nil {
			log.Fatal(err)
		}
		if splitted[0] == "forward" {
			horizontal = horizontal + value
		} else if splitted[0] == "up" {
			depth = depth - value
		} else if splitted[0] == "down" {
			depth = depth + value
		}
	}

	fmt.Println("Depth:", depth, "Horizontal:", horizontal, "Multiplier:", depth*horizontal)
}

func part2(data []string) {
	depth := 0
	horizontal := 0
	aim := 0
	for _, s := range data {
		splitted := strings.Split(s, " ")
		value, err := strconv.Atoi(splitted[1])
		if err != nil {
			log.Fatal(err)
		}
		if splitted[0] == "forward" {
			horizontal = horizontal + value
			depth = depth + aim*value
		} else if splitted[0] == "up" {
			aim = aim - value
		} else if splitted[0] == "down" {
			aim = aim + value
		}
	}

	fmt.Println("Depth:", depth, "Horizontal:", horizontal, "Multiplier:", depth*horizontal)
}

func main() {
	var data []string
	file, err := os.Open("../input_aoc2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		data = append(data, str)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:")
	part1(data)
	fmt.Println("Part 2:")
	part2(data)
}
