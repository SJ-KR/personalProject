package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024)

	inputStr := strings.TrimSpace(readLine(reader))
	inputSlice := strings.Split(inputStr, " ")
	var input []int64
	for _, val := range inputSlice {
		inputInt, err := strconv.ParseInt(val, 10, 64)
		checkError(err)
		input = append(input, inputInt)
	}
	target := input[2] - input[1]
	unit := input[0] - input[1]

	result := target / unit
	if target%unit != 0 {
		result++
	}

	fmt.Printf("%d", result)

	return
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
