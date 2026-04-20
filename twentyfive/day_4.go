package twentyfive

import (
	"fmt"
	"log"
	"strings"
)

func countNeighbours(grid [][]string, i, j int) int {
	offsets := [][2]int {
		{-1,-1},{-1,0},{-1,1},
		{0,-1},{0,1},
		{1,-1},{1,0},{1,1},
	}

	count := 0

	rows := len(grid)
	cols := len(grid[0])

	for _, offset := range offsets {
		delta_i := offset[0]
		delta_j := offset[1]

		k := i + delta_i
		l := j + delta_j

		// boundary check
		if k >= 0 && k < rows && l >= 0 && l < cols {
			if grid[k][l] == "@" {
				count++
			}
		}
	}

	return count
}

func traverse(grid [][]string) int {
	count := 0

	for i := range grid {
		for j := range grid[i] {

			if grid[i][j] != "@" {
				continue
			}

			if countNeighbours(grid, i, j) < 4 {
				count++
			}
		}
	}

	return count
}

func findRemovals(grid [][]string) [][2]int {
	toRemove := make([][2]int, 0)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != "@" {
				continue
			}

			if countNeighbours(grid, i, j) < 4 {
				toRemove = append(toRemove, [2]int{i, j})
			}
		}
	}

	return toRemove
}
func applyRemovals(grid [][]string, cells [][2]int) {
	for _, c := range cells {
		i, j := c[0], c[1]
		grid[i][j] = "x"
	}
}

func simulate(grid [][]string) int {
	totalRemoved := 0

	for {
		removals := findRemovals(grid)

		if len(removals) == 0 {
			break
		}

		applyRemovals(grid, removals)
		totalRemoved += len(removals)
	}

	return totalRemoved
}

func parsePositions(lines []string) [][]string {
	positions := make([][]string, 0, len(lines))
	for _, line := range lines {
		row := strings.Split(line, "")
		positions = append(positions, row)
	}
	return positions
}
func DayFour() {
	lines, err := readFileLineByLine("inputs/2025_day_4.txt")
	if err != nil {
		log.Fatalf("could not read file: %s", err)
	}
	grid := parsePositions(lines)
	reachablePoints := simulate(grid)

	fmt.Printf("Forlift can access: %d\n", reachablePoints)
}