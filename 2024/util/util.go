package util

import "strconv"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func AppendStringToIntArray(array []int, element string) []int {
	el, err := strconv.Atoi(element)
	Check(err)
	array = append(array, el)
	return array
}
