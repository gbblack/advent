package twentyfive

import (
	"fmt"
	"log"
	"strconv"
)

type rotation struct {
	direction rune
	distance  int
}

func parseRotations(input []string) ([]rotation, error) {
	rotations := make([]rotation, 0, len(input))
	for _, line := range input {
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid distance in %q: %w", line, err)
		}
		rotations = append(rotations, rotation{direction: rune(line[0]), distance: dist})
	}
	return rotations, nil
}

func move(position, k, n int) int {
	return (position + k%n + n) % n
}

func dialJump(rotations []rotation, dialSize, start int) int {
	position := start
	count := 0

	for _, r := range rotations {
		sign := 1
		if r.direction == 'L' {
			sign = -1
		}
		position = move(position, sign*r.distance, dialSize)
		if position == 0 {
			count++
		}
	}
	return count
}

func dialWalk(rotations []rotation, dialSize, start int) int {
	position := start
	count := 0

	for _, r := range rotations {
		sign := 1
		if r.direction == 'L' {
			sign = -1
		}
		for range r.distance {
			position = move(position, sign, dialSize)
			if position == 0 {
				count++
			}
		}
	}
	return count
}

func DayOne() {
	lines, err := readFileLineByLine("inputs/2025_day_1.txt")
	if err != nil {
		log.Fatalf("could not read file: %s", err)
	}

	rotations, err := parseRotations(lines)
	if err != nil {
		log.Fatalf("could not parse input: %s", err)
	}
	fmt.Printf("Part 1 Password: %d\n", dialJump(rotations, 100, 50))
	fmt.Printf("Part 2 Password: %d\n", dialWalk(rotations, 100, 50))
}
