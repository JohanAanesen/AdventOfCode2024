package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	input := make([][]string, 130)
	lineCounter := 0
	for r.Scan() {
		line := r.Text()
		input[lineCounter] = strings.Split(line, "")
		lineCounter++
	}

	path := make([][]string, 130)
	x, y := 0, 0
	x2, y2 := 0, 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {

			if input[i][j] == "^" {
				x = j
				x2 = j
				y = i
				y2 = j
				input[y][x] = "X"

			}
		}
		path[i] = append(path[i], input[i]...)
	}

	dx, dy := 0, -1

	for 0 < x+dx && x+dx < len(input) && 0 < y+dy && y+dy < len(input) {

		if input[y+dy][x+dx] == "#" {
			//turn 90
			if dx == 0 && dy == -1 {
				dx, dy = 1, 0
			} else if dx == 1 && dy == 0 {
				dx, dy = 0, 1
			} else if dx == 0 && dy == 1 {
				dx, dy = -1, 0
			} else if dx == -1 && dy == 0 {
				dx, dy = 0, -1
			}
		}
		y += dy
		x += dx
		input[y][x] = "X"
	}

	counter := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {

			if input[i][j] == "X" {
				counter++
			}
		}
	}

	fmt.Println("Part 1: ", counter)

	fmt.Println("Part 2: ", part2(input, x2, y2))

}

func part2(path [][]string, x int, y int) int {

	infiniteCounter := 0

	for i := 0; i < len(path); i++ {
		for j := 0; j < len(path); j++ {

			changeback := false
			if path[i][j] == "X" {
				path[i][j] = "#"
				changeback = true
			} else {
				continue
			}

			//visited path
			visited := make(map[string]bool, 0)
			ix, iy := x, y
			dx, dy := 0, -1
			dir := "U"

			for 0 < ix+dx && ix+dx < len(path) && 0 < iy+dy && iy+dy < len(path) {

				if path[iy+dy][ix+dx] == "#" {
					//turn 90
					if dx == 0 && dy == -1 {
						dx, dy = 1, 0
						dir = "R"
					} else if dx == 1 && dy == 0 {
						dx, dy = 0, 1
						dir = "D"
					} else if dx == 0 && dy == 1 {
						dx, dy = -1, 0
						dir = "L"
					} else if dx == -1 && dy == 0 {
						dx, dy = 0, -1
						dir = "U"
					}
				}

				iy += dy
				ix += dx

				if visited[strconv.Itoa(ix)+"&"+strconv.Itoa(iy)+dir] {
					infiniteCounter++
					break
				}

				visited[strconv.Itoa(ix)+"&"+strconv.Itoa(iy)+dir] = true

			}

			if changeback {
				path[i][j] = "X"
			}

		}
	}

	return infiniteCounter
}
