package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/ctci-merge-sort
func countInversions(arr []int) int {
	t := []int{1, 2, 3, 4}
	r := append([]int{}, t[2:]...)
	r[1] = 5
	fmt.Println(r)
	fmt.Println(t)
	return inversions(arr)
}

func inversions(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}
	n1 := n / 2
	n2 := n - n1
	arr1 := append([]int{}, arr[:n1]...)
	arr2 := append([]int{}, arr[n1:]...)
	count := inversions(arr1) + inversions(arr2)
	i1 := 0
	i2 := 0
	for i := 0; i <= n; i++ {
		if i1 < n1 && (i2 >= n2 || arr1[i1] <= arr2[i2]) {
			arr[i] = arr1[i1]
			count += i2
			i1++
		} else if i2 < n2 {
			arr[i] = arr2[i2]
			i2++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int(nTemp)

		arrTemp := strings.Split(readLine(reader), " ")

		var arr []int

		for i := 0; i < int(n); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arr = append(arr, int(arrItem))
		}

		result := countInversions(arr)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
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
