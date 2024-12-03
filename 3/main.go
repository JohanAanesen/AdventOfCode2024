package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	counter := 0
	counter2 := 0
	do := true

	for r.Scan() {
		line := r.Text()

		//mul\(([0-9]|[1-9][0-9]|[1-9][0-9][0-9]),([0-9]|[1-9][0-9]|[1-9][0-9][0-9])\)

		r := regexp.MustCompile(`mul\(([0-9]|[1-9][0-9]|[1-9][0-9][0-9]),([0-9]|[1-9][0-9]|[1-9][0-9][0-9])\)`)
		matches := r.FindAllStringSubmatch(line, -1)
		for _, v := range matches {
			nrOne, _ := strconv.Atoi(v[1])
			nrTwo, _ := strconv.Atoi(v[2])
			counter += nrOne * nrTwo
		}

		// mul\(([0-9]|[1-9][0-9]|[1-9][0-9][0-9]),([0-9]|[1-9][0-9]|[1-9][0-9][0-9])\)|(do\(\))*(don't\(\))*

		r2 := regexp.MustCompile(`mul\(([0-9]|[1-9][0-9]|[1-9][0-9][0-9]),([0-9]|[1-9][0-9]|[1-9][0-9][0-9])\)|do\(\)|don't\(\)`)
		matches2 := r2.FindAllStringSubmatch(line, -1)
		for _, v := range matches2 {
			if v[0] == "do()" {
				do = true
			} else if v[0] == "don't()" {
				do = false
			} else if do {
				nrOne, _ := strconv.Atoi(v[1])
				nrTwo, _ := strconv.Atoi(v[2])
				counter2 += nrOne * nrTwo
			}
		}
	}

	fmt.Println("Part 1: ", counter)
	fmt.Println("Part 2: ", counter2)

}
