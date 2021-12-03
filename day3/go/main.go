/*
https://adventofcode.com/2021/day/3
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func Btod(row []int, row2 []int) (result int, result2 int) {
	result = 0
	result2 = 0
	for i := len(row) - 1; i >= 0; i-- {
		tmp := (len(row) - 1) - i%len(row)
		result = result + int(math.Pow(2.0, float64(i)))*row[tmp]
		result2 = result2 + int(math.Pow(2.0, float64(i)))*row2[tmp]
	}

	return
}

func GetIndexes(data []string, compare int, order int) (indexes []int) {
	for index, num := range data {
		if int(num[order]%2) == compare {
			indexes = append(indexes, index)
		}
	}
	return
}

func rateVec(data []string) (rate []int) {
	var row []int

	for i := 0; i < len(data[0]); i++ {
		row = append(row, int(data[0][i])%2)
	}

	for _, bits := range data[1:] {
		for i := 0; i < len(bits); i++ {
			row[i] = row[i] + int(bits[i])%2
		}
	}

	for i := 0; i < len(row); i++ {
		if float64(row[i])/float64(len(data)) >= float64(0.5) {
			rate = append(rate, 1)
		} else {
			rate = append(rate, 0)
		}
	}
	return
}

func denyVec(rate []int) (rateres []int) {
	elem_len := len(rate)
	rateres = make([]int, len(rate))
	copy(rateres, rate)
	for i := 0; i < elem_len; i++ {
		if rateres[i] == 1 {
			rateres[i] = 0
		} else {
			rateres[i] = 1
		}
	}
	return
}

func toIntArray(data []string) (rate []int) {
	item_len := len(data[0])
	for i := 0; i < item_len; i++ {
		rate = append(rate, int(data[0][i])%2)
	}
	return
}

func filterData(data []string, rangeData []int, index int, positive bool) (data_res []string) {
	if len(data) < 2 {
		data_res = data
		return
	} else {
		comparer := rangeData[index]
		indexes := GetIndexes(data, rangeData[index], index)
		var data2 []string
		for _, i := range indexes {
			data2 = append(data2, data[i])
		}
		data = data2
		data2 = nil
		fmt.Println("Comparer:", comparer, "Index: ", index, "DATA:", data)
		if positive {
			data_res = filterData(data, rateVec(data), index+1, positive)
		} else {
			data_res = filterData(data, denyVec(rateVec(data)), index+1, positive)
		}
		return
	}

}

func part1(data []string) (gamma_rate []int, epsilon_rate []int) {
	fmt.Println("Part 1:")

	gamma_rate = rateVec(data)
	epsilon_rate = denyVec(gamma_rate)

	gamma, epsilon := Btod(gamma_rate, epsilon_rate)

	fmt.Println("Gamma:", gamma_rate, "(", gamma, ") Epsilon:", epsilon_rate, "(", epsilon, ")")

	power_consumption := gamma * epsilon
	fmt.Println("Power Consumption:", power_consumption)

	return
}

func part2(data []string) {

	fmt.Println("Part 2:")
	rate := rateVec(data)
	o2_gen_rate := toIntArray(filterData(data, rate, 0, true))
	co2_scrubber_rate := toIntArray(filterData(data, denyVec(rate), 0, false))

	o2, co2 := Btod(o2_gen_rate, co2_scrubber_rate)
	fmt.Println("O2:", o2, "CO2:", co2)
	life_support_rating := o2 * co2
	fmt.Println("Life Support Rating:", life_support_rating)
}

func main() {
	var data []string
	file, err := os.Open("../input_aoc3")
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

	part2(data)
}
