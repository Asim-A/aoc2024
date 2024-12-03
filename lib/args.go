package lib

import (
	"fmt"
	"os"
	"strconv"
)

type Range = [2]int

type Arguments struct {
	Day    int
	IsTest bool
}

type Run interface {
	Run() Output
}
type Help interface {
	PrintHelp()
}

func PrintHelp() {
	fmt.Println(`
		Use --day [int] to specify which day you want to execute.
		E.g. --day 2 will give you the result for the second day of AoC

		Use --test to specify that you want to use test input`)
}

func ResolveIntInput(value string) int {
	intVal, err := strconv.Atoi(value)
	Check(err)
	return intVal
}

func ResolveRangeInput(value string) Range {

	r := [2]int{}
	for i, ch := range value {
		if ch == '.' {
			r[0] = ResolveIntInput(value[:i-1])
			r[1] = ResolveIntInput(value[i+2:])
		}
	}

	return r
}

func GetArguments() Arguments {
	lenAllArgs := len(os.Args)
	if lenAllArgs < 2 {
		panic("not enough args")
	}
	arguments := os.Args[1:]
	amountOfArgs := len(arguments)

	a := Arguments{}
	for i := 0; i < amountOfArgs; i++ {
		key := arguments[i]
		switch key {
		case "--help":
			PrintHelp()
			os.Exit(0)
		case "--test":
			a.IsTest = true
		case "--day":
			a.Day = ResolveIntInput(arguments[i+1])
		}

	}
	return a
}
