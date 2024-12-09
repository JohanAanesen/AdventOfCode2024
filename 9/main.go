package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type DiskThing struct {
	length int
	value  string
	moved  bool
}

func main() {

	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	diskmap := make([]string, 0)

	diskmap2 := make([]DiskThing, 0)
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

				diskmap2 = append(diskmap2, DiskThing{value: "free", length: inValue, moved: false})
			} else {
				for j := 0; j < inValue; j++ {
					diskmap = append(diskmap, strconv.Itoa(blockNameCounter))
				}

				diskmap2 = append(diskmap2, DiskThing{value: strconv.Itoa(blockNameCounter), length: inValue, moved: false})
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

	for i := len(diskmap2) - 1; i >= 0; i-- {

		if diskmap2[i].value != "free" && !diskmap2[i].moved {

			for j := 0; j < len(diskmap2); j++ {

				if i <= j {
					break
				}

				//finne første sted hvor det er en verdi som er free og lengde >= diskmap2[i].value
				if diskmap2[j].value == "free" && diskmap2[j].length >= diskmap2[i].length {
					//flytte verdien <-
					if diskmap2[j].length == diskmap2[i].length {
						diskmap2[j].value = diskmap2[i].value
						diskmap2[j].moved = true

						diskmap2[i].value = "free"
						diskmap2[i].moved = false

					} else {
						diskmap2[j].value = diskmap2[i].value
						lengdeDif := diskmap2[j].length - diskmap2[i].length
						diskmap2[j].length = diskmap2[i].length
						diskmap2[j].moved = true

						diskmap2[i].value = "free" //skal stå
						diskmap2[i].moved = false
						//passe på empty spaces (sette inn free verdi med rest lengde midt i slice)
						diskmap2 = slices.Insert(diskmap2, j+1, DiskThing{length: lengdeDif, value: "free", moved: false})
					}

					break

				}

			}

		}

	}

	diskmap2Ferdig := make([]string, 0)
	for i := 0; i < len(diskmap2); i++ {
		for j := 0; j < diskmap2[i].length; j++ {
			diskmap2Ferdig = append(diskmap2Ferdig, diskmap2[i].value)
		}
	}

	checksum2 := 0

	for i := 0; i < len(diskmap2Ferdig); i++ {
		if diskmap2Ferdig[i] != "free" {
			intVal, err := strconv.Atoi(diskmap2Ferdig[i])
			if err != nil {
				fmt.Println(err)
				continue
			}
			checksum2 += intVal * i
		} else {
			diskmap2Ferdig[i] = "."
		}
	}

	//fmt.Println(diskmap2Ferdig)

	fmt.Println("Part 2: ", checksum2)
}
