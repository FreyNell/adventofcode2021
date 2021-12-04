package main

import (
	"testing"

	"adventofcode2021/frey_utils"
)

func TestPart1(t *testing.T) {
	data := frey_utils.ReadInput("example_input")
	want := 150
	if got := Part1(data); got != want {
		t.Errorf("Part1() = %d, want %d", int(got), int(want))
	}
}

func TestPart2(t *testing.T) {
	data := frey_utils.ReadInput("example_input")
	want := 900
	if got := Part2(data); got != want {
		t.Errorf("Part2() = %d, want %d", int(got), int(want))
	}
}
