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

	diskmap := make([]string, 0)
	for r.Scan() {
		line := r.Text()
		input := strings.Split(line, "")

		freeSpace := false
		blockNameCounter := 0

		for i := 0; i < len(input); i++ {

			inValue, _ := strconv.Atoi(input[i])

			if freeSpace {
				for j := 0; j < inValue; j++ {
					diskmap = append(diskmap, "free")
				}
			} else {
				for j := 0; j < inValue; j++ {
					diskmap = append(diskmap, strconv.Itoa(blockNameCounter))
				}
				blockNameCounter++
			}

			freeSpace = !freeSpace
		}

	}

	for i := 0; i < len(diskmap); i++ {

		if diskmap[i] == "free" {
			for j := len(diskmap) - 1; j >= 0; j-- {

				if j <= i {
					break
				}

				if diskmap[j] != "free" {
					diskmap[i] = diskmap[j]
					diskmap[j] = "free"
					break
				}
			}
		}

	}

	checksum := 0

	for i := 0; i < len(diskmap); i++ {
		if diskmap[i] != "free" {
			intVal, err := strconv.Atoi(diskmap[i])
			if err != nil {
				continue
			}
			checksum += intVal * i
		}
	}

	fmt.Println("Part 1: ", checksum)

}
