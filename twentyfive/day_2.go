package twentyfive

import (
	"fmt"
	"log"
	"strings"
	"strconv"
)

type productIDRange struct {
	firstID string
	lastID string
}

func parseRanges(input string) []productIDRange {
	v := strings.Split(input, ",")
	productRange := make([]productIDRange, 0, len(v))
	for _, pidr := range v {
		idrange := strings.Split(pidr, "-")
		productRange = append(productRange, productIDRange{
			firstID: idrange[0], 
			lastID: idrange[1],
		})
	}
	return productRange
}

func isInvalid(id string) bool {
	if len(id) % 2 != 0 {
		return false
	}

	mid := len(id)/2
	left := id[:mid]
	right := id[mid:]

	return left == right
}

func sumInvalidIds(pidr []productIDRange) (int, error) {
	sum := 0

	for _, p := range pidr {
		firstNum, err := strconv.Atoi(p.firstID)
		if err != nil {
			return 0, fmt.Errorf("invalid firstID in %q: %w", p.firstID, err)
		}
		lastNum, err := strconv.Atoi(p.lastID)
		if err != nil {
			return 0, fmt.Errorf("invalid lastID in %q: %w", p.lastID, err)
		}

		for i := firstNum; i < lastNum + 1; i++ {
			s := strconv.Itoa(i)
			if isInvalid(s) {
				sum += i
			}
		}
	}

	return sum, nil
}

func isRepeating(id string) bool {
	doubled := id + id
	n := len(id)
	for i := 1; i < n; i++ {
		substring := doubled[i:i+n]
		if substring == id {
			return true
		}
	}
	return false
}
func sumMoreInvalidIds(pidr []productIDRange) (int, error) {
	sum := 0

	for _, p := range pidr {
		firstNum, err := strconv.Atoi(p.firstID)
		if err != nil {
			return 0, fmt.Errorf("invalid firstID in %q: %w", p.firstID, err)
		}
		lastNum, err := strconv.Atoi(p.lastID)
		if err != nil {
			return 0, fmt.Errorf("invalid lastID in %q: %w", p.lastID, err)
		}

		for i := firstNum; i < lastNum + 1; i++ {
			s := strconv.Itoa(i)
			if isRepeating(s) {
				sum += i
			}
		}
	}

	return sum, nil
}

func DayTwo() {
	text, err := readFileWhole("inputs/2025_day_2.txt")
	if err != nil {
		log.Fatalf("could not read file: %s", err)
	}
	ids := parseRanges(text)
	sum, err := sumMoreInvalidIds(ids)
	if err != nil {
		log.Fatalf("could not sum ids: %s", err)
	}
	fmt.Println(sum)
}
