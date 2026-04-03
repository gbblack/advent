package twentyfive

import (
	"fmt"
	"log"
	"strconv"
)

func maxJoltagePairs(bank []int) int {
	max := 0

	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			val := bank[i]*10 + bank[j]
			if val > max {
				max = val
			}
		}
	}
	return max
}

//greedy stack
func maxJoltage(bank []int, k int) int {
	drop := len(bank) - k
	stack := make([]int, 0, len(bank))

	for _, digit := range bank {
		for len(stack) > 0 && drop > 0 && stack[len(stack) - 1] < digit {
			stack = stack[:len(stack)-1]
			drop--
		}
		stack = append(stack, digit)
	}
	stack = stack[:k]
	result := 0
	for _, d := range stack {
		result = result*10 + d
	}
	return result
}
func findTotalJoltage(banks [][]int) int {
	sum := 0
	for _, bank := range banks {
		sum += maxJoltagePairs(bank)
	}
	return sum
}
func findTotalJoltageBig(banks [][]int, k int) int {
	sum := 0
	for _, bank := range banks {
		sum += maxJoltage(bank, k)
	}
	return sum
}

func stringToInt(line string) ([]int, error) {
	ints := make([]int, 0, len(line))
	for _, i := range line {
		n, err := strconv.Atoi(string(i))
		if err != nil {
			return nil, err
		}
		ints = append(ints, n)
	}
	return ints, nil
}

func parseBanks(lines []string) ([][]int, error) {
	banks := make([][]int, 0, len(lines))
	for _, line := range lines {
		bank, err := stringToInt(line)
		if err != nil {
			return nil, err
		}
		banks = append(banks, bank)
	}
	return banks, nil
}

func DayThree() {
	lines, err := readFileLineByLine("inputs/2025_day_3.txt")
	if err != nil {
		log.Fatalf("could not read file: %s", err)
	}
	banks, err := parseBanks(lines)
	if err != nil {
		log.Fatalf("could not parse banks: %s", err)
	}
	totalJoltagePairs := findTotalJoltage(banks)
	totalJoltage := findTotalJoltageBig(banks, 12)
	fmt.Printf("Part 1 Joltage: %d\n", totalJoltagePairs)
	fmt.Printf("Part 2 Joltage: %d\n", totalJoltage)
}