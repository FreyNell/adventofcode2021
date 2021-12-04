package frey_utils

import (
	"bufio"
	"log"
	"math"
	"os"
)

func BinaryToDecimal(row []int) (result int) {
	result = 0
	for i := len(row) - 1; i >= 0; i-- {
		tmp := (len(row) - 1) - i%len(row)
		result = result + int(math.Pow(2.0, float64(i)))*row[tmp]
	}
	return
}

func DenyBinaryArray(arr []int) (arrn []int) {
	elem_len := len(arr)
	arrn = make([]int, len(arr))
	copy(arrn, arr)
	for i := 0; i < elem_len; i++ {
		if arrn[i] == 1 {
			arrn[i] = 0
		} else {
			arrn[i] = 1
		}
	}
	return
}

func ReadInput(path string) (data []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func StrToIntArray(data string) (arr []int) {
	item_len := len(data)
	for i := 0; i < item_len; i++ {
		arr = append(arr, int(data[i])%2)
	}
	return
}
