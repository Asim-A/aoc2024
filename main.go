package main

import (
	"github.com/asim-a/aoc2024/days"
	"github.com/asim-a/aoc2024/lib"
)

type Context struct {
	Solutions map[int]lib.Output
}

func (c *Context) Run(elements ...int) {
	for _, element := range elements {
		c.Solutions[element].Show()
	}
}

func (c *Context) AddRun(number int, outputer lib.Output) {
	c.Solutions[number] = outputer
}

func main() {
	ctx := Context{
		Solutions: map[int]lib.Output{},
	}
	ctx.AddRun(1, days.InitDayOne(false))
	ctx.Run(1)
}
