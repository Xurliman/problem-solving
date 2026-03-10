package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mainTWO() {
	reader := bufio.NewReader(os.Stdin)

	// Read number of test cases
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read from stdin: %v", err)
	}
	testsNum, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		log.Fatalf("failed to convert to int: %v", err)
	}

	var res strings.Builder

	// Process all test cases recursively
	goThroughTestCases(reader, &res, testsNum)

	// Print all results
	fmt.Print(res.String())
}

// Recursive function to process test cases
func goThroughTestCases(reader *bufio.Reader, res *strings.Builder, num int) {
	if num == 0 {
		return
	}

	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read from stdin: %v", err)
	}

	testsCount, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		log.Fatalf("failed to convert to int: %v", err)
	}

	lineVals, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read integers from stdin: %v", err)
	}

	vals := strings.Split(strings.TrimSpace(lineVals), " ")
	if len(vals) != testsCount {
		log.Fatalf("expected %d integers, got %d", testsCount, len(vals))
	}

	// Calculate the powered sum of each val
	sum := goThroughArr(vals, 0)
	res.WriteString(fmt.Sprintf("%d\n", sum))

	// Recursive call for remaining test cases
	goThroughTestCases(reader, res, num-1)
}

// Recursive function to process an array of numbers
func goThroughArr(nums []string, k int) int {
	kn, err := strconv.Atoi(nums[k])
	if err != nil {
		log.Fatalf("failed to convert to int: %v - %v in [%v]", err, nums[k], k)
		return 0
	}

	sum := powerOfFour(kn)

	if k+1 < len(nums) {
		sum += goThroughArr(nums, k+1)
	}

	return sum
}

// Function to calculate fourth power for negative numbers
func powerOfFour(i int) int {
	if i > 0 {
		return 0
	}
	return i * i * i * i
}
