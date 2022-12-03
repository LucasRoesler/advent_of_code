package main

import (
	"bufio"
	"fmt"
	"os"
)

type Rucksack struct {
	CompartmentOne []Item
	CompartmentTwo []Item
	duplicates     []Item
}

const lowercaseShift = int('a') - 1
const uppercaseShift = int('A') - 27

type Item int

func (i Item) Priority() int {
	return int(i)
}

func (i Item) String() string {
	value := int(i)
	if value >= 27 && value <= 52 {
		return string(rune(value + uppercaseShift))
	}

	if value >= 1 && value <= 26 {
		return string(rune(value + lowercaseShift))
	}

	return ""
}

func main() {
	supplies, err := load("input.txt")
	if err != nil {
		panic(err)
	}

	value := sumDuplicates(supplies)
	fmt.Printf("The value of the duplicates is %d\n", value)
}

func load(src string) ([]Rucksack, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	supplies := []Rucksack{}
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		// split line in half
		// first half is compartment one
		// second half is compartment two
		if len(line)%2 != 0 {
			return nil, fmt.Errorf("invalid line length: %d", len(line))
		}
		left, err := parseItems(line[:len(line)/2])
		if err != nil {
			return nil, err
		}
		right, err := parseItems(line[len(line)/2:])
		if err != nil {
			return nil, err
		}

		r := Rucksack{
			CompartmentOne: left,
			CompartmentTwo: right,
		}
		r.duplicates = findDuplicates(r, 1)

		// fmt.Printf("Rucksack: %v\n", r)

		supplies = append(supplies, r)
	}

	return supplies, nil
}

func parseItems[T string | []byte](value T) ([]Item, error) {
	data := []byte(value)
	items := []Item{}
	for _, c := range data {
		if c >= byte('A') && c <= byte('Z') {
			item := Item(int(c) - uppercaseShift)
			// fmt.Printf("Found uppercase: %s => %s\n", string(c), item)
			items = append(items, item)
			continue
		}

		if c >= byte('a') && c <= byte('z') {
			item := Item(int(c) - lowercaseShift)
			// fmt.Printf("Found lowercase: %s => %s\n", string(c), item)
			items = append(items, item)
			continue
		}

		return nil, fmt.Errorf("invalid character: %s", string(c))
	}

	return items, nil
}

func findDuplicates(r Rucksack, limit int) []Item {
	var duplicates []Item
	for _, item := range r.CompartmentOne {
		// fmt.Printf("Checking %s\n", item)
		if contains(r.CompartmentTwo, item) {
			duplicates = append(duplicates, item)
		}
		if len(duplicates) == limit {
			// fmt.Printf("Found %d duplicates: %v\n", limit, duplicates)
			break
		}
	}
	return duplicates
}

func contains(items []Item, item Item) bool {
	for _, i := range items {
		if i.Priority() == item.Priority() {
			return true
		}
	}
	return false
}

func sumDuplicates(supplies []Rucksack) int {
	sum := 0
	for _, r := range supplies {
		for _, item := range r.duplicates {
			sum += item.Priority()
		}
	}
	return sum
}
