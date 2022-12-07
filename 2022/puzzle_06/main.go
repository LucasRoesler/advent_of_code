package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	chunks, err := load("input.txt", 4)
	if err != nil {
		panic(err)
	}

	index := 0
	for chunk := range chunks {
		if !hasDuplicates(chunk) {
			fmt.Println(string(chunk))
			fmt.Printf("Part 1: the first index with no duplicates is: %d\n", index+4)
			break
		}
		index++
	}

	chunks, err = load("input.txt", 14)
	if err != nil {
		panic(err)
	}

	index = 0
	for chunk := range chunks {
		if !hasDuplicates(chunk) {
			fmt.Println(string(chunk))
			fmt.Printf("Part 2: the first index with no duplicates is: %d\n", index+14)
			break
		}
		index++
	}
}

func load(src string, size int) (<-chan []byte, error) {
	data := make(chan []byte)

	f, err := os.Open(src)
	if err != nil {
		return data, err
	}

	go func() {
		defer f.Close()
		defer close(data)

		scanner := bufio.NewScanner(f)
		scanner.Split(slidingWindowSplitFunc(size))

		for scanner.Scan() {
			data <- scanner.Bytes()
		}

		return
	}()

	return data, nil
}

func slidingWindowSplitFunc(size int) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if len(data) < size {
			return len(data), data, nil
		}

		return 1, data[:size], nil
	}
}
func hasDuplicates(d []byte) bool {
	seen := make(map[byte]bool)
	for _, v := range d {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}
