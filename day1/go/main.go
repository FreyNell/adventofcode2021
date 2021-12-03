package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(data []int) {
	arr_len := len(data)
	increases := 0
	fmt.Println(data[0], " N/A")
	for i := 1; i < arr_len; i++ {
		if data[i-1] < data[i] {
			increases = increases + 1
			fmt.Println(data[i], " increased ", increases, "(", data[i-1], data[i], ")")
		} else {
			fmt.Println(data[i], " decreased")
		}
	}
	fmt.Println("Total len: ", arr_len)
	fmt.Println("Increased times: ", increases)
}

func part2(data []int) {
	previous := data[0] + data[1] + data[2]
	fmt.Println(data[0], data[1], data[2], "=", previous, "N/A")
	sum := 0
	increase := 0
	for i := 1; i < len(data)-2; i++ {
		sum = data[i] + data[i+1] + data[i+2]
		if previous < sum {
			fmt.Println(data[i], data[i+1], data[i+2], "=", sum, "increased")
			increase = increase + 1
		} else {
			fmt.Println(data[i], data[i+1], data[i+2], "=", sum, "decreased")
		}
		previous = sum
	}

	fmt.Println(increase)
}

func main() {
	var data []int
	file, err := os.Open("../input_aoc1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//part1(data)
	part2(data)
}
