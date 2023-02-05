package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type TestSet struct {
	K int64
	N int64
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var set []TestSet
	var maxK, maxN int64
	for i := int64(0); i < arCount; i++ {
		inputK, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		inputN, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		set = append(set, TestSet{
			K: inputK,
			N: inputN,
		})
		if maxK < inputK {
			maxK = inputK
		}
		if maxN < inputN {
			maxN = inputN
		}
	}

	lib := make(map[int64]map[int64]int64)
	for i := int64(0); i <= maxK; i++ {
		lib[i] = make(map[int64]int64)
	}

	for i := int64(1); i <= maxN; i++ {
		lib[0][i] = i
	}

	for k := int64(1); k <= maxK; k++ {
		for n := int64(1); n <= maxN; n++ {
			lib[k][n] = lib[k-1][n] + lib[k][n-1]
		}
	}

	for _, s := range set {
		fmt.Printf("%d\n", lib[s.K][s.N])
	}

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
