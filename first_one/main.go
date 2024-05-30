package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func LineCalibrationNumber(lineContent string) int {
	var first int
	var last int
	for _, symbol := range lineContent {
		integer, err := strconv.Atoi(string(symbol))
		if err != nil {
			continue
		}
		first = integer
		break
	}
	reversed := Reverse(strings.Clone(lineContent))
	for _, symbol := range reversed {
		integer, err := strconv.Atoi(string(symbol))
		if err != nil {
			continue
		}
		last = integer
		break
	}
	return first*10 + last
}

func main() {
	f, _ := os.Open("input.txt")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	scanner := bufio.NewScanner(f)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		calibrationNumber := LineCalibrationNumber(line)
		sum += calibrationNumber
	}
	fmt.Println(sum)
}
