package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// https://www.hackerrank.com/challenges/common-child
func commonChild(s1, s2 string) int {
	l := len(s1)
	lcs := make([][]int, l+1)
	for i := range lcs {
		lcs[i] = make([]int, l+1)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if s1[i] == s2[j] {
				lcs[i+1][j+1] = lcs[i][j] + 1
			} else if lcs[i+1][j] > lcs[i][j+1] {
				lcs[i+1][j+1] = lcs[i+1][j]
			} else {
				lcs[i+1][j+1] = lcs[i][j+1]
			}
		}
	}
	return lcs[l][l]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s1 := readLine(reader)

	s2 := readLine(reader)

	result := commonChild(s1, s2)

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
