package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func SolveMulOps(inputfile string) int {
	mulOpResults := readInput(&inputfile)
	var ans int
	for _, v := range mulOpResults {
		ans += v
	}
	return ans
}

func readInput(inputfile *string) []int {
	file, _ := os.Open(*inputfile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var multOpResults []int

	mulMatchObj, mulMatchErr := regexp.Compile(`mul\(\d+\,\d+\)`)
	if mulMatchErr != nil {
		fmt.Println(mulMatchErr)
	}

	// var left int = 0
	// var right int = 0
	// var op []rune
	// var opening = "mul("
	// var oIndex int = 0
	var matches []string

	for scanner.Scan() {
		item := scanner.Text()
		matches = mulMatchObj.FindAllString(item, -1)
		//fmt.Println(matches)

		for _, mulValue := range matches {
			numMatchObj, numMatchErr := regexp.Compile(`\d+`)
			if numMatchErr != nil {
				fmt.Println(numMatchErr)
			}
			nums := numMatchObj.FindAllString(mulValue, -1)
			left, err := strconv.Atoi(nums[0])
			if err != nil {
				fmt.Println(err)
			}

			right, err := strconv.Atoi(nums[1])
			if err != nil {
				fmt.Println()
			}

			multOpResults = append(multOpResults, left*right)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return multOpResults
}
