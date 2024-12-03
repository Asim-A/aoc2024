package days

import (
	"fmt"
	"github.com/asim-a/aoc2024/lib"
	"sort"
	"strconv"
)

type D1 struct {
	Challenge lib.Challenge
}

func InitDayOne(arguments lib.Arguments) lib.Output {
	return &D1{
		Challenge: lib.Challenge{
			Arguments: arguments,
		},
	}
}

func (d *D1) Show() {
	isTest := d.Challenge.Arguments.IsTest
	var isTestString string
	if isTest {
		isTestString = "[Test]"
	}
	a := d.A()
	b := d.B()
	fmt.Printf("%s A: %d, B: %d\n", isTestString, a, b)
	if (a != 1590491 || b != 22588371) && !isTest {
		panic("a != 1590491 || b != 22588371")
	}
}

func (d *D1) A() int {
	lines := d.Challenge.GetLines()

	firstColumn, secondColumn := toIntColumns(lines)
	if len(firstColumn) != len(secondColumn) {
		panic("different length")
	}
	sort.Ints(firstColumn)
	sort.Ints(secondColumn)

	res := 0
	for i := 0; i < len(firstColumn); i++ {
		left := firstColumn[i]
		right := secondColumn[i]

		if left > right {
			res += (left - right)
		} else {

			res += (right - left)
		}
	}

	return res

}
func (d *D1) B() int {
	lines := d.Challenge.GetLines()

	leftColumn, rightColumn := toIntColumns(lines)
	rightCountMap := map[int]int{}
	sum := 0

	for _, r := range rightColumn {
		_, ok := rightCountMap[r]
		if ok {
			rightCountMap[r] = rightCountMap[r] + 1
			continue
		}
		rightCountMap[r] = 1
	}

	for _, l := range leftColumn {
		v, ok := rightCountMap[l]
		if ok {
			sum += (l * v)
		}
	}

	return sum
}

func toIntColumns(lines []string) ([]int, []int) {

	firstColumn := []int{}
	secondColumn := []int{}
	for _, line := range lines {

		for i, r := range line {
			if r == ' ' {
				first, err1 := strconv.Atoi(string(line[0:i]))
				lib.Check(err1)
				firstColumn = append(firstColumn, first)

				second, err2 := strconv.Atoi(string(line[i+3:]))
				secondColumn = append(secondColumn, second)
				lib.Check(err2)
				break
			}

		}
	}
	return firstColumn, secondColumn
}
