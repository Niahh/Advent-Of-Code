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

type report struct {
	levels []int
}

func (r report) isLevelsIncrOrDecr() bool {
	if r.isLevelIncr() || r.isLevelDecr() {
		return true
	}
	return false
}

func (r report) isLevelIncr() bool {
	return slices.IsSorted(r.levels)
}

func (r report) isLevelDecr() bool {
	return slices.IsSortedFunc(r.levels, func(a, b int) int {
		if a == b {
			return 0
		} else if a < b {
			return 1
		} else {
			return -1
		}
	})
}

func (r report) isAdjacentLevelDifferenceAcceptable() bool {
	for i := 0; i < len(r.levels)-1; i++ {
		differ := math.Abs(float64(r.levels[i] - r.levels[i+1]))
		if !(differ >= 1 && differ <= 3) {
			return false
		}
	}
	return true
}

func (r report) isReportSafe() bool {
	return r.isLevelsIncrOrDecr() && r.isAdjacentLevelDifferenceAcceptable()
}

func (r report) removeLevel(index int) []int {
	var newReport []int
	for i, elem := range r.levels {
		if i != index {
			newReport = append(newReport, elem)
		}
	}
	return newReport
}

func (r report) isReportSafeProblemDampener() bool {
	if r.isReportSafe() {
		return true
	}

	for i := range r.levels {
		levelDanpened := report{r.removeLevel(i)}
		if levelDanpened.isReportSafe() {
			return true
		}
	}

	return false
}

func Part1() string {
	levels := strings.Split(input, "\n")
	reportsSafe := 0
	for _, reportStr := range levels {
		report := report{util.ConvertStringArrayToInt(strings.Fields(reportStr))}
		if reportStr != "" && report.isReportSafe() {
			reportsSafe++
		}
	}
	return strconv.Itoa(reportsSafe)
}

func Part2() string {
	reports := strings.Split(input, "\n")
	reportsSafe := 0

	for _, reportStr := range reports {
		report := report{util.ConvertStringArrayToInt(strings.Fields(reportStr))}
		if reportStr != "" && report.isReportSafeProblemDampener() {
			reportsSafe++
		}
	}

	return strconv.Itoa(reportsSafe)
}
