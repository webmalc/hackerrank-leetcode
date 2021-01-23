package angrychildren

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Return the unfairness of the array
func getUnfairness(array []int) int {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return max - min
}

// https://www.hackerrank.com/challenges/angry-children
func maxMin(k int, arr []int) int {
	sort.Ints(arr)
	max := math.MaxInt32
	i := 0
	for i < len(arr)-k+1 {
		unfairness := getUnfairness(arr[i : i+k])
		if unfairness < max {
			max = unfairness
		}
		i++
	}

	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int(kTemp)

	var arr []int

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := maxMin(k, arr)

	fmt.Fprintf(writer, "%d\n", result)

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
