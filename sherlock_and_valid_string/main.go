package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// https://www.hackerrank.com/challenges/sherlock-and-valid-string/
func isValid(s string) string {
	const yes = "YES"
	const no = "NO"
	max := 1
	hash := make(map[string]int)
	hashCount := make(map[int]int)
	for _, char := range s {
		hash[string(char)]++
	}
	for _, count := range hash {
		hashCount[count]++
	}
	if len(hashCount) == max {
		return yes
	}
	for i, c := range hashCount {
		if c == 1 {
			diff := i - c
			if diff > 0 {
				hashCount[diff]++
			}
			delete(hashCount, i)
			break
		}
	}
	if len(hashCount) == max {
		return yes
	}
	return no
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := isValid(s)

	fmt.Fprintf(writer, "%s\n", result)

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
