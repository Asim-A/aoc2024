package days

import (
	"fmt"

	"github.com/asim-a/aoc2024/lib"
)

type D4 struct {
	Challenge lib.Challenge
}

func InitDayFour(arguments lib.Arguments) lib.Output {
	return &D4{
		Challenge: lib.Challenge{
			Arguments: arguments,
		},
	}
}

func (d *D4) A() int {
	sum := 0
	input := d.Challenge.GetInput()

	for i := 0; i < len(input)-4; i++ {
		x := input[i]
		m := input[i+1]
		a := input[i+2]
		s := input[i+3]
		if x == 'x' && m == 'm' && a == 'a' && s == 's' {
			sum += 1
		} else if x == 's' && m == 'a' && a == 'm' && s == 'x' {
			sum += 1
		}
	}
	//
	//
	//
	//x x x x
	//x m x x
	//x x a x
	//x x x s
	lines := d.Challenge.GetLines()
	for i := 0; i < len(lines); i += 4 {
		for j := 0; j < 4; j++ {
		}
	}

	return sum
}

func (d *D4) B() int {
	sum := 0

	return sum
}

func (d *D4) Show() {

	isTest := d.Challenge.Arguments.IsTest
	var isTestString string
	if isTest {
		isTestString = "[Test]"
	}
	a := d.A()
	b := d.B()
	fmt.Printf("%s A: %d, B: %d\n", isTestString, a, b)
	// if (a != 190604937 || b != 82857512) && !isTest {
	// 	panic("a != 190604937 || b != 82857512")
	// }
}
