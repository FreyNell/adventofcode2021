package main

import (
	"adventofcode2021/frey_utils"
	"fmt"
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	var data []int
	input := frey_utils.ReadInput("example_input")
	for _, elem := range input {
		num, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Println("Error reading the data.")
		}
		data = append(data, num)
	}
	want := 7
	if got := Part1(data); got != want {
		t.Errorf("Part1() = %q, want %q", got, want)
	}
}

func TestPart2(t *testing.T) {
	var data []int
	input := frey_utils.ReadInput("example_input")
	for _, elem := range input {
		num, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Println("Error reading the data.")
		}
		data = append(data, num)
	}
	want := 5
	if got := Part2(data); got != want {
		t.Errorf("Part2() = %q, want %q", int(got), int(want))
	}
}
