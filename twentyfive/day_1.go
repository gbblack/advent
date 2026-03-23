package twentyfive

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func initArr(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func parseLine(line string) (string, string) {
	direction := line[0:1]
	distance := line[1:]
	return direction, distance
}

func move(current_idx, k, n int) int {
	return (current_idx + k%n + n) % n
}

func processMoves(lines []string) (int, error) {
	arr := initArr(100)
	currIdx := 50
	count := 0

	for _, line := range lines {
		direction, distance := parseLine(line)
		k, err := strconv.Atoi(distance)
		if err != nil {
			return 0, err
		}

		var newIdx int
		switch direction {
		case "R":
			newIdx = move(currIdx, k, len(arr))
		case "L":
			newIdx = move(currIdx, -k, len(arr))
		}

		if arr[newIdx] == 0 {
			count++
		}

		currIdx = newIdx
	}
	return count, nil
}

func DayOne() {
	lines, err := readFile("inputs/2025_day_1.txt")
	if err != nil {
		log.Fatalf("could not read file: %s", err)
	}

	count, err := processMoves(lines)
	if err != nil {
		log.Fatalf("coould not process moces: %s", err)
	}

	fmt.Printf("Password: %d\n", count)
}