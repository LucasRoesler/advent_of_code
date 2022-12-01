package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ElfParty struct {
	elves      map[int]Backpack
	currentMax struct {
		elf      int
		calories int
	}
}

func (p *ElfParty) Add(elf int, b Backpack) {
	p.elves[elf] = b
	if b.total > p.currentMax.calories {
		p.currentMax.elf = elf
		p.currentMax.calories = b.total
	}
}

type Backpack struct {
	contents []int
	total    int
}

func newBackpack() Backpack {
	return Backpack{
		contents: []int{},
		total:    0,
	}
}

func main() {
	party, err := load("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Elf: %d has the most calories: %d\n", party.currentMax.elf, party.currentMax.calories)
}

func load(src string) (ElfParty, error) {
	party := ElfParty{
		elves: make(map[int]Backpack),
	}

	f, err := os.Open(src)
	if err != nil {
		return party, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	currentElf := 1
	currentBackpack := newBackpack()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			party.Add(currentElf, currentBackpack)
			currentBackpack = newBackpack()
			currentElf++
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			return party, fmt.Errorf("elf %d had non integer values %w", currentElf, err)
		}
		currentBackpack.contents = append(currentBackpack.contents, calories)
		currentBackpack.total += calories
	}

	party.Add(currentElf, currentBackpack)

	return party, nil
}
