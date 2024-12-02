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

func Solve(inputfile string) int {
	var inputLists Lists = readInput(&inputfile)

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

type Lists struct {
	left  []int
	right []int
}

func readInput(inputfile *string) Lists {
	//takes a pointer to a string and steps through the file contents
	//returning an array of split nums
	file, _ := os.Open(*inputfile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	count := 0

	var l_list []int
	var r_list []int

	for scanner.Scan() {
		item := scanner.Text()
		num, err := strconv.Atoi(item)
		if err != nil {
			fmt.Println(err)
		}

		if count%2 == 0 {
			l_list = append(l_list, num)
		} else {
			r_list = append(r_list, num)
		}
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	var output Lists
	output.left = l_list
	output.right = r_list

	return output
}
