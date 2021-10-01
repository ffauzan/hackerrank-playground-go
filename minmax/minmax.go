package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int64) {
	// Write your code here

	min := int64(0)
	max := int64(0)
	sum := int64(0)

	for i := 0; i < len(arr); i++ {
		sum = 0
		for j := 0; j < len(arr); j++ {
			if i != j {
				sum += arr[j]
			}
		}
		if i == 0 {
			min = sum
			max = sum
		} else {
			if sum < min {
				min = sum
			} else if sum > max {
				max = sum
			}
		}
	}

	fmt.Printf("%d %d", min, max)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int64(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
