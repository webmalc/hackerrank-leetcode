package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// https://www.hackerrank.com/challenges/ctci-making-anagrams/
func makeAnagram(a, b string) int {
	counter := 0
	hash := make(map[rune]int)
	for _, char := range a {
		hash[char]++
	}
	for _, char := range b {
		hash[char]--
	}
	for _, val := range hash {
		if val > 0 {
			counter += val
		} else if val < 0 {
			counter -= val
		}
	}
	return counter
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	a := readLine(reader)

	b := readLine(reader)

	res := makeAnagram(a, b)

	fmt.Fprintf(writer, "%d\n", res)

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
