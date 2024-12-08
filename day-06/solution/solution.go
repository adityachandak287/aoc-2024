package solution

import "log"

type Solution struct {
	grid    [][]rune
	visited [][]bool
	rows    int
	cols    int
	part    string
}

func NewSolution(part string) *Solution {
	return &Solution{
		part: part,
	}
}

func (s *Solution) InitGrid(lines []string) {
	s.rows = len(lines)
	s.cols = len(lines[0])

	s.visited = make([][]bool, s.rows)

	for idx, line := range lines {
		s.grid = append(s.grid, []rune(line))

		s.visited[idx] = make([]bool, s.cols)
	}
}

const (
	UP    = 1
	RIGHT = 2
	DOWN  = 3
	LEFT  = 4
)

func (s *Solution) findGuard() (r, c, dir int) {
	for rowIdx, row := range s.grid {
		for colIdx, col := range row {
			switch col {
			case '^':
				return rowIdx, colIdx, UP
			case '>':
				return rowIdx, colIdx, RIGHT
			case '<':
				return rowIdx, colIdx, LEFT
			case 'v':
				return rowIdx, colIdx, DOWN
			}
		}
	}

	log.Panic("Could not find guard!")
	return -1, -1, -1
}

func (s *Solution) checkBounds(r, c int) bool {
	if r < 0 || r >= s.rows {
		return false
	}

	if c < 0 || c >= s.cols {
		return false
	}

	return true
}

func (s *Solution) markVisited(r, c int) {
	s.visited[r][c] = true
}

func (s *Solution) simulateGuardPatrolling() {
	r, c, dir := s.findGuard()

	s.markVisited(r, c)

	for s.checkBounds(r, c) {
		nextR, nextC := getNext(r, c, dir)

		if s.checkBounds(nextR, nextC) && s.grid[nextR][nextC] == '#' {
			switch dir {
			case UP:
				dir = RIGHT
			case RIGHT:
				dir = DOWN
			case DOWN:
				dir = LEFT
			case LEFT:
				dir = UP
			}

			continue
		}

		r = nextR
		c = nextC

		if s.checkBounds(r, c) {
			s.markVisited(r, c)
		}
	}

}

func (s *Solution) getVisitedCount() int {
	count := 0

	for _, row := range s.visited {
		for _, col := range row {
			if col {
				count += 1
			}
		}
	}

	return count
}

func (s *Solution) Solve() int {
	if s.part == "A" {
		s.simulateGuardPatrolling()
		return s.getVisitedCount()
	} else {
		log.Panic("Not Implemented!")
		return 0
	}
}

func getNext(r, c, dir int) (int, int) {
	switch dir {
	case UP:
		return r - 1, c
	case RIGHT:
		return r, c + 1
	case DOWN:
		return r + 1, c
	case LEFT:
		return r, c - 1
	}

	log.Panicf("Invalid direction! [Pos (%d, %d)] [Dir %d]", r, c, dir)
	return -1, -1
}
