package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func ParseInputReports(input string) []string {
	return strings.Split(input, "\n")
}

func ParseInputReport(input string) ([]int, error) {
	vals := strings.Split(input, " ")
	nums := make([]int, 0)

	for _, val := range vals {
		num, err := strconv.Atoi(val)
		if err != nil {
			return []int{}, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func IsReportSafe(report string) (bool, error) {
	vals, _ := ParseInputReport(report)
	isIncreasing := vals[0] < vals[1]

	for i := 1; i < len(vals); i++ {
		diff := vals[i] - vals[i-1]

		if isIncreasing {
			if !(diff >= 1 && diff <= 3) {
				return false, nil
			}
		} else {
			if !(diff <= -1 && diff >= -3) {
				return false, nil
			}
		}
	}
	return true, nil
}

func CountSafeReports(reports []string) int {
	count := 0

	for _, report := range reports {
		if len(report) == 0 {
			continue
		}
		isValid, _ := IsReportSafe(report)
		if isValid {
			count += 1
		}
	}
	return count
}

func main() {
	fmt.Println("Counting Safe Reports")
	reports := ParseInputReports(input)
	result := CountSafeReports(reports)
	fmt.Println("Total safe reports: ", result)
}
