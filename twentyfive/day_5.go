package twentyfive

import (
	"fmt"
	"log"
	"strings"
	"strconv"
)

type  idRange struct {
	min int
	max int
}

func isFresh(id int, r idRange) bool {
	return id >= r.min && id <= r.max
}

func freshIds(ids []int, idr []idRange) int {
	freshIds := make([]int, 0)
	for _, id := range ids {
		for _, r := range idr {
			if isFresh(id, r) {
				freshIds = append(freshIds, id)
				break
			}
		}
	}
	return len(freshIds)
}

func allFresh(idr []idRange) int {
	setFresh := map[int]struct{}{}
	for _, r := range idr {
		i := r.min
		for i <= r.max {
			if _, ok := setFresh[i]; !ok {
				setFresh[i] = struct{}{}
			}
			i++
		}
	}
	return len(setFresh)
}

func getValues(lines []string) ([]idRange, []int) {
	idRanges := make([]idRange, 0)
	ids := make([]int, 0)
	for _, line := range lines {
		value := strings.Split(line, "-")
		if value[0] == "" {
			continue
		}
		if len(value) < 2 {
			id, err := strconv.Atoi(value[0])
			if err != nil {
				log.Fatalf("string can't be cast into number: %s", err)
			}
			
			ids = append(ids, id)
			continue
		}
		min, err := strconv.Atoi(value[0])
		if err != nil {
			log.Fatalf("string can't be cast into number: %s", err)
		}
		max, err := strconv.Atoi(value[1])
		if err != nil {
			log.Fatalf("string can't be cast into number: %s", err)
		}
		idRanges = append(idRanges, idRange{min, max})
	}

	return idRanges, ids
}


func DayFive() {
	// lines, err := readFileLineByLine("inputs/2025_day_5.txt")
	lines, err := readFileLineByLine("inputs/example.txt")
	if err != nil {
		log.Fatalf("could not read file: %s", err)
	}
	idRanges, ids := getValues(lines)
	numberFresh := freshIds(ids, idRanges)
	allFresh := allFresh(idRanges)
	fmt.Println(numberFresh)
	fmt.Println(allFresh)
}