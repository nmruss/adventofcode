package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//build a function that ingests a list of rules
//and creates a rule map where you can easily look up
//rules for each page

// loop through each update, and for each page to print
// check that to make sure that you havent alredy printed
// one of the pages that it needs to be printed before
func ValidUpdatesList(inputfile string, ruleMap *map[string][]string) []bool {
	//takes in an ruleMap and a set of update data
	//determines if the updates are valid
	file, _ := os.Open(inputfile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	valids := []bool{}

	for scanner.Scan() {
		item := scanner.Text()
		items := strings.Split(item, ",")
		seen := make(map[string]bool)
		//if any alredy seen item
		//is in the current items list of rules
		rules := *ruleMap
		valid := true
		for _, v := range items {
			for _, r := range rules[v] {
				if seen[r] {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
			seen[v] = true
		}
		valids = append(valids, valid)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return valids
}

func SolvePrintQueueNum(inputfile string, valids []bool) int {

	file, _ := os.Open(inputfile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	answer := 0
	i := 0

	for scanner.Scan() {
		item := scanner.Text()
		items := strings.Split(item, ",")

		if valids[i] {
			mid := len(items) / 2
			midval, err := strconv.Atoi(items[mid])
			if err != nil {
				fmt.Println(err)
			}
			answer += midval
		}

		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return answer
}

func BuildRules(inputfile string) map[string][]string {
	//returns a one to many int mapping of rules based on an input
	ruleMap := make(map[string][]string)
	ReadInputToMap(&inputfile, &ruleMap)
	return ruleMap
}

func ReadInputToMap(inputfile *string, ruleMap *map[string][]string) {
	file, _ := os.Open(*inputfile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		item := scanner.Text()
		items := strings.Split(item, "|")
		m := *ruleMap
		m[items[0]] = append(m[items[0]], items[1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
