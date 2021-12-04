/*
https://adventofcode.com/2021/day/1
*/
package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"adventofcode2021/frey_utils"
)

func Part1(data []int) (increases int) {
	arr_len := len(data)
	increases = 0
	for i := 1; i < arr_len; i++ {
		if data[i-1] < data[i] {
			increases = increases + 1
		} else {
		}
	}
	return
}

func Part2(data []int) (increase int) {
	previous := data[0] + data[1] + data[2]
	sum := 0
	increase = 0
	for i := 1; i < len(data)-2; i++ {
		sum = data[i] + data[i+1] + data[i+2]
		if previous < sum {
			increase = increase + 1
		} else {
		}
		previous = sum
	}
	return
}

func main() {
	start := time.Now()

	var data []int
	input := frey_utils.ReadInput("real_input")
	for _, elem := range input {
		num, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Println("Error reading the data.")
		}
		data = append(data, num)
	}

	fmt.Println(Part1(data))
	fmt.Println(Part2(data))

	elapsed := time.Since(start)
	log.Printf("Day1 Took %s", elapsed)
}
