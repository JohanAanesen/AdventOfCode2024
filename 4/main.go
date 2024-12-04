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

	input := make([][]string, 140)
	lineCounter := 0
	for r.Scan() {
		line := r.Text()
		input[lineCounter] = strings.Split(line, "")
		lineCounter++
	}

	xmasCounter := 0
	xmasCounter2 := 0
	// ->

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-3; j++ {
			//høyre
			if input[i][j] == "X" && input[i][j+1] == "M" && input[i][j+2] == "A" && input[i][j+3] == "S" {
				xmasCounter++
			}
			//venstre
			if input[i][j] == "S" && input[i][j+1] == "A" && input[i][j+2] == "M" && input[i][j+3] == "X" {
				xmasCounter++
			}

		}
	}

	for i := 0; i < len(input)-3; i++ {
		for j := 0; j < len(input); j++ {
			//ned
			if input[i][j] == "X" && input[i+1][j] == "M" && input[i+2][j] == "A" && input[i+3][j] == "S" {
				xmasCounter++
			}
			//opp
			if input[i][j] == "S" && input[i+1][j] == "A" && input[i+2][j] == "M" && input[i+3][j] == "X" {
				xmasCounter++
			}
		}
	}

	for i := 0; i < len(input)-3; i++ {
		for j := 0; j < len(input)-3; j++ {
			//høyre og ned
			if input[i][j] == "X" && input[i+1][j+1] == "M" && input[i+2][j+2] == "A" && input[i+3][j+3] == "S" {
				xmasCounter++
			}

			//venstre og opp
			if input[i][j] == "S" && input[i+1][j+1] == "A" && input[i+2][j+2] == "M" && input[i+3][j+3] == "X" {
				xmasCounter++
			}
		}
	}

	for i := 3; i < len(input); i++ {
		for j := 0; j < len(input)-3; j++ {

			if input[i][j] == "X" && input[i-1][j+1] == "M" && input[i-2][j+2] == "A" && input[i-3][j+3] == "S" {
				xmasCounter++
			}

			if input[i][j] == "S" && input[i-1][j+1] == "A" && input[i-2][j+2] == "M" && input[i-3][j+3] == "X" {
				xmasCounter++
			}

		}
	}

	fmt.Println("Part 1: ", xmasCounter)

	for i := 0; i < len(input)-2; i++ {
		for j := 0; j < len(input)-2; j++ {

			//MAS
			if input[i][j] == "M" && input[i][j+2] == "S" && input[i+1][j+1] == "A" && input[i+2][j] == "M" && input[i+2][j+2] == "S" {
				xmasCounter2++
			}

			if input[i][j] == "S" && input[i][j+2] == "S" && input[i+1][j+1] == "A" && input[i+2][j] == "M" && input[i+2][j+2] == "M" {
				xmasCounter2++
			}

			if input[i][j] == "S" && input[i][j+2] == "M" && input[i+1][j+1] == "A" && input[i+2][j] == "S" && input[i+2][j+2] == "M" {
				xmasCounter2++
			}

			if input[i][j] == "M" && input[i][j+2] == "M" && input[i+1][j+1] == "A" && input[i+2][j] == "S" && input[i+2][j+2] == "S" {
				xmasCounter2++
			}

		}
	}

	fmt.Println("Part 2: ", xmasCounter2)
}
