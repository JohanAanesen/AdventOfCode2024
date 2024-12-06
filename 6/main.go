package main

import (
	"bufio"
	"fmt"
	"os"
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

	x, y := 0, 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {

			if input[i][j] == "^" {
				x = j
				y = i
				input[y][x] = "X"
				break
			}
		}
	}

	dx, dy := 0, -1

	for 0 < x+dx && x+dy < len(input) && 0 < y+dx && y+dy < len(input) {

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

}
