package main

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

// Generate the sequence of numbers
func generateRange(max int32) []int32 {
	result := make([]int32, max)
	for i := int32(0); i < max; i++ {
		result[i] = i + 1
	}
	return result
}

// https://www.hackerrank.com/challenges/new-year-chaos/
func minimumBribes(queue []int32) string {
	const step = 2
	bribes := 0
	queueLen := len(queue)
	source := generateRange(int32(queueLen))

	for i := 0; i < queueLen; i++ {
		queueValue := queue[i]

		if source[i] != queueValue {
			diff := bits.UintSize
			for k := 1; k <= step; k++ {
				if source[i+k] == queueValue {
					diff = k
					break
				}
			}
			if diff > step {
				return "Too chaotic"
			}
			for diff > 0 {
				source[i+diff], source[i+diff-1] = source[i+diff-1], source[i+diff]
				bribes++
				diff--
			}
		}
	}
	return strconv.Itoa(bribes)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		fmt.Println(minimumBribes(q))
	}
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
