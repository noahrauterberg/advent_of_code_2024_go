package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_int(in string) int {
	ret, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return ret
}

func parse_rules(rules []string) map[int][]int {
	ret := make(map[int][]int)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		before := parse_int(parts[0])
		after := parse_int(parts[1])
		ret[after] = append(ret[after], before)
	}

	return ret
}

func parse_schedules(schedules []string) [][]int {
	ret := make([][]int, len(schedules))
	for i, schedule := range schedules {
		parts := strings.Split(schedule, ",")
		for _, num := range parts {
			ret[i] = append(ret[i], parse_int(num))
		}
	}
	
	return ret
}

// If not valid, returns an array where the first number is the index of the first violation
func check_schedule(schedule []int, rules map[int][]int) (bool, []int) {
	for i, page_num := range schedule {
		cur_rules := rules[page_num]
		for j:=i; j < len(schedule); j++ {
			for _, rule := range cur_rules {
				if (schedule[j] == rule) {
					return false, []int{i, j}
				}
			}
		}
	}
	
	return true, nil
}

func fix_schedule(schedule []int, violation []int) []int {
	correct := make([]int, len(schedule))
	copy(correct, schedule)
	correct[violation[0]] = schedule[violation[1]]
	correct[violation[1]] = schedule[violation[0]]

	return correct
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	in := strings.SplitN(string(data), "\n\n", 2)
	
	rules := parse_rules(strings.Split(in[0], "\n"))
	schedules := parse_schedules(strings.Split(in[1], "\n"))

	sum := 0
	sum_reordered := 0
	for _, schedule := range schedules {
		valid, violation := check_schedule(schedule, rules)
		if (valid) { // part one
			middle_page := schedule[(len(schedule)-1)/2]
			sum += middle_page
		} else { // part two
			new_schedule := fix_schedule(schedule, violation)
			valid, violation := check_schedule(new_schedule, rules)
			for !valid {
				new_schedule = fix_schedule(new_schedule, violation)
				valid, violation = check_schedule(new_schedule, rules)
			}
			middle_page := new_schedule[(len(schedule)-1)/2]
			sum_reordered += middle_page
		}
	}

	fmt.Printf("The sum of the middle page numbers for correctly-ordered schedules is %d\n", sum)
	fmt.Printf("The sum of the middle page numbers for re-ordered schedules is %d\n", sum_reordered)
}
