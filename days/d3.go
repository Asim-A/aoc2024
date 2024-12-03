package days

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/asim-a/aoc2024/lib"
)

type D3 struct {
	Challenge lib.Challenge
}

func InitDayThree(arguments lib.Arguments) lib.Output {
	return &D3{
		Challenge: lib.Challenge{
			Arguments: arguments,
		},
	}
}

func (d *D3) A() int {
	input := d.Challenge.GetInput()
	multiplications := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := multiplications.FindAllStringSubmatch(input, -1)

	sum := 0

	for _, m := range matches {
		left, err := strconv.Atoi(m[1])
		lib.Check(err)
		right, err := strconv.Atoi(m[2])
		lib.Check(err)
		sum += left * right
	}

	return sum
}

func (d *D3) B() int {
	input := string(d.Challenge.GetInput())

	cleaner := regexp.MustCompile(`(?s)don't\(\).*?(?:do\(\)|$)`)
	input = cleaner.ReplaceAllString(input, "")

	multiplications := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := multiplications.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, m := range matches {
		left, err := strconv.Atoi(m[1])
		lib.Check(err)
		right, err := strconv.Atoi(m[2])
		lib.Check(err)
		sum += left * right
	}

	return sum
}

func (d *D3) Show() {
	isTest := d.Challenge.Arguments.IsTest
	var isTestString string
	if isTest {
		isTestString = "[Test]"
	}
	a := d.A()
	b := d.B()
	fmt.Printf("%s A: %d, B: %d\n", isTestString, a, b)
	if (a != 190604937 || b != 82857512) && !isTest {
		panic("a != 190604937 || b != 82857512")
	}
}
