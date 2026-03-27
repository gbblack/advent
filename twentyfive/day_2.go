package twentyfive

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type productIDRange struct {
	first, last int
}

func parseRanges(input string) ([]productIDRange, error) {
	parts := strings.Split(input, ",")
	ranges := make([]productIDRange, 0, len(parts))
	for _, p := range parts {
		ends := strings.Split(p, "-")
		first, err := strconv.Atoi(ends[0])
		if err != nil {
			return nil, fmt.Errorf("invalid range start in %q: %w", p, err)
		}
		last, err := strconv.Atoi(ends[1])
		if err != nil {
			return nil, fmt.Errorf("invalid range end in %q: %w", p, err)
		}
		ranges = append(ranges, productIDRange{first, last})
	}
	return ranges, nil
}

func isDuplicated(id string) bool {
	n := len(id)
	if n%2 != 0 {
		return false
	}
	return id[:n/2] == id[n/2:]
}

func isRepeating(id string) bool {
	doubled := id + id
	pos := strings.Index(doubled[1:], id)
	if pos == -1 {
		return false
	}
	return pos < len(id)-1
}

func sumIDs(ranges []productIDRange, invalid func(string) bool) int {
	sum := 0
	for _, r := range ranges {
		for i := r.first; i <= r.last; i++ {
			if invalid(strconv.Itoa(i)) {
				sum += i
			}
		}
	}
	return sum
}

func DayTwo() {
	text, err := readFileWhole("inputs/2025_day_2.txt")
	if err != nil {
		log.Fatalf("could not read file: %s", err)
	}
	ranges, err := parseRanges(text)
	if err != nil {
		log.Fatalf("could not parse input: %s", err)
	}
	fmt.Printf("Part 1 Sum: %d\n", sumIDs(ranges, isDuplicated))
	fmt.Printf("Part 2 Sum: %d\n", sumIDs(ranges, isRepeating))
}
