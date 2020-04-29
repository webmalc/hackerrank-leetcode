package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/special-palindrome-again/
func substrCount(n int, s string) int64 {
	var j, i, c, total int
	same := make(map[int]int)
	for i < n {
		j = i + 1
		c = 1

		for j < n && s[i] == s[j] {
			j++
			c++
		}
		total += (c * (c + 1)) / 2
		same[i] = c
		i = j
	}

	for j := 1; j < n-1; j++ {
		if s[j] == s[j-1] {
			same[j] = same[j-1]
		}

		if s[j-1] == s[j+1] && s[j] != s[j-1] {
			if same[j-1] > same[j+1] {
				total += same[j+1]
			} else {
				total += same[j-1]
			}
		}
	}

	return int64(total)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)

	s := readLine(reader)

	result := substrCount(n, s)

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
