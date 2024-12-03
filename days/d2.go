package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/asim-a/aoc2024/lib"
)

type D2 struct {
	Challenge lib.Challenge
}

type ChangeDirection int

const (
	Decreasing ChangeDirection = iota
	Increasing
	NotSet
)

func InitDayTwo(arguments lib.Arguments) lib.Output {
	return &D2{
		Challenge: lib.Challenge{
			Arguments: arguments,
		},
	}
}

func (e ChangeDirection) String() string {
	switch e {
	case NotSet:
		return "NotSet"
	case Decreasing:
		return "Decreasing"
	case Increasing:
		return "Increasing"
	default:
		return "Not supported"
	}
}

func IsRowSafe(numbers []string) bool {
	direction := NotSet

	for i := 0; i <= len(numbers)-2; i++ {
		left, err := strconv.Atoi(numbers[i])
		lib.Check(err)
		right, err := strconv.Atoi(numbers[i+1])
		lib.Check(err)

		if right > left {
			if direction == Decreasing {
				return false
			}
			direction = Increasing
		} else {
			if direction == Increasing {
				return false
			}
			direction = Decreasing
		}

		diff := math.Abs(float64(right - left))
		if diff > 3.0 || diff < 1.0 {
			return false
		}
	}
	return true
}

func (d *D2) A() int {
	lines := d.Challenge.GetLines()

	safeCount := 0
	for i, line := range lines {
		numbers := strings.Split(line, " ")
		isRowSafe := IsRowSafe(numbers)
		if isRowSafe {
			safeCount += 1
			fmt.Printf("(OK) %s >[%d]\n", line, i)
		}
	}
	fmt.Println("=== A DONE ===")

	return safeCount
}

func Dampener(numbers []string) bool {

	nLen := len(numbers)
	for j := 0; j < nLen; j += 1 {
		temp := make([]string, nLen)
		copy(temp, numbers)
		if IsRowSafe(lib.RemoveIndexString(temp, j)) {
			return true
		}
	}

	return false
}

func (d *D2) B() int {
	lines := d.Challenge.GetLines()

	safeCount := 0
	for i, line := range lines {
		numbers := strings.Split(line, " ")
		isRowSafe := IsRowSafe(numbers)
		if isRowSafe {
			safeCount += 1
			fmt.Printf("(OK) %s >[%d]\n", line, i)
		} else {
			if Dampener(numbers) {
				safeCount += 1
			}
		}
	}

	return safeCount
}

func (d *D2) Show() {
	isTest := d.Challenge.Arguments.IsTest
	var isTestString string
	if isTest {
		isTestString = "[Test]"
	}
	a := d.A()
	b := d.B()
	fmt.Printf("%s A: %d, B: %d\n", isTestString, a, b)
	if (a != 242 || b != 311) && !isTest {
		panic("a != 242 || b != 311")
	}
}
