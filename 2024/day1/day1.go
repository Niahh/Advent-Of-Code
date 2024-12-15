package day1

import (
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Niahh/Advent-Of-Code/2024/util"
)

// Relative path to where the code is executed.
const inputFilepath = "day1/input-part1.txt"

func parseInputFile() ([]int, []int) {
	util.CheckIfFileExists(inputFilepath)

	// This methods retrive all the text from the file.
	data, err := os.ReadFile(inputFilepath)
	util.Check(err)

	// We get all the location ID in one array.
	locationIds := strings.Fields(string(data))

	var locationIdsLeft []int
	var locationIdsRight []int

	// We organize this giant array into left and right arrays
	for i := 0; i < len(locationIds); i++ {
		number, err := strconv.Atoi(locationIds[i])
		util.Check(err)

		if i%2 == 0 {
			locationIdsLeft = append(locationIdsLeft, number)
		} else {
			locationIdsRight = append(locationIdsRight, number)
		}

	}
	return locationIdsLeft, locationIdsRight
}

func Part1() string {
	leftIds, rightIds := parseInputFile()
	sort.Ints(leftIds)
	sort.Ints(rightIds)

	sum := 0
	for i := 0; i < len(leftIds); i++ {
		sum += absolute(leftIds[i] - rightIds[i])
	}
	return strconv.Itoa(sum)
}

func Part2() string {
	leftIds, rightIds := parseInputFile()
	sum := 0
	for i := 0; i < len(leftIds); i++ {
		sum += similarityScore(leftIds[i], rightIds)
	}

	return strconv.Itoa(sum)
}

func similarityScore(leftId int, rightIds []int) int {
	occurence := 0
	for i := 0; i < len(rightIds); i++ {
		if leftId == rightIds[i] {
			occurence++
		}
	}
	return leftId * occurence
}

func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
