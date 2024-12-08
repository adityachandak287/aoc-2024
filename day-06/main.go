package main

import (
	"flag"
	"log"
	"os"
	"strings"

	solution "github.com/adityachandak287/aoc-2024/day-06/solution"
)

func main() {
	inputFile := flag.String("input", "input", "Input file relative path")
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

func SolvePart(part string, lines []string) int {
	soln := solution.NewSolution(part)

	soln.InitGrid(lines)

	return soln.Solve()
}
