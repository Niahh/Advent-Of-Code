package day2

import (
	_ "embed"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/Niahh/Advent-Of-Code/2024/util"
)

//go:embed input.txt
var input string

func isLevelIncrOrDecr(level []int) bool {
	if isLevelIncr(level) || isLevelDecr(level) {
		return true
	}

	return false
}
func isLevelIncr(level []int) bool {
	return slices.IsSorted(level)
}

func isLevelDecr(level []int) bool {
	return slices.IsSortedFunc(level, func(a, b int) int {
		if a == b {
			return 0
		} else if a < b {
			return 1
		} else {
			return -1
		}
	})
}

func isAdjacentLevelDifferenceAcceptable(intLevels []int) bool {
	for i := 0; i < len(intLevels)-1; i++ {
		differ := math.Abs(float64(intLevels[i] - intLevels[i+1]))
		if !(differ >= 1 && differ <= 3) {
			return false
		}
	}
	return true
}

func isReportSafe(report []int) bool {
	return isLevelIncrOrDecr(report) && isAdjacentLevelDifferenceAcceptable(report)
}

func removeLevel(index int, levels []int) []int {
	var newReport []int
	for i, elem := range levels {
		if i != index {
			newReport = append(newReport, elem)
		}
	}
	return newReport
}

func isReportSafeProblemDampener(levels []int) bool {
	if isReportSafe(levels) {
		return true
	}

	for i := range levels {
		levelDanpened := removeLevel(i, levels)
		if isReportSafe(levelDanpened) {
			return true
		}
	}

	return false
}

func Part1() string {
	levels := strings.Split(input, "\n")
	reportsSafe := 0
	for _, report := range levels {
		if report != "" && isReportSafe(util.ConvertStringArrayToInt(strings.Fields(report))) {
			reportsSafe++
		}
	}
	return strconv.Itoa(reportsSafe)
}

func Part2() string {
	reports := strings.Split(input, "\n")
	reportsSafe := 0

	for _, report := range reports {
		if report != "" && isReportSafeProblemDampener(util.ConvertStringArrayToInt(strings.Fields(report))) {
			reportsSafe++
		}
	}

	return strconv.Itoa(reportsSafe)
}
