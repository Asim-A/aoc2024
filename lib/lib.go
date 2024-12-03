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
	if err != nil {
		return []byte{}
	}

	return data
}

func ReadFileAsString(filename string) string {
	return string(ReadFileAsBytes(filename))
}

func ReadLines(filename string, opts opts.ReadOptions) []string {

	input := ReadFileAsString(filename)

	result := []string{}
	s := ""

	for _, ch := range input {
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

func RemoveIndexInt(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}
func RemoveIndexString(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func (c *Challenge) GetFilename() string {
	day := c.Arguments.Day
	if c.Arguments.IsTest {
		return fmt.Sprintf("data/%dtest", day)
	} else {
		return fmt.Sprintf("data/%d", day)
	}
}

func (c *Challenge) GetInput() string {
	filename := c.GetFilename()
	return ReadFileAsString(filename)
}

func (c *Challenge) GetLines() []string {
	filename := c.GetFilename()

	return ReadLines(filename, opts.ReadOptions{})
}
