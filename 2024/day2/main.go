package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func SolveSafeLevels(inputfile string) int {
	var l Levels
	var answer int = 0
	l = readInput(&inputfile)

	for _, v := range l.levelSafety {
		if v {
			answer++
		}
	}

	return answer
}

type Levels struct {
	levelSafety []bool
}

func readInput(inputfile *string) Levels {
	file, _ := os.Open(*inputfile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var l Levels

	for scanner.Scan() {
		item := scanner.Text()
		safe := true
		direction := ""
		var lastDirection string

		strArr := strings.Split(item, " ")

		last, err := strconv.Atoi(strArr[0])
		if err != nil {
			fmt.Println(err)
		}

		for _, v := range strArr[1:] {
			//fmt.Println(v)
			n, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
			}

			if n-last > 0 {
				direction = "up"
			}

			if n-last < 0 {
				direction = "down"
			}

			if lastDirection != "" && lastDirection != direction {
				safe = false
			}

			if math.Abs(float64(n-last)) > 3 || math.Abs(float64(n-last)) == 0 {
				safe = false
			}

			lastDirection = direction
			last = n
		}

		l.levelSafety = append(l.levelSafety, safe)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return l
}
