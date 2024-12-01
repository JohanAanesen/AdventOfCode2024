package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	listOne := make([]int, 0)
	listTwo := make([]int, 0)

	for r.Scan() {
		line := r.Text()
		tempArr := strings.Split(line, "   ")

		nrOne, _ := strconv.Atoi(tempArr[0])
		nrTwo, _ := strconv.Atoi(tempArr[1])
		listOne = append(listOne, nrOne)
		listTwo = append(listTwo, nrTwo)

	}

	sort.Ints(listOne)
	sort.Ints(listTwo)

	counter := 0

	for i := 0; i < len(listOne); i++ {
		counter += absDiffInt(listTwo[i], listOne[i])
	}

	fmt.Println("Part 1: ", counter)

	counter2 := 0

	for _, one := range listOne {
		tempCounter := 0
		for _, two := range listTwo {
			if two == one {
				tempCounter++
			}

		}
		counter2 += one * tempCounter
	}

	fmt.Println("Part 2: ", counter2)

}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
