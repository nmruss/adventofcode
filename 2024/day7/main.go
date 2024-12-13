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

func ReadBridgeRepairInput(inputfile *string) int64 {
	//read each line item into a binary tree
	//as you move along, add each result, if the next node is ever == the test value
	//return as list of valid calibrations
	file, _ := os.Open(*inputfile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var answer int64

	for scanner.Scan() {
		item := scanner.Text()
		items := strings.Split(item, ":")
		testValue := items[0]
		items[1] = strings.Trim(items[1], " ")
		itemValues := strings.Split(items[1], " ")
		firstItemVal, err := strconv.ParseInt(itemValues[0], 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		root := Node{Value: firstItemVal}

		for _, value := range itemValues[1:] {
			currValue, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				fmt.Println(err)
			}

			root.Insert(currValue)
		}

		testVal, err := strconv.ParseInt(testValue, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		//	fmt.Println(testVal, root.IsValid(&testVal))
		if root.IsValid(&testVal) {
			answer += testVal
		}
		root.InOrder()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return answer
}
