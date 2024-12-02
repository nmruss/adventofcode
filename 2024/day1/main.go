// read from the file, separate each of the lists into left and right arrays, sort the arrays by size
// do the additions, and produce the answer
package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

type Lists struct {
	left      []int
	right     []int
	leftFreq  map[int]int
	rightFreq map[int]int
}

func sortLists(inputLists *Lists) {
	//sorts a list object's left[] and right[] values in place
	slices.SortFunc(inputLists.left, func(a int, b int) int {
		if a > b {
			return 1
		} else {
			return -1
		}
	})
	slices.SortFunc(inputLists.right, func(a int, b int) int {
		if a > b {
			return 1
		} else {
			return -1
		}
	})
}

func SolveTotalDistance(inputfile string) int {
	//solves Day 1: Part 1 (Total Distance)
	var inputLists Lists = readInput(&inputfile)
	sortLists(&inputLists)

	i := 0
	ans := 0
	for i < len(inputLists.left) {
		num := math.Abs(float64(inputLists.left[i]) - float64(inputLists.right[i]))
		ans += int(num)

		i++
	}

	fmt.Println(inputLists)
	return ans
}

func SolveSimilarityScore(inputfile string) int {
	//Solve similarity score by stepping through the left list and comparing
	//each number to a frequency table on the right list

	var inputLists Lists = readInput(&inputfile)
	sortLists(&inputLists)
	var output int

	for _, v := range inputLists.left {
		rightVal, ok := inputLists.rightFreq[v]

		if ok {
			output += v * rightVal
		}
	}

	return output
}

func readInput(inputfile *string) Lists {
	//takes a pointer to a string and steps through the file contents
	//returning a 'Lists' struct with frequency and L/R array data
	file, _ := os.Open(*inputfile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	count := 0

	var l_list []int
	var r_list []int
	l_freq := make(map[int]int)
	r_freq := make(map[int]int)

	for scanner.Scan() {
		item := scanner.Text()
		num, err := strconv.Atoi(item)
		if err != nil {
			fmt.Println(err)
		}

		if count%2 == 0 {
			l_list = append(l_list, num)
			l_freq[num]++
		} else {
			r_list = append(r_list, num)
			r_freq[num]++
		}
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	var output Lists
	output.left = l_list
	output.right = r_list
	output.leftFreq = l_freq
	output.rightFreq = r_freq

	return output
}
