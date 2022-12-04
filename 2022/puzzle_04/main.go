package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Assignment struct {
	start int
	end   int
}

type Elves [][2]Assignment

func main() {
	elves, err := load("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	containmentCount := 0
	overlapCount := 0
	for _, pair := range elves {
		if contained(pair) {
			containmentCount++
		}

		if overlap(pair) {
			overlapCount++
		}
	}

	fmt.Printf("Contained: %d / %d\n", containmentCount, len(elves))
	fmt.Printf("Overlap: %d / %d\n", overlapCount, len(elves))
}

func load(src string) (Elves, error) {
	elves := Elves{}

	f, err := os.Open(src)
	if err != nil {
		return elves, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		pair := [2]Assignment{}
		line := scanner.Text()
		left, right, _ := strings.Cut(line, ",")
		_, err := fmt.Sscanf(left, "%d-%d", &pair[0].start, &pair[0].end)
		if err != nil {
			return elves, err
		}

		_, err = fmt.Sscanf(right, "%d-%d", &pair[1].start, &pair[1].end)
		if err != nil {
			return elves, err
		}

		// fmt.Printf("parsed pair: %v\n", pair)

		elves = append(elves, pair)
	}
	return elves, nil
}

// return true if the first assignment is contained within the second
// or if the second assignment is contained within the first
func contained(pair [2]Assignment) bool {
	left, right := pair[0], pair[1]

	// left is contained within right
	if left.start >= right.start && left.end <= right.end {
		return true
	}

	// right is contained within left
	if right.start >= left.start && right.end <= left.end {
		return true
	}

	return false
}

func overlap(pair [2]Assignment) bool {
	left, right := pair[0], pair[1]

	// left has any overlap with right
	if left.start <= right.start && left.end >= right.start {
		return true
	}

	// right has any overlap with left
	if right.start <= left.start && right.end >= left.start {
		return true
	}

	return false
}
