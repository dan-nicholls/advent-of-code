package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func ParseInputLists(input string) ([]int, []int, error) {
	listA := make([]int, 0)
	listB := make([]int, 0)

	for i, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		vals := strings.SplitN(line, "   ", 2)
		if len(vals) == 2 {
			val1, err := strconv.Atoi(vals[0])
			if err != nil {
				return []int{}, []int{}, fmt.Errorf("Unable to conv '%s' to int: %v", vals[0], err)
			}
			val2, err := strconv.Atoi(vals[1])
			if err != nil {
				return []int{}, []int{}, fmt.Errorf("Unable to conv '%s' to int: %v", vals[1], err)
			}

			listA = append(listA, val1)
			listB = append(listB, val2)
		} else {
			return []int{}, []int{}, fmt.Errorf("Seperated not found (line %v)", i)
		}
	}

	return listA, listB, nil
}

func FindTotalDif(listA, listB []int) int {
	sort.Ints(listA)
	sort.Ints(listB)

	totalDif := 0
	for i := range len(listA) {
		if listA[i] > listB[i] {
			dif := listA[i] - listB[i]
			totalDif += dif
		} else {
			dif := listB[i] - listA[i]
			totalDif += dif
		}
	}

	return totalDif
}

func main() {
	listA, listB, err := ParseInputLists(input)
	if err != nil {
		fmt.Errorf("Unable to parse Input Lists: %v", err)
	}

	result := FindTotalDif(listA, listB)
	fmt.Println("Final Result is: ", result)
}
