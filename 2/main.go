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

	listOne := make([]string, 0)

	for r.Scan() {
		line := r.Text()
		listOne = append(listOne, line)

	}

	counter := 0

	for i := 0; i < len(listOne); i++ {
		tempArr := strings.Split(listOne[i], " ")
		nrOne, _ := strconv.Atoi(tempArr[0])
		nrTwo, _ := strconv.Atoi(tempArr[1])

		safe := true

		if nrOne < nrTwo { //increasing
			for i := 0; i < len(tempArr)-1; i++ {
				if !checkDif(tempArr[i], tempArr[i+1]) {
					safe = false
				}
			}
		} else { //decreasing
			for i := len(tempArr) - 1; i > 0; i-- {
				if !checkDif(tempArr[i], tempArr[i-1]) {
					safe = false
				}
			}
		}

		if safe {
			counter++
		}

	}

	fmt.Println("Part 1: ", counter)

}

func checkDif(one string, two string) bool {
	nrOne, _ := strconv.Atoi(one)
	nrTwo, _ := strconv.Atoi(two)

	dif := nrTwo - nrOne

	if dif >= 1 && dif <= 3 {
		return true
	}

	return false
}
