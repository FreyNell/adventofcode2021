/*
https://adventofcode.com/2021/day/2
*/
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"adventofcode2021/frey_utils"
)

func Part1(data []string) (multiplier int) {
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

	multiplier = depth * horizontal
	return
}

func Part2(data []string) (multiplier int) {
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

	multiplier = depth * horizontal
	return
}

func main() {
	start := time.Now()

	data := frey_utils.ReadInput("real_input")
	fmt.Println(Part1(data))
	fmt.Println(Part2(data))

	elapsed := time.Since(start)
	log.Printf("Day2 Took %s", elapsed)
}
