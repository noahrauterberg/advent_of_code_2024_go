package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInt(str string) int {
	ret, err := strconv.Atoi(str)
	checkErr(err)
	return ret
}

func abs(num int) int {
	if (num < 0) {
		return -num
	}
	return num
}

func part_one(l []int, r []int) {
	// copy arrays for further use
	left_arr := l
	right_arr := r
	
	// pairing numbers
	sort.Ints(left_arr)
	sort.Ints(right_arr)
	
	// compute total distance
	total_dist := 0
	for i := 0; i < len(left_arr); i++ {
		total_dist += abs(left_arr[i] - right_arr[i])
	}

	fmt.Printf("The total distance between the lists is %d\n", total_dist)
}

func part_two(l []int, r []int) {
	// copy arrays for further use
	left_arr := l
	right_arr := r

	count_map := make(map[int]int)
	for _, num := range right_arr {
		count_map[num]++
	}

	sim_score := 0
	for _, num := range left_arr {
		sim_score += num * count_map[num]
	}

	fmt.Printf("The similarity score is %d\n", sim_score)
}

func main() {
	data, err := os.ReadFile("input.txt")
	checkErr(err)
	lines := strings.Split(string(data), "\n")

	// read numbers into two lists
	var left_arr []int
	var right_arr []int
	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		if (len(numbers) == 2) {
			left_arr = append(left_arr, parseInt(numbers[0]))
			right_arr = append(right_arr, parseInt(numbers[1]))
		}
	}
	
	part_one(left_arr, right_arr)
	part_two(left_arr, right_arr)
}