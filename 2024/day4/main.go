package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// type Loc struct {
// 	x    int
// 	y    int
// 	dist int
// }

func SolveFindXmas(inputfile string) int {
	data := readInputToArray(&inputfile)
	numXmas := 0

	for row := range len(data) {
		for col := range len(data[0]) {
			currentItem := data[row][col]
			if currentItem == "X" {
				//move in each direction, stepping along 'check' variable
				location := []int{row, col}
				if checkDirection("up", &data, location) {
					numXmas++
				}
				if checkDirection("left", &data, location) {
					numXmas++
				}
				if checkDirection("right", &data, location) {
					numXmas++
				}
				if checkDirection("down", &data, location) {
					numXmas++
				}
				if checkDirection("upleft", &data, location) {
					numXmas++
				}
				if checkDirection("upright", &data, location) {
					numXmas++
				}
				if checkDirection("downleft", &data, location) {
					numXmas++
				}
				if checkDirection("downright", &data, location) {
					numXmas++
				}
			}
		}
	}
	fmt.Println(numXmas)

	return 0
}

func checkDirection(direction string, data *[][]string, location []int) bool {
	var checkArray []string
	for _, v := range "XMAS" {
		checkArray = append(checkArray, string(v))
	}

	checkIndex := 0
	i := 0
	d := *data

	if direction == "up" {
		for i < len(checkArray) {
			if location[0]-i < 0 {
				return false
			}
			//fmt.Println(location)
			//fmt.Println(checkArray[checkIndex], d[location[0]-i][location[1]])
			if checkArray[checkIndex] != d[location[0]-i][location[1]] {
				return false
			}
			i++
			checkIndex++
		}
	}
	if direction == "left" {
		for i < len(checkArray) {
			if location[1]-i < 0 {
				return false
			}
			if checkArray[checkIndex] != d[location[0]][location[1]-i] {
				return false
			}
			i++
			checkIndex++
		}
	}

	if direction == "right" {
		for i < len(checkArray) {
			if location[1]+i > len(d[location[0]])-1 {
				return false
			}
			if checkArray[checkIndex] != d[location[0]][location[1]+i] {
				return false
			}
			i++
			checkIndex++
		}
	}

	if direction == "down" {
		for i < len(checkArray) {
			if location[0]+i > len(d)-1 {
				return false
			}
			if checkArray[checkIndex] != d[location[0]+i][location[1]] {
				return false
			}
			i++
			checkIndex++
		}
	}

	if direction == "upleft" {
		for i < len(checkArray) {
			if location[0]-i < 0 || location[1]-i < 0 {
				return false
			}
			if checkArray[checkIndex] != d[location[0]-i][location[1]-i] {
				return false
			}
			i++
			checkIndex++
		}
	}

	if direction == "upright" {
		for i < len(checkArray) {
			if location[0]-i < 0 || location[1]+i > len(d[location[0]])-1 {
				return false
			}
			if checkArray[checkIndex] != d[location[0]-i][location[1]+i] {
				return false
			}
			i++
			checkIndex++
		}
	}

	if direction == "upleft" {
		for i < len(checkArray) {
			if location[0]-i < 0 || location[1]-i < 0 {
				return false
			}
			if checkArray[checkIndex] != d[location[0]-i][location[1]-i] {
				return false
			}
			i++
			checkIndex++
		}
	}

	if direction == "downleft" {
		for i < len(checkArray) {
			if location[0]+i > len(d)-1 || location[1]-i < 0 {
				return false
			}
			if checkArray[checkIndex] != d[location[0]+i][location[1]-i] {
				return false
			}
			i++
			checkIndex++
		}
	}

	if direction == "downright" {
		for i < len(checkArray) {
			if location[0]+i > len(d)-1 || location[1]+i > len(d[location[0]])-1 {
				return false
			}
			if checkArray[checkIndex] != d[location[0]+i][location[1]+i] {
				return false
			}
			i++
			checkIndex++
		}
	}

	fmt.Println(location, direction)
	return true
	//returns whether direction contains "MAS"

}

func readInputToArray(inputfile *string) [][]string {
	file, _ := os.Open(*inputfile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var output [][]string

	for scanner.Scan() {
		item := scanner.Text()
		items := strings.Split(item, "")
		output = append(output, items)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return output
}
