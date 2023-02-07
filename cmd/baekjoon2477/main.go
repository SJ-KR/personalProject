package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type T struct {
	Face int64
	Size int64
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	k, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var input []T
	for i := 0; i < 6; i++ {
		inputStr := strings.TrimSpace(readLine(reader))
		inputArray := strings.Split(inputStr, " ")
		f, err := strconv.ParseInt(inputArray[0], 10, 64)
		checkError(err)
		s, err := strconv.ParseInt(inputArray[1], 10, 64)
		checkError(err)
		input = append(input, T{
			Face: f,
			Size: s,
		})
	}
	for {
		if input[0].Face == input[2].Face && input[1].Face == input[3].Face {
			break
		}
		input = append(input[1:], input[0])
	}
	large := input[4].Size * input[5].Size
	small := input[1].Size * input[2].Size
	fmt.Printf("%d", k*(large-small))
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
