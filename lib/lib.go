package lib

import (
	"fmt"
	"os"

	opts "github.com/asim-a/aoc2024/opts"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileAsBytes(filename string) []byte {

	data, err := os.ReadFile(filename)
	Check(err)

	return data
}

func ReadFileAsString(filename string) string {
	return string(ReadFileAsBytes(filename))
}

func ReadLines(filename string, opts opts.ReadOptions) []string {

	string_content := ReadFileAsString(filename)

	result := []string{}
	s := ""

	for _, ch := range string_content {
		if ch == '\n' {
			result = append(result, s)
			if opts.IncludeNewLine {
				result = append(result, "\n")
			}
			s = ""
			continue
		}
		s += string(ch)
	}

	return result
}

func Config(day int, isTest bool) DayConfig {
	filenameLines := fmt.Sprintf("data/%d", day)
	filenameTestLines := fmt.Sprintf("%stest", filenameLines)
	lines := ReadLines(filenameLines, opts.ReadOptions{})
	testLines := ReadLines(filenameTestLines, opts.ReadOptions{})

	return DayConfig{
		Day:       day,
		IsTest:    isTest,
		Lines:     lines,
		TestLines: testLines,
	}
}

func (c *Challenge) GetLines(day int, isTest bool) []string {

	if c.Configuration.IsTest {
		return c.Configuration.TestLines
	} else {
		return c.Configuration.Lines
	}
}
