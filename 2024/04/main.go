package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Direction struct {
	DX, DY int
}

var (
	Up        = Direction{0, -1}
	Down      = Direction{0, 1}
	Left      = Direction{-1, 0}
	Right     = Direction{1, 0}
	UpLeft    = Direction{-1, -1}
	UpRight   = Direction{1, -1}
	DownLeft  = Direction{-1, 1}
	DownRight = Direction{1, 1}
)

var Directions = []Direction{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight}

func ParsingInputToGrid(i string) [][]rune {
	lines := strings.Split(strings.TrimSpace(i), "\n")
	grid := make([][]rune, len(lines))

	for i, lines := range lines {
		grid[i] = []rune(lines)
	}
	return grid
}

func CheckMatchInDirection(
	data [][]rune,
	startX, startY int,
	matchedWord string,
	direction Direction,
) bool {
	x, y := startX, startY

	for _, c := range matchedWord {
		if x < 0 || y < 0 || y >= len(data) || x >= len(data[y]) || data[y][x] != c {
			return false
		}
		x += direction.DX
		y += direction.DY
	}
	return true
}

func CountMatchesAtPoint(data [][]rune, x, y int, matchedWord string) int {
	count := 0

	for _, d := range Directions {
		if CheckMatchInDirection(data, x, y, matchedWord, d) {
			count++
		}
	}

	return count
}

func EvaluateWordSearch(input string) (int, error) {
	grid := ParsingInputToGrid(input)
	word := "XMAS"
	totalMatches := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			totalMatches += CountMatchesAtPoint(grid, x, y, word)
		}
	}

	return totalMatches, nil
}

func main() {
	fmt.Println("Elf Word Search Evaluator")
	result, err := EvaluateWordSearch(input)
	if err != nil {
		fmt.Println("Error evaluating input: ", err)
	}
	fmt.Println("Total occurances: ", result)
}
