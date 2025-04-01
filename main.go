package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	col1, col2, err := loadInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	distances := calculateDistances(col1, col2)
	fmt.Println("Total distance:", distances)

	similarity := calculateSimilarity(col1, col2)
	fmt.Println("Total similarity:", similarity)
}

func calculateDistances(col1 []int, col2 []int) int {
	sort.Ints(col1)
	sort.Ints(col2)

	var totalDistance int
	for i := 0; i < len(col1); i++ {
		distance := col1[i] - col2[i]
		if distance < 0 {
			distance *= -1
		}
		totalDistance += distance
	}

	return totalDistance
}

func calculateSimilarity(col1 []int, col2 []int) int {
	occurrencies := make(map[int]int)

	for _, num := range col2 {
		occurrencies[num]++
	}

	similarity := 0

	for _, num := range col1 {
		if occ, exists := occurrencies[num]; exists {
			similarity += num * occ
		}
	}

	return similarity
}

func loadInput(fileName string) ([]int, []int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var col1, col2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		num1, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing first column: %v", err)
		}

		// Parse second column
		num2, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing second column: %v", err)
		}

		// Append to slices
		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %s", fileName)
	}

	if len(col1) != len(col2) {
		return nil, nil, fmt.Errorf("columns have different lengths")
	}

	return col1, col2, nil
}
