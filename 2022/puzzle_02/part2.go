package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Strategy struct {
	rounds []Round
	score  int
}

func (s *Strategy) Add(r Round) {
	s.rounds = append(s.rounds, r)
	s.score += r.Score()
}

type Round struct {
	opponent      Move
	desiredResult Result
}

func (r Round) Score() int {

	player := determineMove(r.opponent, r.desiredResult)

	return r.desiredResult.Value() + player.Value()
}

type Result int

const (
	Win  Result = 6
	Draw Result = 3
	Lose Result = 0
)

func (r Result) Value() int {
	return int(r)
}

type Move string

const (
	Rock     Move = "rock"
	Paper    Move = "paper"
	Scissors Move = "scissors"
)

func (m Move) Value() int {
	switch m {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		return 0
	}
}

func (m Move) Result(opponent Move) Result {
	if m == opponent {
		return Draw
	}

	switch m {
	case Rock:
		if opponent == Scissors {
			return Win
		}
	case Paper:
		if opponent == Rock {
			return Win
		}
	case Scissors:
		if opponent == Paper {
			return Win
		}
	}

	return Lose
}

func main() {

	strategy, err := load("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("The strategy has a score of %d\n", strategy.score)
}

func load(filename string) (Strategy, error) {

	f, err := os.Open(filename)
	if err != nil {
		return Strategy{}, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	strategy := Strategy{
		rounds: []Round{},
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		opponent, player, ok := strings.Cut(line, " ")
		if !ok {
			return strategy, fmt.Errorf("invalid line: %s", line)
		}

		round := Round{
			opponent:      ParseMove(opponent),
			desiredResult: ParseDesiredResult(player),
		}

		strategy.Add(round)
	}

	return strategy, nil
}

func ParseMove(s string) Move {
	switch strings.ToLower(s) {
	case "rock", "x", "a":
		return Rock
	case "paper", "y", "b":
		return Paper
	case "scissors", "z", "c":
		return Scissors
	}
	panic("invalid move")
}

func ParseDesiredResult(s string) Result {
	switch strings.ToLower(s) {
	case "win", "z":
		return Win
	case "draw", "y":
		return Draw
	case "lose", "x":
		return Lose
	}
	panic("invalid result")
}

func determineMove(opponent Move, r Result) Move {
	switch r {
	case Win:
		switch opponent {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	case Lose:
		switch opponent {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	}

	return opponent
}
