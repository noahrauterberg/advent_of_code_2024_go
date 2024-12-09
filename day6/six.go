package main

import (
	"fmt"
	"os"
	"strings"
)

var grid [][]byte
// depends on the input
var guard = byte('^')
var visited_positions = make(map[int][]byte)

func check_valid_position(next_x int, next_y int) bool {
	if (next_x >= 0 ) && (next_x < len(grid[0])) && (next_y >= 0) && (next_y < len(grid)) {
		return true
	}
	return false
}

func move_guard(x int, y int) (int, int, bool) {
	grid[y][x] = 'X'
	add_position(x, y)

	switch guard {
	case '^':
		if !check_valid_position(x, y-1) {
			return x, y-1, false
		}
		if (grid[y-1][x] == '#' || grid[y-1][x] == 'O') {
			guard = '>'
			return x, y, true
		}
		return x, y-1, true
	case '<':
		if !check_valid_position(x-1, y) {
			return x-1, y, false
		}
		if (grid[y][x-1] == '#' || grid[y][x-1] == 'O') {
			guard = '^'
			return x, y, true
		}
		return x-1, y, true
	case '>':
		if !check_valid_position(x+1, y) {
			return x+1, y, false
		}
		if (grid[y][x+1] == '#' || grid[y][x+1] == 'O') {
			guard = 'v'
			return x, y, true
		}
		return x+1, y, true
	case 'v':
		if !check_valid_position(x, y+1) {
			return x, y+1, false
		}
		if (grid[y+1][x] == '#' || grid[y+1][x] == 'O') {
			guard = '<'
			return x, y, true
		}
		return x, y+1, true
	}
	panic("Guard value not valid")
}

func add_position(x int, y int) {
	key := len(grid[0])*y + x
	direction, ok := visited_positions[key]
	if (ok) {
		visited_positions[key] = append(direction, guard)
	} else {
		visited_positions[key] = []byte{guard}
	}
}

func check_loop(x int, y int) bool {
	directions, _ := visited_positions[len(grid[0])*y+x]
	for _, dir := range directions {
		if guard == dir {
			return true
		}
	}

	return false
}

func block_position(x int, y int, start_x int, start_y int) bool {
	if (x == start_x && y == start_y) {
		return false
	}
	guard = '^'
	grid[y][x] = '#'
	visited_positions = make(map[int][]byte)

	next_x, next_y, continue_move := move_guard(start_x, start_y)
	for continue_move {
		if (check_loop(next_x, next_y)) {
			grid[y][x] = 'X'
			return true
		}
		next_x, next_y, continue_move = move_guard(next_x, next_y)
	}

	grid[y][x] = 'X'
	return false
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	grid = make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	// find starting position
	var start_x int
	var start_y int
	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[i]); j++ {
			if (grid[i][j] == guard) {
				start_y, start_x = i, j
			}
		}
	}

	next_x, next_y, continue_move := move_guard(start_x, start_y)
	for continue_move {
		next_x, next_y, continue_move = move_guard(next_x, next_y)
	}

	// part one
	distinct_positions := 0
	original_positions := make([][]int, 0)
	for y, row := range grid {
		for x, char := range row {
			if (char == 'X') {
				distinct_positions++
				original_positions = append(original_positions, []int{x, y})
			}
		}
	}
	
	// part two
	possible_loops := 0
	for _, pos := range original_positions {
		if block_position(pos[0], pos[1], start_x, start_y) {
			possible_loops++
		}
	}
	
	fmt.Printf("The guard visits %d unique positions before leaving the map\n", distinct_positions)
	fmt.Printf("There are %d many possibilities to send the guard to a loop\n", possible_loops)
}