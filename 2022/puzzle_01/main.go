package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type ElfParty struct {
	elves      Backpacks
	currentMax struct {
		elf      int
		calories int
	}
}

func (p *ElfParty) Add(elf int, b Backpack) {
	if b.total > p.currentMax.calories {
		p.currentMax.elf = elf
		p.currentMax.calories = b.total
	}
	p.elves = append(p.elves, b)
	sort.Sort(p.elves)
}

type Backpacks []Backpack

func (b Backpacks) Len() int           { return len(b) }
func (b Backpacks) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b Backpacks) Less(i, j int) bool { return b[i].total > b[j].total }

type Backpack struct {
	owner    int
	contents []int
	total    int
}

func newBackpack(owner int) Backpack {
	return Backpack{
		owner:    owner,
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

	topN := 3
	total := 0
	fmt.Printf("The top %d elves are:\n", topN)
	for i := 0; i < topN; i++ {
		fmt.Printf("Elf: %d has %d calories\n", party.elves[i].owner, party.elves[i].total)
		total += party.elves[i].total
	}
	fmt.Printf("Total: %d\n", total)
}

func load(src string) (ElfParty, error) {
	party := ElfParty{
		elves: []Backpack{},
	}

	f, err := os.Open(src)
	if err != nil {
		return party, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	currentElf := 1
	currentBackpack := newBackpack(currentElf)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			party.Add(currentElf, currentBackpack)
			currentElf++
			currentBackpack = newBackpack(currentElf)
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
