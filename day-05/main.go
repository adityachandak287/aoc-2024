package main

import (
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

	var answer int

	switch *part {
	case "A":
		answer = PartA(lines)
	case "B":
		answer = PartB(lines)
	default:
		log.Panic("Invalid input for part")
	}

	log.Println("Answer", answer)
}

const (
	PAGE_ORDERING_RULES = 1
	PAGES_IN_UPDATE     = 2
)

func splitAndParse(line, sep string) []int {
	var parts []int
	for _, intStr := range strings.Split(line, sep) {
		num, err := strconv.Atoi(intStr)
		if err != nil {
			panic(err)
		}
		parts = append(parts, num)
	}
	return parts
}

type OrderingRule struct {
	before int
	after  int
}

func parseLines(lines []string) ([]OrderingRule, [][]int) {
	var rules []OrderingRule
	var updatePages [][]int

	mode := PAGE_ORDERING_RULES

	for _, line := range lines {
		if line == "" {
			mode = PAGES_IN_UPDATE
			continue
		}

		if mode == PAGE_ORDERING_RULES {
			parts := splitAndParse(line, "|")
			rules = append(rules, OrderingRule{before: parts[0], after: parts[1]})
		} else if mode == PAGES_IN_UPDATE {
			parts := splitAndParse(line, ",")
			updatePages = append(updatePages, parts)
		} else {
			log.Panic("Invalid mode!")
		}
	}

	return rules, updatePages
}

func PartA(lines []string) int {
	rules, updatePages := parseLines(lines)

	answer := 0
	for _, pages := range updatePages {
		isValid := true
		for _, rule := range rules {
			if !validateRule(pages, rule) {
				isValid = false
				break
			}
		}

		if isValid {
			log.Printf("Valid page updates %v", pages)
			if len(pages)%2 == 0 {
				log.Panicf("Even number of pages! %v", pages)
			}
			midPageNum := pages[len(pages)/2]
			answer += midPageNum
		}
	}

	return answer
}

func validateRule(pages []int, rule OrderingRule) bool {
	beforeIdx := -1
	afterIdx := -1

	for idx, page := range pages {
		if page == rule.before {
			beforeIdx = idx
		}

		if page == rule.after {
			afterIdx = idx
		}
	}

	if beforeIdx == -1 || afterIdx == -1 {
		return true
	}

	return beforeIdx < afterIdx
}

func PartB(lines []string) int {
	rules, updatePages := parseLines(lines)

	answer := 0
	for _, pages := range updatePages {
		isValid := true
		for _, rule := range rules {
			if !validateRule(pages, rule) {
				isValid = false
				break
			}
		}

		if !isValid {
			log.Printf("Invalid page updates %v", pages)

			sort.SliceStable(pages, func(i, j int) bool {
				a := pages[i]
				b := pages[j]

				matchingRule := OrderingRule{before: -1, after: -1}
				for _, r := range rules {
					if (r.before == a && r.after == b) || (r.before == b && r.after == a) {
						matchingRule = r
					}
				}

				if matchingRule.before == -1 || matchingRule.after == -1 {
					log.Panic("Matching rule not found!")
				}

				if a == matchingRule.before {
					return true
				}
				return false
			})

			log.Printf("Invalid pages sorted %v", pages)

			if len(pages)%2 == 0 {
				log.Panicf("Even number of pages! %v", pages)
			}

			midPageNum := pages[len(pages)/2]
			answer += midPageNum
		}
	}

	return answer
}
