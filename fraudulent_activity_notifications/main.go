package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const max = 201

func get2xMedian(count []int, d int) int {
	sum := make([]int, max)
	copy(sum, count)
	for i := 1; i < max; i++ {
		sum[i] += sum[i-1]
	}
	var a, b int
	if d%2 == 0 {
		first := d / 2
		second := first + 1
		for i := 0; i < 201; i++ {
			if first <= sum[i] {
				a = i
				break
			}
		}
		for i := 0; i < 201; i++ {
			if second <= sum[i] {
				b = i
				break
			}
		}
	} else {
		first := d/2 + 1
		for i := 0; i < 201; i++ {
			if first <= sum[i] {
				a = 2 * i
				break
			}
		}
	}
	return a + b
}

// https://www.hackerrank.com/challenges/fraudulent-activity-notifications/
func activityNotifications(expenditure []int, d int) int {
	notifications := 0
	count := make([]int, max)
	n := len(expenditure)
	for i := 0; i < d; i++ {
		count[expenditure[i]]++
	}
	for i := d; i < n; i++ {
		median := get2xMedian(count, d)
		if expenditure[i] >= median {
			notifications++
		}
		count[expenditure[i]]++
		count[expenditure[i-d]]--
	}

	return notifications
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int(dTemp)

	expenditureTemp := strings.Split(readLine(reader), " ")

	var expenditure []int

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

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
