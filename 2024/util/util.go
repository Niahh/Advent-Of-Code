package util

import (
	"errors"
	"fmt"
	"os"
)

func CheckIfFileExists(filename string) {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("File %s does not exists\n", filename)
		os.Exit(1)
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
