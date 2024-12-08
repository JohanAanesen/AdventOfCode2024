package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func main() {
	ROWLEN := 50

	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	input := make([][]string, ROWLEN)
	lineCounter := 0
	for r.Scan() {
		line := r.Text()
		input[lineCounter] = strings.Split(line, "")
		lineCounter++
	}

	antennas := make(map[string][]Coordinate, 0)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if input[i][j] != "." {
				antennas[input[i][j]] = append(antennas[input[i][j]], Coordinate{x: i, y: j})
			}
		}
	}

	antiNodes := make(map[Coordinate]bool)

	for _, antenna := range antennas {
		for i := 0; i < len(antenna); i++ {
			for j := 0; j < len(antenna); j++ {
				if i == j {
					continue
				}

				antiNodes[Coordinate{x: antenna[i].x, y: antenna[i].y}] = true
				antiNodes[Coordinate{x: antenna[j].x, y: antenna[j].y}] = true

				difX := antenna[j].x - antenna[i].x
				difY := antenna[j].y - antenna[i].y

				if 0 <= antenna[i].x-difX && antenna[i].x-difX < ROWLEN && 0 <= antenna[i].y-difY && antenna[i].y-difY < ROWLEN {
					antiNodes[Coordinate{x: antenna[i].x - difX, y: antenna[i].y - difY}] = true
				}

				if 0 <= antenna[j].x+difX && antenna[j].x+difX < ROWLEN && 0 <= antenna[j].y+difY && antenna[j].y+difY < ROWLEN {
					antiNodes[Coordinate{x: antenna[j].x + difX, y: antenna[j].y + difY}] = true
				}

				newLocationX := antenna[i].x - difX
				newLocationY := antenna[i].y - difY

				for 0 <= newLocationX-difX && newLocationX-difX < ROWLEN && 0 <= newLocationY-difY && newLocationY-difY < ROWLEN {
					newLocationX = newLocationX - difX
					newLocationY = newLocationY - difY
					antiNodes[Coordinate{x: newLocationX, y: newLocationY}] = true
				}

				newLocationX = antenna[j].x + difX
				newLocationY = antenna[j].y + difY

				for 0 <= newLocationX+difX && newLocationX+difX < ROWLEN && 0 <= newLocationY+difY && newLocationY+difY < ROWLEN {
					newLocationX = newLocationX + difX
					newLocationY = newLocationY + difY
					antiNodes[Coordinate{x: newLocationX, y: newLocationY}] = true
				}

			}
		}
	}

	fmt.Println("Part 1: ", len(antiNodes))

}
