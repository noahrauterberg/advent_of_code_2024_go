package main

import (
	"fmt"
	"os"
	"regexp"
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

func part_one(data string) int {
	mul_regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	digits_regex := regexp.MustCompile(`\d+,\d+`)

	operations := mul_regex.FindAllString(data, -1)
	sum := 0
	for _, operation := range operations {
		match := digits_regex.FindAllString(operation, 1)[0]
		num1 := parseInt(strings.Split(match, ",")[0])
		num2 := parseInt(strings.Split(match, ",")[1])

		sum += num1 * num2
	}
	return sum
}

func part_two(input string) int {
	// prefix input with do()
	var sb strings.Builder
	sb.WriteString("do()")
	sb.WriteString(input)
	data := sb.String()

	activated_regex := regexp.MustCompile(`do\(\)|don't\(\)`)
	splits := activated_regex.FindAllStringIndex(data, -1)

	sum := 0
	for i:=0; i < len(splits)-1; i++ {
		sub_string := data[splits[i][0]:splits[i+1][0]]
		if strings.HasPrefix(sub_string, "don't()") {
			continue
		}
		sum += part_one(sub_string)
	}
	// check last match outside the loop
	last_sub_string := data[splits[len(splits)-1][0]:]
	if (strings.HasPrefix(last_sub_string, "do()")) {
		return sum + part_one(last_sub_string)
	}
	return sum
}

func main () {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	sum := part_one(string(data))

	activated_sum := part_two(string(data))

	fmt.Printf("The sum of all multiplications is %d\n", sum)
	fmt.Printf("The sum of all activated multiplications is %d\n", activated_sum)
}

