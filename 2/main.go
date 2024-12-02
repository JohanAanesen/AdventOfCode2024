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
	counter2 := 0

	for i := 0; i < len(listOne); i++ {
		tempArr := strings.Split(listOne[i], " ")

		safe, _ := partOne(tempArr)
		if safe {
			counter++
		}

		if !safe {
			safesafe := false

			for j := 0; j < len(tempArr); j++ {
				s2 := make([]string, len(tempArr))
				_ = copy(s2, tempArr)
				safe2, _ := partOne(deleteElement(s2, j))

				if safe2 {
					safesafe = true
					continue
				}
			}

			if safesafe {
				counter2++
			}
		}

	}

	fmt.Println("Part 1: ", counter)
	fmt.Println("Part 2: ", counter2+counter)

}

func partOne(tempArr []string) (bool, int) {
	nrOne, _ := strconv.Atoi(tempArr[0])
	nrTwo, _ := strconv.Atoi(tempArr[len(tempArr)-1])

	safe := true
	unsafeCount := 0

	if nrOne < nrTwo { //increasing
		for i := 0; i < len(tempArr)-1; i++ {
			if !checkDif(tempArr[i], tempArr[i+1]) {
				safe = false
				unsafeCount++
			}
		}
	} else { //decreasing
		for i := len(tempArr) - 1; i > 0; i-- {
			if !checkDif(tempArr[i], tempArr[i-1]) {
				safe = false
				unsafeCount++
			}
		}
	}
	return safe, unsafeCount
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

func deleteElement(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
