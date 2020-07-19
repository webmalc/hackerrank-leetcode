package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/ctci-bubble-sort
func countSwaps(a []int) [3]int {
	result := [3]int{}
	n := len(a)
	for range a {
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				result[0]++
			}
		}
	}
	result[1] = a[0]
	result[2] = a[len(a)-1]
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int(aItemTemp)
		a = append(a, aItem)
	}

	result := countSwaps(a)
	fmt.Printf("Array is sorted in %d swaps.\n", result[0])
	fmt.Printf("First Element: %d\n", result[1])
	fmt.Printf("Last Element: %d\n", result[2])
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
