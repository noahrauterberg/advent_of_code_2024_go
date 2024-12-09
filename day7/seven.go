package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_int(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func test_operators(target int, start int, numbers []int) bool {
	if (len(numbers) == 0) {
		if (target == start) {
			return true
		}
		return false
	}
	return test_operators(target, start + numbers[0], numbers[1:]) || test_operators(target, start*numbers[0], numbers[1:])
}

func int_concat(i int, j int) int {
	str_i := strconv.Itoa(i)
	str_j := strconv.Itoa(j)

	var sb strings.Builder

	sb.WriteString(str_i)
	sb.WriteString(str_j)
	return parse_int(sb.String())
}

func test_operators_concat(target int, start int, numbers []int) bool {
	if (len(numbers) == 0) {
		if (target == start) {
			return true
		}
		return false
	}
	return test_operators_concat(target, start + numbers[0], numbers[1:]) || test_operators_concat(target, start*numbers[0], numbers[1:]) || test_operators_concat(target, int_concat(start, numbers[0]), numbers[1:])
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum_possible := 0
	sum_concat := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		result := parse_int(line[0])
		numbers := make([]int, len(strings.Split(line[1], " ")))
		for i, num := range strings.Split(line[1], " ") {
			numbers[i] = parse_int(num)
		}
		
		// part one
		if (test_operators(result, numbers[0], numbers[1:])){
			sum_possible += result
		}

		// part two
		if (test_operators_concat(result, numbers[0], numbers[1:])) {
			sum_concat += result
		}
	}

	fmt.Printf("The sum of all equations that could possibly be true is %d\n", sum_possible)
	fmt.Printf("The sum of all possibly correct equations including concatenation is %d\n", sum_concat)
}