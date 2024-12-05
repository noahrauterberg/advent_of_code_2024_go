package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInt(str string) int {
	ret, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return ret
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// the return value of type int indicates the position of the first error encountered 
// -> not needed for part one 
func isReportSafe(levels []int) (bool, int) {
	increasing := false
	decreasing := false
	first_dif := levels[0] - levels[1]
	if first_dif == 0 {return false, 0}
	if first_dif < 0 {
		decreasing = true
	} else {
		increasing = true
	}

	for i := 0; i < len(levels)-1; i++ {
		cur_diff := levels[i] - levels[i+1]

		if abs(cur_diff) > 3 || cur_diff == 0 || (increasing && cur_diff < 0) || (decreasing && cur_diff > 0) {
			return false, i
		}
	}
	return true, -1
}

func listMissingI(arr []int, i int) []int {
	// slicing always leads to ugly code due to additional boundary checks
	var list_missing_i []int
	for j := 0; j < len(arr); j++ {
		if (j != i) {
			list_missing_i = append(list_missing_i, arr[j])
		}
	}

	return list_missing_i
}

func isReportSafeModified(levels []int) bool {
	safe, first_err := isReportSafe(levels)
	if safe {
		return true
	}

	// sufficient to test around i
	levels_im1 := listMissingI(levels, first_err-1)
	levels_i := listMissingI(levels, first_err)
	levels_ip1 := listMissingI(levels, first_err+1)

	safe_im1, _ := isReportSafe(levels_im1)
	safe_i, _ := isReportSafe(levels_i)
	safe_ip1, _ := isReportSafe(levels_ip1)

	if safe_im1 || safe_i || safe_ip1 {
		return true
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	safe_reports := 0
	safe_reports_modified := 0
	for scanner.Scan() {
		line := scanner.Text()
		var numbers []int
		for _, numStr := range strings.Split(line, " "){
			numbers = append(numbers, parseInt(numStr))
		}
		
		// part one
		safe, _ := isReportSafe(numbers) 
		if safe {
			safe_reports++
		}

		// part two
		if isReportSafeModified(numbers) {
			safe_reports_modified++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safe_reports)
	fmt.Printf("Number of safe reports with possibly one bad level: %d\n", safe_reports_modified)
}