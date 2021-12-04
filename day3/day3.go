/*
https://adventofcode.com/2021/day/3
*/
package main

import (
	"adventofcode2021/frey_utils"
	"fmt"
	"log"
	"time"
)

func GetIndexes(data []string, compare int, order int) (indexes []int) {
	for index, num := range data {
		if int(num[order]%2) == compare {
			indexes = append(indexes, index)
		}
	}
	return
}

func MostCommon(data []string) (rate []int) {
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

func FilterData(data []string, rangeData []int, index int, positive bool) (data_res []string) {
	if len(data) < 2 {
		data_res = data
		return
	} else {
		indexes := GetIndexes(data, rangeData[index], index)
		var data2 []string
		for _, i := range indexes {
			data2 = append(data2, data[i])
		}
		data = data2
		data2 = nil
		if positive {
			data_res = FilterData(data, MostCommon(data), index+1, positive)
		} else {
			data_res = FilterData(data, frey_utils.DenyBinaryArray(MostCommon(data)), index+1, positive)
		}
		return
	}
}

func Part1(data []string) (power_consumption int) {
	gamma_rate := MostCommon(data)
	epsilon_rate := frey_utils.DenyBinaryArray(gamma_rate)
	gamma := frey_utils.BinaryToDecimal(gamma_rate)
	epsilon := frey_utils.BinaryToDecimal(epsilon_rate)
	power_consumption = gamma * epsilon
	return
}

func Part2(data []string) (life_support_rating int) {
	arr := MostCommon(data)

	data_res := FilterData(data, arr, 0, true)
	o2_gen_arr := frey_utils.StrToIntArray(data_res[0])

	data_res = FilterData(data, frey_utils.DenyBinaryArray(arr), 0, false)
	co2_scrubber_arr := frey_utils.StrToIntArray(data_res[0])

	o2 := frey_utils.BinaryToDecimal(o2_gen_arr)
	co2 := frey_utils.BinaryToDecimal(co2_scrubber_arr)

	life_support_rating = o2 * co2
	return
}

func main() {
	start := time.Now()

	data := frey_utils.ReadInput("real_input")
	fmt.Println(Part1(data))
	fmt.Println(Part2(data))

	elapsed := time.Since(start)
	log.Printf("Day3 Took %s", elapsed)
}
