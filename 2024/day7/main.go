package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int64
	Left  *Node
	Right *Node
}

// const (
// 	left  = iota
// 	right = iota
// )

type Direction int

func (n *Node) Insert(value int64) {
	if n.Left == nil {
		n.Left = &Node{Value: n.Value + value}
	} else {
		n.Left.Insert(value)
	}

	if n.Right == nil {
		n.Right = &Node{Value: n.Value * value}
	} else {
		n.Right.Insert(value)
	}
}

func (n *Node) InOrder() {
	if n == nil {
		return
	}

	n.Left.InOrder()
	fmt.Println(n.Value)
	n.Right.InOrder()
}

func (n *Node) IsValid(target *int64) bool {
	if n == nil {
		return false
	}

	if n.Value == *target {
		return true
	}

	return n.Left.IsValid(target) || n.Right.IsValid(target)
}

func isPossible(target int64, nums []int64, current int64, index int) bool {
	//step along the input array recursively
	if index == 0 {
		return isPossible(target, nums, nums[0], 1)
	}

	if index == len(nums)-1 {
		return (current+nums[index] == target) || (current*nums[index] == target)
	}

	return isPossible(target, nums, current+nums[index], index+1) || isPossible(target, nums, current*nums[index], index+1)
}

func ReadBridgeRepairInput(inputfile *string) int64 {
	//return a sum of valid calibrations
	file, _ := os.Open(*inputfile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var answer int64

	for scanner.Scan() {
		item := scanner.Text()
		items := strings.Split(item, ":")
		target := items[0]
		items[1] = strings.Trim(items[1], " ")
		itemValues := strings.Split(items[1], " ")

		var itemNums []int64

		for _, value := range itemValues {
			currValue, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			itemNums = append(itemNums, currValue)
		}

		targetVal, err := strconv.ParseInt(target, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(targetVal, itemNums)

		if isPossible(targetVal, itemNums, 0, 0) {
			answer += targetVal
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return answer
}
