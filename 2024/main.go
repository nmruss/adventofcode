package main

import (
	"fmt"
	"nmruss/adventofcode24/day5"
)

func main() {
	//fmt.Println(day1.SolveTotalDistance("day1/inputfiles/testinput"))
	// fmt.Println(day1.SolveSimilarityScore("day1/inputfiles/input"))

	// fmt.Println(day2.SolveSafeLevels("day2/inputfiles/testinput"))
	// fmt.Println(day2.SolveSafeLevels("day2/inputfiles/input"))

	//fmt.Println(day3.SolveMulOps("day3/inputfiles/testinput"))
	//fmt.Println(day3.SolveMulOps("day3/inputfiles/input"))

	//fmt.Println(day4.SolveFindXmas("day4/inputfiles/testinput"))
	//fmt.Println(day4.SolveFindXmas("day4/inputfiles/input"))

	ruleMap := day5.BuildRules("day5/inputfiles/ruleinput")
	validList := day5.ValidUpdatesList("day5/inputfiles/updateinput", &ruleMap)
	fmt.Println(day5.SolvePrintQueueNum("day5/inputfiles/updateinput", validList))
}
