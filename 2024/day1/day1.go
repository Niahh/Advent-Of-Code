package day1

import (
	_ "embed"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/Niahh/Advent-Of-Code/2024/util"
)

//go:embed input-part1.txt
var input string

func parseInputFile() ([]int, []int) {
	locationIds := strings.Split(input, "\n")
	var locationIdsLeft, locationIdsRight []int

	for _, element := range locationIds {
		locations := strings.Fields(element)
		locationIdsLeft = util.AppendStringToIntArray(locationIdsLeft, locations[0])
		locationIdsRight = util.AppendStringToIntArray(locationIdsRight, locations[1])
	}
	slices.Sort(locationIdsLeft)
	slices.Sort(locationIdsRight)
	return locationIdsLeft, locationIdsRight
}

func Part1() string {
	leftIds, rightIds := parseInputFile()
	sum := 0
	for i := 0; i < len(leftIds); i++ {

		sum += int(math.Abs(float64(leftIds[i] - rightIds[i])))
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
