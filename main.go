package main

import (
	"github.com/asim-a/aoc2024/days"
	"github.com/asim-a/aoc2024/lib"
)

var dayMap = map[int]func(lib.Arguments) lib.Output{
	1: days.InitDayOne,
	2: days.InitDayTwo,
	3: days.InitDayThree,
}

func main() {
	arguments := lib.GetArguments()
	fp, ok := dayMap[arguments.Day]
	if !ok {
		panic("I have not got that far yet")
	}
	ctx := Context{
		Solution: fp(arguments),
	}

	if arguments.Day <= 0 || arguments.Day > 25 {
		panic("invalid day, only valid days are between 1 and 25")
	}

	ctx.Solution.Show()
}

type Context struct {
	Solution lib.Output
}
