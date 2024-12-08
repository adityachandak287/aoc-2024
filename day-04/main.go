package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	inputFile := flag.String("input", "", "Input file relative path")
	part := flag.String("part", "A", "Implementation of part A or B of the problem")

	flag.Parse()

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		panic(err)
	}

	input := string(data)

	lines := strings.Split(input, "\n")

	answer := SolvePart(*part, lines)

	log.Println("Answer", answer)
}

type Solution struct {
	grid [][]rune
	part string
}

func NewSolution(part string) *Solution {
	return &Solution{
		part: part,
	}
}

func (s *Solution) initGrid(lines []string) {
	for _, line := range lines {
		s.grid = append(s.grid, []rune(line))
	}
}

func (s *Solution) solve() int {
	count := 0

	for r := 0; r < len(s.grid); r++ {
		for c := 0; c < len(s.grid[r]); c++ {
			char := s.grid[r][c]

			if s.part == "A" {
				if char != 'X' {
					continue
				}
				count += s.checkInAllDirections(r, c)
			} else if s.part == "B" {
				if char != 'A' {
					continue
				}
				count += s.checkInXDirections(r, c)
			} else {
				log.Panic("Invalid part!")
			}
		}
	}

	return count
}

func (s *Solution) checkInAllDirections(r, c int) int {
	cellCount := 0

	targets := []rune{'X', 'M', 'A', 'S'}

	for dirR := -1; dirR <= 1; dirR++ {
		for dirC := -1; dirC <= 1; dirC++ {
			if dirR == 0 && dirC == 0 {
				continue
			}

			if s.checkInDirection(r, c, dirR, dirC, targets) {
				log.Printf("Found match [Start (%d, %d)] [Dir (%d, %d)]", r, c, dirR, dirC)
				cellCount += 1
			}
		}
	}

	return cellCount
}

func (s *Solution) checkInDirection(r, c, dirR, dirC int, targets []rune) bool {
	for idx, target := range targets {
		row := r + idx*dirR
		col := c + idx*dirC

		if row < 0 || row >= len(s.grid) {
			return false
		}

		if col < 0 || col >= len(s.grid[row]) {
			return false
		}

		if target != s.grid[row][col] {
			return false
		}
	}

	return true
}

func (s *Solution) checkInXDirections(r, c int) int {
	if r < 1 || r >= len(s.grid)-1 {
		return 0
	}

	if c < 1 || c >= len(s.grid[r])-1 {
		return 0
	}

	diagLR := string([]rune{s.grid[r-1][c-1], s.grid[r][c], s.grid[r+1][c+1]})
	diagRL := string([]rune{s.grid[r-1][c+1], s.grid[r][c], s.grid[r+1][c-1]})

	if (diagLR == "MAS" || diagLR == "SAM") && (diagRL == "MAS" || diagRL == "SAM") {
		log.Printf("Found match [Start (%d, %d)] ", r, c)
		return 1
	}
	return 0
}

func SolvePart(part string, lines []string) int {
	solution := NewSolution(part)

	solution.initGrid(lines)

	return solution.solve()
}
