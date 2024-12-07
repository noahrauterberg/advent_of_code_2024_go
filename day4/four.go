package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func search_in_string(data string, query string) int {
	regex := regexp.MustCompile(query)
	return len(regex.FindAllString(data, -1))
}

func part_one(lines []string) int {
	count := 0

	// horizontal
	for _, line := range lines {
		count += search_in_string(line, "XMAS") + search_in_string(line, "SAMX")
	}

	// vertical
	var vertical_strings []strings.Builder
	for i := 0; i < len(lines[0]); i++ {
		vertical_strings = append(vertical_strings, strings.Builder{})
	}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			vertical_strings[j].WriteString(string(lines[i][j]))
		}
	}

	for _, line := range vertical_strings {
		count += search_in_string(line.String(), "XMAS") + search_in_string(line.String(), "SAMX")
	}

	// diagonal NW <-> SE
	var diagnonal []strings.Builder
	for i := 0; i < len(lines); i++ {
		var sb strings.Builder
		for char_x, char_y := 0, i; char_x < len(lines[0]) && char_y < len(lines); char_x, char_y = char_x+1, char_y+1 {
			sb.WriteString(string(lines[char_y][char_x]))
		}
		diagnonal = append(diagnonal, sb)
	}

	for i := 0; i < len(lines[0]); i++ {
		var sb strings.Builder
		for char_x, char_y := i+1, 0; char_x < len(lines[0]) && char_y < len(lines); char_x, char_y = char_x+1, char_y+1 {
			sb.WriteString(string(lines[char_y][char_x]))
		}
		diagnonal = append(diagnonal, sb)
	}

	// diagonal SW <-> NE
	for i := len(lines) - 1; i >= 0; i-- {
		var sb strings.Builder
		for char_x, char_y := 0, i; char_x < len(lines[0]) && char_y >= 0; char_x, char_y = char_x+1, char_y-1 {
			sb.WriteString(string(lines[char_y][char_x]))
		}
		diagnonal = append(diagnonal, sb)
	}

	for i := 0; i < len(lines[0]); i++ {
		var sb strings.Builder
		for char_x, char_y := i+1, len(lines)-1; char_x < len(lines[0]) && char_y >= 0; char_x, char_y = char_x+1, char_y-1 {
			sb.WriteString(string(lines[char_y][char_x]))
		}
		diagnonal = append(diagnonal, sb)
	}

	for _, line := range diagnonal {
		count += search_in_string(line.String(), "XMAS") + search_in_string(line.String(), "SAMX")
	}
	return count
}

func check_X(x int, y int, grid [][]byte) bool {
	width := len(grid[0])
	height := len(grid)
	if (y == 0 || x == 0 || y+1 >= height || x+1 >= width) {
		return false
	}
	
	var NE_SW strings.Builder
	var NW_SE strings.Builder
	NE_SW.WriteByte(grid[y-1][x+1])
	NE_SW.WriteByte(grid[y][x])
	NE_SW.WriteByte(grid[y+1][x-1])

	NW_SE.WriteByte(grid[y-1][x-1])
	NW_SE.WriteByte(grid[y][x])
	NW_SE.WriteByte(grid[y+1][x+1])

	if (NE_SW.String() == "MAS" || NE_SW.String() == "SAM") && (NW_SE.String() == "MAS" || NW_SE.String() == "SAM") {
		return true
	}
	return false
}

func part_two(grid [][]byte) int {
	count := 0
	for y, row := range grid {
		for x, char := range row {
			if (char == 'A' && check_X(x, y, grid)) {
				count++
			}
		}
	}
	return count
}

func main()  {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	count_one := part_one(lines)
	count_two := part_two(grid)

	fmt.Printf("There is a total of %d occurrences of 'XMAS'\n", count_one)
	fmt.Printf("There is a total of %d occurrences of 'MAS' forming an X\n", count_two)
}
