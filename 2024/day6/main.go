package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	Up Direction = iota + 0
	Right
	Down
	Left
)

func SolveGuardPositions(inputfile string) int {
	//returns the number of visited guard positions based on input string
	var area [][]string
	var startingPos []int
	ReadInputToArray(&inputfile, &area, &startingPos)

	fmt.Println(startingPos)
	//from the starting position, move in the current direction until you hit out of bounds, or "#"
	//if you havent hit "#" or out of bounds, and space is ".": add 1 to visited spaces, and mark the space with an "X"
	//if you have hit "#", add one to direction (turning right) continue moving in that direction
	var currentPos []int = []int{startingPos[0], startingPos[1]}
	var d Direction = Up
	tilesCrossed := 1

	//while in bounds
	for currentPos[0] >= 0 && currentPos[0] < len(area) && currentPos[1] >= 0 && currentPos[1] < len(area[0]) {
		if d == Up {
			if currentPos[0]-1 < 0 {
				area[currentPos[0]][currentPos[1]] = "X"
				tilesCrossed++
				break
			}
			if area[currentPos[0]][currentPos[1]] == "." {
				tilesCrossed++
			}
			area[currentPos[0]][currentPos[1]] = "X"
			nextTile := area[currentPos[0]-1][currentPos[1]]
			if nextTile == "." {
				currentPos[0]--
			} else if nextTile == "#" {
				d = (d + 1) % 4
			} else if nextTile == "X" {
				currentPos[0]--
			}
		}

		if d == Right {
			if currentPos[1]+1 > len(area[0])-1 {
				area[currentPos[0]][currentPos[1]] = "X"
				tilesCrossed++
				break
			}
			if area[currentPos[0]][currentPos[1]] == "." {
				tilesCrossed++
			}
			nextTile := area[currentPos[0]][currentPos[1]+1]
			area[currentPos[0]][currentPos[1]] = "X"
			if nextTile == "." {
				currentPos[1]++
			} else if nextTile == "#" {
				d = (d + 1) % 4
			} else if nextTile == "X" {
				currentPos[1]++
			}
		}

		if d == Down {
			if currentPos[0]+1 > len(area)-1 {
				area[currentPos[0]][currentPos[1]] = "X"
				tilesCrossed++
				break
			}
			if area[currentPos[0]][currentPos[1]] == "." {
				tilesCrossed++
			}
			nextTile := area[currentPos[0]+1][currentPos[1]]
			area[currentPos[0]][currentPos[1]] = "X"
			if nextTile == "." {
				currentPos[0]++
			} else if nextTile == "#" {
				d = (d + 1) % 4
			} else if nextTile == "X" {
				currentPos[0]++
			}
		}

		if d == Left {
			if currentPos[1]-1 < 0 {
				area[currentPos[0]][currentPos[1]] = "X"
				tilesCrossed++
				break
			}
			if area[currentPos[0]][currentPos[1]] == "." {
				tilesCrossed++
			}
			nextTile := area[currentPos[0]][currentPos[1]-1]
			area[currentPos[0]][currentPos[1]] = "X"
			if nextTile == "." {
				currentPos[1]--
			} else if nextTile == "#" {
				d = (d + 1) % 4
			} else if nextTile == "X" {
				currentPos[1]--
			}
		}

	}
	PrintArea(&area)

	return tilesCrossed
}

func PrintArea(area *[][]string) {

	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, v := range *area {
		fmt.Fprint(file, v)
		fmt.Fprint(file, "\n")
	}

}

func ReadInputToArray(inputfile *string, area *[][]string, startingPos *[]int) {
	//reads input to 2d array as well as starting position
	file, _ := os.Open(*inputfile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	row := 0
	col := 0

	for scanner.Scan() {
		item := scanner.Text()
		items := strings.Split(item, "")
		for _, v := range items {
			if v == "^" {
				*startingPos = append(*startingPos, row)
				*startingPos = append(*startingPos, col)
			}
			col++
		}
		col = 0
		row++
		*area = append(*area, items)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
