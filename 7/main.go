package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	result      int
	calibration []int
}

func main() {

	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	counter := 0

	var equations []Equation

	for r.Scan() {
		line := r.Text()
		lineSplit := strings.Split(line, ":")
		result, _ := strconv.Atoi(lineSplit[0])

		var listOfNums []int
		listOfNumsAsSting := strings.Split(strings.TrimSpace(lineSplit[1]), " ")
		for i := range listOfNumsAsSting {
			num, _ := strconv.Atoi(listOfNumsAsSting[i])

			listOfNums = append(listOfNums, num)
		}
		equations = append(equations, Equation{result: result, calibration: listOfNums})
	}

	for e := range equations {
		if isCalibrationValid(equations[e]) {
			counter += equations[e].result
		}
	}

	fmt.Println(counter)
}

func isCalibrationValid(e Equation) bool {
	addResult := getResult(e.calibration, 0, e.result, "+")
	mulResult := getResult(e.calibration, 0, e.result, "*")
	concatResult := getResult(e.calibration, 0, e.result, "||")

	return addResult || mulResult || concatResult
}

func getResult(equationNums []int, total int, target int, operation string) bool {

	var newTotal int
	switch operation {
	case "+":
		newTotal = total + equationNums[0]
	case "*":
		newTotal = total * equationNums[0]
	case "||":
		p1 := strconv.Itoa(total)
		p2 := strconv.Itoa(equationNums[0])
		p3, _ := strconv.Atoi(p1 + p2)
		newTotal = p3
	default:
		panic("Not a valid operation")
	}

	if len(equationNums) == 1 {
		return newTotal == target
	}

	addResult := getResult(equationNums[1:], newTotal, target, "+")
	mulResult := getResult(equationNums[1:], newTotal, target, "*")
	concatResult := getResult(equationNums[1:], newTotal, target, "||")

	return addResult || mulResult || concatResult
}
